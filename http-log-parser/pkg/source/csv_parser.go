package source

import (
	"encoding/csv"
	"io"
	"strconv"

	"companyx-http-log-parser/internal/logger"
	"companyx-http-log-parser/pkg/types"
)

// CSVParser is a input parser to process each line as a LogLine.
type CSVParser struct {
	logger  logger.Logger
	reader  io.Reader
	tsIndex int
}

// NewCSVParser returns a new instance of CSVParser.
func NewCSVParser(logger logger.Logger, r io.Reader, tsIndex int) *CSVParser {
	return &CSVParser{
		logger,
		r,
		tsIndex,
	}
}

// Run setups the Reader, iterates over each line and send them to the line channel.
func (p CSVParser) Run(linesCh chan types.LogLine, doneCh chan struct{}) {
	r := csv.NewReader(p.reader)

	// Skip header
	_, err := r.Read()
	if err != nil {
		p.logger.Fatalf("error reading header: %s", err)
		doneCh <- struct{}{}

		return
	}

	lineNum := 1

	for {
		rec, err := r.Read()
		if err != nil {
			if err == io.EOF {
				p.logger.Debugf("all lines read")

				doneCh <- struct{}{}

				return
			}

			p.logger.Errorf("error reading line: %s", err)

			break
		}
		lineNum++

		ts, err := strconv.ParseInt(rec[p.tsIndex], 10, 64)
		if err != nil {
			p.logger.Errorf("error reading line: bad timestamp on line %d: %v\n", lineNum, err)
			break
		}

		linesCh <- types.LogLine{
			TS:   int(ts),
			Data: rec,
		}
	}
}
