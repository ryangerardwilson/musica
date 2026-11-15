package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultConfigPathUsesXDGConfigHome(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", filepath.Join(tmp, "xdg"))
	t.Setenv("HOME", filepath.Join(tmp, "home"))

	got := DefaultConfigPath()
	want := filepath.Join(tmp, "xdg", "musica", "config.json")
	if got != want {
		t.Fatalf("DefaultConfigPath() = %q, want %q", got, want)
	}
}

func TestSaveLoadAndBuildConfig(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "config.json")
	musicDir := filepath.Join(tmp, "Music")
	if err := os.MkdirAll(musicDir, 0o755); err != nil {
		t.Fatal(err)
	}

	if err := SaveUserConfig(path, UserConfig{MusicDir: musicDir, FPS: 3}); err != nil {
		t.Fatal(err)
	}

	loaded := LoadUserConfig(path)
	if loaded.MusicDir != musicDir {
		t.Fatalf("MusicDir = %q, want %q", loaded.MusicDir, musicDir)
	}
	cfg := BuildAppConfig("", nil, path)
	if cfg.MusicDir != musicDir {
		t.Fatalf("app MusicDir = %q, want %q", cfg.MusicDir, musicDir)
	}
	if cfg.FPS != 5 {
		t.Fatalf("FPS = %d, want floor 5", cfg.FPS)
	}
}

func TestEnsureConfigFileWritesBootstrap(t *testing.T) {
	path := filepath.Join(t.TempDir(), "musica", "config.json")
	if err := EnsureConfigFile(path); err != nil {
		t.Fatal(err)
	}
	body, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != string(BootstrapText()) {
		t.Fatalf("bootstrap body = %q", string(body))
	}
}
