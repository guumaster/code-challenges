package main

// NOTE: These parameters are not exposed for simplicity.
const (
	// TSColumnIndex constant indicating the position of the timestamp in each line.
	TSColumnIndex = 3
)

// Config for section aggregation processor.
const (
	// SectionColumnIndex constant indicating the position of section data in each line.
	SectionColumnIndex = 4
	// AggregationInterval constant indicating the position of section data in each line.
	AggregationInterval = 10
	// TopSections indicates how many sections to show on each line.
	TopSections = 2
)

// Config for alerting aggregation processor.
const (
	// ThresholdInterval indicates the seconds to check when triggering an alert.
	ThresholdInterval = 120
	// MaxRequestsPerSecond indicates the average req/s that would trigger an alert.
	MaxRequestsPerSecond = 10
)
