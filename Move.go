package main

import "fmt"

func main() {
	fmt.Println("Enter the number of miles driven")
	var milesMove float64
	fmt.Scanf("%f", &milesMove)
	fmt.Println("Enter the number of gallon used")
	var numberOfGallon float64
	fmt.Scanf("%f", &numberOfGallon)
	totalMilesPerGallon := milesMove / numberOfGallon
	fmt.Printf("The total miles driven is per gallon is %.2fn\n", totalMilesPerGallon)
}
