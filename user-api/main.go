package main

import "user-api/router"

func main() {
	r := router.InitRouter()
	r.Run(":3000")
}