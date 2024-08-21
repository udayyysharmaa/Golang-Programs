package main

import "fmt"

func main() {
    // Create a 3D array with dimensions 2x3x4
    var array [2][3][4]int

    // Assign values to the 3D array
    value := 1
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            for k := 0; k < 4; k++ {
                array[i][j][k] = value
                value++
            }
        }
    }

    // Print the 3D array
    fmt.Println("The 3D array is:")
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            for k := 0; k < 4; k++ {
                fmt.Printf("array[%d][%d][%d] = %d\n", i, j, k, array[i][j][k])
            }
        }
    }
}
