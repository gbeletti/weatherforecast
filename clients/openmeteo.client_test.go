package clients_test

import (
	"context"
	"testing"

	"github.com/gbeletti/weatherforecast/clients"
	"github.com/gbeletti/weatherforecast/config"
	testutil "github.com/gbeletti/weatherforecast/test/util"
)

func TestGetForecast(t *testing.T) {
	configApp := config.NewConfig()
	openCli := clients.NewOpenmeteoClient(configApp)
	forecast, err := openCli.GetForecast(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	testutil.CreateJSONFile(t, "forecast.json", forecast)
}
