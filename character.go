// Package tlfgen implements a character generator for the SotDL RPG.
package tlfgen

import (
	"encoding/json"
	"fmt"
	"sort"

	log "github.com/sirupsen/logrus"
)

// Declare logger.
var (
	logLevels = map[string]log.Level{
		"INFO":    log.InfoLevel,
		"ERROR":   log.ErrorLevel,
		"WARNING": log.WarnLevel,
	}
	threeD6 = Die{code: 3, sides: 6}
	twoD6   = Die{code: 2, sides: 6}
	oneD3   = Die{code: 1, sides: 3}
)

const skillmax = 75

// Sets the random seed from a hex hash string.
func (c *Character) setCharSeed(charSeed string) *Character {
	var err error
	c.Seed, err = setSeed(charSeed)
	if err != nil {
		log.Error("Failed to set character hash: ", err)
	}
	return c
}

// Declare various character data lists.
var (
	genders          = []string{"Male", "Female"}
	assignments      = ListAssignmentKeys()
	personalityTypes = ListPersonalityTypeKeys()
	professions      = ListProfessionKeys()
)

// Character represents the primary features of the character.
type Character struct {
	Name            string                 `json:"name"`
	Gender          string                 `json:"gender"`
	Age             int                    `json:"age"`
	PersonalityType string                 `json:"personality_type"`
	Assignment      string                 `json:"assignment"`
	Profession      string                 `json:"profession"`
	Wealth          string                 `json:"wealth"`
	Base            BaseCharacteristics    `json:"base"`
	Derived         DerivedCharacteristics `json:"derived"`
	Skills          map[string]int         `json:"skills"`
	Seed            string                 `json:"seed"`
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

func (c *Character) rollBaseCharacteristics(bonus string) *Character {
	c.Base.Strength = threeD6.roll()
	c.Base.Constitution = threeD6.roll()
	c.Base.Power = threeD6.roll()
	c.Base.Dexterity = threeD6.roll()
	c.Base.Charisma = threeD6.roll()
	c.Base.Size = twoD6.roll() + 6
	c.Base.Intelligence = twoD6.roll() + 6
	c.Base.Education = threeD6.roll() + 3
	log.Info("Rolled base characterstics: ", c.Base)
	switch bonus {
	case "smart":
		c.Base.Intelligence += oneD3.roll()
		c.Base.Education += oneD3.roll()
		c.Base.Dexterity += oneD3.roll()
	case "tough":
		c.Base.Strength += oneD3.roll()
		c.Base.Constitution += oneD3.roll()
		c.Base.Size += oneD3.roll()
	case "mystical":
		c.Base.Power += oneD3.roll()
		c.Base.Charisma += oneD3.roll()
	}
	if bonus != "" {
		log.Info("Applied '", bonus, "' bonus: ", c.Base)
	}
	return c
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

func (c *Character) calcDerivedCharacteristics() *Character {
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
	log.Info("Added derived characteristics: ", c.Derived)
	return c
}

func (c *Character) setPersonalityType(opt string) *Character {
	if !arrayContains(personalityTypes, opt) {
		log.Warning("Personality type not found. Randomizing.")
		opt = ""
	}
	if opt != "" {
		c.PersonalityType = opt
	} else {
		c.PersonalityType = randomChoice(personalityTypes)
	}
	log.Info("Set personality type: ", c.PersonalityType)
	return c
}

func (c *Character) setProfession(opt string) *Character {
	if !arrayContains(professions, opt) {
		log.Warning("Profession not found. Randomizing.")
		opt = ""
	}
	if opt != "" {
		c.Profession = opt
	} else {
		c.Profession = randomChoice(professions)
	}
	log.Info("Set profession: ", c.Profession)
	c.Wealth = Professions[c.Profession].wealth
	return c
}

func (c *Character) setAssignment(opt string) *Character {
	if !arrayContains(assignments, opt) {
		log.Warning("Assignment not found. Randomizing.")
		opt = ""
	}
	if opt != "" {
		c.Assignment = opt
	} else {
		c.Assignment = randomChoice(assignments)
	}
	log.Info("Set assignment: ", c.Assignment)
	return c
}

func (c *Character) calcBaseSkills() *Character {
	c.Skills = make(map[string]int)
	for name, skill := range DefaultSkills {
		c.Skills[name] = skill.base
		log.Info("Added base skill: ", name)
	}
	c.Skills["Language: Own"] = c.Base.Intelligence * 5
	AllSkills["Language: Own"] = Skill{base: c.Skills["Language: Own"],
		weight: -c.Skills["Language: Own"]}
	log.Info("Added base skill: Language: Own")
	c.Skills["Dodge"] = c.Base.Dexterity * 2
	log.Info("Added base skill: Dodge")
	return c
}

func (c *Character) calcPersonalitySkills() *Character {
	for _, name := range PersonalityTypes[c.PersonalityType].skills {
		name, newSkill := getSkill(name)
		if _, ok := c.Skills[name]; !ok {
			c.Skills[name] = newSkill.base
		}
		c.Skills[name] += PersonalityTypes[c.PersonalityType].bonus
			log.Info("Improved personality type skill: ", name)
	}
	return c
}

func (c *Character) calcAssignmentSkills() *Character {
	// Primary assignment skills
	for _, name := range Assignments[c.Assignment].skills {
		name, newSkill := getSkill(name)
		if _, ok := c.Skills[name]; !ok {
			c.Skills[name] = newSkill.base
		}
		c.Skills[name] += Assignments[c.Assignment].bonus
			log.Info("Improved assignment skill: ", name)
	}
	// All assignments skills
	for _, name := range Assignments["all"].skills {
		name, newSkill := getSkill(name)
		if _, ok := c.Skills[name]; !ok {
			c.Skills[name] = newSkill.base
		}
		c.Skills[name] += Assignments["all"].bonus
			log.Info("Improved assignment skill: ", name)
	}
	return c
}

func (c *Character) rollProfessionSkills() *Character {
	prof := Professions[c.Profession]
	skills := []string{}
	// Set randomized skills
	if prof.n > 0 {
		skills = prof.skills[0:prof.offset]
		randSkills := prof.skills[prof.offset:]
		for _, s := range sampleWithoutReplacement(randSkills, prof.n) {
			skills = append(skills, s)
		}
	} else {
		skills = prof.skills
	}
	//for i := 0; i < prof.n; i++ {
	//	skills = append(skills, randomChoice(randSkills))
	//}
	// Resolve skill names
	for i, s := range skills {
		skills[i], _ = getSkill(s)
		log.Info("Added professional skill: ", skills[i])
	}
	// Roll skill points
	log.Info("Rolling professional skill points.")
	c.rollSkillPoints(skills, c.Base.Education*20, skillmax)
	return c
}

func (c *Character) rollAdditionalSkillPoints(points int) *Character {
	skills := []string{}
	for skill := range c.Skills {
		skills = append(skills, skill)
	}
	log.Info("Rolling additional skill points.")
	c.rollSkillPoints(skills, points, 95)
	return c
}

func (c *Character) rollSkillPoints(skills []string, points, max int) *Character {
	sort.Strings(skills)
	for points > 0 {
		newSkills := []string{}
		weights := []int{}
		weightTotal := 0
		for _, s := range skills {
			weightTotal += c.Skills[s]
		}
		for _, s := range skills {
			w := 0
			if c.Skills[s] < max {
				w = int(float64((c.Skills[s])+AllSkills[s].weight+1) / float64(weightTotal) * 50)
				weights = append(weights, w)
				newSkills = append(newSkills, s)
			} else {
				log.Warning("Maxed out skill: ", s)
			}
		}
		if arraySum(weights) <= 0 {
			log.Warning("No where else to put skill points!")
			return c
		}
		skill := weightedRandomChoice(newSkills, weights)
		//skill := randomChoice(newSkills)
		c.Skills[skill]++
		points--
		skills = newSkills
	}
	for k, v := range c.Skills {
		if v != AllSkills[k].base {
			log.Info(fmt.Sprintf("Set skill: %s, %d%%.", k, v))
		}
	}
	return c
}

// Randomly sample from gender list.
func (c *Character) setGender(gender string) *Character {
	if gender != "" {
		c.Gender = gender
	} else {
		c.Gender = randomChoice(genders)
	}
	log.Info("Set gender: ", c.Gender)
	return c
}

func (c *Character) setAge(age int) *Character {
	if age != 0 {
		c.Age = age
	} else {
		c.Age = twoD6.roll() + 17
	}
	c.Base.Education += int(float64(c.Age) / 10.0)
	log.Info("Set age and Edu bonus: ", c.Age, ", ", c.Base.Education)
	for a := c.Age - 40; a > 40; a += 10 {
		r := randomInt(5)
		switch r {
		case 0:
			log.Info("Penalized strength.")
			c.Base.Strength--
		case 1:
			log.Info("Penalized constitution.")
			c.Base.Constitution--
		case 2:
			log.Info("Penalized dexterity.")
			c.Base.Dexterity--
		case 3:
			log.Info("Penalized charisma.")
			c.Base.Charisma--
		}
	}
	return c
}

func (c *Character) setName(name string) *Character {
	if name != "" {
		c.Name = name
	} else {
		c.Name = randomName(c.Gender)
	}
	log.Info("Set name: ", c.Name)
	return c
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
	List            bool   `docopt:"list"`
}

// NewCharacter generates a SotDL character given a set of user options.
func NewCharacter(opts Opts) (c Character, err error) {

	log.SetLevel(logLevels[opts.LogLevel])

	c.setCharSeed(opts.Seed).
		setGender(opts.Gender).
		setName(opts.Name).
		rollBaseCharacteristics(opts.AttributeBonus).
		setAge(opts.Age).
		calcDerivedCharacteristics().
		calcBaseSkills().
		setPersonalityType(opts.PersonalityType).
		calcPersonalitySkills().
		setProfession(opts.Profession).
		rollProfessionSkills().
		setAssignment(opts.Assignment).
		calcAssignmentSkills().
		rollAdditionalSkillPoints(opts.SkillPoints)

	// Generate stuff
	//c.setWeapons()
	//c.setEqugipment()

	// Generate fluff
	//c.setDescription(opts.Description)
	//c.setBackground(opts.Background)
	//c.setLanguages(opts.Languages)

	return c, nil
}
