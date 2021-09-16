package logger

import (
	"os"
	"time"
)

func Logger(message string) {
	f, _ := os.OpenFile("logfile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f.WriteString("\n" + time.Now().Format("2006-01-02 15:04:05.000") + "\t|\t" + message)
	f.Close()
}
