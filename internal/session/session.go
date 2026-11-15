package session

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

const (
	controlSocketName = "control.sock"
	controlLockName   = "control.lock"
	endCommand        = "end"
)

var ErrAlreadyRunning = errors.New("musica is already running")

type Server struct {
	listener net.Listener
	lock     *os.File
	requests chan struct{}
	path     string
}

func Listen() (*Server, error) {
	path, err := socketPath()
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return nil, fmt.Errorf("create musica runtime directory: %w", err)
	}

	lock, err := acquireLock(filepath.Dir(path))
	if err != nil {
		return nil, err
	}
	if connection, err := net.DialTimeout("unix", path, 100*time.Millisecond); err == nil {
		_ = connection.Close()
		releaseLock(lock)
		return nil, ErrAlreadyRunning
	}
	if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
		releaseLock(lock)
		return nil, fmt.Errorf("remove stale musica control socket: %w", err)
	}

	listener, err := net.Listen("unix", path)
	if err != nil {
		releaseLock(lock)
		return nil, fmt.Errorf("listen on musica control socket: %w", err)
	}
	if err := os.Chmod(path, 0o600); err != nil {
		_ = listener.Close()
		_ = os.Remove(path)
		releaseLock(lock)
		return nil, fmt.Errorf("secure musica control socket: %w", err)
	}

	server := &Server{
		listener: listener,
		lock:     lock,
		requests: make(chan struct{}, 1),
		path:     path,
	}
	go server.serve()
	return server, nil
}

func (s *Server) EndRequests() <-chan struct{} {
	return s.requests
}

func (s *Server) Close() {
	if s == nil {
		return
	}
	_ = s.listener.Close()
	_ = os.Remove(s.path)
	releaseLock(s.lock)
}

func acquireLock(directory string) (*os.File, error) {
	lock, err := os.OpenFile(filepath.Join(directory, controlLockName), os.O_CREATE|os.O_RDWR, 0o600)
	if err != nil {
		return nil, fmt.Errorf("open musica session lock: %w", err)
	}
	if err := syscall.Flock(int(lock.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		_ = lock.Close()
		if errors.Is(err, syscall.EWOULDBLOCK) {
			return nil, ErrAlreadyRunning
		}
		return nil, fmt.Errorf("lock musica session: %w", err)
	}
	return lock, nil
}

func releaseLock(lock *os.File) {
	if lock == nil {
		return
	}
	_ = syscall.Flock(int(lock.Fd()), syscall.LOCK_UN)
	_ = lock.Close()
}

func (s *Server) serve() {
	for {
		connection, err := s.listener.Accept()
		if err != nil {
			return
		}
		go s.handle(connection)
	}
}

func (s *Server) handle(connection net.Conn) {
	defer connection.Close()
	_ = connection.SetDeadline(time.Now().Add(2 * time.Second))
	command, err := bufio.NewReader(connection).ReadString('\n')
	if err != nil {
		return
	}
	if strings.TrimSpace(command) != endCommand {
		_, _ = fmt.Fprintln(connection, "unknown command")
		return
	}
	select {
	case s.requests <- struct{}{}:
	default:
	}
	_, _ = fmt.Fprintln(connection, "ok")
}

func End() error {
	path, err := socketPath()
	if err != nil {
		return err
	}
	connection, err := net.DialTimeout("unix", path, 500*time.Millisecond)
	if err != nil {
		return errors.New("musica is not running")
	}
	defer connection.Close()
	_ = connection.SetDeadline(time.Now().Add(2 * time.Second))
	if _, err := fmt.Fprintln(connection, endCommand); err != nil {
		return fmt.Errorf("send musica end command: %w", err)
	}
	response, err := bufio.NewReader(connection).ReadString('\n')
	if err != nil {
		return fmt.Errorf("read musica end response: %w", err)
	}
	if strings.TrimSpace(response) != "ok" {
		return fmt.Errorf("musica rejected end command: %s", strings.TrimSpace(response))
	}
	return nil
}

func socketPath() (string, error) {
	runtimeDir := strings.TrimSpace(os.Getenv("XDG_RUNTIME_DIR"))
	if runtimeDir == "" {
		runtimeDir = filepath.Join(os.TempDir(), fmt.Sprintf("musica-%d", os.Getuid()))
	}
	if !filepath.IsAbs(runtimeDir) {
		return "", errors.New("XDG_RUNTIME_DIR must be an absolute path")
	}
	return filepath.Join(runtimeDir, "musica", controlSocketName), nil
}
