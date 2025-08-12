package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTimePerLayer int) int{
    if avgPrepTimePerLayer == 0 {
		avgPrepTimePerLayer = 2
	}
	return len(layers) * avgPrepTimePerLayer
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0

	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	if len(friendsList) == 0 || len(myList) == 0 {
		return
	}
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
    scaled := make([] float64, len(quantities))
    factor := float64(portions)/2

    for i, amount := range quantities{
        scaled[i] = amount * factor
    }

    return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
