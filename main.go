package main

import (
	"fmt"
	"os"
)
import "WheelChair-tiktok/config"

func main() {
	config.LoadEnv()

	fmt.Println(os.Getenv("MYSQL_DSN"))
}
