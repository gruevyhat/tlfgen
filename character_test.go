package tlfgen

import (
	"testing"
)

func TestProfessions(t *testing.T) {
	logLevel := "ERROR"
	if testing.Verbose() {
		logLevel = "INFO"
	}
	opts := []Opts{}
	for _, profession := range professions {
		opts = append(opts, Opts{
			Profession: profession,
			LogLevel:   logLevel,
		})
	}
	for _, o := range opts {
		NewCharacter(o)
	}
}

func TestPersonalityTypes(t *testing.T) {
	logLevel := "ERROR"
	if testing.Verbose() {
		logLevel = "INFO"
	}
	opts := []Opts{}
	for _, personalityType := range personalityTypes {
		opts = append(opts, Opts{
			PersonalityType: personalityType,
			LogLevel:        logLevel,
		})
	}
	for _, o := range opts {
		NewCharacter(o)
	}
}

func TestAssignments(t *testing.T) {
	logLevel := "ERROR"
	if testing.Verbose() {
		logLevel = "INFO"
	}
	opts := []Opts{}
	for _, assignment := range assignments {
		opts = append(opts, Opts{
			Assignment: assignment,
			LogLevel:   logLevel,
		})
	}
	for _, o := range opts {
		NewCharacter(o)
	}
}
