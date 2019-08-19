// Package tlfgen implements a character generator for the SotDL RPG.
package tlfgen

import (
	"encoding/json"
	"fmt"
	"sort"

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
	threeD6 = Die{code: 3, sides: 6}
	twoD6   = Die{code: 2, sides: 6}
	oneD3   = Die{code: 1, sides: 3}
)

const skillmax = 75

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
	assignments      = assignmentKeys()
	personalityTypes = personalityTypeKeys()
	professions      = professionKeys()
)

// Character represents the primary features of the character.
type Character struct {
	Name            string                 `json:"name"`
	Gender          string                 `json:"gender"`
	Age             int                    `json:"age"`
	PersonalityType string                 `json:"personality_type"`
	Assignment      string                 `json:"assignment"`
	Profession      string                 `json:"profession"`
	Skills          map[string]int         `json:"skills"`
	Seed            string                 `json:"seed"`
	Base            BaseCharacteristics    `json:"base"`
	Derived         DerivedCharacteristics `json:"derived"`
	Wealth          string                 `json:"wealth"`
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
	DamageBonus     string `json:"damage_bonus"`
	HitPoints       int    `json:"hit_points"`
	MajorWoundLevel int    `json:"major_wound_level"`
	ExperienceBonus int    `json:"experience_bonus"`
	Move            int    `json:"move"`
	Sanity          int    `json:"sanity"`
	Effort          int    `json:"effort"`
	Endurance       int    `json:"endurance"`
	DamageMod       int    `json:"damage_mod"`
	Idea            int    `json:"idea"`
	Luck            int    `json:"luck"`
	Agility         int    `json:"agility"`
	Know            int    `json:"know"`
}

// Weapon represents properties of a given weapon.
type Weapon struct {
	// TODO: Implement weapons.
	Name   string `json:"name"`
	Range  string `json:"range"`
	Damage Die    `json:"damage"`
}

func (c *Character) rollBaseCharacteristics(bonus string) {
	log.Info("Rolling base characteristics.")
	c.Base.Strength = threeD6.roll()
	c.Base.Constitution = threeD6.roll()
	c.Base.Power = threeD6.roll()
	c.Base.Dexterity = threeD6.roll()
	c.Base.Charisma = threeD6.roll()
	c.Base.Size = twoD6.roll() + 6
	c.Base.Intelligence = twoD6.roll() + 6
	c.Base.Education = threeD6.roll() + 3
	switch bonus {
	case "smart":
		log.Info("Increasing bonus characteristics:", bonus)
		c.Base.Intelligence += oneD3.roll()
		c.Base.Education += oneD3.roll()
		c.Base.Dexterity += oneD3.roll()
	case "tough":
		log.Info("Increasing bonus characteristics:", bonus)
		c.Base.Strength += oneD3.roll()
		c.Base.Constitution += oneD3.roll()
		c.Base.Size += oneD3.roll()
	case "mystical":
		log.Info("Increasing bonus characteristics:", bonus)
		c.Base.Power += oneD3.roll()
		c.Base.Charisma += oneD3.roll()
	}
}

func getDamageBonus(n int) (code string) {
	switch {
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
	default:
		return "+2D6"
	}
}

func (c *Character) calcDerivedCharacteristics() {
	log.Info("Adding derived characteristics.")
	c.Derived.DamageBonus = getDamageBonus(c.Base.Strength + c.Base.Size)
	c.Derived.HitPoints = (c.Base.Constitution + c.Base.Size) / 2
	c.Derived.MajorWoundLevel = c.Derived.HitPoints / 2
	c.Derived.ExperienceBonus = c.Base.Intelligence / 2
	c.Derived.Move = 10
	c.Derived.Sanity = c.Base.Power * 5
	c.Derived.Effort = c.Base.Strength * 5
	c.Derived.Endurance = c.Base.Constitution * 5
	c.Derived.Idea = c.Base.Intelligence * 5
	c.Derived.Luck = c.Base.Power * 5
	c.Derived.Agility = c.Base.Dexterity * 5
	c.Derived.Know = c.Base.Education * 5
}

func (c *Character) getPersonalityType(opt string) *Character {
	log.Info("Setting personality type.")
	if !arrayContains(personalityTypes, opt) {
		log.Error("Personality type not found. Randomizing.")
		opt = ""
	}
	if opt != "" {
		c.PersonalityType = opt
	} else {
		c.PersonalityType = randomChoice(personalityTypes)
	}
	return c
}

func (c *Character) getProfession(opt string) *Character {
	log.Info("Setting profession.")
	if !arrayContains(professions, opt) {
		log.Error("Profession not found. Randomizing.")
		opt = ""
	}
	if opt != "" {
		c.Profession = opt
	} else {
		c.Profession = randomChoice(professions)
	}
	c.Wealth = Professions[c.Profession].wealth
	return c
}

func (c *Character) getAssignment(opt string) *Character {
	log.Info("Setting assignment.")
	if !arrayContains(assignments, opt) {
		log.Error("Assignment not found. Randomizing.")
		opt = ""
	}
	if opt != "" {
		c.Assignment = opt
	} else {
		c.Assignment = randomChoice(assignments)
	}
	return c
}

func (c *Character) calcBaseSkills() {
	c.Skills = make(map[string]int)
	for name, skill := range DefaultSkills {
		log.Info("Adding base skill:", name)
		c.Skills[name] = skill.base
	}
	log.Info("Adding base skill: Language: Own")
	c.Skills["Language: Own"] = c.Base.Intelligence * 5
	log.Info("Adding base skill: Dodge")
	c.Skills["Dodge"] = c.Base.Dexterity * 2
}

