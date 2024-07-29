package cars

const IndividualProductionCost uint = 10_000
const GroupProductionCost uint = 95_000
const GroupSize uint = 10

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	if productionRate < 0 {
		panic("production rate must be greater than or equal to zero")
	}

	if successRate < 0 || successRate > 100 {
		panic("success rate must be between 0 and 100")
	}

	return float64(productionRate) * successRate / 100

}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {

	if carsCount < 0 {
		panic("cars count must be greater than or equal to zero")
	}

	var numberOfGroups = uint(carsCount) / GroupSize
	var remainingCars = uint(carsCount) % GroupSize

	return numberOfGroups*GroupProductionCost + remainingCars*IndividualProductionCost
}
