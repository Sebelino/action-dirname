package main

import "fmt"
import "log"
import "os"
import "encoding/json"
import "path/filepath"

func main() {
	myJsonString := os.Args[1]

	var files []string
	err := json.Unmarshal([]byte(myJsonString), &files)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	dirs := make([]string, 0)
	for _, file := range files {
		dirs = append(dirs, filepath.Dir(file))
	}
	dirBytes, err := json.Marshal(dirs)

	fmt.Println(string(dirBytes))
}
