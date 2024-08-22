package main
import "fmt"

func subnumber (a int, b int) int{
	sum:= a - b
	return sum
}

func main(){
	sumofnumber:= subnumber(16,2)
	fmt.Println("Subtraction of the number", sumofnumber)
}