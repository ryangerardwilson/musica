package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

const DefaultFPS = 20

type AppConfig struct {
	MusicDir     string
	FPS          int
	StartupTrack string
}

type UserConfig struct {
	MusicDir string `json:"music_dir,omitempty"`
	FPS      int    `json:"fps,omitempty"`
}

func DefaultConfigPath() string {
	if base := os.Getenv("XDG_CONFIG_HOME"); base != "" {
		return filepath.Join(ExpandPath(base), "musica", "config.json")
	}
	home, err := os.UserHomeDir()
	if err != nil || home == "" {
		home = "."
	}
	return filepath.Join(home, ".config", "musica", "config.json")
}

func LoadUserConfig(path string) UserConfig {
	if path == "" {
		path = DefaultConfigPath()
	}
	body, err := os.ReadFile(path)
	if err != nil {
		return UserConfig{}
	}
	var cfg UserConfig
	if err := json.Unmarshal(body, &cfg); err != nil {
		return UserConfig{}
	}
	return cfg
}

func SaveUserConfig(path string, updates UserConfig) error {
	if path == "" {
		path = DefaultConfigPath()
	}
	cfg := LoadUserConfig(path)
	if updates.MusicDir != "" {
		cfg.MusicDir = updates.MusicDir
	}
	if updates.FPS != 0 {
		cfg.FPS = updates.FPS
	}
	body, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	body = append(body, '\n')
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, body, 0o644)
}

func BootstrapText() []byte {
	return []byte("{\n  \"music_dir\": \"\",\n  \"fps\": 20\n}\n")
}

func EnsureConfigFile(path string) error {
	if path == "" {
		path = DefaultConfigPath()
	}
	if _, err := os.Stat(path); err == nil {
		return nil
	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, BootstrapText(), 0o644)
}

func AdoptLegacyConfig(path string) error {
	if path == "" {
		path = DefaultConfigPath()
	}
	if _, err := os.Stat(path); err == nil {
		return nil
	} else if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	legacy := filepath.Join(filepath.Dir(filepath.Dir(path)), "rlc", "config.json")
	body, err := os.ReadFile(legacy)
	if err != nil {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, body, 0o644)
}

func BuildAppConfig(cliMusicDir string, cliFPS *int, configPath string) AppConfig {
	userCfg := LoadUserConfig(configPath)
	musicDir := ResolveMusicDir(cliMusicDir, userCfg.MusicDir)

	fps := userCfg.FPS
	if cliFPS != nil {
		fps = *cliFPS
	}
	if fps == 0 {
		fps = DefaultFPS
	} else if fps < 5 {
		fps = 5
	}

	return AppConfig{
		MusicDir: musicDir,
		FPS:      fps,
	}
}

func ResolveMusicDir(cliValue, fileValue string) string {
	if cliValue != "" {
		return AbsPath(cliValue)
	}
	if fileValue != "" {
		return AbsPath(fileValue)
	}
	return AbsPath("~/Music")
}

func AbsPath(path string) string {
	expanded := ExpandPath(path)
	abs, err := filepath.Abs(expanded)
	if err != nil {
		return expanded
	}
	return abs
}

func ExpandPath(path string) string {
	if path == "" {
		return path
	}
	if path == "~" || len(path) > 2 && path[:2] == "~/" {
		home, err := os.UserHomeDir()
		if err == nil && home != "" {
			if path == "~" {
				return home
			}
			return filepath.Join(home, path[2:])
		}
	}
	return os.ExpandEnv(path)
}
