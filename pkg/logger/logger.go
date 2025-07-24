package logger

import (
	"github.com/rs/zerolog"
	"os"
)

type Logger struct {
	zerolog.Logger
	logFile *os.File
}

func New(consoleLevel, fileLevel string) (*Logger, error) {

}
