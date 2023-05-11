package monitor

import (
	"sync"

	"companyx-http-log-parser/internal/logger"
	"companyx-http-log-parser/pkg/processor"
	"companyx-http-log-parser/pkg/source"
	"companyx-http-log-parser/pkg/types"
)

// Monitor object to manage log parsing and processing.
type Monitor struct {
	wg         *sync.WaitGroup
	linesCh    chan types.LogLine
	doneCh     chan struct{}
	logger     logger.Logger
	input      source.Source
	processors []processor.LineProcessor
}

// New returns a new instance of the Monitor to manage log parsing and processing.
func New(logger logger.Logger, input source.Source, processors []processor.LineProcessor) Monitor {
	linesCh := make(chan types.LogLine, 1)
	doneCh := make(chan struct{})

	return Monitor{
		&sync.WaitGroup{},
		linesCh,
		doneCh,
		logger,
		input,
		processors,
	}
}

// Run setup all processors and then starts the input reading process.
func (m Monitor) Run() {
	// prepare line channels per processor.
	chs := makeChannelsForProcessors(len(m.processors))
	go multiplexChannels(m.linesCh, chs)

	for i, pr := range m.processors {
		go m.startProcessor(pr, chs[i])
	}

	go m.startInputProcess()

	// Wait until the reading is done, and close all after all messages are processed.
	<-m.doneCh
	close(m.linesCh)
	m.wg.Wait()
}

// startInputProcess runs the input stream process.
func (m Monitor) startInputProcess() {
	m.wg.Add(1)
	defer m.wg.Done()
	m.input.Run(m.linesCh, m.doneCh)
}

// startProcessor by ranging over its own line channel.
func (m Monitor) startProcessor(pr processor.LineProcessor, linesCh chan types.LogLine) {
	m.wg.Add(1)
	defer m.wg.Done()

	for l := range linesCh {
		pr.OnLine(l)
	}

	// Launch OnComplete only if the processor implements that interface.
	if pr, ok := pr.(processor.LineProcessorComplete); ok {
		pr.OnComplete()
	}
}

// makeChannelsForProcessors creates one channel for each registered processor.
func makeChannelsForProcessors(total int) []chan types.LogLine {
	chs := make([]chan types.LogLine, total)
	for i := 0; i < total; i++ {
		chs[i] = make(chan types.LogLine, 1)
	}

	return chs
}

// multiplexChannels repeat lines for each registered processor.
func multiplexChannels(src chan types.LogLine, dst []chan types.LogLine) {
	for v := range src {
		for _, cons := range dst {
			cons <- v
		}
	}
	// close all after src channel is closed
	for _, cons := range dst {
		close(cons)
	}
}
