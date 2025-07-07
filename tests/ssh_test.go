package tests

import (
	"testing"
	"os/exec"
	"time"
	"bytes"
	"strings"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/jean0t/testing-fail2ban/internal/logging"
)

func GetLogLastMessage() (string, error) {
	var err error	
	var output bytes.Buffer
	var command *exec.Cmd = exec.Command("sh", "-c", "journalctl -g ssh2 | tail -n 1")

	command.Stdout = &output
	err = command.Run()
	if err != nil {
		return "", fmt.Errorf("journalctl command didn't work out")
	}

	return output.String(), nil
}

func TestSSHFailedToAuthenticate(t *testing.T) {
	var err error
	var timeNow string
	t.Run("Testing Logging for Failed Password", func(t *testing.T) {
		var sshLog *logging.SSHLogger = logging.NewSSHLogger()
		var logMessage string = ""
		var ip_test, username_test, port_test string = fmt.Sprintf("192.108.3.%s", strconv.Itoa(rand.Intn(255))), "test", strconv.Itoa(rand.Intn(25)) 
		timeNow = time.Now().Format("15:04:05")
		err = sshLog.LogFailedAttempt(ip_test, username_test, port_test)
		if err != nil {
			t.Error("Error when writing to journal")
		}
		logMessage, err = GetLogLastMessage()
		if err != nil {
			t.Error("Couldn't get the journalctl entry")
		}

		if !strings.Contains(logMessage, timeNow) {
			t.Errorf("The time for the journal entry is wrong")
		}

		if !strings.Contains(logMessage, fmt.Sprintf("Failed password for test from %s port %s ssh2", ip_test, port_test)) {
			t.Errorf("Journalctl wasn't logged for ip: %s, user: %s, port: %s\nOutput is: %s", ip_test, username_test, port_test, logMessage)
		}

	})
}
