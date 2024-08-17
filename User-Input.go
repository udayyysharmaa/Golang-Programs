package main
import "fmt"

func main(){
	var user string
	fmt.Print("Enter the Name you have!")
	fmt.Scanf("%s", &user)
	fmt.Println("Hello", user, "I hope you are doing well!")
}