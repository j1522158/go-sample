package gomap

import "fmt"

func RunMap() {

	fmt.Println("########## map ##########")

	myMap := make(map[string]string)

	myMap["name"] = "Alice"
	myMap["country"] = "Japan"

	fmt.Println(myMap)
	fmt.Println(myMap["name"])

	value, exists := myMap["city"]
	if exists {
		fmt.Println(value)
	} else {
		fmt.Println("Key 'city' does not exist")
	}

	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}
