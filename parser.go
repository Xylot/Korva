package main

import (
    "bytes"
    "encoding/binary"
    "fmt"
    "log"
    "os"
)

func main() {    
    path := "resources/Replays/2015_08.replay"
    if len(os.Args) >= 2 {
        path = os.Args[1]
    }

    parse(path)
}

func parse(fileName string) {

    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal("Error while opening file", err)
    }

    defer file.Close()

    fmt.Printf("%s opened\n", fileName)
    
    hp := getHeaderProperties(file)
    buf := hp.ContentBuffer
    headerContents := FullHeaderContents{}
    goals := Goals{}
    highlights := HighLights{}
    playerStats := PlayerStats{}

    vi := getVersionInfo(buf)
    gc := getGameConstant(buf)

    for i := 0; i < 21; i++ {
        getPropertyAttributes(buf, &headerContents, &goals, &highlights, &playerStats)
    }

    headerContents.Size = hp.Size
    headerContents.CRC = int32(hp.CRC)
    headerContents.MajorVersion = vi.MajorVersion
    headerContents.MinorVersion = vi.MinorVersion
    headerContents.NetVersion = vi.NetVersion
    headerContents.GameConstant = gc.Name
    writeJSON(headerContents)
}

func getContentProperties(file *os.File) ContentProperties {
	contentProps := ContentProperties{}
    binary.Read(bytes.NewBuffer(readNextBytesFromFile(file, 4)), binary.LittleEndian, &contentProps.Size)
    binary.Read(bytes.NewBuffer(readNextBytesFromFile(file, 4)), binary.LittleEndian, &contentProps.CRC)
    contentProps.ByteStream = readNextBytesFromFile(file, int(contentProps.Size))
    contentProps.ContentBuffer = bytes.NewBuffer(contentProps.ByteStream)
    return contentProps
}

func getLevels(buf *bytes.Buffer) {
	count := int(getStringLength(buf))
	for i:=0;i<count;i++ {
		s := getNextString(buf)
		fmt.Printf("???: %+v\n", s)
	}
}

func getKeyframes(buf *bytes.Buffer, count int) {
	skipBytes(buf, 4)
	for i:=0;i<count;i++ {
		time := getFloatValue(buf)
		frame := getIntValue(buf)
		position := getIntValue(buf)
		fmt.Printf("???: %+v %+v %+v\n", time, frame, position)
	}
}
