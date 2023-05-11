package processor

import (
	"time"

	"companyx-http-log-parser/internal/logger"
	"companyx-http-log-parser/pkg/types"
)

type dataMap map[int]int

// AlertAggregator object to handle alerting based on aggregated data.
type AlertAggregator struct {
	logger               logger.Logger
	maxRequestsPerSecond int
	thresholdInterval    int
	alertActive          bool
	dataMap              dataMap
}

// NewAlertAggregator creates a new instance of AlertAggregator.
func NewAlertAggregator(logger logger.Logger, maxRequestsPerSecond, thresholdInterval int) *AlertAggregator {
	return &AlertAggregator{
		logger,
		maxRequestsPerSecond,
		thresholdInterval,
		false,
		dataMap{},
	}
}

// OnLine process one log line. Triggers an alert when req/s is over threshold.
func (a *AlertAggregator) OnLine(l types.LogLine) {
	if _, ok := a.dataMap[l.TS]; !ok {
		a.dataMap[l.TS] = 0
	}

	a.dataMap[l.TS]++

	filtered, totalHits := filterData(a.dataMap, l.TS, a.thresholdInterval)
	isOverThreshold := (totalHits / a.thresholdInterval) > a.maxRequestsPerSecond

	a.dataMap = filtered

	if isOverThreshold {
		if !a.alertActive {
			unixTimeUTC := time.Unix(int64(l.TS), 0)
			a.logger.Infof("High traffic generated an alert - hits = %d, triggered at %s\n", totalHits, unixTimeUTC.Format(time.RFC3339))
		}

		a.alertActive = true
	} else if !isOverThreshold && a.alertActive {
		unixTimeUTC := time.Unix(int64(l.TS), 0)
		a.logger.Infof("Alert recovered at %s\n", unixTimeUTC.Format(time.RFC3339))

		a.alertActive = false
	}
}

// filterData returns only the data in the interval and the total requests.
func filterData(lines dataMap, currentTs, interval int) (dataMap, int) {
	lastToKeep := currentTs - interval
	newList := dataMap{}
	totalRequests := 0

	for k, v := range lines {
		if k < lastToKeep {
			continue
		}

		newList[k] = v
		totalRequests += v
	}

	return newList, totalRequests
}
