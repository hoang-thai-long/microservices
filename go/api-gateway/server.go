package main

import (
	"api-gateway/utils"
	"fmt"
	"os"
)

func main() {
	utils.LoadEnvVars()

	fmt.Println("Server start ", os.Getenv("MONGODB_URL"))
}
