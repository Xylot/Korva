package main

import (
    "bytes"
    "encoding/binary"
    "os"
)

func getHeaderProperties(file *os.File) HeaderProperties {
	headerProps := HeaderProperties{}
    binary.Read(bytes.NewBuffer(readNextBytesFromFile(file, 4)), binary.LittleEndian, &headerProps.Size)
    binary.Read(bytes.NewBuffer(readNextBytesFromFile(file, 4)), binary.LittleEndian, &headerProps.CRC)
    headerProps.ByteStream = readNextBytesFromFile(file, int(headerProps.Size))
    headerProps.ContentBuffer = bytes.NewBuffer(headerProps.ByteStream)
    return headerProps
}

func getVersionInfo(buf *bytes.Buffer) VersionInfo {
	versionInfo := VersionInfo{}
	count := 0
	bs := make([]byte, 4)
    binary.LittleEndian.PutUint32(bs, 24)
    data := readNextBytes(buf, 4)
	for bs[0] != data[0] {
		buffer := bytes.NewBuffer(data)
		if count == 0 {
			binary.Read(buffer, binary.LittleEndian, &versionInfo.MajorVersion)
		} else if count == 1 {
			binary.Read(buffer, binary.LittleEndian, &versionInfo.MinorVersion)
		} else if count ==2 {
			binary.Read(buffer, binary.LittleEndian, &versionInfo.NetVersion)
		}
		data = readNextBytes(buf, 4)
		count++
	}

	return versionInfo
}

func getGameConstant(buf *bytes.Buffer) GameConstant {
	gameConstant := GameConstant{}
	gameConstant.Length = 24
	gameConstant.Name = convertToString(getStringBytes(buf, gameConstant.Length))

	return gameConstant
}
