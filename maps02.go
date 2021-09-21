
package main

import "fmt"

func main() {
	// create a map type
	hostResolution := map[string]string{
		"prom": "10.0.22.32:9090",
		"kafka": "10.7.8.99:9092",
		"minecraft": "192.168.59.11:25565",
	}

	// name to lookup
	hostName := "hugo"

	value, ok := hostResolution[hostName]
	if ok == true {
		fmt.Println("The value of ", hostName, "is", value)
		return
	}
	fmt.Println("Socket Information for", hostName, "not found")

}
