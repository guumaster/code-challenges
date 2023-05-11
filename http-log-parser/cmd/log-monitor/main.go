package main

import (
	"os"

	"companyx-http-log-parser/internal"
	"companyx-http-log-parser/internal/logger"
	"companyx-http-log-parser/pkg/monitor"
	"companyx-http-log-parser/pkg/processor"
	"companyx-http-log-parser/pkg/source"
)

func main() {
	// Reads config from flags
	cfg := internal.ParseConfigFromFlags()

	// Create a log level
	log := logger.NewLogger(cfg.LogLevel)

	// Create a reader from file or STDIN
	r, err := internal.GetInputStream(os.Stdin, cfg.InputFile)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// Create the input parser
	input := source.NewCSVParser(log, r, TSColumnIndex)

	// Prepare all line processors
	fileStats := processor.NewGlobalStats(log)
	pList := []processor.LineProcessor{
		processor.NewAlertAggregator(log, MaxRequestsPerSecond, ThresholdInterval),
		processor.NewSectionAggregator(log, AggregationInterval, TopSections, SectionColumnIndex),
		fileStats,
	}

	// Add line printer for debugging only
	if cfg.LogLevel == logger.Debug {
		pList = append(pList, processor.NewLinePrinter(log))
	}

	// Create and run the log monitor
	m := monitor.New(log, input, pList)
	m.Run()

	// Show some stats before exit
	fileStats.ShowStats()
	log.Infof("done")
}
