package processor

import (
	"companyx-http-log-parser/pkg/types"
)

// LineProcessor is the base interface to implement a new processor that reads logs lines.
type LineProcessor interface {
	OnLine(line types.LogLine)
}

// LineProcessorComplete extends LineProcessor interface with OnComplete that triggers after all log lines are read.
type LineProcessorComplete interface {
	LineProcessor
	OnComplete()
}
