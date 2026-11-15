package store

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

func sqliteBin() (string, error) {
	if path, err := exec.LookPath("sqlite3"); err == nil {
		return path, nil
	}
	return "", fmt.Errorf("sqlite3 is not installed")
}

func sqliteExec(path, sql string) error {
	_, err := sqliteRun(path, sql, false)
	return err
}

func sqliteRun(path, sql string, asJSON bool) (string, error) {
	bin, err := sqliteBin()
	if err != nil {
		return "", err
	}
	args := []string{path}
	if asJSON {
		args = append(args, "-json")
	}
	args = append(args, sql)
	cmd := exec.Command(bin, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		detail := strings.TrimSpace(stderr.String())
		if detail == "" {
			detail = err.Error()
		}
		return "", fmt.Errorf("sqlite3: %s", detail)
	}
	return stdout.String(), nil
}

func queryJobs(path, sql string) ([]Job, error) {
	raw, err := sqliteRun(path, sql, true)
	if err != nil {
		return nil, err
	}
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, nil
	}
	var jobs []Job
	if err := json.Unmarshal([]byte(raw), &jobs); err != nil {
		return nil, fmt.Errorf("sqlite json: %w", err)
	}
	return jobs, nil
}

func quote(value string) string {
	return "'" + strings.ReplaceAll(value, "'", "''") + "'"
}
