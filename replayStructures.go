package main

import (
    "bytes"
)

type HeaderProperties struct {
	Size uint32
	CRC uint32
	ByteStream []byte
	ContentBuffer *bytes.Buffer
}

type HeaderContents struct {
	BuildID uint32
	BuildVersion string
	ChangeList uint32
	Date string
	GameVersion uint32
	Goals []string
	HighLights []string
	ID string
	KeyFrameDelay uint32
	MapName string
	MatchType string
	MaxChannels uint32
	MaxReplaySize uint32
	NumFrames uint32
	PlayerStats []string
	RecordFPS uint32
	ReplayVersion uint32
	Team0Score uint32
	Team1Score uint32
	TeamSize uint32
}

type HContents struct {
	BuildID string
	BuildVersion string
	ChangeList string
	Date string
	GameVersion string
	ID string
	KeyFrameDelay string
	MapName string
	MatchType string
	MaxChannels string
	MaxReplaySize string
	NumFrames string
	RecordFPS string
	ReplayVersion string
	Team0Score string
	Team1Score string
	TeamSize string
}

type FullHeaderContents struct {
	Size uint32
	CRC int32
	MajorVersion uint32
    MinorVersion uint32
    NetVersion uint32
    GameConstant string
	BuildID string
	BuildVersion string
	ChangeList string
	Date string
	GameVersion string
	GoalsList []Goals
	HighlightsList []HighLights
	ID string
	KeyFrameDelay string
	MapName string
	MatchType string
	MaxChannels string
	MaxReplaySize string
	NumFrames string
	PlayerStatsList []PlayerStats
	RecordFPS string
	ReplayVersion string
	Team0Score string
	Team1Score string
	TeamSize string
	PlayerName string
}

type ContentProperties struct {
	Size uint32
	CRC int32
	ByteStream []byte
	ContentBuffer *bytes.Buffer
}

type VersionInfo struct {
    MajorVersion uint32
    MinorVersion uint32
    NetVersion uint32
}

type GameConstant struct {
	Length int
	Name string
}

type Property struct {
	Name string
	Type string
	Value string
}

type Goals struct {
	Frame string
	PlayerName string
	PlayerTeam string
}

type GoalsList struct {
	GL []Goals
}

type HighLights struct {
	Frame string
	CarName string
	BallName string
}

// type PlayerStats struct {
// 	Name string
// 	Platform string
// 	OnlinePlatform string
// 	OnlineID string
// 	Team string
// 	Score string
// 	Goals string
// 	Assists string
// 	Saves string
// 	Shots string
// 	BBot string
// }

type PlayerStats struct {
	Name string
	Platform PlatformProperty
	Team string
	Score string
	Goals string
	Assists string
	Saves string
	Shots string
	BBot string
}

type IntProperty struct {
	Name_length uint32
	Name []byte
	Type_length uint32
	Type []byte
	Nullspace uint32
	Value uint32
}

type PlatformProperty struct {
	OnlinePlatform string
	OnlineID string
}

type ActorState int

const (
	DELETED ActorState = 0
	NEW ActorState = 1
	EXISTING ActorState = 2
)

type Vector3D struct {
	DX uint32
	DY uint32
	DZ uint32
	X float32
	Y float32
	Z float32
}