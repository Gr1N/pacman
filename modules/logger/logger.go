package logger

import (
	"bytes"
	"log"
	"os"
	"time"

	"github.com/hashicorp/logutils"

	"github.com/Gr1N/pacman/modules/settings"
)

func Init() {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"debug", "warn", "error"},
		MinLevel: logutils.LogLevel(settings.S.Logger.MinLevel),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)
}

func Debug(message string) {
	l("debug", message)
}

func Warn(message string) {
	l("warn", message)
}

func Error(message string) {
	l("error", message)
}

func l(level, message string) {
	var buf bytes.Buffer

	buf.WriteRune('[')
	buf.WriteString(level)
	buf.WriteString("] ")
	buf.WriteString(time.Now().Format("2006/01/02 - 15:04:05"))
	buf.WriteString(" | ")
	buf.WriteString(message)

	log.Print(buf.String())
}
