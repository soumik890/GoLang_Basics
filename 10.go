package main

import "fmt"

func main() {

	fmt.Println("Tenth File")

	userInfo := map[string]int{
		"soumik": 30,
		"piyali": 31,
		"soma":   35,
	}

	fmt.Println(userInfo)
	fmt.Println(userInfo["soumik"])
	userInfo["santosh"] = 67
	fmt.Println(userInfo)

	for key, value := range userInfo {
		fmt.Println(key, value, "****one by one iterarting*****")
	}

}
