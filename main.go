package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
)

type Program struct {
    Name    string
    URL     string
    Platform string
    Active  bool
    Private bool
    Scope   []Asset
}

type Asset struct {
    Asset  string
    Type   string
    Wildcard bool
    Exclude map[string][]string
}

func main() {
	
	if len(os.Args) == 0 {
		fmt.Println("No JSON file provided..")
		os.Exit(1)
	}

	jsonFile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading input JSON file: ", err)
	}

	// Unmarshal the JSON data into a jsonFile
	var prog Program
	err = json.Unmarshal(jsonFile, &prog)
	if err != nil {
		fmt.Println("Error parsing the JSON object: ", err)
	}
	
	if prog.Active {
		for _, item := range prog.Scope {
			if item.Type == "url" && !item.Wildcard {
				fmt.Printf("osmedeus -f recon -t %s -p'program=%s' -p'wildcard=false' \n", item.Asset, prog.Name)
			}
		}

		for _, item := range prog.Scope {
			if item.Type == "url" && item.Wildcard {
				if len(item.Exclude) == 0 {
					fmt.Printf("osmedeus -f recon -t %s -p'program=%s' \n", item.Asset, prog.Name)
				} else {
					excludeObj, _ := json.Marshal(item.Exclude)
					fmt.Printf("osmedeus -f recon -t %s -p'program=%s' -p'exclude=%s' \n", item.Asset, prog.Name, string(excludeObj))
				}
			}
		}
	}

}
