package analyzer_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/GaijinEntertainment/go-exhaustruct/v3/analyzer"
)

var testdataPath, _ = filepath.Abs("./testdata/") //nolint:gochecknoglobals

func TestAnalyzer(t *testing.T) {
	t.Parallel()

	a, err := analyzer.NewAnalyzer([]string{""}, nil)
	assert.Nil(t, a)
	assert.Error(t, err)

	a, err = analyzer.NewAnalyzer([]string{"["}, nil)
	assert.Nil(t, a)
	assert.Error(t, err)

	a, err = analyzer.NewAnalyzer(nil, []string{""})
	assert.Nil(t, a)
	assert.Error(t, err)

	a, err = analyzer.NewAnalyzer(nil, []string{"["})
	assert.Nil(t, a)
	assert.Error(t, err)

	a, err = analyzer.NewAnalyzer(
		[]string{`.*[Tt]est.*`, `.*External`, `.*Embedded`, `.*\.<anonymous>`},
		[]string{`.*Excluded$`, `e\.<anonymous>`},
	)
	require.NoError(t, err)

	analysistest.Run(t, testdataPath, a, "i", "e")

	a, err = analyzer.NewWithOptions(
		analyzer.Include(`.*[T]est.*`),
		analyzer.OnlyExported,
	)
	require.NoError(t, err)

	analysistest.Run(t, testdataPath, a, "exported")

	a, err = analyzer.NewWithOptions(
		analyzer.Include(`.*[T]est.*`),
		analyzer.SkipEmpty,
	)
	require.NoError(t, err)

	analysistest.Run(t, testdataPath, a, "skipempty")
}
