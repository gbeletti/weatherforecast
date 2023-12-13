package mocks

import (
	"context"
	"testing"

	"github.com/gbeletti/weatherforecast/contracts"
	testutil "github.com/gbeletti/weatherforecast/test/util"
)

type MockForecastClient struct {
	t        *testing.T
	filename string
}

func NewMockForecastClient(t *testing.T, filename ...string) *MockForecastClient {
	if len(filename) == 0 {
		filename = []string{""}
	}
	return &MockForecastClient{
		t:        t,
		filename: filename[0],
	}
}

func (m *MockForecastClient) GetForecast(ctx context.Context) (contracts.ForecastModel, error) {
	result := testutil.ReadJSONFile[contracts.ForecastModel](m.t, m.filename)
	return result, nil
}
