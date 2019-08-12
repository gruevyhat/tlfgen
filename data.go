package tlfgen

// Skill is a single skill.
type Skill struct {
	Value  int `json:"value"`
	weight int
}

func (s Skill) add(n int) {
	s.Value += n
}

// PersonalityType is a list of skill names and a starting bonus.
type PersonalityType struct {
	bonus   int
	desc    string
	skills  []string
	special string
}

// Assignment is a list of skill names and a starting bonus.
type Assignment struct {
	bonus  int
	skills []string
}

// Profession is a list of skills and an [offset, num] slice of
// optional skills from which num should be sampled.
type Profession struct {
	skills    []string
	optOffset int
	optNum    int
}

// Professions is a map of professions.
var Professions = map[string]Profession{
	"Occultist": Profession{
		skills: []string{
			"Fast Talk",
			"Insight",
			"Knowledge: Anthropology",
			"Knowledge: History",
			"Knowledge: Occult",
			"LANGUAGE",
			"Language: Own",
			"Research",
			"Computer Use: *",
			"CRAFT",
			"Knowledge: Archaeology",
			"Medicine",
			"Science: any",
		},
		optOffset: 8,
		optNum:    2,
	},
	"Philosopher": Profession{
		skills: []string{
			"Insight",
			"Knowledge: History",
			"Knowledge: Philosophy",
			"LANGUAGE",
			"Language: Own",
			"Persuade",
			"Research",
			"Teach",
			"KNOWLEDGE",
			"SCIENCE",
		},
		optOffset: 8,
		optNum:    2,
	},
}

// PersonalityTypes is a map of PTypes.
var PersonalityTypes = map[string]PersonalityType{
	"Bruiser": PersonalityType{
		bonus: 20,
		desc:  "Your character believes that solving problems is best handled through quick application of physical force.",
		skills: []string{
			"Brawl",
			"Climb",
			"COMBAT",
			"COMBAT",
			"Dodge",
			"Grapple",
			"Insight",
			"Jump",
			"Ride",
			"Sense",
			"Stealth",
			"Swim",
			"Throw",
		},
		special: "",
	},
	"Nutter": PersonalityType{
		bonus: 20,
		desc:  "Your character can safely be categorised as insane, though they are functional and able to work within the organisation of the Laundry. Rational thought and problem-solving methods are neglected: insane leaps of logic are the primary means of attaining goals.",
		skills: []string{
			"Command",
			"Fast Talk",
			"Hide",
			"Insight",
			"KNOWLEDGE",
			"KNOWLEDGE",
			"Research",
			"SCIENCE",
			"SCIENCE",
			"Sense",
			"Spot",
			"Stealth",
			"Strategy",
		},
		special: "Reduce starting SAN by 20 points and assign an appropriate mental disorder. Depending on the disorder, your character may be in therapy or on medication to deal with the symptoms.",
	},
}

// Assignments is a map of assignments.
var Assignments = map[string]Assignment{
	"Archives": Assignment{
		bonus: 10,
		skills: []string{
			"Bureaucracy",
			"KNOWLEDGE",
			"KNOWLEDGE",
			"Navigate",
			"Research",
			"Stealth",
		},
	},
	"Computational Demonology": Assignment{
		bonus: 10,
		skills: []string{
			"COMPUTER USE",
			"Computer Use: Magic",
			"Science: Mathematics",
			"Science: Thaumaturgy",
			"Sorcery",
		},
	},
	"all": Assignment{
		bonus: 5,
		skills: []string{
			"Bureaucracy",
			"Computer Use",
			"Fine Manipulation",
			"FIREARM",
			"Knowledge: Accounting",
			"Knowledge: Espionage",
			"Knowledge: Law",
			"Knowledge: Occult",
			"Knowledge: Politics",
			"Spot",
		},
	},
}

func randomSkillChoice(skills map[string]Skill) (string, Skill) {
	names := make([]string, len(skills))
	for k := range skills {
		names = append(names, k)
	}
	name := randomChoice(names)
	return name, skills[name]
}

