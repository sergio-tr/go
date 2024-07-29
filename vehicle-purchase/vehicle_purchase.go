package purchase

import (
	"fmt"
	"strings"
)

const Car string = "car"
const Truck string = "truck"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	lowerCase := strings.ToLower(kind)
	return lowerCase == Car || lowerCase == Truck
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {

	chosenVehicle := option2

	if option1 < option2 {
		chosenVehicle = option1
	}

	return fmt.Sprintf("%s is clearly the better choice.", chosenVehicle)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {

	switch {
	case age < 3:
		return originalPrice * 0.8
	case age >= 3 && age < 10:
		return originalPrice * 0.7
	case age >= 10:
		return originalPrice * 0.5
	default:
		panic("The car's age must be greater than or equal to zero")
	}

}
