package yr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	yrUrl    = "https://api.met.no/weatherapi/locationforecast/2.0/complete"
	altitude = "480"
	lat      = "59.980171"
	lon      = "10.663870"
)

func GetForecast(max int) (Forecast, error) {
	var f Forecast

	// Construct the URL with query parameters
	url := fmt.Sprintf("%s?altitude=%s&lat=%s&lon=%s", yrUrl, altitude, lat, lon)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Forecast{}, fmt.Errorf("could not create request: %w", err)
	}
	req.Header.Add("User-Agent", "perbu/yr")
	req.Header.Add("Accept", "application/json")
	// Perform the HTTP GET request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Forecast{}, fmt.Errorf("could not send request: %w", err)
	}
	defer resp.Body.Close()

	// Check HTTP response status
	if resp.StatusCode != http.StatusOK {
		return Forecast{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read and parse the JSON response
	err = json.NewDecoder(resp.Body).Decode(&f)
	if err != nil {
		return Forecast{}, fmt.Errorf("could not parse response: %w", err)
	}

	// Truncate timeseries data as required
	now := time.Now()
	i := 0
	for ; i < len(f.Properties.TimeSeries) && !f.Properties.TimeSeries[i].Time.Before(now); i++ {
	}
	f.Properties.TimeSeries = f.Properties.TimeSeries[i:]
	if len(f.Properties.TimeSeries) > max {
		f.Properties.TimeSeries = f.Properties.TimeSeries[:max]
	}

	return f, nil
}
