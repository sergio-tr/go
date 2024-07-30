package birdwatcher

const DaysPerWeek = 7

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {

	totalBirdCount := 0
	for i := 0; i < len(birdsPerDay); i++ {
		totalBirdCount += birdsPerDay[i]
	}

	/*
		Another way:

		for _, elem := range birdsPerDay {
				totalBirdCount += elem
			}
	*/

	return totalBirdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {

	firstIndex := (week - 1) * DaysPerWeek

	if firstIndex < 0 {
		firstIndex = 0
	}

	lastIndex := firstIndex + DaysPerWeek

	/*
		Not necessary due to exercise wording
		- We should check whether the indices are out of range given the slice length
		- If the first index is in range, we should assign the last index to the last position inside the slice
	*/

	return TotalBirdCount(birdsPerDay[firstIndex:lastIndex])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i] = birdsPerDay[i] + 1
		}
	}

	return birdsPerDay
}
