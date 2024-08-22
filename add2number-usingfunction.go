package main
import "fmt"
func addnumber(a int, b int) int{
	sum := a + b
	return sum
}


func main(){
	sumofnumber := addnumber(14,15)
	fmt.Println("Addition of the number", sumofnumber)
}

