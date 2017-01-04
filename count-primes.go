package main

import ("fmt"
    "math"
)

func isPrime(n int) (bool) {
    limit := int(math.Sqrt(float64(n)))
    for d := 2; d <= limit; d++ {
        if n % d == 0 {
            return false
        }
    }
    return true
}

func main() {
    for i:=2; i < 20; i++ {
        fmt.Println(i, isPrime(i))
    }
}