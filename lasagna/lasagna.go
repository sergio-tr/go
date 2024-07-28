package lasagna

const OvenTime int = 40
const MinutesPerLayer int = 2

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	if actualMinutesInOven > OvenTime {
		panic("The lasagna has already been overcooked")
	}

	return OvenTime - actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	if numberOfLayers <= 0 {
		panic("The lasagna must have at least one layer")
	}

	return numberOfLayers * MinutesPerLayer
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	if actualMinutesInOven < 0 {
		panic("The actual minimum minutes that the lasagna has been inside the oven must be 0")
	}

	return PreparationTime(numberOfLayers) + actualMinutesInOven
}
