package source_test

import (
	"os"
	"strings"
	"testing"

	"companyx-http-log-parser/pkg/source"
	"companyx-http-log-parser/pkg/types"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"
)

func TestNewCSVParser(t *testing.T) {
	logger, _ := test.NewNullLogger()
	r := strings.NewReader("testing")
	p := source.NewCSVParser(logger, r, 3)
	require.NotNil(t, p)
}

func TestCSVParser_Run(t *testing.T) {
	logger, hook := test.NewNullLogger()
	logger.SetLevel(logrus.DebugLevel)

	filename := "../../test/data/test_alert_logs.csv"
	r, err := os.Open(filename)
	require.NoError(t, err)

	p := source.NewCSVParser(logger, r, 3)
	require.NotNil(t, p)

	linesCh := make(chan types.LogLine, 100)
	doneCh := make(chan struct{})

	go p.Run(linesCh, doneCh)

	totalLines := 0

	<-doneCh
	close(linesCh)

	for l := range linesCh {
		require.IsType(t, types.LogLine{}, l)
		totalLines++
	}

	require.Equal(t, 20, totalLines)
	require.Len(t, hook.Entries, 1)
	require.Equal(t, "all lines read", hook.Entries[0].Message)
}
