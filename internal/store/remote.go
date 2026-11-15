package store

import (
	"context"
	"os/exec"
	"time"
)

// Nebula sql is the Postgres interface and is not called. Nebula s3 backs
// this application's sqlite file up at s3://backups/musica.sqlite.

func restoreObject(path string) error {
	bin, args, err := s3Args("get", "s3://backups/musica.sqlite", path)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return exec.CommandContext(ctx, bin, args...).Run()
}

func publishObject(path string) {
	if path == "" {
		return
	}
	bin, args, err := s3Args("put", path, "s3://backups/musica.sqlite")
	if err != nil {
		return
	}
	go runS3(bin, args)
}

func runS3(bin string, args []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_ = exec.CommandContext(ctx, bin, args...).Run()
}

func s3Args(action string, rest ...string) (string, []string, error) {
	if bin, err := exec.LookPath("s3"); err == nil {
		return bin, append([]string{action}, rest...), nil
	}
	bin, err := exec.LookPath("nebula")
	if err != nil {
		return "", nil, err
	}
	return bin, append([]string{"s3", action}, rest...), nil
}
