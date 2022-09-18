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
	dirMap := make(map[string]bool)
	for _, file := range files {
        dir := filepath.Dir(file)
		dirMap[dir] = true
	}
    dirs := make([]string, len(dirMap))
    i := 0
    for d := range dirMap {
        dirs[i] = d
        i++
    }
    dirBytes, err := json.Marshal(dirs)

	fmt.Println(string(dirBytes))
}
