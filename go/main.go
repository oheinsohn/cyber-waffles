package main

import (
	"fmt"
)

func main() {
	
	// mix
		fmt.Println("### CyberWaffles ###")
		// getting the scope from resources
		scope := getScope()
		// use config from env or default
		config := configWaffle()
	
	// form
		jenkinsWaffle(scope, config)

		
	// bake

	// eat
		lazarusCyberWafflesUI()
	
}

