package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type idx struct {
	Files []node `json:"files"`
}

type node struct {
	Filename     string `json:"filename"`
	ResourceType string `json:"resourceType"`
}

type fhirInput struct {
	ResourceType string `json:"resourceType"`
}

func main() {
	nodes := []node{}
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.Contains(path, "/") {
			return nil
		}
		if filepath.Ext(path) != ".json" {
			return nil
		}
		if strings.Contains(path, "-example") {
			return nil
		}
		if path == "package.json" {
			return nil
		}

		raw, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		var input fhirInput
		json.Unmarshal(raw, &input)

		nodes = append(nodes, node{
			Filename:     info.Name(),
			ResourceType: input.ResourceType,
		})
		return nil
	})
	i := idx{
		Files: nodes,
	}
	json, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("json: %s", json)

	fp, err := os.Create(".index.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	data := []byte(json)
	_, _ = fp.Write(data)
}
