// Package weather provides functionality to forecast the current weather condition
// of various cities in Goblinocus. It includes variables to store the current
// weather condition and location, as well as a function to generate the weather forecast.
package weather

// CurrentCondition stores the current weather condition of the specified location.
var CurrentCondition string

// CurrentLocation stores the name of the location for which the weather condition is being forecasted.
var CurrentLocation string

// Forecast returns a string that describes the current weather condition in the specified city.
// It sets the global variables CurrentLocation and CurrentCondition based on the input parameters.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
