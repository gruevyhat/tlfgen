// Package tlfgen implements a character generator for the SotDL RPG.
package tlfgen

import (
	"encoding/json"
	"fmt"

	logging "github.com/op/go-logging"
)

// Declare logger.
var (
	log       = logging.MustGetLogger("sotdl")
	logLevels = map[string]logging.Level{
		"INFO":    logging.INFO,
		"ERROR":   logging.ERROR,
		"WARNING": logging.WARNING,
	}
)

// Sets the random seed from a hex hash string.
func (c *Character) setCharSeed(charSeed string) {
	var err error
	c.Seed, err = setSeed(charSeed)
	if err != nil {
		log.Error("Failed to set character hash:", err)
	}
}

// Declare various character data lists.
var (
	genders          = []string{"Male", "Female"}
	assignments      = []string{"Archives", "Computational Demonology"}
	personalityTypes = []string{"Bruiser", "Nutter"}
	professions      = []string{"Occultist", "Philosopher"}
	languages        = []string{"English", "French", "Spanish", "German", "Latin", "Ancient Greek", "Arabic", "Enochian"}
)

// Character represents the primary features of the character.
type Character struct {
	Name                   string                 `json:"name"`
	Gender                 string                 `json:"gender"`
	Age                    int                    `json:"age"`
	BaseCharacteristics    BaseCharacteristics    `json:"base_characteristics"`
	DerivedCharacteristics DerivedCharacteristics `json:"derived_characteristics"`
	Skills                 map[string]int         `json:"skills"`
	Seed                   string                 `json:"seed"`
	//Description string     `json:"description"`
	//Weapons     []Weapon   `json:"weapons"`
	//Equipment   []string   `json:"equipment"`
}

// BaseCharacteristics represents character statistics.
type BaseCharacteristics struct {
	Strength     int `json:"strength"`
	Constitution int `json:"constitution"`
	Size         int `json:"size"`
	Intelligence int `json:"intelligence"`
	Power        int `json:"power"`
	Dexterity    int `json:"dexterity"`
	Charisma     int `json:"charisma"`
	Education    int `json:"education"`
}

// DerivedCharacteristics are derived from base characteristics.
type DerivedCharacteristics struct {
	DamageBonus     int `json:"damage_bonus"`
	HitPoints       int `json:"hit_points"`
	MajorWoundLevel int `json:"major_wound_level"`
	ExperienceBonus int `json:"experience_bonus"`
	Move            int `json:"move"`
	Sanity          int `json:"sanity"`
	Effort          int `json:"effort"`
	Endurance       int `json:"endurance"`
	DamageMod       int `json:"damage_mod"`
	Idea            int `json:"idea"`
	Luck            int `json:"luck"`
	Agility         int `json:"agility"`
	Know            int `json:"know"`
}

// Weapon represents properties of a given weapon.
type Weapon struct {
	// TODO: Implement weapons.
	Name   string `json:"name"`
	Range  string `json:"range"`
	Damage Die    `json:"damage"`
}

func (c *Character) rollBaseCharacteristics() {
	c.Strength = Die(3, 6).roll()
	c.Constitution = Die(3, 6).roll()
	c.Size = Die(2, 6).roll() + 6
	c.Intelligence = Die(2, 6).roll() + 6
	c.Power = Die(3, 6).roll()
	c.Dexterity = Die(3, 6).roll()
	c.Charisma = Die(3, 6).roll()
	c.Education = Die(3, 6).roll() + 3
}

func getDamageBonus(n int) (code string) {
	switch n {
	case n <= 12:
		return "-1D6"
	case n <= 16:
		return "-1D4"
	case n <= 24:
		return "None"
	case n <= 32:
		return "+1D4"
	case n <= 40:
		return "+1D6"
	}
	return "+2D6"
}

func (c *Character) calcDerivedCharacteristics() {
	c.DamageBonus = getDamageBonus(c.Strength + c.Size)
	c.HitPoints = (c.Constitution + c.Size) / 2
	c.MajorWoundLevel = c.HitPoints / 2
	c.ExperienceBonus = c.Intelligence / 2
	c.Move = 10
	c.Sanity = c.Power
	c.Effort = c.Strenth * 5
	c.Endurance = c.Consitution * 5
	c.Idea = c.Intelligence * 5
	c.Luck = c.Power * 5
	c.Agility = c.Dexterity * 5
	c.Know = c.Education * 5
}

func (c *Character) getPersonalityType(opt string) {
	if opt != nil {
		c.personalityType = opt
	} else {
		c.personalityType = randomChoice(personalityTypes)
	}
}

func (c *Character) getProfession(opt string) {
	if opt != nil {
		c.Profession = opt
	} else {
		c.Profession = randomChoice(professions)
	}
}

func (c *Character) getAssignment(opt string) {
	if opt != nil {
		c.Assignment = opt
	} else {
		c.Assignment = randomChoice(assignments)
	}
}

func (c *Character) rollSkills() {}

// Randomly sample from gender list.
func (c *Character) setGender(gender string) {
	if gender != "" {
		c.Gender = gender
	} else {
		c.Gender = randomChoice(genders)
	}
}

// TODO: Additional character data functions.
func (c *Character) setWeapons()                       {}
func (c *Character) setEquipment()                     {}
func (c *Character) setDescription(description string) {}
func (c *Character) setBackground(background string)   {}

// Print writes tab-delimited character details to STDOUT.
func (c Character) Print() {
	fmt.Println("Name\t" + c.Name)
	fmt.Println("Gender\t" + c.Gender)
	fmt.Println("Level\t", c.Level)
	fmt.Println("Character Seed\t", c.Seed)
	fmt.Println()
}

// ToJSON writes JSON character details to STDOUT.
func (c Character) ToJSON(pretty bool) string {
	var j []byte
	if pretty {
		j, _ = json.MarshalIndent(c, "  ", "  ")
	} else {
		j, _ = json.Marshal(c)
	}
	fmt.Println(string(j))
	//err := ioutil.WriteFile(fn, j, 0644)
	//if err != nil {
	//	panic(err)
	//}
	return string(j)
}

// Opts contains user input optionsr; used in CLI implementations.
type Opts struct {
	Name     string `docopt:"--name"`
	Age      string
	Gender   string `docopt:"--gender"`
	LogLevel string `docopt:"--log-level"`
	Seed     string `docopt:"--seed"`
}

// NewCharacter generates a SotDL character given a set of user options.
func NewCharacter(opts Opts) (c Character, err error) {

	logging.SetLevel(logLevels[opts.LogLevel], "")

	// Load the character db if empty.
	if len(db.Paths) == 0 {
		log.Info("Loading Character DB.")
		db, err = NewCharDB(opts.DataFile, false)
		if err != nil {
			return c, err
		}
	}

	// Initialize character and set random seed from hash
	c.setCharSeed(opts.Seed)

	// Generate character
	log.Info("Generating characteristics.")
	c.setGender(opts.Gender)
	c.setName(opts.Name)
	c.rollBaseCharacteristics()
	c.calcDerivedCharacteristics()
	c.getPersonalityType(opts.PersonalityType)
	c.getProfession(opts.Profession)
	c.getAssignment(opts.Assignment)

	// Generate stuff
	//c.setWeapons()
	//c.setEquipment()

	// Generate fluff
	//c.setDescription(opts.Description)
	//c.setBackground(opts.Background)
	//c.setProfessions(opts.Professions)
	//c.setLanguages(opts.Languages)

	return c, nil
}