func randomWeightedSkillChoice(skills map[string]Skill) (string, Skill) {
	names := make([]string, len(skills))
	weights := make([]int, len(skills))
	for name, skill := range skills {
		names = append(names, name)
		weights = append(weights, skill.weight)
	}
	name := weightedRandomChoice(names, weights)
	return name, skills[name]
}

func getSkill(name string) (string, Skill) {
	var newSkill Skill
	switch name {
	case "COMBAT":
		name, newSkill = randomWeightedSkillChoice(CombatSkills)
	case "FIREARM":
		name, newSkill = randomWeightedSkillChoice(FirearmSkills)
	case "KNOWLEDGE":
		name, newSkill = randomWeightedSkillChoice(KnowledgeSkills)
	case "CRAFT":
		name, newSkill = randomWeightedSkillChoice(CraftSkills)
	case "LANGUAGE":
		name, newSkill = randomWeightedSkillChoice(LanguageSkills)
	case "SCIENCE":
		name, newSkill = randomWeightedSkillChoice(ScienceSkills)
	case "PERFORM":
		name, newSkill = randomWeightedSkillChoice(PerformSkills)
	case "ART":
		name, newSkill = randomWeightedSkillChoice(ArtSkills)
	case "DRIVE":
		name, newSkill = randomWeightedSkillChoice(DriveSkills)
	case "":
		fallthrough
	default:
		newSkill = GeneralSkills[name]
	}
	return name, newSkill
}

func joinMaps(maps ...map[string]Skill) map[string]Skill {
	newMap := make(map[string]Skill)
	for _, m := range maps {
		for name, skill := range m {
			newMap[name] = skill
		}
	}
	return newMap
}

// ArtillerySkills is a map of skills.
var ArtillerySkills = map[string]Skill{
	"Artillery: Cannon":           Skill{Value: 0, weight: 1},
	"Artillery: Other":            Skill{Value: 0, weight: 1},
	"Artillery: Rocket Launcher":  Skill{Value: 0, weight: 1},
	"Artillery: Turret":           Skill{Value: 0, weight: 1},
	"Artillery: Vehicular Weapon": Skill{Value: 0, weight: 1},
}

// FirearmSkills is a map of skills.
var FirearmSkills = map[string]Skill{
	"Firearm: Assault Rifle":  Skill{Value: 15, weight: 5},
	"Firearm: Esoteric":       Skill{Value: 0, weight: 5},
	"Firearm: Exotic":         Skill{Value: 5, weight: 1},
	"Firearm: Pistol":         Skill{Value: 20, weight: 50},
	"Firearm: Rifle":          Skill{Value: 25, weight: 20},
	"Firearm: Shotgun":        Skill{Value: 30, weight: 20},
	"Firearm: Submachine Gun": Skill{Value: 15, weight: 1},
}

// HeavyWeaponSkills is a map of skills.
var HeavyWeaponSkills = map[string]Skill{
	"Heavy Weapon: Bazooka":           Skill{Value: 0, weight: 1},
	"Heavy Weapon: Grenade Launcher":  Skill{Value: 0, weight: 1},
	"Heavy Weapon: Heavy Machine Gun": Skill{Value: 0, weight: 1},
	"Heavy Weapon: Minigun":           Skill{Value: 0, weight: 1},
	"Heavy Weapon: Other":             Skill{Value: 0, weight: 1},
	"Heavy Weapon: Rocket Launcher":   Skill{Value: 0, weight: 1},
}

