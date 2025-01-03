package calculator

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func RunCalculator() {
	fmt.Println("########## calculator ##########")

	// envの利用
	godotenv.Load()
	fmt.Println("env: " + os.Getenv("GO_ENV"))

	fmt.Println(Offset)
	fmt.Println(Sum(1, 2))
	fmt.Println(Multiply(1, 2))
}
