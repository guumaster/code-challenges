package processor_test

import (
	"testing"

	"companyx-http-log-parser/pkg/processor"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"
)

func TestNewSectionAggregator(t *testing.T) {
	logger, _ := test.NewNullLogger()

	agg := processor.NewSectionAggregator(logger, 1, 3, 4)

	require.NotNil(t, agg)
}

func TestSectionAggregator_Run(t *testing.T) {
	logger, hook := test.NewNullLogger()

	agg := processor.NewSectionAggregator(logger, 1, 2, 4)

	require.NotNil(t, agg)

	file := "../../test/data/test_section_logs.csv"
	lines := ProcessCSVFile(t, file, agg)

	require.Len(t, lines, 18)
	require.Len(t, hook.Entries, 4)
	require.Equal(t, "[2019-02-07T22:11:03+01:00] avg 4 req/s. top sections [/api=2, /user=2] \n", hook.Entries[0].Message)
	require.Equal(t, "[2019-02-07T22:11:14+01:00] avg 5 req/s. top sections [/help=5] \n", hook.Entries[1].Message)
	require.Equal(t, "[2019-02-07T22:11:15+01:00] avg 5 req/s. top sections [/report=3, /api=1] \n", hook.Entries[2].Message)
	require.Equal(t, "[2019-02-07T22:11:16+01:00] avg 4 req/s. top sections [/api=2, /help=2] \n", hook.Entries[3].Message)
}

func TestSectionAggregator_RunSlow(t *testing.T) {
	logger, hook := test.NewNullLogger()

	agg := processor.NewSectionAggregator(logger, 10, 2, 4)

	require.NotNil(t, agg)

	file := "../../test/data/test_section_slow_logs.csv"
	lines := ProcessCSVFile(t, file, agg)

	require.Len(t, lines, 8)
	require.Len(t, hook.Entries, 2)
	require.Equal(t, "[2019-02-07T22:11:10+01:00] avg <1 req/s. top sections [/help=3, /report=1] \n", hook.Entries[0].Message)
	require.Equal(t, "[2019-02-07T22:11:30+01:00] avg <1 req/s. top sections [/api=2, /report=1] \n", hook.Entries[1].Message)
}
