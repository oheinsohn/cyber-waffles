package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Scope struct {
	Scope []Target `json:"scope"`
}

type Target struct {
	Name string `json:"name"`
	IP   string `json:"IP"`
}

func main() {
	fmt.Println("### CyberWaffles ###")
	readScope()
}

// there needs to be a read from a list on which targets it should run
func readScope() {
	scopeFile, err := os.Open("scope.json")
	if err != nil {
		fmt.Println("Could not read json file", err)
	}
	defer scopeFile.Close()
	byteValue, _ := ioutil.ReadAll(scopeFile)
	var scope Scope
	json.Unmarshal(byteValue, &scope)

	fmt.Println("Entries: ", len(scope.Scope))
	for i := 0; i < len(scope.Scope); i++ {
		fmt.Println("Name: " + scope.Scope[i].Name)
		fmt.Println("IP: " + scope.Scope[i].IP)
	}
}
