package main

import (
    "bytes"
    "encoding/binary"
    "fmt"
    "log"
    "os"
    //"strconv"
    //"math"
    //"strings"
	//"github.com/kennygrant/sanitize"
)

func main() {
    //files := gf()
    // fmt.Println("%+v", files)
    
    path := "resources/Replays/2016_09.replay"
    if len(os.Args) >= 2 {
        path = os.Args[1]
    }

    // fmt.Printf("Path: %+v", path)

    // for _, filePath := range files {
    //     parse(filePath)
        
    // }
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
    //fmt.Printf("hp: %+v\n", hp)
    //headerContents := HContents{}
    headerContents := FullHeaderContents{}
    goals := Goals{}
    highlights := HighLights{}
    playerStats := PlayerStats{}


    //writeJSON(getVersionInfo(buf))
    //writeJSON(getGameConstant(buf))

    vi := getVersionInfo(buf)
    fmt.Printf("Version info: %+v\n", vi)
    gc := getGameConstant(buf)
    fmt.Printf("Game constant: %+v\n", gc)

    for i := 0; i < 21; i++ {
        getPropertyAttributes(buf, &headerContents, &goals, &highlights, &playerStats)
        // err2 := getPropertyAttributes(buf, &headerContents, &goals, &highlights, &playerStats)
        // if err2 != nil {
        //     log.Fatal(err2)
        //     return
        // }
    }

    // getPropertyAttributes(buf, &headerContents, &goals, &highlights, &playerStats)
    // getPropertyAttributes(buf, &headerContents, &goals, &highlights, &playerStats)
    // getPropertyAttributes(buf, &headerContents, &goals, &highlights, &playerStats)
    // getPropertyAttributes(buf, &headerContents, &goals, &highlights, &playerStats)
    // getPropertyAttributes(buf, &headerContents, &goals, &highlights, &playerStats)
    // getPropertyAttributes(buf, &headerContents, &goals, &highlights, &playerStats)
    // getPropertyAttributes(buf, &headerContents, &goals, &highlights, &playerStats)
    // getPropertyAttributes(buf, &headerContents, &goals, &highlights, &playerStats)
    // getPropertyAttributes(buf, &headerContents, &goals, &highlights, &playerStats)
    // getPropertyAttributes(buf, &headerContents, &goals, &highlights, &playerStats)
    // getPropertyAttributes(buf, &headerContents, &goals, &highlights, &playerStats)
    // getPropertyAttributes(buf, &headerContents, &goals, &highlights, &playerStats)
    // getPropertyAttributes(buf, &headerContents, &goals, &highlights, &playerStats)
    // getPropertyAttributes(buf, &headerContents, &goals, &highlights, &playerStats)
    // getPropertyAttributes(buf, &headerContents, &goals, &highlights, &playerStats)
    // getPropertyAttributes(buf, &headerContents, &goals, &highlights, &playerStats)
    // getPropertyAttributes(buf, &headerContents, &goals, &highlights, &playerStats)
    // getPropertyAttributes(buf, &headerContents, &goals, &highlights, &playerStats)
    // getPropertyAttributes(buf, &headerContents, &goals, &highlights, &playerStats)
    // getPropertyAttributes(buf, &headerContents, &goals, &highlights, &playerStats)
    // getPropertyAttributes(buf, &headerContents, &goals, &highlights, &playerStats)

    //fmt.Printf("DID IT WORK: %+v\n", headerContents)
    headerContents.Size = hp.Size
    headerContents.CRC = int32(hp.CRC)
    headerContents.MajorVersion = vi.MajorVersion
    headerContents.MinorVersion = vi.MinorVersion
    headerContents.NetVersion = vi.NetVersion
    headerContents.GameConstant = gc.Name
    //writeJSON(headerContents)
    //writeJSON(goals)
    // cp := getContentProperties(file)
    // buf = cp.ContentBuffer
    // getLevels(buf)
    // getKeyframes(buf, 40)
    //parseNetworkStream(buf)
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



// func getBoolValue(buf *bytes.Buffer) bool {
// 	var value uint
// 	data := readNextBytes(buf, 1)
// 	buffer := bytes.NewBuffer(data)
//     binary.Read(buffer, binary.LittleEndian, &value)
//     return bool(value);
// }

func getKeyframes(buf *bytes.Buffer, count int) {
	skipBytes(buf, 4)
	for i:=0;i<count;i++ {
		time := getFloatValue(buf)
		frame := getIntValue(buf)
		position := getIntValue(buf)
		fmt.Printf("???: %+v %+v %+v\n", time, frame, position)
	}
}


// func readUint32Max(buf *bytes.Buffer, maxValue int) uint32 {
// 	maxBits := math.Floor(math.Log10(float64(maxValue)) / math.Log10(2)) + 1
// 	fmt.Printf("Max bits: %+v\n", maxBits)
// 	value := 0
// 	for i := 0; i < int(maxBits) && (value + (1 << uint(i))) < maxValue; i++ {
// 		if getBoolValue(buf) {
// 			value += uint32(1) << i
// 		} else{
// 			value += uint32(0) << i
// 		}
// 	}
// 	return value;
// }




// func parseNetworkStream(buf *bytes.Buffer, maxChannels int, engineVersion uint32, licenseeVersion uint32) {
// 	skipBytes(buf, 4)
// 	actorState := 1
// 	time := getFloatValue(buf)
// 	delta := getFloatValue(buf)
// 	fmt.Printf("Frame propeties: %+v %+v\n", time, delta)
// 	actorID := readUint32Max(maxChannels)
// 	if readBool(buf) {
// 		if readBool(buf) {
// 			actorState = 1
// 			if engineVersion > 868 || (engineVersion == 868 && licenseeVersion >= 14) {
// 				nameID := readNextBytes(buf, 4)
// 			}
// 			unknown1 := readBool(buf)
// 			typeID := readUInt32(buf)
// 			typeName := "" //objectIndexToName(typeID)
// 			classNetCache := "" //objectNameToClassNetCache(typeName, classNetCacheByName)
// 			classID := "" //classNetCache.objectIndex
// 			if !initialPos() {
// 				return
// 			}
// 			pos := readVector(buf, netVersion)
// 		}
// 	}
// }