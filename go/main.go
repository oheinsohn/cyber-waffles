package main

import (
	"fmt"
	"os"
	"net/http"
	"io"
	"io/fs"
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
		os.Mkdir("./Jenkins", 0700)
		err := downloadJenkins("https://get.jenkins.io/war-stable/2.277.4/jenkins.war", "./Jenkins/jenkins_2_277_4.war")
		if err != nil {
			fmt.Println("Jenkins WAR file download not possible as ", err)
		}
		diretory_entries, err := discoverDirectory("./Jenkins")
		if err != nil {
			fmt.Println("Could not read directory ./Jenkins", err)
		}
		fmt.Println("Files in ./Jenkins:")
		for i := 0; i < len(diretory_entries); i++ {
			fmt.Println("	", diretory_entries[i].Name())
		}
		return
	
	// bake

	// eat
		lazarusCyberWafflesUI()
	
}

func downloadJenkins(url string, path string) (err error) {
	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("Jenkins WAR file %s already exists.", path)
	}
	fmt.Println("Download Jenkins")
	warFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer warFile.Close()
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP status not okay: %s", response.Status)
	}
	_, err = io.Copy(warFile, response.Body)
	if err != nil {
		return err
	}
	fmt.Println("Download Jenkins Completed")
	return nil
}

func discoverDirectory(path string) ([]fs.DirEntry, error) {
	list_of_directory_entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	} else {
		return list_of_directory_entries, nil
	}
}