// MeleeWeaponSkills is a map of skills.
var MeleeWeaponSkills = map[string]Skill{
	"Melee Weapon: Axe":     Skill{Value: 5, weight: 5},
	"Melee Weapon: Club":    Skill{Value: 5, weight: 10},
	"Melee Weapon: Garrote": Skill{Value: 5, weight: 1},
	"Melee Weapon: Knife":   Skill{Value: 5, weight: 20},
	"Melee Weapon: Other":   Skill{Value: 5, weight: 1},
	"Melee Weapon: Spear":   Skill{Value: 5, weight: 1},
	"Melee Weapon: Staff":   Skill{Value: 5, weight: 5},
	"Melee Weapon: Sword":   Skill{Value: 5, weight: 5},
	"Melee Weapon: Whip":    Skill{Value: 5, weight: 1},
}

// MissileWeaponSkills is a map of skills.
var MissileWeaponSkills = map[string]Skill{
	"Missile Weapon: Blowgun":        Skill{Value: 5, weight: 1},
	"Missile Weapon: Boomerang":      Skill{Value: 5, weight: 1},
	"Missile Weapon: Bow":            Skill{Value: 5, weight: 20},
	"Missile Weapon: Crossbow":       Skill{Value: 5, weight: 1},
	"Missile Weapon: Dart":           Skill{Value: 5, weight: 1},
	"Missile Weapon: Javelin":        Skill{Value: 5, weight: 1},
	"Missile Weapon: Other":          Skill{Value: 5, weight: 1},
	"Missile Weapon: Shuriken":       Skill{Value: 5, weight: 1},
	"Missile Weapon: Sling":          Skill{Value: 5, weight: 1},
	"Missile Weapon: Spear":          Skill{Value: 5, weight: 1},
	"Missile Weapon: Throwing Axe":   Skill{Value: 5, weight: 1},
	"Missile Weapon: Throwing Knife": Skill{Value: 5, weight: 10},
}

// CombatSkills list of all the possible combat skills.
var CombatSkills = joinMaps(
	ArtillerySkills,
	FirearmSkills,
	HeavyWeaponSkills,
	MeleeWeaponSkills,
	MissileWeaponSkills,
)

// DefaultSkills is a map of default skills.
var DefaultSkills = map[string]Skill{
	"Appraise":          Skill{Value: 15, weight: 1},
	"Art":               Skill{Value: 5, weight: 1},
	"Athletics":         Skill{Value: 10, weight: 1},
	"Bargain":           Skill{Value: 5, weight: 1},
	"Brawl":             Skill{Value: 25, weight: 1},
	"Bureaucracy":       Skill{Value: 5, weight: 1},
	"Climb":             Skill{Value: 40, weight: 1},
	"Command":           Skill{Value: 5, weight: 1},
	"Computer Use":      Skill{Value: 5, weight: 1},
	"Craft":             Skill{Value: 5, weight: 1},
	"Demolition":        Skill{Value: 1, weight: 1},
	"Disguise":          Skill{Value: 5, weight: 1},
	"Dodge":             Skill{Value: -1, weight: 10},
	"Drive":             Skill{Value: 20, weight: 1},
	"Etiquette":         Skill{Value: 5, weight: 1},
	"Fast Talk":         Skill{Value: 5, weight: 1},
	"Fine Manipulation": Skill{Value: 5, weight: 1},
	"First Aid":         Skill{Value: 30, weight: 1},
	"Gaming":            Skill{Value: 10, weight: 1},
	"Geology":           Skill{Value: 0, weight: 1},
	"Grapple":           Skill{Value: 25, weight: 1},
	"Heavy Machine":     Skill{Value: 5, weight: 1},
	"Hide":              Skill{Value: 10, weight: 1},
	"Insight":           Skill{Value: 5, weight: 1},
	"Jump":              Skill{Value: 25, weight: 1},
	"Language: Own":     Skill{Value: -1, weight: 1},
	"Listen":            Skill{Value: 25, weight: 1},
	"Medicine":          Skill{Value: 5, weight: 1},
	"Navigate":          Skill{Value: 10, weight: 1},
	"Perform":           Skill{Value: 5, weight: 1},
	"Persuade":          Skill{Value: 15, weight: 1},
	"Research":          Skill{Value: 25, weight: 10},
	"Ride":              Skill{Value: 5, weight: 1},
	"Science":           Skill{Value: 1, weight: 1},
	"Sense":             Skill{Value: 10, weight: 1},
	"Sleight of Hand":   Skill{Value: 5, weight: 1},
	"Spot":              Skill{Value: 25, weight: 10},
	"Status":            Skill{Value: 15, weight: 1},
	"Stealth":           Skill{Value: 10, weight: 1},
	"Strategy":          Skill{Value: 5, weight: 1},
	"Swim":              Skill{Value: 25, weight: 1},
	"Teach":             Skill{Value: 10, weight: 1},
	"Technology Use":    Skill{Value: 5, weight: 1},
	"Throw":             Skill{Value: 25, weight: 1},
	"Track":             Skill{Value: 10, weight: 1},
}

