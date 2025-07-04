package logging

import (
    "fmt"
    
    "github.com/ssgreg/journald"
)


type SSHLogger struct {}

func NewSSHLogger() *SSHLogger {
    return &SSHLogger{}
}

func (sl *SSHLogger) LogFailedAttempt(ip, username, port string) error {
    var err error
    
    var message string = fmt.Sprintf("Failed password for %s from %s port %s ssh2", username, ip, port)
    err = journald.Send(message, journald.PriorityInfo,
            map[string]interface{}{
                "SSHD_IP_ADDRESS": ip,
                "SSHD_USERNAME": username,
                "SSHD_EVENT": "FailedPassword",
                "SYSLOG_IDENTIFIER": "sshd",
		"_SYSTEMD_UNIT": "sshd.service",
		"_COMM": "sshd",
		"PRIORITY": 6,
            })
    if err != nil {
        return fmt.Errorf("failed to log to journal: %w", err)
    }

    return nil
}


func (sl *SSHLogger) LogMessage(message string, priority journald.Priority) error {
    var err error = journald.Print(priority, message)
    if err != nil {
        return fmt.Errorf("[!] Failed to log message to journald: %w", err)
    }

    fmt.Printf("Logged message successfully: %s\n", message)
    return nil
}
