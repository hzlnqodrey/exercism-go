// Package weather provides weather condition in current location.
package weather

// CurrentCondition represents a certain weather condition in string type.
var CurrentCondition string
// CurrentLocation represents a certain location in string type.
var CurrentLocation string

// Forecast function will need city and condition variable as a parameter to return the concatenated string that will tell certain location has certain weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
