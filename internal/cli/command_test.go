package cli

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/ryangerardwilson/musica/internal/config"
)

func TestBareMusicaExecsR2d2Musica(t *testing.T) {
	restore := SetRunAppForTest(func(config.AppConfig) int {
		t.Fatal("bare musica should not open the player itself")
		return 1
	})
	defer restore()
	oldLook, oldExec := lookPath, execProcess
	defer func() { lookPath, execProcess = oldLook, oldExec }()
	lookPath = func(name string) (string, error) {
		if name != "r2d2" {
			t.Fatalf("looked up %q", name)
		}
		return "/opt/bin/r2d2", nil
	}
	var gotPath string
	var gotArgs []string
	execProcess = func(path string, args []string, env []string) error {
		gotPath, gotArgs = path, args
		return nil
	}
	var stdout, stderr bytes.Buffer
	if code := Command(nil, strings.NewReader(""), &stdout, &stderr); code != 0 {
		t.Fatalf("code %d %s", code, stderr.String())
	}
	if gotPath != "/opt/bin/r2d2" || len(gotArgs) != 2 || gotArgs[0] != "r2d2" || gotArgs[1] != "musica" {
		t.Fatalf("exec %q %q", gotPath, gotArgs)
	}
}

func TestBareMusicaFailsWithoutR2d2(t *testing.T) {
	oldLook, oldExec := lookPath, execProcess
	defer func() { lookPath, execProcess = oldLook, oldExec }()
	lookPath = func(string) (string, error) { return "", errors.New("not on PATH") }
	execProcess = func(string, []string, []string) error {
		t.Fatal("exec without a path")
		return nil
	}
	var stdout, stderr bytes.Buffer
	if code := Command(nil, strings.NewReader(""), &stdout, &stderr); code != 1 || !strings.Contains(stderr.String(), "r2d2 musica") {
		t.Fatalf("code %d %s", code, stderr.String())
	}
}

func TestMusicaWithArgumentsStaysInTheBinary(t *testing.T) {
	oldLook := lookPath
	defer func() { lookPath = oldLook }()
	lookPath = func(string) (string, error) {
		t.Fatal("a command with arguments must not look up r2d2")
		return "", nil
	}
	var stdout, stderr bytes.Buffer
	if code := Command([]string{"help"}, strings.NewReader(""), &stdout, &stderr); code != 0 || !strings.Contains(stdout.String(), "r2d2 musica") {
		t.Fatalf("code %d %s", code, stdout.String())
	}
}
