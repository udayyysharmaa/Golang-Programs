package main
import "fmt"

func factorial(n int) int {
	if (n==1){
		return 1
	}
	return n * factorial(n-1)
}


func main(){
	var n int
	fmt.Println("Enter tha number")
	fmt.Scanf("%d", &n)

	output :=  factorial(n)
	fmt.Println(n, "Factorial of:", output)
}