func (c *Character) calcPersonalitySkills() {
	for _, name := range PersonalityTypes[c.PersonalityType].skills {
		name, newSkill := getSkill(name)
		if _, ok := c.Skills[name]; !ok {
			log.Info("Adding personality type skill:", name)
			c.Skills[name] = newSkill.base
		}
		c.Skills[name] += PersonalityTypes[c.PersonalityType].bonus
	}
}

func (c *Character) calcAssignmentSkills() {
	// Primary assignment skills
	for _, name := range Assignments[c.Assignment].skills {
		name, newSkill := getSkill(name)
		if _, ok := c.Skills[name]; !ok {
			log.Info("Adding assignment skill:", name)
			c.Skills[name] = newSkill.base
		}
		c.Skills[name] += Assignments[c.Assignment].bonus
	}
	// All assignments skills
	for _, name := range Assignments["all"].skills {
		name, newSkill := getSkill(name)
		if _, ok := c.Skills[name]; !ok {
			log.Info("Adding assignment skill:", name)
			c.Skills[name] = newSkill.base
		}
		c.Skills[name] += Assignments["all"].bonus
	}
}

func (c *Character) rollProfessionSkills() {
	prof := Professions[c.Profession]
	skills := []string{}
	if prof.n > 0 {
		skills = prof.skills[0:prof.offset]
	} else {
		skills = prof.skills
	}
	// Set randomized skills
	randSkills := prof.skills[prof.offset:]
	sort.Strings(randSkills)
	for i := 0; i < prof.n; i++ {
		skills = append(skills, randomChoice(randSkills))
	}
	// Resolve skill names
	for i, s := range skills {
		skills[i], _ = getSkill(s)
		log.Info("Adding professional skill:", skills[i])
	}
	// Roll skill points
	//fmt.Println(skills)
	log.Info("Rolling professional skill points.")
	c.rollSkillPoints(skills, c.Base.Education*20, skillmax)
}

func (c *Character) rollAdditionalSkillPoints(points int) {
	skills := []string{}
	for skill := range c.Skills {
		skills = append(skills, skill)
	}
	log.Info("Rolling additional skill points.")
	c.rollSkillPoints(skills, points, 95)
}

func (c *Character) rollSkillPoints(skills []string, points, max int) {
	sort.Strings(skills)
	for points := points; points > 0; points-- {
		weights := []int{}
		weightTotal := 0
		for _, s := range skills {
			weightTotal += c.Skills[s]
		}
		for _, s := range skills {
			w := int(float64((c.Skills[s])+AllSkills[s].weight+10) / float64(weightTotal) * 100)
			if c.Skills[s] >= max {
				w = 0
			}
			weights = append(weights, w)
		}
		//fmt.Println(weights)
		skill := weightedRandomChoice(skills, weights)
		c.Skills[skill]++
	}
	for k, v := range c.Skills {
		if v != AllSkills[k].base {
			log.Info(fmt.Sprintf("Set skill: %s, %d%%.", k, v))
		}
	}
}

// Randomly sample from gender list.
func (c *Character) setGender(gender string) {
	log.Info("Setting gender.")
	if gender != "" {
		c.Gender = gender
	} else {
		c.Gender = randomChoice(genders)
	}
}

func (c *Character) setAge(age int) {
	log.Info("Setting age and Edu bonus.")
	if age != 0 {
		c.Age = age
	} else {
		c.Age = twoD6.roll() + 17
	}
	c.Base.Education += int(float64(c.Age) / 10.0)
	for a := c.Age - 40; a > 40; a += 10 {
		log.Info("Calculating age penalty.")
		r := randomInt(3)
		switch r {
		case 0:
			c.Base.Strength--
		case 1:
			c.Base.Constitution--
		case 2:
			c.Base.Dexterity--
		case 3:
			c.Base.Charisma--
		}
	}
}

func (c *Character) setName(name string) {
	log.Info("Setting Name.")
	if name != "" {
		c.Name = name
	} else {
		c.Name = randomName(c.Gender)
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
	//fmt.Println(string(j))
	//err := ioutil.WriteFile(fn, j, 0644)
	//if err != nil {
	//	panic(err)
	//}
	return string(j)
}

// Opts contains user input optionsr; used in CLI implementations.
type Opts struct {
	Name            string `docopt:"--name"`
	Age             int    `docopt:"--age"`
	Gender          string `docopt:"--gender"`
	PersonalityType string `docopt:"--personality"`
	Assignment      string `docopt:"--assignment"`
	Profession      string `docopt:"--profession"`
	SkillPoints     int    `docopt:"--skill-points"`
	AttributeBonus  string `docopt:"--attr-bonus"`
	LogLevel        string `docopt:"--log-level"`
	Seed            string `docopt:"--seed"`
}

// NewCharacter generates a SotDL character given a set of user options.
func NewCharacter(opts Opts) (c Character, err error) {

	logging.SetLevel(logLevels[opts.LogLevel], "")

	// Initialize character and set random seed from hash
	c.setCharSeed(opts.Seed)

	// Generate character
	log.Info("Generating characteristics.")
	c.setGender(opts.Gender)
	c.setName(opts.Name)

	c.rollBaseCharacteristics(opts.AttributeBonus)
	c.setAge(opts.Age)
	c.calcDerivedCharacteristics()
	c.calcBaseSkills()

	c.getProfession(opts.Profession).rollProfessionSkills()
	c.getPersonalityType(opts.PersonalityType).calcPersonalitySkills()
	c.getAssignment(opts.Assignment).calcAssignmentSkills()

	c.rollAdditionalSkillPoints(opts.SkillPoints)

	// Generate stuff
	//c.setWeapons()
	//c.setEqugipment()

	// Generate fluff
	//c.setDescription(opts.Description)
	//c.setBackground(opts.Background)
	//c.setLanguages(opts.Languages)

	return c, nil
}
