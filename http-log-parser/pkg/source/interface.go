package source

import (
	"companyx-http-log-parser/pkg/types"
)

// Source interface to be implemented by all log line reader.
type Source interface {
	Run(chan types.LogLine, chan struct{})
}
