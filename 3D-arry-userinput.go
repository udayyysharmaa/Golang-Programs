package main

import "fmt"

func main() {
    var x, y, z int

    // Input dimensions for the 3D array
    fmt.Print("Enter the dimensions (x, y, z): ")
    fmt.Scan(&x, &y, &z)

    // Create a 3D array with user-specified dimensions
    array := make([][][]int, x)
    for i := range array {
        array[i] = make([][]int, y)
        for j := range array[i] {
            array[i][j] = make([]int, z)
        }
    }

    // Input values for the 3D array
    fmt.Println("Enter the elements of the array:")
    for i := 0; i < x; i++ {
        for j := 0; j < y; j++ {
            for k := 0; k < z; k++ {
                fmt.Printf("Element [%d][%d][%d]: ", i, j, k)
                fmt.Scan(&array[i][j][k])
            }
        }
    }

    // Print the 3D array
    fmt.Println("The 3D array is:")
    for i := 0; i < x; i++ {
        for j := 0; j < y; j++ {
            for k := 0; k < z; k++ {
                fmt.Printf("array[%d][%d][%d] = %d\n", i, j, k, array[i][j][k])
            }
        }
    }
}
