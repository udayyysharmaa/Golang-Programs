package main
import "fmt"

func main(){
	var user string
	var is_monkey bool
	fmt.Println("Enter your Name!")
	fmt.Scanf("%s", &user, "\n")
	fmt.Println("You are moneky or not (False and True)")
	fmt.Scanf("%t", &is_monkey, "\n")
	fmt.Println("Hello",  user, "You are monkey:", is_monkey)

}