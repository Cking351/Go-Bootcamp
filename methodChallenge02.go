package main

import "fmt"

type Virtmach struct {
	ip string
	hostname string
	diskgb int
	ram int
}

func (vir Virtmach) DisplayInfo() {
	fmt.Println(vir)
}

func (vir *Virtmach) IncreaseRam() {
	vir.ram++
}

func main() {
	machine1 := Virtmach{"192.168.1.1", "linux/amd", 234453, 234430}
	machine1.DisplayInfo()
	machine1.IncreaseRam()
	machine1.DisplayInfo()
}
