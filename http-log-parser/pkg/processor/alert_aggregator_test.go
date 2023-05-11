package processor_test

import (
	"testing"

	"companyx-http-log-parser/pkg/processor"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"
)

func TestNewAlertAggregator(t *testing.T) {
	logger, _ := test.NewNullLogger()

	agg := processor.NewAlertAggregator(logger, 1, 3)

	require.NotNil(t, agg)
}

func TestAlertAggregator_NoAlerts(t *testing.T) {
	logger, hook := test.NewNullLogger()

	agg := processor.NewAlertAggregator(logger, 1, 3)

	require.NotNil(t, agg)

	file := "../../test/data/test_normal_logs.csv"
	lines := ProcessCSVFile(t, file, agg)

	require.Len(t, lines, 10)
	require.Len(t, hook.Entries, 0)
}

func TestAlertAggregator_WithAlertsAndRecover(t *testing.T) {
	logger, hook := test.NewNullLogger()
	agg := processor.NewAlertAggregator(logger, 1, 3)
	require.NotNil(t, agg)

	file := "../../test/data/test_alert_and_recover_logs.csv"
	lines := ProcessCSVFile(t, file, agg)

	require.Len(t, lines, 25)
	require.Len(t, hook.Entries, 2)
	require.Equal(t, "High traffic generated an alert - hits = 6, triggered at 2019-02-07T22:11:03+01:00\n", hook.Entries[0].Message)
	require.Equal(t, "Alert recovered at 2019-02-07T22:11:10+01:00\n", hook.Entries[1].Message)
}

func TestAlertAggregator_WithAlerts(t *testing.T) {
	logger, hook := test.NewNullLogger()
	agg := processor.NewAlertAggregator(logger, 1, 3)
	require.NotNil(t, agg)

	file := "../../test/data/test_alert_logs.csv"
	lines := ProcessCSVFile(t, file, agg)

	require.Len(t, lines, 20)
	require.Len(t, hook.Entries, 1)
	require.Equal(t, "High traffic generated an alert - hits = 6, triggered at 2019-02-07T22:11:03+01:00\n", hook.Entries[0].Message)
}
