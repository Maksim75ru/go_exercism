package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTime int) int {
    var t int 
    if avgTime == 0 {
        t = 2
    } else {
        t = avgTime
    }

    return len(layers) * t 
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    var noodles int
    var sauce float64
    
    for _, layer := range(layers) {
        if layer == "noodles" {
            noodles += 50
        } else if layer == "sauce" {
            sauce += 0.2
        }
    }
    
    return noodles, sauce
}
                     
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendLayers, myLayers []string) {
    myLayers[len(myLayers) - 1] = friendLayers[len(friendLayers) - 1]
}


// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portionNum int) []float64  {
    result := make([]float64, 0, len(quantities))
    for _, q := range quantities {
        current := q * (float64(portionNum) / 2.0)
        result = append(result, current)
    }
    return result
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
