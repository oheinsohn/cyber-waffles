package main

import (
	"fmt"
)

func main() {
	
	// mix
		fmt.Println("### CyberWaffles ###")
		// getting the scope from resources
		displayScope()
		// use config from env or default
		config := configWaffle()
	
	// form
		fmt.Println("Jenkins installation in: ", EnvWafflePath, config.Path)
		// TODO download Jenkins versions
	
	// bake

	// eat
		lazarusCyberWafflesUI()
	
}