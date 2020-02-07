package main

import "log"

func main() {
	r := SetupRouter()
	if err := r.Run(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("service run...")
	}
}
