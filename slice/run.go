package slice

import (
	"fmt"

	"github.com/joho/godotenv"
)

func RunSlice() {

	fmt.Println("########## slice ##########")

	godotenv.Load()
	var list []string

	fmt.Println(list)

	list = append(list, "Apple")
	list = append(list, "Banana", "Cherry")

	fmt.Println(list)
	fmt.Println(list[0])
	fmt.Println(len(list))

	for i, item := range list {
		fmt.Printf("Index %d: %s\n", i, item)
	}
}
