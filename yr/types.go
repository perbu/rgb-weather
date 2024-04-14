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

type Next_12_hours struct {
	Summary Summary `json:"summary"`
	Details Details `json:"details"`
}

type Next_1_hours struct {
	Summary Summary          `json:"summary"`
	Details Next1HourDetails `json:"details"`
}

type Next_6_hours struct {
	Summary Summary `json:"summary"`
	Details Details `json:"details"`
}

type Summary struct {
	Symbol_code string `json:"symbol_code"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Properties struct {
	Meta       Meta         `json:"meta"`
	Timeseries []Timeseries `json:"timeseries"`
}

type Meta struct {
	Updated_at time.Time `json:"updated_at"`
	Units      Units     `json:"units"`
}

type Units struct {
	Air_temperature_max           string `json:"air_temperature_max"`
	Air_temperature_min           string `json:"air_temperature_min"`
	Cloud_area_fraction_low       string `json:"cloud_area_fraction_low"`
	Cloud_area_fraction_medium    string `json:"cloud_area_fraction_medium"`
	Fog_area_fraction             string `json:"fog_area_fraction"`
	Precipitation_amount_max      string `json:"precipitation_amount_max"`
	Probability_of_precipitation  string `json:"probability_of_precipitation"`
	Wind_speed_percentile_10      string `json:"wind_speed_percentile_10"`
	Precipitation_amount          string `json:"precipitation_amount"`
	Air_temperature               string `json:"air_temperature"`
	Cloud_area_fraction           string `json:"cloud_area_fraction"`
	Dew_point_temperature         string `json:"dew_point_temperature"`
	Relative_humidity             string `json:"relative_humidity"`
	Wind_speed                    string `json:"wind_speed"`
	Air_pressure_at_sea_level     string `json:"air_pressure_at_sea_level"`
	Air_temperature_percentile_10 string `json:"air_temperature_percentile_10"`
	Air_temperature_percentile_90 string `json:"air_temperature_percentile_90"`
	Cloud_area_fraction_high      string `json:"cloud_area_fraction_high"`
	Precipitation_amount_min      string `json:"precipitation_amount_min"`
	Probability_of_thunder        string `json:"probability_of_thunder"`
	Ultraviolet_index_clear_sky   string `json:"ultraviolet_index_clear_sky"`
	Wind_from_direction           string `json:"wind_from_direction"`
	Wind_speed_of_gust            string `json:"wind_speed_of_gust"`
	Wind_speed_percentile_90      string `json:"wind_speed_percentile_90"`
}

type Timeseries struct {
	Time time.Time `json:"time"`
	Data Data      `json:"data"`
}

type Data struct {
	Instant Instant `json:"instant"`
	// Next_12_hours Next_12_hours `json:"next_12_hours"`
	Next_1_hours Next_1_hours `json:"next_1_hours"`
	// Next_6_hours  Next_6_hours  `json:"next_6_hours"`
}

type Details struct {
	AirPressureAtSeaLevel      float64 `json:"air_pressure_at_sea_level"`
	AirTemperature             float64 `json:"air_temperature"`
	AirTemperaturPercentile10  float64 `json:"air_temperature_percentile_10"`
	AirTemperaturePercentile90 float64 `json:"air_temperature_percentile_90"`
	CloudAreaFraction          float64 `json:"cloud_area_fraction"`
	CloudAreaFractionHigh      float64 `json:"cloud_area_fraction_high"`
	CloudAreaFractionLow       float64 `json:"cloud_area_fraction_low"`
	CloudAreaFractionMedium    float64 `json:"cloud_area_fraction_medium"`
	DewPointTemperature        float64 `json:"dew_point_temperature"`
	FogAreaFraction            float64 `json:"fog_area_fraction"`
	RelativeHumidity           float64 `json:"relative_humidity"`
	UltravioletIndexClearSky   float64 `json:"ultraviolet_index_clear_sky"`
	WindSpeed                  float64 `json:"wind_speed"`
	WindSpeedOfGust            float64 `json:"wind_speed_of_gust"`
	WindSpeedPercentile10      float64 `json:"wind_speed_percentile_10"`
	WindSpeedPercentile90      float64 `json:"wind_speed_percentile_90"`
}

type Next1HourDetails struct {
	PrecipitationAmount        float64 `json:"precipitation_amount"`
	PrecipitationAmountMax     float64 `json:"precipitation_amount_max"`
	PrecipitationAmountMin     float64 `json:"precipitation_amount_min"`
	ProbabilityOfPrecipitation float64 `json:"probability_of_precipitation"`
	ProbabilityOfThunder       float64 `json:"probability_of_thunder"`
}
