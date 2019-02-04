package main

import (
    "bytes"
    "encoding/binary"
    "strconv"
	"github.com/kennygrant/sanitize"
	"strings"
	"fmt"
)

func readIntProperty(buf *bytes.Buffer) int {
	var value uint32
	skipBytes(buf, 8)
	data := readNextBytes(buf, 4)
	buffer := bytes.NewBuffer(data)
    binary.Read(buffer, binary.LittleEndian, &value)
    return int(value);
}

func readStrProperty(buf *bytes.Buffer) string {
	skipBytes(buf, 8)
	return sanitize.Name(getNextString(buf));
}

func readByteProperty(buf *bytes.Buffer) int {
	var value uint32
	skipBytes(buf, 4)
	data := readNextBytes(buf, 4)
	buffer := bytes.NewBuffer(data)
    binary.Read(buffer, binary.LittleEndian, &value)
    return int(value);
}

func readQWordProperty(buf *bytes.Buffer) int {
	var value uint64
	skipBytes(buf, 8)
	data := readNextBytes(buf, 8)
	buffer := bytes.NewBuffer(data)
    binary.Read(buffer, binary.LittleEndian, &value)
    return int(value);
}

func readBooleanProperty(buf *bytes.Buffer) int {
	var value uint8
	skipBytes(buf, 8)
	data := readNextBytes(buf, 1)
	buffer := bytes.NewBuffer(data)
    binary.Read(buffer, binary.LittleEndian, &value)
    return int(value);
}

func readFloatProperty(buf *bytes.Buffer) int {
	var value float32
	skipBytes(buf, 8)
	data := readNextBytes(buf, 4)
	buffer := bytes.NewBuffer(data)
    binary.Read(buffer, binary.LittleEndian, &value)
    return int(value);
}

func readArrayProperty(buf *bytes.Buffer, arrayLength int, dataLength int, name string, headerContents *FullHeaderContents, goals *Goals, highlights *HighLights, playerStats *PlayerStats) {
	//var arr [][]string
	var goa []Goals
	var hl []HighLights
	var ps []PlayerStats
	for i:=0; i<arrayLength; i++ {
		for j:=0; j<dataLength; j++ {
			readArrayValue(buf, name, headerContents, goals, highlights, playerStats)
			//arr[i] = append(arr[i], readArrayValue(buf))
		}
		if name == "Goals" {
			goa = append(goa, *goals)
		} else if name == "HighLights" {
			hl = append(hl, *highlights)
		} else if name == "PlayerStats" {
			ps = append(ps, *playerStats)
		}
		//arr = append(arr, currentFrame)
		getNextString(buf)
	}
	if name == "Goals" {
		//fmt.Printf("Goals: %+v\n", goa)
		//writeJSON(goa)
		headerContents.GoalsList = goa
	} else if name == "HighLights" {
		//fmt.Printf("HighLights: %+v\n", hl)
		//writeJSON(hl)
		headerContents.HighlightsList = hl
	} else if name == "PlayerStats" {
		//fmt.Printf("Player Stats: %+v\n", ps)
		//writeJSON(ps)
		headerContents.PlayerStatsList = ps
	}
	
	//return arr
}

