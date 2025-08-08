// Package weather provides tools to track and report the current weather condition for a given location.
package weather

// CurrentCondition stores the latest reported weather condition, such as "sunny", "rainy", or "cloudy".
var CurrentCondition string

// CurrentLocation stores the name of the location for which the current weather condition is reported.
var CurrentLocation string

// Forecast sets the current location and weather condition, and returns a formatted forecast string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
