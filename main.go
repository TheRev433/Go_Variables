package main

import "fmt"
import "math"

func main() {

	var opt int
	
	
	fmt.Println("Welcome to my programm")
	fmt.Println("Hello user!")
	fmt.Println("Enter the number of the option you want")
	fmt.Println("1. Square area")
	fmt.Println("2. Triangle area")
	fmt.Println("3. Circle area")
	fmt.Println("4. F to C degrees")
	fmt.Println("Press any other number to exit")

	fmt.Scan(&opt)	
}

func operations(opt int) {
	var (
		a float64
		b float64
	)
	switch opt {
	case 1:
		fmt.Println("***Square Area***")
		fmt.Println("Please enter the lenght of side:")
		fmt.Scan(&a)
		fmt.Println("The area is: ", a*a)
		main()
	case 2:
		fmt.Println("***Triangle Area***")
		fmt.Println("Please enter the base:")
		fmt.Scan(&a)
		fmt.Println("Please enter the height:")
		fmt.Scan(&b)
		fmt.Println("The area is: ", a*b*0.5)
		main()
	case 3:
		fmt.Println("***Circle Area***")
		fmt.Println("Please enter the radius:")
		fmt.Scan(&a)
		fmt.Println("The area is: ", (a*a)*math.Pi)
		main()
	case 4:
		fmt.Println("***Fahrenheit to Celsius degress***")
		fmt.Println("Please enter the Fahrenheit degrees:")
		fmt.Scan(&a)
		fmt.Println("It's ", (a-32)*5/9, " Celsius degrees")
		main()
	default: 
		fmt.Println("Thanks for using me")
	}
}