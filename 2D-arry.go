package main
import "fmt"
func main(){
	var array [3][4] int

	array [0][0] = 1
	array [0][1] = 30
	array [0][2] = 90
	array [0][3] = 200
	array [1][0] = 1
	array [1][1] = 30
	array [1][2] = 90
	array [1][3] = 200
	array [2][0] = 1
	array [2][1] = 30
	array [2][2] = 90
	array [2][3] = 200

	for i:=0; i<3; i++{
		for j:=0; j<4; j++{
			fmt.Print(array[i][j]," ")

		}
		fmt.Println()
	}
}