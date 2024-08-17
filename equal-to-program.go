package main
import "fmt"

func main(){
	var city1, city2 string 
	fmt.Println("Enter the First city name!")
	fmt.Scanf("%s", &city1)
	fmt.Println("Enter the Second City name!")
	fmt.Scanf("%s", &city2)
	var result strings
	result = (city1 == city2)
	fmt.Println("City1 && City2 is equal:", result)
}