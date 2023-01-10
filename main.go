package main

import (
	"fmt"
	"time"
	"os"
	"os/exec"
	"io/ioutil"
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
				cmd := fmt.Sprintf("osmedeus scan -f recon -t %s -p'program=%s,wildcard=false' \n", item.Asset, prog.Name)
				fmt.Printf(cmd)
				
				exec.Command("bash", "-c", cmd).Output()
			}
		}

		for _, item := range prog.Scope {
			if item.Type == "url" && item.Wildcard {
				if len(item.Exclude) == 0 {
					cmd := fmt.Sprintf("osmedeus scan -f recon -t %s -p'program=%s' \n", item.Asset, prog.Name)
					fmt.Printf(cmd)

					exec.Command("bash", "-c", cmd).Output()
				} else {
					excludeJson, _ := json.Marshal(item.Exclude)
					exFileName := fmt.Sprintf("EXFILE-%s-%s-%d.json",prog.Name, item.Asset, time.Now().Unix())
					err = ioutil.WriteFile("/tmp/" + exFileName, excludeJson, 0777)
					if err != nil { fmt.Println("Error writing exclude file: ", err) }
					cmd := fmt.Sprintf("osmedeus scan -f recon -t %s -p'program=%s,excludeFile=%s' \n", item.Asset, prog.Name, "/tmp/"+exFileName)
					fmt.Printf(cmd)

					exec.Command("bash", "-c", cmd).Output()
				}
			}
		}
	}
}
