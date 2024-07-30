package lasagna

import "strings"

const DefaultAveragePreparationTime = 2
const Noodles = "noodles"
const NoodlesGrams = 50
const Sauce = "sauce"
const SauceLiters = 0.2

const QuestionMark = "?"

// PreparationTime returns the expected preparation time for the given lasagna layers
func PreparationTime(layers []string, averagePreparationTime int) int {
	if averagePreparationTime < 0 {
		panic("Preparation time must be a positive number")
	}

	if averagePreparationTime == 0 {
		averagePreparationTime = DefaultAveragePreparationTime
	}

	return len(layers) * averagePreparationTime
}

// Quantities For each noodle layer in your lasagna, you will need 50 grams of noodles.
// For each sauce layer in your lasagna, you will need 0.2 liters of sauce.
func Quantities(layers []string) (noodlesGrams int, sauceLiters float64) {
	for i := 0; i < len(layers); i++ {
		elem := strings.ToLower(layers[i])

		switch elem {
		case Noodles:
			noodlesGrams += NoodlesGrams
		case Sauce:
			sauceLiters += SauceLiters
		}
	}

	return
}

// AddSecretIngredient  The first parameter is the list your friend sent you,
// the second is the ingredient list of your own recipe.
// The last element in your ingredient list is always "?".
// The function should replace it with the last item from your friends list.
func AddSecretIngredient(friendsList []string, myList []string) {

	if len(myList) == 0 || myList[len(myList)-1] != QuestionMark {
		panic("Your list need to specify a secret ingredient in the last element")
	}

	if len(friendsList) == 0 {
		panic("Your friend need to specify at least one ingredient")
	}

	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// ScaleRecipe returns a slice of the amounts needed for the desired number of portions.
func ScaleRecipe(amounts []float64, numberOfPortions int) []float64 {
	newAmounts := make([]float64, len(amounts))
	groupsOfTwoPortions := numberOfPortions / 2
	isHalfMoreNeededToBeAdded := numberOfPortions%2 != 0

	for i, elem := range amounts {
		newAmounts[i] = elem * float64(groupsOfTwoPortions)

		if isHalfMoreNeededToBeAdded {
			newAmounts[i] += elem / 2
		}
	}

	return newAmounts

}
