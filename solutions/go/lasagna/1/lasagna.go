package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime int = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualRemainingTime int) int {
    return OvenTime - actualRemainingTime
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
    return 2 * numberOfLayers;
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers int, actualMinutesInOven int) int{
    return (2 * numberOfLayers) + actualMinutesInOven
}

func main(){
    RemainingOvenTime(30)
    PreparationTime(2)
    ElapsedTime(3, 20)
}