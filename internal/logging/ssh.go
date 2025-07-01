package logging

import (
    "bufio"
    "fmt"
    
    "github.com/ssgreg/journald"
)


type SSHLogger struct {}

func NewSSHLogger() *SSHLogger {
    return &SSHLogger{}
}

func (sl *SSHLogger) LogFailedAttempt(ip, username, port string) error {}


func (sl *SSHLogger) LogMessage(message string, priority journald.Priority) error {
    var err error
    err = journald.Print(priority, message); if err != nil {
        return fmt.Errorf("[!] Failed to log message to journald: %w\n", err)
    }

    fmt.Printf("Logged message successfully: %s\n", message)
    return nil
}