func readArrayValue(buf *bytes.Buffer, name string, headerContents *FullHeaderContents, goals *Goals, highlights *HighLights, playerStats *PlayerStats) {
	//prop := Property{}
	// prop.Name = getNextString(buf)
	// prop.Type = getNextString(buf)
	// prop.Value = readNextPropertyValue(buf, prop.Type)
	propName := getNextString(buf)
	propType := getNextString(buf)
	propValue := readNextPropertyValue(buf, propType)
	fmt.Printf("	%+v: %+v\n", propName, propValue)
	//fmt.Printf("Array value: %+v\n", prop)
	//return prop.Value

	if name == "Goals" {
		if strings.ToLower(propName) == "frame" {
			goals.Frame = propValue
			//fmt.Printf("%+v: %+v\n", propName, propValue)
		} else if strings.ToLower(propName) == "playername" {
			goals.PlayerName = propValue
			//fmt.Printf("%+v: %+v\n", propName, propValue)
		} else if strings.ToLower(propName) == "playerteam" {
			goals.PlayerTeam = propValue
			//fmt.Printf("%+v: %+v\n", propName, propValue)
		}
	} else if name == "HighLights" {
		if strings.ToLower(propName) == "frame" {
			highlights.Frame = propValue
			//fmt.Printf("%+v: %+v\n", propName, propValue)
		} else if strings.ToLower(propName) == "carname" {
			highlights.CarName = propValue
			//fmt.Printf("%+v: %+v\n", propName, propValue)
		} else if strings.ToLower(propName) == "ballname" {
			highlights.BallName = propValue
			//fmt.Printf("%+v: %+v\n", propName, propValue)
		}
	} else if name == "PlayerStats" {
		if strings.ToLower(propName) == "name" {
			playerStats.Name = propValue
			//fmt.Printf("%+v: %+v\n", propName, propValue)
		} else if strings.ToLower(propName) == "platform" {
			playerStats.Platform = propValue
			//fmt.Printf("%+v: %+v\n", propName, propValue)
		} else if strings.ToLower(propName) == "onlineplatform" {
			playerStats.OnlinePlatform = propValue
			//fmt.Printf("%+v: %+v\n", propName, propValue)
		} else if strings.ToLower(propName) == "onlineid" {
			playerStats.OnlineID = propValue
			//fmt.Printf("%+v: %+v\n", propName, propValue)
		} else if strings.ToLower(propName) == "team" {
			playerStats.Team = propValue
			//fmt.Printf("%+v: %+v\n", propName, propValue)
		} else if strings.ToLower(propName) == "score" {
			playerStats.Score = propValue
			//fmt.Printf("%+v: %+v\n", propName, propValue)
		} else if strings.ToLower(propName) == "goals" {
			playerStats.Goals = propValue
			//fmt.Printf("%+v: %+v\n", propName, propValue)
		} else if strings.ToLower(propName) == "assists" {
			playerStats.Assists = propValue
			//fmt.Printf("%+v: %+v\n", propName, propValue)
		} else if strings.ToLower(propName) == "saves" {
			playerStats.Saves = propValue
			//fmt.Printf("%+v: %+v\n", propName, propValue)
		} else if strings.ToLower(propName) == "shots" {
			playerStats.Shots = propValue
			//fmt.Printf("%+v: %+v\n", propName, propValue)
		} else if strings.ToLower(propName) == "bbot" {
			playerStats.BBot = propValue
			//fmt.Printf("%+v: %+v\n", propName, propValue)
		}
	}

}

func readNextPropertyValue(buf *bytes.Buffer, propertyType string) string {
	if propertyType == "IntProperty" {
		return strconv.Itoa(readIntProperty(buf));
	} else if propertyType == "StrProperty" || propertyType == "NameProperty" {
		return readStrProperty(buf);
	} else if propertyType == "ArrayProperty" {
		return strconv.Itoa(readIntProperty(buf));
	} else if propertyType == "ByteProperty" {
		return strconv.Itoa(readByteProperty(buf));
	} else if propertyType == "QWordProperty" {
		return strconv.Itoa(readQWordProperty(buf));
	} else if propertyType == "BoolProperty" {
		return strconv.Itoa(readBooleanProperty(buf));
	} else if propertyType == "FloatProperty" {
		return strconv.Itoa(readFloatProperty(buf));
	}
	return "Useless piece of software"
}

