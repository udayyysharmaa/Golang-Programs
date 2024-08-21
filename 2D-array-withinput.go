package main

import (
    "fmt"
)

func main() {
    var rows, cols int

    // Taking the number of rows and columns as input from the user
    fmt.Print("Enter the number of rows: ")
    fmt.Scan(&rows)
    fmt.Print("Enter the number of columns: ")
    fmt.Scan(&cols)

    // Creating a 2D array with user-specified dimensions
    array := make([][]int, rows)
    for i := range array {
        array[i] = make([]int, cols)
    }

    // Taking input for each element of the 2D array
    fmt.Println("Enter the elements of the array:")
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            fmt.Printf("Element [%d][%d]: ", i, j)
            fmt.Scan(&array[i][j])
        }
    }

    // Printing the 2D array
    fmt.Println("The 2D array is:")
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            fmt.Print(array[i][j], " ")
        }
        fmt.Println()
    }
}