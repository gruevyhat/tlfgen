package tlfgen

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewCharacter(t *testing.T) {
	logLevel := "ERROR"
	if testing.Verbose() {
		logLevel = "INFO"
	}
	opts := []Opts{
		{
			LogLevel: logLevel,
		},
		{
			Profession: "Clerical Worker",
		},
		{
			Profession: "Computer Hacker or Technician",
		},
		{
			Profession: "Antiquarian",
		},
		{
			Name:            "Borkenhekenaken",
			Gender:          "Male",
			PersonalityType: "Bruiser",
			Profession: "Occultist",
			Seed:            "1575d911f49e59ee",
			LogLevel:        logLevel,
		},
	}
	for _, o := range opts {
		c, _ := NewCharacter(o)
		if c.Name == "" {
			t.Error("Missing name.")
		}
		if c.Seed == "" {
			t.Error("Incorrect Seed. No value assigned")
		}
		if !arrayContains(genders, c.Gender) {
			g := strings.Join(genders, ", ")
			t.Errorf("Incorrect gender. Expected '%s' in '%s'.", c.Gender, g)
		}
		fmt.Println(c.ToJSON(true))
	}
}
