package main

import (
    "fmt"
    "encoding/json"
)

type User struct {
    Name string
}

func testJSON() {
    user := &User{Name: "Frank"}
    b, err := json.Marshal(user)
    if err != nil {
        return
    }
    fmt.Printf(string(b))
}

func writeJSON(v interface{}) {
	b, err := json.MarshalIndent(v, "", "	")
	//b, err := json.Marshal(v)
    fmt.Printf(string(b))
    if err != nil {}
}