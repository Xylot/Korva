package main

import (
    "io/ioutil"
    "log"
)

func gf() []string {
    var fileList []string
    files, err := ioutil.ReadDir("./Resources/Replays/")
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
        fileList = append(fileList, "resources/Replays/" + f.Name())
        //fmt.Println("resources/Replays/" + f.Name())
    }

    return fileList
}