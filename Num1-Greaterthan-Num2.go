package main
import "fmt"

func main(){
	var num1,num2 int
	fmt.Println("Enter the First NUmber")
	fmt.Scanf("%d", &num1)
	fmt.Println("Enter the Second NUmber")
	fmt.Scanf("%d", &num2)
	if (num1 > num2) {
		fmt.Println("The Greater Number is:", num1)

	} else {
		fmt.Println("The Greater Number is:", num2)
	}
}