// ArtSkills is a map of skills.
var ArtSkills = map[string]Skill{
	"Art: Calligraphy": Skill{Value: 5, weight: 1},
	"Art: Drawing":     Skill{Value: 5, weight: 1},
	"Art: Other":       Skill{Value: 5, weight: 1},
	"Art: Painting":    Skill{Value: 5, weight: 1},
	"Art: Photography": Skill{Value: 5, weight: 1},
	"Art: Sculpture":   Skill{Value: 5, weight: 1},
	"Art: Writing":     Skill{Value: 5, weight: 1},
}

// AthleticsSkills is a map of skills.
var AthleticsSkills = map[string]Skill{
	"Athletics: Acrobatics":                     Skill{Value: 10, weight: 1},
	"Athletics: American and Canadian Football": Skill{Value: 10, weight: 1},
	"Athletics: Baseball":                       Skill{Value: 10, weight: 1},
	"Athletics: Basketball":                     Skill{Value: 10, weight: 1},
	"Athletics: Bowling":                        Skill{Value: 10, weight: 1},
	"Athletics: Cricket":                        Skill{Value: 10, weight: 1},
	"Athletics: Cycling":                        Skill{Value: 10, weight: 1},
	"Athletics: Golf":                           Skill{Value: 10, weight: 1},
	"Athletics: Hockey":                         Skill{Value: 10, weight: 1},
	"Athletics: Rugby":                          Skill{Value: 10, weight: 1},
	"Athletics: Skating":                        Skill{Value: 10, weight: 1},
	"Athletics: Skiing":                         Skill{Value: 10, weight: 1},
	"Athletics: Soccer":                         Skill{Value: 10, weight: 1},
	"Athletics: Tennis":                         Skill{Value: 10, weight: 1},
	"Athletics: Track & Field":                  Skill{Value: 10, weight: 1},
}

// ComputerUseSkills is a map of skills.
var ComputerUseSkills = map[string]Skill{
	"Computer Use: Art":         Skill{Value: 5, weight: 1},
	"Computer Use: Design":      Skill{Value: 5, weight: 1},
	"Computer Use: Gaming":      Skill{Value: 5, weight: 1},
	"Computer Use: Hacking":     Skill{Value: 5, weight: 1},
	"Computer Use: Magic":       Skill{Value: 5, weight: 1},
	"Computer Use: Maintenance": Skill{Value: 5, weight: 1},
	"Computer Use: Other":       Skill{Value: 5, weight: 1},
	"Computer Use: Programming": Skill{Value: 5, weight: 1},
	"Computer Use: Repair":      Skill{Value: 5, weight: 1},
}

// CraftSkills is a map of skills.
var CraftSkills = map[string]Skill{
	"Craft: Carpentry":      Skill{Value: 5, weight: 1},
	"Craft: Cooking":        Skill{Value: 5, weight: 1},
	"Craft: Leatherworking": Skill{Value: 5, weight: 1},
	"Craft: Pottery":        Skill{Value: 5, weight: 1},
	"Craft: Sewing":         Skill{Value: 5, weight: 1},
	"Craft: Woodworking":    Skill{Value: 5, weight: 1},
}

