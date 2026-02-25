func countPrimes(n int) int {
    if n <= 2 {
        return 0
    }

    numbers := make([]bool, n)

    for p := 2; p <= int(math.Sqrt(float64(n))); p++ {
        if numbers[p] == false {
            for j := p * p; j < n; j += p {
                numbers[j] = true
            }
        }
    }

    numberOfPrimes := 0
    for i := 2; i < n; i++ {
        if numbers[i] == false {
            numberOfPrimes += 1
        }
    }
    return numberOfPrimes
}
