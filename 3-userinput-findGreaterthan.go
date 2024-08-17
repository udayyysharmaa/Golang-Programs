package main
import "fmt"

func main(){
	var num1,num2,num3 int
	fmt.Println("Enter the First Number")
	fmt.Scanf("%d", &num1)
	fmt.Println("Enter the Second Number")
	fmt.Scanf("%d", &num2)
	fmt.Println("Enter the Third Number")
	fmt.Scanf("%d", &num3)

	if (num1 >= num2) && (num1 >= num3){
		fmt.Println("The Greater Number is:", num1)
	} else if (num2 >= num1) && (num2 >= num3){
		fmt.Println("The Greater Number is:", num2)
	} else {
		fmt.Println("The Greater Number is:", num3)
	}
}