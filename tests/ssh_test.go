package tests

import (
	"testing"
	"os/exec"
	"time"
	"bytes"
	"strings"

	"github.com/jean0t/testing-fail2ban/internal/logging"
)

func TestSSHFailedToAuthenticate(t *testing.T) {
	var err error
	var timeNow string
	t.Run("Testing Logging", func(t *testing.T) {
		var sshLog *logging.SSHLogger = logging.NewSSHLogger()
		var ip, username, port string = "192.108.3.2", "test", "22" 
		timeNow = time.Now().Format("15:04:05")
		err = sshLog.LogFailedAttempt(ip, username, port)
		if err != nil {
			t.Error("Error when writing to journal")
		}

		var command *exec.Cmd = exec.Command("sh", "-c", "journalctl -u ssh -g ssh2 | tail -n 1")
		var output bytes.Buffer
		command.Stdout = &output
		err = command.Run()
		if err != nil {
			t.Error("journalctl command didn't work out")
		}

		if strings.Contains(output.String(), timeNow) {
			t.Errorf("The time for the journal entry is wrong")
		}
		if strings.Contains(output.String(), "Failed password for test from 192.108.3.2 port 22 ssh2") {
			t.Errorf("Journalctl wasn't logged for ip: %s, user: %s, port: %s\nOutput is: %s", ip, username, port, output.String())
		}

	})
}
