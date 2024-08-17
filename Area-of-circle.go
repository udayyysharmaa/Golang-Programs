package main
import "fmt"

const PI float64 = 3.14
func main(){
	var radius float64
	var area float64
	fmt.Println("Enter the number")
	fmt.Scanf("%f", &radius)
	area = PI*radius*radius
	fmt.Println("Area of Circle:", area)

}