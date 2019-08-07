package tlfgen

import (
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
			DataFile: "./assets/Shadow_of_the_Demon_Lord.pdf",
			LogLevel: logLevel,
		},
		{
			LogLevel: logLevel,
		},
		{
			Name:       "Borkenhekenaken",
			Gender:     "Male",
			Ancestry:   "Goblin",
			NovicePath: "Magician",
			ExpertPath: "Wizard",
			Seed:       "1575d911f49e59ee",
			Level:      "3",
			LogLevel:   logLevel,
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
		if c.Level < 0 {
			t.Errorf("Incorrect Level. '%d' is less than zero.", c.Level)
		}
	}
}
