package internal

import (
	"flag"
	"strings"

	"companyx-http-log-parser/internal/logger"
)

// Config contains all the program parameters and flags.
type Config struct {
	LogLevel  logger.LogLevel
	InputFile string
}

// ParseConfigFromFlags gets input parameters into Config struct.
func ParseConfigFromFlags() Config {
	inputFile := flag.String("from", "", "The input log file to be parsed.")
	logLevel := flag.String("log-level", "", "The verbosity of the command's output.")

	flag.Parse()

	return Config{
		logger.LogLevel(strings.ToLower(*logLevel)),
		*inputFile,
	}
}
