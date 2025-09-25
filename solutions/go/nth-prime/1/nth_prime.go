package prime
import(
    "errors"
    "math"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n must be greater than 0")
	}

    count := 0
    num := 1

    for {
        num++
        if isPrime(num){
            count++
            if count == n {
            	return num, nil
            }
        }
    }
}

func isPrime(n int) bool{
    if n < 2 {
        return false
    }

    sqrtN := int(math.Sqrt(float64(n)))

    for i:=2; i <= sqrtN; i++ {
        if n % i == 0 {
            return false
        }
    }

    return true
}