// DriveSkills is a map of skills.
var DriveSkills = map[string]Skill{
	"Drive: Automobile":       Skill{Value: 20, weight: 1},
	"Drive: Industrial Mover": Skill{Value: 20, weight: 1},
	"Drive: Motorcycle":       Skill{Value: 20, weight: 1},
	"Drive: Other":            Skill{Value: 20, weight: 1},
	"Drive: Tank":             Skill{Value: 20, weight: 1},
}

// HeavyMachineSkills is a map of skills.
var HeavyMachineSkills = map[string]Skill{
	"Heavy Machine: Boiler":    Skill{Value: 5, weight: 1},
	"Heavy Machine: Bulldozer": Skill{Value: 5, weight: 1},
	"Heavy Machine: Crane":     Skill{Value: 5, weight: 1},
	"Heavy Machine: Engine":    Skill{Value: 5, weight: 1},
	"Heavy Machine: Other":     Skill{Value: 5, weight: 1},
	"Heavy Machine: Turbine":   Skill{Value: 5, weight: 1},
	"Heavy Machine: Wrecker":   Skill{Value: 5, weight: 1},
}

// KnowledgeSkills is a map of skills.
var KnowledgeSkills = map[string]Skill{
	"Knowledge: Accounting":      Skill{Value: 10, weight: 1},
	"Knowledge: Anthropology":    Skill{Value: 1, weight: 1},
	"Knowledge: Archaeology":     Skill{Value: 1, weight: 1},
	"Knowledge: Art History":     Skill{Value: 1, weight: 1},
	"Knowledge: Business":        Skill{Value: 1, weight: 1},
	"Knowledge: Espionage":       Skill{Value: 0, weight: 1},
	"Knowledge: Folklore":        Skill{Value: 5, weight: 1},
	"Knowledge: Group":           Skill{Value: 0, weight: 1},
	"Knowledge: History":         Skill{Value: 20, weight: 1},
	"Knowledge: Law":             Skill{Value: 5, weight: 1},
	"Knowledge: Linguistics":     Skill{Value: 0, weight: 1},
	"Knowledge: Literature":      Skill{Value: 5, weight: 1},
	"Knowledge: Natural History": Skill{Value: 10, weight: 1},
	"Knowledge: Occult":          Skill{Value: 5, weight: 1},
	"Knowledge: Philosophy":      Skill{Value: 1, weight: 1},
	"Knowledge: Politics":        Skill{Value: 5, weight: 1},
	"Knowledge: Region":          Skill{Value: 0, weight: 1},
	"Knowledge: Religion":        Skill{Value: 5, weight: 1},
	"Knowledge: Streetwise":      Skill{Value: 5, weight: 1},
}

// MedicineSkills is a map of skills.
var MedicineSkills = map[string]Skill{
	"Medicine: Dermatology":       Skill{Value: 5, weight: 1},
	"Medicine: Family Medicine":   Skill{Value: 5, weight: 1},
	"Medicine: Immunology":        Skill{Value: 5, weight: 1},
	"Medicine: Internal Medicine": Skill{Value: 5, weight: 1},
	"Medicine: Neurology":         Skill{Value: 5, weight: 1},
	"Medicine: Nuclear Medicine":  Skill{Value: 5, weight: 1},
	"Medicine: Oncology":          Skill{Value: 5, weight: 1},
	"Medicine: Other":             Skill{Value: 5, weight: 1},
	"Medicine: Pathology":         Skill{Value: 5, weight: 1},
	"Medicine: Pediatrics":        Skill{Value: 5, weight: 1},
	"Medicine: Radiology":         Skill{Value: 5, weight: 1},
	"Medicine: Surgery":           Skill{Value: 5, weight: 1},
}

// PerformSkills is a map of skills.
var PerformSkills = map[string]Skill{
	"Perform: Act":             Skill{Value: 5, weight: 1},
	"Perform: Dance":           Skill{Value: 5, weight: 1},
	"Perform: Juggle":          Skill{Value: 5, weight: 1},
	"Perform: Other":           Skill{Value: 5, weight: 1},
	"Perform: Play Instrument": Skill{Value: 5, weight: 1},
	"Perform: Recite":          Skill{Value: 5, weight: 1},
	"Perform: Sing":            Skill{Value: 5, weight: 1},
}

