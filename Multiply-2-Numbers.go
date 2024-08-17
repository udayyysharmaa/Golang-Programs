package main
import "fmt"

func main(){

	var a,b int
	fmt.Println("Enter the FirstNumber!")
	fmt.Scanf("%d", &a)
	fmt.Println("Enter the SecondNumber!")
	fmt.Scanf("%d", &b)

	var c int
	c = a * b

	fmt.Println("Result", c)
}