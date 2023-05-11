package monitor_test

import (
	"strings"
	"testing"

	"companyx-http-log-parser/pkg/monitor"
	"companyx-http-log-parser/pkg/processor"
	"companyx-http-log-parser/pkg/source"
	"companyx-http-log-parser/pkg/types"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"
)

type TestCollector struct {
	lines []types.LogLine
}

func (t *TestCollector) OnLine(line types.LogLine) {
	t.lines = append(t.lines, line)
}

func TestMonitor_Run(t *testing.T) {
	log, _ := test.NewNullLogger()

	r := strings.NewReader(`test,test,test,ts,last
"10.0.0.1","-","apache",1549574327,"GET /report HTTP/1.0"
"10.0.0.2","-","apache",1549574330,"GET /report HTTP/1.0"
`)
	input := source.NewCSVParser(log, r, 3)

	testProcessor := &TestCollector{}
	m := monitor.New(log, input, []processor.LineProcessor{
		testProcessor,
	})

	m.Run()

	require.Len(t, testProcessor.lines, 2)
	require.Equal(t, 1549574327, testProcessor.lines[0].TS)
	require.Equal(t, 1549574330, testProcessor.lines[1].TS)
	require.Equal(t, "10.0.0.1,-,apache,1549574327,GET /report HTTP/1.0", strings.Join(testProcessor.lines[0].Data, ","))
	require.Equal(t, "10.0.0.2,-,apache,1549574330,GET /report HTTP/1.0", strings.Join(testProcessor.lines[1].Data, ","))
}
