package main

import (
    "bytes"
    "encoding/binary"
    "log"
    "os"
    "unicode/utf16"
    "fmt"
)


func readNextBytesFromFile(file *os.File, number int) []byte {
    bytes := make([]byte, number)

    _, err := file.Read(bytes)
    if err != nil {
        log.Fatal(err)
    }

    return bytes
}

func readNextBytes(buf *bytes.Buffer, number int) []byte {
    bytes := make([]byte, number)

    _, err := buf.Read(bytes)
    if err != nil {
        log.Fatal(err)
    }

    return bytes
}

func skipBytes(buf *bytes.Buffer, number int) {
	bytes := make([]byte, number)

    _, err := buf.Read(bytes)
    if err != nil {
        log.Fatal(err)
    }
}

func getStringLength(buf *bytes.Buffer) int32 {
	var length int32
	data := readNextBytes(buf, 4)
	buffer := bytes.NewBuffer(data)
    binary.Read(buffer, binary.LittleEndian, &length)

    return length;
}

func getStringBytes(buf *bytes.Buffer, length int) []byte {
	var byteArray []byte
	byteArray = readNextBytes(buf, length)
    return byteArray[:len(byteArray)-1];
}

func getStringBytes1(buf *bytes.Buffer, length int) []uint16 {
    var byteArray []byte
    var utf16Array []uint16
    byteArray = readNextBytes(buf, length)
    fmt.Printf("%v", byteArray)
    if length % 2 == 0 {
        byteArray = byteArray[:len(byteArray)-1]
        for i := 0; i < length/2; i++ {

            utf16Array = append(utf16Array, binary.LittleEndian.Uint16(byteArray[i:i+1]))
            fmt.Printf("%+v", utf16Array)
        }
    } else {
        for i := 0; i < length/2; i++ {
            utf16Array = append(utf16Array, binary.LittleEndian.Uint16(byteArray[i:i+1]))
            fmt.Printf("%+v", utf16Array)            
        }
    }
    
    return utf16Array
}

func convertToString(byteArray []byte) string {
	return string(byteArray[:]);
}

func getNextString(buf *bytes.Buffer) string {
    length := int(getStringLength(buf))
    //fmt.Printf("%+v", length)
    if length < 0 {
        return string(utf16.Decode(getStringBytes1(buf, -1*length)))
    }
	return convertToString(getStringBytes(buf, length))
}

func getFloatValue(buf *bytes.Buffer) float32 {
    var value float32
    data := readNextBytes(buf, 4)
    buffer := bytes.NewBuffer(data)
    binary.Read(buffer, binary.LittleEndian, &value)
    return value;
}

func getIntValue(buf *bytes.Buffer) int {
    var value uint32
    data := readNextBytes(buf, 4)
    buffer := bytes.NewBuffer(data)
    binary.Read(buffer, binary.LittleEndian, &value)
    return int(value);
}

func readUint32(buf *bytes.Buffer) uint32 {
    var value uint32
    buffer := bytes.NewBuffer(readNextBytes(buf, 4))
    binary.Read(buffer, binary.LittleEndian, &value)
    return value;
}