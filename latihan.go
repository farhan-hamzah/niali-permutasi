package main
import "fmt"
func faktorial(n int) int {
    hasil := 1
    for i := 1; i <= n; i++ {
        hasil *= i
    }
    return hasil
}

func main() {
    var x, y int
    fmt.Print("Masukkan x dan y: ")
    fmt.Scan(&x, &y)

    xFact := faktorial(x)  
    yFact := faktorial(y)  

    var hasilKetiga int
    if x >= y {
        hasilKetiga = xFact / faktorial(x-y) 
    } else {
        hasilKetiga = xFact
    }

    fmt.Println(xFact, yFact, hasilKetiga)
}