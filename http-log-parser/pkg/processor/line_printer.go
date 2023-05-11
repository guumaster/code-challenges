package processor

import (
	"strings"

	"companyx-http-log-parser/internal/logger"
	"companyx-http-log-parser/pkg/types"
)

// LinePrinter is a simple debugging processor to print all lines.
type LinePrinter struct {
	counter int
	logger  logger.Logger
}

// NewLinePrinter returns a new instance of LinePrinter.
func NewLinePrinter(logger logger.Logger) *LinePrinter {
	return &LinePrinter{
		1,
		logger,
	}
}

// OnLine process each log line by printing it as a debug message.
func (l *LinePrinter) OnLine(line types.LogLine) {
	l.logger.Debugf("LINE [%d]: [%s]", l.counter, strings.Join(line.Data, ","))
	l.counter++
}
