package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
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
	displayScope()
}

// TODO: output to logging
func displayScope() {
	scope, err := readScopeFromJSON()
	if err == nil {
		fmt.Println("Entries: ", len(scope.Scope))
		for i := 0; i < len(scope.Scope); i++ {
			fmt.Println("Name: " + scope.Scope[i].Name)
			fmt.Println("IP: " + scope.Scope[i].IP)
		}
	} else {
		fmt.Println("An error occured: " + err.Error())
	}
}

// there needs to be a read from a list on which targets it should run
func readScopeFromJSON() (Scope, error) {
	scopeFile, err := os.Open("scope.json")
	if err != nil {
		fmt.Println("Could not read json file", err)
	}
	defer scopeFile.Close()
	byteValue, _ := ioutil.ReadAll(scopeFile)
	var scope Scope
	json.Unmarshal(byteValue, &scope)
	for i := 0; i < len(scope.Scope); i++ {
		err := checkIP(scope.Scope[i].IP)
		if err != nil {
			return scope, err
		}
	}
	return scope, nil
}

// sanity check for IPs
func checkIP(ip string) error {

	// net.ParseIP gives nil if not valid; accepts IPv4, IPv6 and mapped IPv6to4
	ipCheck := net.ParseIP(ip)
	if ipCheck != nil {
		return nil
	} else {
		return errors.New("could not parse IP")
	}
}
