
package main

import "os"

func main() {

	_, err := os.Create("/carolDanvers/msMarvel")

	// if we have an unexepected error
	if err != nil {
		panic(err) // fast fail and dump the error to the screen
	}
}
