package main

import "fmt"

func main() {
	fmt.Println("How much is your monthly salary")
	var monthlyAlawee float64
	fmt.Scanf("%fn", &monthlyAlawee)
	if monthlyAlawee <= 30000 {
		totalPaymentTax := monthlyAlawee * 0.15
		fmt.Printf("Your payable tax for this month is %.2f\n", totalPaymentTax)
	}
	if monthlyAlawee > 30000 {
		totalPaymentTax := monthlyAlawee * 0.20
		fmt.Printf("Your payable tax for this month is %.2f\n", totalPaymentTax)
	}
}
