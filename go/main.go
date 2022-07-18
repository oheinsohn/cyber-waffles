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
	Port string `json:"Port"`
}

func main() {
	// mix
		fmt.Println("### CyberWaffles ###")
		displayScope()
		config := configWaffle()
	// form
	fmt.Println("Jenkins installation in: ", EnvWafflePath, config.Path)
	// bake
	// eat
	lazarusCyberWafflesUI()
	
}

// TODO: output to logging
func displayScope() {
	scope, err := readScopeFromJSON()
	if err == nil {
		fmt.Println("Entries: ", len(scope.Scope))
		for i := 0; i < len(scope.Scope); i++ {
			fmt.Println("Name: " + scope.Scope[i].Name)
			fmt.Println("IP: " + scope.Scope[i].IP)
			fmt.Println("Port: " + scope.Scope[i].Port)
			fmt.Println(discoverHost(scope.Scope[i].IP, scope.Scope[i].Port))
		}
	} else {
		fmt.Println("An error occured: " + err.Error())
	}
}

// there needs to be a read from a list on which targets it should run
func readScopeFromJSON() (Scope, error) {
	scopeFile, err := os.Open("./resources/scope.json")
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

// trying to connect to IP with tcp on 8080
func discoverHost(ip string, port string) (string, error) {
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		return "", errors.New("ERROR - could not connect to " + ip + " due to: " + err.Error())
	} else {
		conn.Close()
		return "Connection to host established.", nil
	}
}
