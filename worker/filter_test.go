package worker_test

import (
	"context"
	"flag"
	"testing"

	"github.com/gbeletti/weatherforecast/api/model"
	"github.com/gbeletti/weatherforecast/test/mocks"
	testutil "github.com/gbeletti/weatherforecast/test/util"
	"github.com/gbeletti/weatherforecast/worker"
	"github.com/google/go-cmp/cmp"
)

var (
	update = flag.Bool("update", false, "update golden files")
)

func TestFilterForecast(t *testing.T) {
	tcases := []struct {
		name       string
		inputFile  string
		resultFile string
	}{
		{
			name:       "01 - filter forecast",
			inputFile:  "../clients/testdata/openmeteo/01/forecast.json",
			resultFile: "./testdata/filter/01/result.json",
		},
	}

	for _, tcase := range tcases {
		t.Run(tcase.name, func(t *testing.T) {
			testFilterForecast(t, tcase.name, tcase.inputFile, tcase.resultFile)
		})
	}
}

func testFilterForecast(t *testing.T, name, inputFile, resultFile string) {
	mockClient := mocks.NewMockForecastClient(t, inputFile)
	gotForecast, err := worker.FilterForecast(context.Background(), mockClient)
	if err != nil {
		t.Error(err)
		return
	}
	if *update {
		testutil.CreateJSONFile(t, resultFile, gotForecast)
		return
	}
	expectedForecast := testutil.ReadJSONFile[[]model.ForecastModel](t, resultFile)
	if diff := cmp.Diff(expectedForecast, gotForecast); diff != "" {
		t.Errorf("test [%s] mismatch (-want +got):\n%s", name, diff)
	}
}
