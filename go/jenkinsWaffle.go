package main

import (
	"fmt"
	"os"
	"net/http"
	"io"
	"io/fs"
	"strings"
)

func jenkinsWaffle(scope Scope, config *ConfigWaffle) {
	fmt.Println("Jenkins installation in: ", EnvWafflePath, config.Path)
	os.Mkdir("./Jenkins", 0700)
	var version = ""
	for i := 0; i < len(scope.Scope); i++ {
		version = strings.Replace(scope.Scope[i].App,"jenkins_", "", -1)
		version = strings.Replace(version,"_", ".", -1)
		version = strings.Replace(version,".war", "", -1)
		fmt.Println("Downloading Jenkins: " + scope.Scope[i].App + " with version " + version)
		err := downloadJenkins("https://get.jenkins.io/war-stable/"+version+"/jenkins.war", "./Jenkins/"+scope.Scope[i].App)
		if err != nil {
			fmt.Println("Jenkins WAR file download not possible as ", err)
		}
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
}

func downloadJenkins(url string, path string) (err error) {
	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("Jenkins WAR file %s already exists.", path)
	}
	fmt.Println("Download Jenkins Started")
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