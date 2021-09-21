package main

import "fmt"

func main() {
	langs := map[string]string{
		"Python": ".py",
		"Golang": ".go",
		"Java": ".java",
		"Ansible": ".yml",
		"C++": ".cpp",
	}

	fmt.Println(langs)

	targetName := "C++"

	_, found := langs[targetName]
	if found == true {
		delete(langs, targetName)
	} else {
		fmt.Println("Could not locate ", targetName)

	}
	langs["Julia"] = ".jl"

	fmt.Println(langs)
}