func getPropertyAttributes(buf *bytes.Buffer, headerContents *FullHeaderContents, goals *Goals, highlights *HighLights, playerStats *PlayerStats) {
	//headerContents := HContents{}
	propName := getNextString(buf)
	propType := getNextString(buf)
	propValue := readNextPropertyValue(buf, propType)
	//fmt.Printf("%+v: %+v\n", propName, propValue)

	fmt.Printf("???: %+v %+v %+v\n", propName, propType, propValue)
	if propType == "ArrayProperty" {
		count, _ := strconv.Atoi(propValue)
		if propName == "Goals" || propName == "HighLights" {
			readArrayProperty(buf, count, 3, propName, headerContents, goals, highlights, playerStats)
			// arr := readArrayProperty(buf, count, 3)
			// fmt.Printf("???: %+v", arr)
		} else if propName == "PlayerStats" {
			readArrayProperty(buf, count, 11, propName, headerContents, goals, highlights, playerStats)
			// arr := readArrayProperty(buf, count, 11)
			// fmt.Printf("???: %+v", arr)
		}
	}

	//fmt.Printf("???: %+v 	", strings.ToLower(propName))

	if strings.ToLower(propName) == "buildid" {
		headerContents.BuildID = propValue
		//fmt.Printf("%+v: %+v\n", propName, propValue)
	} else if strings.ToLower(propName) == "buildversion" {
		headerContents.BuildVersion = propValue
		//fmt.Printf("%+v: %+v\n", propName, propValue)
	} else if strings.ToLower(propName) == "changelist" {
		headerContents.ChangeList = propValue
		//fmt.Printf("%+v: %+v\n", propName, propValue)
	} else if strings.ToLower(propName) == "date" {
		headerContents.Date = propValue
		//fmt.Printf("%+v: %+v\n", propName, propValue)
	} else if strings.ToLower(propName) == "gameversion" {
		headerContents.GameVersion = propValue
		//fmt.Printf("%+v: %+v\n", propName, propValue)
	}  else if strings.ToLower(propName) == "id" {
		headerContents.ID = propValue
		//fmt.Printf("%+v: %+v\n", propName, propValue)
	} else if strings.ToLower(propName) == "keyframedelay" {
		headerContents.KeyFrameDelay = propValue
		//fmt.Printf("%+v: %+v\n", propName, propValue)
	} else if strings.ToLower(propName) == "mapname" {
		headerContents.MapName = propValue
		//fmt.Printf("%+v: %+v\n", propName, propValue)
	} else if strings.ToLower(propName) == "matchtype" {
		headerContents.MatchType = propValue
		//fmt.Printf("%+v: %+v\n", propName, propValue)
	} else if strings.ToLower(propName) == "maxchannels" {
		headerContents.MaxChannels = propValue
		//fmt.Printf("%+v: %+v\n", propName, propValue)
	} else if strings.ToLower(propName) == "maxreplaysize" {
		headerContents.MaxReplaySize = propValue
		//fmt.Printf("%+v: %+v\n", propName, propValue)
	} else if strings.ToLower(propName) == "numframes" {
		headerContents.NumFrames = propValue
		//fmt.Printf("%+v: %+v\n", propName, propValue)
	} else if strings.ToLower(propName) == "recordfps" {
		headerContents.RecordFPS = propValue
		//fmt.Printf("%+v: %+v\n", propName, propValue)
	} else if strings.ToLower(propName) == "replayversion" {
		headerContents.ReplayVersion = propValue
		//fmt.Printf("%+v: %+v\n", propName, propValue)
	} else if strings.ToLower(propName) == "team0score" {
		headerContents.Team0Score = propValue
		//fmt.Printf("%+v: %+v\n", propName, propValue)
	} else if strings.ToLower(propName) == "team1score" {
		headerContents.Team1Score = propValue
		//fmt.Printf("%+v: %+v\n", propName, propValue)
	} else if strings.ToLower(propName) == "teamsize" {
		headerContents.TeamSize = propValue
		//fmt.Printf("%+v: %+v\n", propName, propValue)
	}

	
	// else if strings.ToLower(propName) == "goals" {
	// 	headerContents.Goals = propValue
	// } else if strings.ToLower(propName) == "highlights" {
	// 	headerContents.HighLights = propValue
	// } else if strings.ToLower(propName) == "playerstats" {
	// 	headerContents.PlayerStats = propValue
	// }
	//fmt.Printf("Header contents: %+v\n", headerContents)
	//writeJSON(headerContents)
}
