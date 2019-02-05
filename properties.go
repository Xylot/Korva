package main

import (
    "bytes"
    "encoding/binary"
    "strconv"
	//"github.com/kennygrant/sanitize"
	"strings"
	//"fmt"
	"os"
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
	//return sanitize.Name(getNextString(buf));
	return getNextString(buf)
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
	var goa []Goals
	var hl []HighLights
	var ps []PlayerStats
	for i:=0; i<arrayLength; i++ {
		for j:=0; j<dataLength; j++ {
			readArrayValue(buf, name, headerContents, goals, highlights, playerStats)
		}
		if name == "Goals" {
			goa = append(goa, *goals)
		} else if name == "HighLights" {
			hl = append(hl, *highlights)
		} else if name == "PlayerStats" {
			ps = append(ps, *playerStats)
		}

		getNextString(buf)
	}
	if name == "Goals" {
		headerContents.GoalsList = goa
	} else if name == "HighLights" {
		headerContents.HighlightsList = hl
	} else if name == "PlayerStats" {
		headerContents.PlayerStatsList = ps
	}
}

func readArrayValue(buf *bytes.Buffer, name string, headerContents *FullHeaderContents, goals *Goals, highlights *HighLights, playerStats *PlayerStats) {

	propName := getNextString(buf)
	propType := getNextString(buf)
	propValue := readNextPropertyValue(buf, propType)
	//fmt.Printf("	%+v: %+v\n", propName, propValue)


	if name == "Goals" {
		if strings.ToLower(propName) == "frame" {
			goals.Frame = propValue
		} else if strings.ToLower(propName) == "playername" {
			goals.PlayerName = propValue
		} else if strings.ToLower(propName) == "playerteam" {
			goals.PlayerTeam = propValue
		}
	} else if name == "HighLights" {
		if strings.ToLower(propName) == "frame" {
			highlights.Frame = propValue
		} else if strings.ToLower(propName) == "carname" {
			highlights.CarName = propValue
		} else if strings.ToLower(propName) == "ballname" {
			highlights.BallName = propValue
		}
	} else if name == "PlayerStats" {
		if strings.ToLower(propName) == "name" {
			playerStats.Name = propValue
		} else if strings.ToLower(propName) == "platform" {
			playerStats.Platform = propValue
		} else if strings.ToLower(propName) == "onlineplatform" {
			playerStats.OnlinePlatform = propValue
		} else if strings.ToLower(propName) == "onlineid" {
			playerStats.OnlineID = propValue
		} else if strings.ToLower(propName) == "team" {
			playerStats.Team = propValue
		} else if strings.ToLower(propName) == "score" {
			playerStats.Score = propValue
		} else if strings.ToLower(propName) == "goals" {
			playerStats.Goals = propValue
		} else if strings.ToLower(propName) == "assists" {
			playerStats.Assists = propValue
		} else if strings.ToLower(propName) == "saves" {
			playerStats.Saves = propValue
		} else if strings.ToLower(propName) == "shots" {
			playerStats.Shots = propValue
		} else if strings.ToLower(propName) == "bbot" {
			playerStats.BBot = propValue
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
	propName := getNextString(buf)
	propType := getNextString(buf)
	propValue := readNextPropertyValue(buf, propType)
	//fmt.Printf("???: %+v %+v %+v\n", propName, propType, propValue)
	if propType == "ArrayProperty" {
		count, _ := strconv.Atoi(propValue)
		if propName == "Goals" || propName == "HighLights" {
			readArrayProperty(buf, count, 3, propName, headerContents, goals, highlights, playerStats)
		} else if propName == "PlayerStats" {
			readArrayProperty(buf, count, 11, propName, headerContents, goals, highlights, playerStats)
		}
	}

	if strings.ToLower(propName) == "buildid" {
		headerContents.BuildID = propValue
	} else if strings.ToLower(propName) == "buildversion" {
		headerContents.BuildVersion = propValue
	} else if strings.ToLower(propName) == "changelist" {
		headerContents.ChangeList = propValue
	} else if strings.ToLower(propName) == "date" {
		headerContents.Date = propValue
	} else if strings.ToLower(propName) == "gameversion" {
		headerContents.GameVersion = propValue
	}  else if strings.ToLower(propName) == "id" {
		headerContents.ID = propValue
	} else if strings.ToLower(propName) == "keyframedelay" {
		headerContents.KeyFrameDelay = propValue
	} else if strings.ToLower(propName) == "mapname" {
		headerContents.MapName = propValue
	} else if strings.ToLower(propName) == "matchtype" {
		headerContents.MatchType = propValue
	} else if strings.ToLower(propName) == "maxchannels" {
		headerContents.MaxChannels = propValue
	} else if strings.ToLower(propName) == "maxreplaysize" {
		headerContents.MaxReplaySize = propValue
	} else if strings.ToLower(propName) == "numframes" {
		headerContents.NumFrames = propValue
	} else if strings.ToLower(propName) == "recordfps" {
		headerContents.RecordFPS = propValue
	} else if strings.ToLower(propName) == "replayversion" {
		headerContents.ReplayVersion = propValue
	} else if strings.ToLower(propName) == "team0score" {
		headerContents.Team0Score = propValue
	} else if strings.ToLower(propName) == "team1score" {
		headerContents.Team1Score = propValue
	} else if strings.ToLower(propName) == "teamsize" {
		headerContents.TeamSize = propValue
		return
	} else if strings.ToLower(propName) == "playername" {
		headerContents.PlayerName = propValue
		writeJSON(headerContents)
		os.Exit(3)
	}
}
