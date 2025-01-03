package enumeration

import "fmt"

type HTTPMethod string

const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	DELETE HTTPMethod = "DELETE"
)

func RunEnum() {

	fmt.Println("########## enum ##########")

	var list []HTTPMethod
	list = append(list, GET, POST, PUT, DELETE)

	for i, item := range list {
		fmt.Printf("Index %d: %s\n", i, item)
	}
}
