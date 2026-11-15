package musica_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestInstallerLocalBinaryWritesManagedLauncher(t *testing.T) {
	home := t.TempDir()
	sourceBinary := filepath.Join(home, "source-musica")
	body := "#!/usr/bin/env bash\nif [[ \"${1:-}\" == \"version\" ]]; then printf '0.0.0\\n'; exit 0; fi\n"
	if err := os.WriteFile(sourceBinary, []byte(body), 0o755); err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("bash", "./install.sh", "from", sourceBinary)
	cmd.Env = append(os.Environ(), "HOME="+home)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("installer failed: %v\n%s", err, string(out))
	}

	launcher := filepath.Join(home, ".local", "bin", "musica")
	launcherText, err := os.ReadFile(launcher)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(launcherText), "# Managed by musica installer local-bin launcher") {
		t.Fatalf("launcher missing managed marker: %s", string(launcherText))
	}
	if !strings.Contains(string(out), "0.0.0") {
		t.Fatalf("installer output missing installed version: %s", string(out))
	}
}
