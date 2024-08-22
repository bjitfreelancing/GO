package main

import "fmt"

func main() {
	if no := 20; no%2 == 0 {
		fmt.Printf("%d is an even number\n", no)
	} else {
		fmt.Printf("%d is an odd number\n", no)
	}

	score := 10
	switch score {
	case 0, 1, 2, 3:
		fmt.Printf("Poor")
	case 4, 5, 6, 7:
		fmt.Printf("Good")
	case 8, 9:
		fmt.Printf("Very Good")
	case 10:
		fmt.Printf("Excellent")
	default:
		fmt.Printf("Invalid score")
	}

	switch no := 25; {
	case no%2 == 0:
		fmt.Printf("\n%d is an even number\n", no)
	case no%3 == 0:
		fmt.Printf("\n%d is divisible by 3\n", no)
	case no%2 == 1:
		fmt.Printf("\n%d is an odd number\n", no)
	}

	// fallthrough applied
	switch spotifySubscription := "super"; spotifySubscription {
	case "super":
		fmt.Println("Family of 3")
		fallthrough
	case "pro":
		fmt.Println("Private play list")
		fallthrough
	case "premium":
		fmt.Println("No ads")
		fallthrough
	case "free":
		fmt.Println("Listen to music")
	}

	// I_LOOP:
	//
	//	for i := 0; i < 5; i++ {
	//		for j := 0; j < 5; j++ {
	//			fmt.Printf("i = %d, j = %d\n", i, j)
	//			if i == j {
	//				fmt.Printf("==============\n")
	//				continue I_LOOP
	//			}
	//		}
	//	}
	//	var x, y int
	//	fmt.Println("Enter 2 numbers :")
	//	fmt.Scanln(&x, &y)
	//	fmt.Println(x + y)
	// LOOP:
	// 	for no := 2; no <= 100; no++ {
	// 		for i := 2; i <= (no / 2); i++ {
	// 			if no%i == 0 {
	// 				continue LOOP
	// 			}
	// 		}
	// 		fmt.Printf("Prime : %d\n", no)
	// 	}
	no := 9
	cnt := 0
	for i := 2; i <= (no / 2); i++ {
		fmt.Printf("\ni: %d", i)
		cnt++
	}

}
