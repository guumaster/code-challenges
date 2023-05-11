package processor

import (
	"companyx-http-log-parser/internal/logger"
	"companyx-http-log-parser/pkg/types"
)

// NOTE: This processor is not on the requirements, but its purpose is to show the flexibility of the design.

// GlobalStats object to count.
type GlobalStats struct {
	counter int
	logger  logger.Logger
}

// NewGlobalStats returns a new instance of GlobalStats.
func NewGlobalStats(logger logger.Logger) *GlobalStats {
	return &GlobalStats{
		1,
		logger,
	}
}

// OnLine process a log line by counting it.
func (g *GlobalStats) OnLine(_ types.LogLine) {
	g.counter++
}

// ShowStats prints the total lines counted.
func (g *GlobalStats) ShowStats() {
	g.logger.Infof("TOTAL LINES: %d", g.counter)
}
