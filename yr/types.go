package yr

import "time"

type Forecast struct {
	Properties Properties `json:"properties"`
	Type       string     `json:"type"`
	Geometry   Geometry   `json:"geometry"`
}

type Instant struct {
	Details Details `json:"details"`
}

type Next12Hours struct {
	Summary Summary `json:"summary"`
	Details Details `json:"details"`
}

type Next1Hours struct {
	Summary Summary          `json:"summary"`
	Details Next1HourDetails `json:"details"`
}

type Next6Hours struct {
	Summary Summary `json:"summary"`
	Details Details `json:"details"`
}

type Summary struct {
	SymbolCode string `json:"symbol_code"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float32 `json:"coordinates"`
}

type Properties struct {
	Meta       Meta         `json:"meta"`
	TimeSeries []Timeseries `json:"timeseries"`
}

type Meta struct {
	UpdatedAt time.Time `json:"updated_at"`
	Units     Units     `json:"units"`
}

type Units struct {
	AirTemperatureMax          string `json:"air_temperature_max"`
	AirTemperatureMin          string `json:"air_temperature_min"`
	CloudAreaFractionLow       string `json:"cloud_area_fraction_low"`
	CloudAreaFractionMedium    string `json:"cloud_area_fraction_medium"`
	FogAreaFraction            string `json:"fog_area_fraction"`
	PrecipitationAmountMax     string `json:"precipitation_amount_max"`
	ProbabilityOfPrecipitation string `json:"probability_of_precipitation"`
	WindSpeedPercentile10      string `json:"wind_speed_percentile_10"`
	PrecipitationAmount        string `json:"precipitation_amount"`
	AirTemperature             string `json:"air_temperature"`
	CloudAreaFraction          string `json:"cloud_area_fraction"`
	DewPointTemperature        string `json:"dew_point_temperature"`
	RelativeHumidity           string `json:"relative_humidity"`
	WindSpeed                  string `json:"wind_speed"`
	AirPressureAtSeaLevel      string `json:"air_pressure_at_sea_level"`
	AirTemperaturePercentile10 string `json:"air_temperature_percentile_10"`
	AirTemperaturePercentile90 string `json:"air_temperature_percentile_90"`
	CloudAreaFractionHigh      string `json:"cloud_area_fraction_high"`
	PrecipitationAmountMin     string `json:"precipitation_amount_min"`
	ProbabilityOfThunder       string `json:"probability_of_thunder"`
	UltravioletIndexClearSky   string `json:"ultraviolet_index_clear_sky"`
	WindFromDirection          string `json:"wind_from_direction"`
	WindSpeedOfGust            string `json:"wind_speed_of_gust"`
	WindSpeedPercentile90      string `json:"wind_speed_percentile_90"`
}

type Timeseries struct {
	Time time.Time `json:"time"`
	Data Data      `json:"data"`
}

type Data struct {
	Instant Instant `json:"instant"`
	// Next12hours Next12Hours `json:"next_12_hours"`
	Next1Hours Next1Hours `json:"next_1_hours"`
	// Next6hours  Next_6_hours  `json:"next_6_hours"`
}

type Details struct {
	AirPressureAtSeaLevel      float32 `json:"air_pressure_at_sea_level"`
	AirTemperature             float32 `json:"air_temperature"`
	AirTemperaturePercentile10 float32 `json:"air_temperature_percentile_10"`
	AirTemperaturePercentile90 float32 `json:"air_temperature_percentile_90"`
	CloudAreaFraction          float32 `json:"cloud_area_fraction"`
	CloudAreaFractionHigh      float32 `json:"cloud_area_fraction_high"`
	CloudAreaFractionLow       float32 `json:"cloud_area_fraction_low"`
	CloudAreaFractionMedium    float32 `json:"cloud_area_fraction_medium"`
	DewPointTemperature        float32 `json:"dew_point_temperature"`
	FogAreaFraction            float32 `json:"fog_area_fraction"`
	RelativeHumidity           float32 `json:"relative_humidity"`
	UltravioletIndexClearSky   float32 `json:"ultraviolet_index_clear_sky"`
	WindSpeed                  float32 `json:"wind_speed"`
	WindSpeedOfGust            float32 `json:"wind_speed_of_gust"`
	WindSpeedPercentile10      float32 `json:"wind_speed_percentile_10"`
	WindSpeedPercentile90      float32 `json:"wind_speed_percentile_90"`
}

type Next1HourDetails struct {
	PrecipitationAmount        float32 `json:"precipitation_amount"`
	PrecipitationAmountMax     float32 `json:"precipitation_amount_max"`
	PrecipitationAmountMin     float32 `json:"precipitation_amount_min"`
	ProbabilityOfPrecipitation float32 `json:"probability_of_precipitation"`
	ProbabilityOfThunder       float32 `json:"probability_of_thunder"`
}
