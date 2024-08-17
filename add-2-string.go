package main
import "fmt"

func main(){
	var FirstName, SecondName string
	fmt.Println("Enter the FirstName!")
	fmt.Scanf("%s", &FirstName)
	fmt.Println("Enter the SecondName!")
	fmt.Scanf("%s", &SecondName)

	var FullName string

	FullName = FirstName + SecondName

	fmt.Println("Your Full Name is:", FullName)
}