// PilotSkills is a map of skills.
var PilotSkills = map[string]Skill{
	"Pilot: Battleship":      Skill{Value: 0, weight: 1},
	"Pilot: Helicopter":      Skill{Value: 0, weight: 1},
	"Pilot: Hot Air Balloon": Skill{Value: 0, weight: 1},
	"Pilot: Hovercraft":      Skill{Value: 0, weight: 1},
	"Pilot: Jet Airliner":    Skill{Value: 0, weight: 1},
	"Pilot: Jet Boat":        Skill{Value: 0, weight: 1},
	"Pilot: Jet Fighter":     Skill{Value: 0, weight: 1},
	"Pilot: Ocean Liner":     Skill{Value: 0, weight: 1},
	"Pilot: Propeller Plane": Skill{Value: 0, weight: 1},
}

// LanguageSkills is a map of skills.
var LanguageSkills = map[string]Skill{
	"Language: (various non-native)": Skill{Value: 0, weight: 1},
	"Language: Own":                  Skill{Value: -1, weight: 1},
}

// ScienceSkills is a map of skills.
var ScienceSkills = map[string]Skill{
	"Science: Astronomy":   Skill{Value: 1, weight: 1},
	"Science: Biology":     Skill{Value: 1, weight: 1},
	"Science: Botany":      Skill{Value: 1, weight: 1},
	"Science: Chemistry":   Skill{Value: 1, weight: 1},
	"Science: Cyptography": Skill{Value: 1, weight: 1},
	"Science: Forensics":   Skill{Value: 1, weight: 1},
	"Science: Genetics":    Skill{Value: 1, weight: 1},
	"Science: Geology":     Skill{Value: 1, weight: 1},
	"Science: Mathematics": Skill{Value: 10, weight: 1},
	"Science: Meteorology": Skill{Value: 1, weight: 1},
	"Science: Pharmacy":    Skill{Value: 1, weight: 1},
	"Science: Physics":     Skill{Value: 1, weight: 1},
	"Science: Planetology": Skill{Value: 1, weight: 1},
	"Science: Psychology":  Skill{Value: 5, weight: 1},
	"Science: Thaumaturgy": Skill{Value: 0, weight: 1},
	"Science: Zoology":     Skill{Value: 5, weight: 1},
}

// TechnologyUseSkills is a map of skills.
var TechnologyUseSkills = map[string]Skill{
	"Technology Use: Communications":      Skill{Value: 5, weight: 1},
	"Technology Use: Electronic Security": Skill{Value: 5, weight: 1},
	"Technology Use: Electronics":         Skill{Value: 5, weight: 1},
	"Technology Use: Other":               Skill{Value: 5, weight: 1},
	"Technology Use: Sensor Systems":      Skill{Value: 5, weight: 1},
	"Technology Use: Surveillance":        Skill{Value: 5, weight: 1},
	"Technology Use: Traps":               Skill{Value: 5, weight: 1},
}

// EsotericSkills is a map of the creepy stuff.
var EsotericSkills = map[string]Skill{
	"Cthulhu Mythos": Skill{Value: 0, weight: 1},
	"Sorcery":        Skill{Value: 0, weight: 1},
}

// GeneralSkills is a map of everything except Combat skills.
var GeneralSkills = joinMaps(
	DefaultSkills,
	ArtSkills,
	AthleticsSkills,
	ComputerUseSkills,
	CraftSkills,
	DriveSkills,
	HeavyMachineSkills,
	KnowledgeSkills,
	MedicineSkills,
	PerformSkills,
	PilotSkills,
	LanguageSkills,
	ScienceSkills,
	TechnologyUseSkills,
	EsotericSkills,
)
