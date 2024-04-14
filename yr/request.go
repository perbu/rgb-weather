package yr

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	yrUrl    = "https://api.met.no/weatherapi/locationforecast/2.0/complete"
	altitude = "480"
	lat      = "59.980171"
	lon      = "10.663870"
)

func GetForecast() (Forecast, error) {
	req, err := http.NewRequest(http.MethodGet, yrUrl, nil)
	if err != nil {
		return Forecast{}, fmt.Errorf("could not create request: %w", err)
	}

	q := req.URL.Query()
	q.Add("altitude", altitude)
	q.Add("lat", lat)
	q.Add("lon", lon)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("User-Agent", "perbu/yr")
	req.Header.Add("Accept", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Forecast{}, fmt.Errorf("could not send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Forecast{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	// parse response into Forecast struct:
	var f Forecast
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return Forecast{}, fmt.Errorf("could not read response: %w", err)
	}
	// os.WriteFile("response.json", bytes, 0644)

	err = json.Unmarshal(bytes, &f)
	if err != nil {
		return Forecast{}, fmt.Errorf("could not parse response: %w", err)
	}

	// Remove timeseries in the future
	for f.Properties.Timeseries[0].Time.After(time.Now().Add(time.Hour)) {
		// Chop off the first timeseries if it is in the future
		f.Properties.Timeseries = f.Properties.Timeseries[1:]
		fmt.Println("chopping off first timeseries")
	}
	return f, nil
}
