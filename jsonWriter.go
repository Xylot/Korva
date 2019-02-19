package main

import (
    "fmt"
    "encoding/json"
    //"bufio"
    "io/ioutil"
)

func writeJSON(v interface{}, id string) {
	b, err := json.MarshalIndent(v, "", "	")
	//b, err := json.Marshal(v)
    fmt.Printf(string(b))
    if err != nil {}

    fileName := "Resources/Headers/" + id + ".json"

    err = ioutil.WriteFile(fileName, b, 0644)
}
