package processor_test

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"testing"

	"companyx-http-log-parser/pkg/processor"
	"companyx-http-log-parser/pkg/types"

	"github.com/stretchr/testify/require"
)

func ProcessCSVFile(t *testing.T, filename string, agg processor.LineProcessor) [][]string {
	t.Helper()

	file, err := os.Open(filename)
	require.NoError(t, err)

	r := csv.NewReader(file)

	// Skip header
	_, err = r.Read()
	require.NoError(t, err)

	var lines [][]string

	for {
		line, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			t.Fatal(err)
		}

		require.NotEmpty(t, line)
		lines = append(lines, line)
		ts, err := strconv.ParseInt(line[3], 10, 64)

		agg.OnLine(types.LogLine{
			TS:   int(ts),
			Data: line,
		})
	}

	if pr, ok := agg.(processor.LineProcessorComplete); ok {
		pr.OnComplete()
	}
	return lines
}
