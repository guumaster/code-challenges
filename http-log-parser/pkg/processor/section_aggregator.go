package processor

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"companyx-http-log-parser/internal/logger"
	"companyx-http-log-parser/pkg/types"
)

// UnknownSection name to group requests when the section parser fails.
const UnknownSection = "/unknown"

// SectionAggregator processor to aggregate data per section and interval.
type SectionAggregator struct {
	logger              logger.Logger
	aggregationInterval int
	lastIdx             int
	dataMap             map[string]int
	SectionColumnIndex  int
	topSections         int
}

// sortable simple struct to sort sections by hits.
type sortable struct {
	Key   string
	Value int
}

// sortableList contains a list of sortable objects and implement sorting interface.
type sortableList []sortable

func (p sortableList) Len() int      { return len(p) }
func (p sortableList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p sortableList) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		return p[i].Key < p[j].Key
	}

	return p[i].Value > p[j].Value
}

// NewSectionAggregator returns a new  SectionAggregator instance.
func NewSectionAggregator(logger logger.Logger, interval, topSections, sectionColumnIndex int) *SectionAggregator {
	return &SectionAggregator{
		logger,
		interval,
		0,
		map[string]int{},
		sectionColumnIndex,
		topSections,
	}
}

// OnLine process each log line shows aggregated stats on every interval.
func (s *SectionAggregator) OnLine(l types.LogLine) {
	idx := l.TS
	if s.lastIdx == 0 {
		s.lastIdx = idx
	}

	if idx-s.lastIdx >= s.aggregationInterval {
		s.logger.Infof(getIntervalStats(s.dataMap, s.lastIdx, s.aggregationInterval, s.topSections))
		s.lastIdx = (idx / s.aggregationInterval) * s.aggregationInterval
		s.dataMap = map[string]int{}
	}

	section := parseSection(l.Data[s.SectionColumnIndex])

	if _, ok := s.dataMap[section]; !ok {
		s.dataMap[section] = 0
	}

	s.dataMap[section]++
}

// OnComplete shows the last interval if there is any data left.
func (s *SectionAggregator) OnComplete() {
	if len(s.dataMap) == 0 {
		return
	}

	s.logger.Infof(getIntervalStats(s.dataMap, s.lastIdx, s.aggregationInterval, s.topSections))
}

// getIntervalStats calculates the current interval stats.
func getIntervalStats(stats map[string]int, ts, interval, topSections int) string {
	l := make(sortableList, len(stats))

	i := 0

	for k, v := range stats {
		l[i] = sortable{Key: k, Value: v}
		i++
	}

	sort.Sort(l)

	var top []string

	totalRequests := 0

	for i, k := range l {
		if i < topSections {
			top = append(top, fmt.Sprintf("%s=%d", k.Key, k.Value))
		}

		totalRequests += k.Value
	}

	reqsPerSecond := strconv.Itoa(totalRequests / interval)

	if reqsPerSecond == "0" {
		reqsPerSecond = "<1"
	}

	unixTimeUTC := time.Unix(int64(ts), 0)

	return fmt.Sprintf(
		"[%s] avg %s req/s. top sections [%s] \n",
		unixTimeUTC.Format(time.RFC3339),
		reqsPerSecond,
		strings.Join(top, ", "),
	)
}

// parseSection returns the section name from apache format.
func parseSection(r string) string {
	s := strings.Split(r, " ")
	if len(s) < 1 {
		return UnknownSection
	}

	s = strings.Split(s[1], "/")
	if len(s) < 1 {
		return UnknownSection
	}

	return fmt.Sprintf("/%s", s[1])
}
