package main
import "fmt"

func main(){
	var num1, num2 int
	fmt.Println("Enter the First Number")
	fmt.Scanf("%d", &num1)
	fmt.Println("Enter the Second Number")
	fmt.Scanf("%d", &num2)
	var operators string
	fmt.Println("Enter the Operators")
	fmt.Scanf("%s", &operators)
	switch operators {
	case "+":
		fmt.Println("Addition is:", (num1 + num2))
	case "-":
		fmt.Println("Subtraction is:", (num1 - num2))
	case "*":
		fmt.Println("multiply is:", (num1 * num2))
	case "/":
		fmt.Println("Devision is:", (num1 / num2))

	default:
		fmt.Println("Enter the Right Operators")
		
	}
	
}