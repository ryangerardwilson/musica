package cli

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ryangerardwilson/musica/internal/config"
)

func TestHelpAndVersionDoNotStartApp(t *testing.T) {
	restore := SetRunAppForTest(func(config.AppConfig) int {
		t.Fatal("runApp should not be called")
		return 1
	})
	defer restore()

	var stdout, stderr bytes.Buffer
	if code := Main([]string{"help"}, strings.NewReader(""), &stdout, &stderr); code != 0 {
		t.Fatalf("help exit = %d, stderr = %s", code, stderr.String())
	}
	if !strings.Contains(stdout.String(), "musica terminal music player") {
		t.Fatalf("help output missing expected text: %s", stdout.String())
	}

	stdout.Reset()
	if code := Main([]string{"version"}, strings.NewReader(""), &stdout, &stderr); code != 0 {
		t.Fatalf("version exit = %d", code)
	}
	if strings.TrimSpace(stdout.String()) == "" {
		t.Fatal("version output was empty")
	}
}

func TestOpenAndPlayValidation(t *testing.T) {
	var stdout, stderr bytes.Buffer
	if code := Main([]string{"open", filepath.Join(t.TempDir(), "missing")}, strings.NewReader(""), &stdout, &stderr); code == 0 {
		t.Fatal("open missing directory should fail")
	}
	if !strings.Contains(stderr.String(), "Music directory not found") {
		t.Fatalf("unexpected stderr: %s", stderr.String())
	}

	stderr.Reset()
	if code := Main([]string{"play", filepath.Join(t.TempDir(), "missing.mp3")}, strings.NewReader(""), &stdout, &stderr); code == 0 {
		t.Fatal("play missing track should fail")
	}
	if !strings.Contains(stderr.String(), "Track not found") {
		t.Fatalf("unexpected stderr: %s", stderr.String())
	}
}

func TestDefaultRunBootstrapsConfigAndStartsApp(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("HOME", filepath.Join(tmp, "home"))
	t.Setenv("XDG_CONFIG_HOME", filepath.Join(tmp, "xdg"))
	t.Setenv("MUSICA_DB", filepath.Join(tmp, "musica.sqlite"))

	called := false
	restore := SetRunAppForTest(func(cfg config.AppConfig) int {
		called = true
		if cfg.MusicDir != filepath.Join(tmp, "home", "Music") {
			t.Fatalf("MusicDir = %q", cfg.MusicDir)
		}
		return 0
	})
	defer restore()

	var stdout, stderr bytes.Buffer
	if code := Main(nil, strings.NewReader(""), &stdout, &stderr); code != 0 {
		t.Fatalf("default run exit = %d, stderr = %s", code, stderr.String())
	}
	if !called {
		t.Fatal("runApp was not called")
	}
	if _, err := os.Stat(filepath.Join(tmp, "xdg", "musica", "config.json")); err != nil {
		t.Fatalf("config was not written: %v", err)
	}
}

func TestEndStopsRunningSessionWithoutStartingApp(t *testing.T) {
	restore := SetRunAppForTest(func(config.AppConfig) int {
		t.Fatal("runApp should not be called")
		return 1
	})
	defer restore()

	previous := endSession
	defer func() { endSession = previous }()
	called := false
	endSession = func() error {
		called = true
		return nil
	}

	var stdout, stderr bytes.Buffer
	if code := Main([]string{"end"}, strings.NewReader(""), &stdout, &stderr); code != 0 {
		t.Fatalf("end exit = %d, stderr = %s", code, stderr.String())
	}
	if !called {
		t.Fatal("endSession was not called")
	}
	if strings.TrimSpace(stdout.String()) != "musica ended" {
		t.Fatalf("end output = %q", stdout.String())
	}
}

func TestEndReportsMissingSession(t *testing.T) {
	previous := endSession
	defer func() { endSession = previous }()
	endSession = func() error { return errors.New("musica is not running") }

	var stdout, stderr bytes.Buffer
	if code := Main([]string{"end"}, strings.NewReader(""), &stdout, &stderr); code == 0 {
		t.Fatal("end succeeded without a running session")
	}
	if !strings.Contains(stderr.String(), "musica is not running") {
		t.Fatalf("end stderr = %q", stderr.String())
	}
}
