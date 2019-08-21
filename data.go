package tlfgen

import "sort"

// Skill is a single skill.
type Skill struct {
	base   int `json:"value"`
	weight int
}

func (s Skill) add(n int) {
	s.base += n
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
	skills []string
	offset int
	n      int
	wealth string
}

// Professions is a map of professions.
var Professions = map[string]Profession{
	"Student": {
		skills: []string{
			"Language: Own",
			"Research",
			"ART",
			"Athletics",
			"COMPUTER USE",
			"CRAFT",
			"First Aid",
			"Insight",
			"KNOWLEDGE",
			"LANGUAGE",
			"Listen",
			"Medicine",
			"Perform",
			"Persuade",
			"Medicine: Psychotherapy",
			"REPAIR",
			"SCIENCE",
			"TECHNOLOGY USE",
		},
		offset: 2,
		n:      8,
		wealth: "Poor or Average",
	},
	"Spy": {
		skills: []string{
			"Dodge",
			"Fast Talk",
			"Hide",
			"Listen",
			"Research",
			"Spot",
			"Stealth",
			"Art: Photography",
			"Brawl",
			"Bureaucracy",
			"COMPUTER USE",
			"Disguise",
			"Etiquette",
			"FIREARM",
			"Grapple",
			"KNOWLEDGE",
			"LANGUAGE",
			"Language: Own",
			"Navigate",
			"PILOT",
			"Knowledge: Psychology",
			"Repair: Electronics",
			"Repair: Mechanical",
			"Ride",
			"Swim",
			"TECHNOLOGY USE",
			"Throw",
			"Track",
		},
		offset: 7,
		n:      3,
		wealth: "Average",
	},
	"Scientist": {
		skills: []string{
			"COMPUTER USE",
			"CRAFT",
			"Persuade",
			"Research",
			"TECHNOLOGY USE",
			"KNOWLEDGE",
			"SCIENCE",
		},
		offset: 5,
		n:      5,
		wealth: "Average",
	},
	"Professor, Scholar or Teacher": {
		skills: []string{
			"Insight",
			"Persuade",
			"Research",
			"Teach",
			"Appraise",
			"ART",
			"COMPUTER USE",
			"CRAFT",
			"First Aid",
			"Insight",
			"KNOWLEDGE",
			"LANGUAGE",
			"Language: Own",
			"Listen",
			"MEDICINE",
			"PSYCHOLOGY",
			"REPAIR",
			"SCIENCE",
			"TECHNOLOGY USE",
		},
		offset: 4,
		n:      6,
		wealth: "Average",
	},
	"Police Officer": {
		skills: []string{
			"Brawl",
			"Knowledge: Law",
			"Listen",
			"Persuade",
			"Spot",
			"Brawl",
			"COMPUTER USE",
			"Disguise",
			"Dodge",
			"DRIVE",
			"Fast Talk",
			"FIREARM",
			"Grapple",
			"Hide",
			"Insight",
			"Knowledge: Streetwise",
			"LANGUAGE",
			"Stealth",
		},
		offset: 5,
		n:      5,
		wealth: "Average",
	},
	"Parapsychologist": {
		skills: []string{
			"Fast Talk",
			"Hide",
			"Insight",
			"KNOWLEDGE",
			"Knowledge: Occult",
			"Listen",
			"Research",
			"Sense",
			"Spot",
			"Stealth",
			"PSYCHOLOGY",
		},
		offset: 0,
		n:      0,
		wealth: "Average",
	},
	"Military Officer": {
		skills: []string{
			"ARTILLERY",
			"Command",
			"DRIVE",
			"FIREARM",
			"Navigate",
			"Persuade",
			"Strategy",
			"Teach",
			"ANY",
		},
		offset: 8,
		n:      2,
		wealth: "Average",
	},
	"Linguist": {
		skills: []string{
			"Etiquette",
			"Insight",
			"KNOWLEDGE",
			"LANGUAGE",
			"Language: Own",
			"Listen",
			"Persuade",
			"KNOWLEDGE",
			"LANGUAGE",
		},
		offset: 7,
		n:      3,
		wealth: "Average",
	},
	"Lawyer": {
		skills: []string{
			"Bargain",
			"Bureaucracy",
			"Fast Talk",
			"Insight",
			"KNOWLEDGE",
			"Knowledge: Law",
			"Language: Own",
			"Perform: Oratory",
			"Persuade",
			"Research",
		},
		offset: 10,
		n:      0,
		wealth: "Affluent",
	},
	"Labourer": {
		skills: []string{
			"Brawl",
			"Climb",
			"CRAFT",
			"Drive",
			"Grapple",
			"HEAVY MACHINE",
			"Appraise",
			"COMPUTER USE",
			"Fine Manipulation",
			"LANGUAGE",
			"Repair: Mechanical",
			"Repair: Structural",
			"TECHNOLOGY USE",
		},
		offset: 6,
		n:      4,
		wealth: "Average",
	},
	"Journalist": {
		skills: []string{
			"Fast Talk",
			"Insight",
			"Language: Own",
			"Listen",
			"Persuade",
			"Research",
			"Spot",
			"Art: Photography",
			"COMPUTER USE",
			"Craft: Photography",
			"Disguise",
			"Hide",
			"KNOWLEDGE",
			"LANGUAGE",
			"Stealth",
		},
		offset: 7,
		n:      3,
		wealth: "Average",
	},
	"Engineer or Technician": {
		skills: []string{
			"COMPUTER USE",
			"CRAFT",
			"Fine Manipulation",
			"REPAIR",
			"TECHNOLOGY USE",
			"Art: Drafting",
			"ARTILLERY",
			"Demolition",
			"Drive",
			"HEAVY MACHINE",
			"KNOWLEDGE",
			"PILOT",
			"REPAIR",
			"Research",
			"SCIENCE",
		},

		offset: 5,
		n:      5,
		wealth: "Average",
	},
	"Doctor": {
		skills: []string{
			"First Aid",
			"Language: Own",
			"MEDICINE",
			"Persuade",
			"Research",
			"Spot",
			"Insight",
			"LANGUAGE",
			"Medicine: Psychotherapy",
			"SCIENCE",
			"Science: Biology",
		},
		offset: 6,
		n:      4,
		wealth: "Affluent",
	},
	"Dilettante": {
		skills: []string{
			"Appraise",
			"Etiquette",
			"ART",
			"Athletics",
			"CRAFT",
			"DRIVE",
			"Gaming",
			"KNOWLEDGE",
			"LANGUAGE",
			"PERFORM",
			"Research",
			"SCIENCE",
			"TECHNOLOGY USE",
		},
		offset: 2,
		n:      6,
		wealth: "Affluent",
	},
	"Consultant": {
		skills: []string{
			"Bureaucracy",
			"COMPUTER USE",
			"Fast Talk",
			"Insight",
			"Listen",
			"Language: Own",
			"Persuade",
			"Research",
			"Appraise",
			"Bargain",
			"Etiquette",
			"KNOWLEDGE",
			"SCIENCE",
			"TECHNICAL",
		},
		offset: 10,
		n:      2,
		wealth: "Affluent",
	},
	"Computer Hacker or Technician": {
		skills: []string{
			"COMPUTER USE",
			"COMPUTER USE",
			"KNOWLEDGE",
			"KNOWLEDGE",
			"Repair: Electrical",
			"Repair: Electronics",
			"Research",
			"Science: Mathematics",
			"TECHNOLOGY USE",
			"Bureaucracy",
			"Hide",
			"Knowledge: Law",
		},
		offset: 9,
		n:      1,
		wealth: "Average",
	},
	"Clerical Worker": {
		skills: []string{
			"Bargain",
			"Bureaucracy",
			"COMPUTER USE",
			"Etiquette",
			"Knowledge: Accounting",
			"KNOWLEDGE",
			"Knowledge: Law",
			"Language: Own",
			"Persuade",
			"Research",
		},
		offset: 0,
		n:      0,
		wealth: "Average",
	},
	"Artist or Designer": {
		skills: []string{
			"CRAFT",
			"Insight",
			"KNOWLEDGE",
			"LANGUAGE",
			"Language: Own",
			"Listen",
			"Research",
			"Spot",
			"ART",
		},
		offset: 8,
		n:      2,
		wealth: "Average",
	},
	"Antiquarian": {
		skills: []string{
			"Appraise",
			"ART",
			"Bargain",
			"CRAFT",
			"Fine Manipulation",
			"KNOWLEDGE",
			"Knowledge: History",
			"Research",
			"ART",
			"KNOWLEDGE",
			"SCIENCE",
		},
		offset: 8,
		n:      2,
		wealth: "Average",
	},
	"Occultist": {
		skills: []string{
			"Fast Talk",
			"Insight",
			"Knowledge: Anthropology",
			"Knowledge: History",
			"Knowledge: Occult",
			"LANGUAGE",
			"Language: Own",
			"Research",
			"COMPUTER USE",
			"CRAFT",
			"Knowledge: Archaeology",
			"Medicine",
			"SCIENCE",
		},
		offset: 8,
		n:      2,
		wealth: "Average",
	},
	"Philosopher": {
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
		offset: 8,
		n:      2,
		wealth: "Average",
	},
}

// ListProfessionKeys lists all the profession types.
func ListProfessionKeys() []string {
	keys := []string{}
	for key := range Professions {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// PersonalityTypes is a map of PTypes.
var PersonalityTypes = map[string]PersonalityType{
	"Slacker": {
		bonus: 20,
		desc:  "Your character has spent their lifetime dodging responsibility and believes that problems are best avoided altogether.",
		skills: []string{
			"Bargain",
			"Bureaucracy",
			"Disguise",
			"Dodge",
			"Fast Talk",
			"Gaming",
			"Hide",
			"Insight",
			"LANGUAGE",
			"Persuade",
			"Sense",
			"Sleight of Hand",
			"Spot",
		},
		special: "",
	},
	"Leader": {
		bonus: 20,
		desc:  "Your character enjoys calling the shots and persuading others to work.",
		skills: []string{
			"Appraise",
			"Bargain",
			"COMBAT",
			"Command",
			"Etiquette",
			"Fast Talk",
			"Insight",
			"KNOWLEDGE",
			"LANGUAGE",
			"Language: Own",
			"PERFORM",
			"Persuade",
			"Sense",
		},
		special: "",
	},
	"Thinker": {
		bonus: 20,
		desc:  "When confronted with opposition, your character’s first instinct is to outsmart their opponent to gain an advantage.",
		skills: []string{
			"Appraise",
			"Bargain",
			"COMBAT",
			"Disguise",
			"Insight",
			"KNOWLEDGE",
			"Listen",
			"Research",
			"Sense",
			"Spot",
			"Stealth",
			"TECHNICAL",
		},
		special: "",
	},
	"Master": {
		bonus: 21,
		desc:  "Your character believes that technique, craft and expertise are the keys to success.",
		skills: []string{
			"Appraise",
			"COMBAT",
			"CRAFT",
			"Disguise",
			"Dodge",
			"Fine Manipulation",
			"First Aid",
			"KNOWLEDGE",
			"Navigate",
			"PILOT",
			"Ride",
			"Sleight of Hand",
			"Stealth",
		},
		special: "",
	},
	"Bruiser": {
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
	"Nutter": {
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

// ListPersonalityTypeKeys lists all the personality types.
func ListPersonalityTypeKeys() []string {
	keys := []string{}
	for key := range PersonalityTypes {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// Assignments is a map of assignments.
var Assignments = map[string]Assignment{
	"Zombie Wrangler": {
		bonus: 10,
		skills: []string{
			"Command",
			"Grapple",
			"Knowledge: Occult",
			"Sense",
			"Sorcery",
		},
	},
	"Researcher": {
		bonus: 10,
		skills: []string{
			"Bureaucracy",
			"KNOWLEDGE",
			"KNOWLEDGE",
			"Listen",
			"Research",
		},
	},
	"Tosher": {
		bonus: 10,
		skills: []string{
			"Climb",
			"FIREARM",
			"SCIENCE",
			"Spot",
			"Technology Use: Survival gear",
		},
	},
	"Mad Boffin": {
		bonus: 10,
		skills: []string{
			"Computer Use: Magic",
			"CRAFT",
			"Firearms: Esoteric",
			"REPAIR",
			"SCIENCE",
		},
	},
	"Computational Demonology Researcher": {
		bonus: 10,
		skills: []string{
			"COMPUTER USE",
			"Computer Use: Magic",
			"Science: Mathematics",
			"Science: Thaumaturgy",
			"Sorcery",
		},
	},
	"Cultural Attache": {
		bonus: 10,
		skills: []string{
			"Etiquette",
			"Knowledge: History",
			"Knowledge: Occult",
			"Knowledge: Politics",
			"Persuade",
		},
	},
	"Counter-Possession Exorcist": {
		bonus: 10,
		skills: []string{
			"Bureaucracy",
			"Insight",
			"Knowledge: Occult",
			"Sorcery",
			"Stealth",
		},
	},
	"Translator": {
		bonus: 10,
		skills: []string{
			"Appraise",
			"Fine Manipulation",
			"Knowledge: History",
			"Knowledge: Occult",
			"Language: Any",
			"Sorcery",
		},
	},
	"Apprentice Demonologist": {
		bonus: 10,
		skills: []string{
			"Knowledge: Law",
			"Knowledge: Occult",
			"Persuade",
			"Research",
			"Sorcery",
		},
	},
	"Plumber": {
		bonus: 10,
		skills: []string{
			"FIREARM",
			"Knowledge: Occult",
			"Science: Thaumaturgy",
			"Sorcery",
			"Stealth",
		},
	},
	"Occult Forensics": {
		bonus: 10,
		skills: []string{
			"SCIENCE",
			"SCIENCE",
			"SCIENCE",
			"Sense",
			"Spot",
		},
	},
	"Medical and Psychological": {
		bonus: 10,
		skills: []string{
			"First Aid",
			"MEDICINE",
			"MEDICINE",
			"Research",
			"Science: Biology",
		},
	},
	"Media Relations": {
		bonus: 10,
		skills: []string{
			"Computer Use: Hacking",
			"Fast Talk",
			"Knowledge: Occult",
			"Knowledge: Politics",
			"Research",
		},
	},
	"Counter-Subversion": {
		bonus: 10,
		skills: []string{
			"Insight",
			"Knowledge: Espionage",
			"Knowledge: Politics",
			"Research",
			"Technology Use: Surveillance",
		},
	},
	"Information Technology": {
		bonus: 10,
		skills: []string{
			"COMPUTER USE",
			"COMPUTER USE",
			"COMPUTER USE",
			"Knowledge: Occult",
			"TECHNOLOGY USE",
		},
	},
	"Counter-Possession": {
		bonus: 10,
		skills: []string{
			"Bureaucracy",
			"Insight",
			"Knowledge: Occult",
			"Stealth",
			"Sorcery",
		},
	},
	"Contracts and Bindings": {
		bonus: 10,
		skills: []string{
			"Knowledge: Law",
			"Knowledge: Occult",
			"Persuade",
			"Research",
			"Sorcery",
		},
	},
	"Archives": {
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
	"Computational Demonology": {
		bonus: 10,
		skills: []string{
			"COMPUTER USE",
			"Computer Use: Magic",
			"Science: Mathematics",
			"Science: Thaumaturgy",
			"Sorcery",
		},
	},
	"all": {
		bonus: 5,
		skills: []string{
			"Bureaucracy",
			"COMPUTER USE",
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

// ListAssignmentKeys lists all the assignment types.
func ListAssignmentKeys() []string {
	keys := []string{}
	for key := range Assignments {
		if key != "all" {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	return keys
}

func getSkillNames(skills map[string]Skill) []string {
	names := []string{}
	for k := range skills {
		names = append(names, k)
	}
	sort.Strings(names)
	return names
}

func randomSkillChoice(skills map[string]Skill) (string, Skill) {
	names := getSkillNames(skills)
	name := randomChoice(names)
	return name, skills[name]
}

func randomWeightedSkillChoice(skills map[string]Skill) (string, Skill) {
	names := getSkillNames(skills)
	weights := []int{}
	for _, name := range names {
		weights = append(weights, skills[name].weight)
	}
	choice := weightedRandomChoice(names, weights)
	return choice, skills[choice]
}

func getSkill(name string) (string, Skill) {
	var newSkill Skill
	switch name {
	case "ARTILLERY":
		name, newSkill = randomWeightedSkillChoice(ArtillerySkills)
	case "COMBAT":
		name, newSkill = randomWeightedSkillChoice(CombatSkills)
	case "FIREARM":
		name, newSkill = randomWeightedSkillChoice(FirearmSkills)
	case "KNOWLEDGE":
		name, newSkill = randomWeightedSkillChoice(KnowledgeSkills)
	case "CRAFT":
		name, newSkill = randomWeightedSkillChoice(CraftSkills)
	case "PSYCHOLOGY":
		name, newSkill = randomWeightedSkillChoice(PsychologySkills)
	case "MEDICINE":
		name, newSkill = randomWeightedSkillChoice(MedicineSkills)
	case "LANGUAGE":
		name, newSkill = randomWeightedSkillChoice(LanguageSkills)
	case "SCIENCE":
		name, newSkill = randomWeightedSkillChoice(ScienceSkills)
	case "PERFORM":
		name, newSkill = randomWeightedSkillChoice(PerformSkills)
	case "PILOT":
		name, newSkill = randomWeightedSkillChoice(PilotSkills)
	case "REPAIR":
		name, newSkill = randomWeightedSkillChoice(RepairSkills)
	case "ART":
		name, newSkill = randomWeightedSkillChoice(ArtSkills)
	case "TECHNOLOGY USE":
		name, newSkill = randomWeightedSkillChoice(TechnologyUseSkills)
	case "TECHNICAL":
		name, newSkill = randomWeightedSkillChoice(TechnicalSkills)
	case "COMPUTER USE":
		name, newSkill = randomWeightedSkillChoice(ComputerUseSkills)
	case "HEAVY MACHINE":
		name, newSkill = randomWeightedSkillChoice(HeavyMachineSkills)
	case "DRIVE":
		name, newSkill = randomWeightedSkillChoice(DriveSkills)
	case "ANY":
		name, newSkill = randomWeightedSkillChoice(AllSkills)
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
	"Artillery: Cannon":           {base: 0, weight: 10},
	"Artillery: Other":            {base: 0, weight: 10},
	"Artillery: Rocket Launcher":  {base: 0, weight: 10},
	"Artillery: Turret":           {base: 0, weight: 10},
	"Artillery: Vehicular Weapon": {base: 0, weight: 10},
}

// FirearmSkills is a map of skills.
var FirearmSkills = map[string]Skill{
	"Firearm: Assault Rifle":  {base: 15, weight: 5},
	"Firearm: Esoteric":       {base: 0, weight: 20},
	"Firearm: Exotic":         {base: 5, weight: 10},
	"Firearm: Pistol":         {base: 20, weight: 40},
	"Firearm: Rifle":          {base: 25, weight: 30},
	"Firearm: Shotgun":        {base: 30, weight: 10},
	"Firearm: Submachine Gun": {base: 15, weight: 5},
}

// HeavyWeaponSkills is a map of skills.
var HeavyWeaponSkills = map[string]Skill{
	"Heavy Weapon: Bazooka":           {base: 0, weight: 10},
	"Heavy Weapon: Grenade Launcher":  {base: 0, weight: 10},
	"Heavy Weapon: Heavy Machine Gun": {base: 0, weight: 10},
	"Heavy Weapon: Minigun":           {base: 0, weight: 10},
	"Heavy Weapon: Other":             {base: 0, weight: 10},
	"Heavy Weapon: Rocket Launcher":   {base: 0, weight: 10},
}

// MeleeWeaponSkills is a map of skills.
var MeleeWeaponSkills = map[string]Skill{
	"Melee Weapon: Axe":     {base: 5, weight: 15},
	"Melee Weapon: Club":    {base: 5, weight: 10},
	"Melee Weapon: Garrote": {base: 5, weight: 10},
	"Melee Weapon: Knife":   {base: 5, weight: 30},
	"Melee Weapon: Other":   {base: 5, weight: 10},
	"Melee Weapon: Spear":   {base: 5, weight: 10},
	"Melee Weapon: Staff":   {base: 5, weight: 15},
	"Melee Weapon: Sword":   {base: 5, weight: 15},
	"Melee Weapon: Whip":    {base: 5, weight: 10},
}

// MissileWeaponSkills is a map of skills.
var MissileWeaponSkills = map[string]Skill{
	"Missile Weapon: Blowgun":        {base: 5, weight: 10},
	"Missile Weapon: Boomerang":      {base: 5, weight: 10},
	"Missile Weapon: Bow":            {base: 5, weight: 30},
	"Missile Weapon: Crossbow":       {base: 5, weight: 10},
	"Missile Weapon: Dart":           {base: 5, weight: 10},
	"Missile Weapon: Javelin":        {base: 5, weight: 10},
	"Missile Weapon: Other":          {base: 5, weight: 10},
	"Missile Weapon: Shuriken":       {base: 5, weight: 10},
	"Missile Weapon: Sling":          {base: 5, weight: 10},
	"Missile Weapon: Spear":          {base: 5, weight: 10},
	"Missile Weapon: Throwing Axe":   {base: 5, weight: 10},
	"Missile Weapon: Throwing Knife": {base: 5, weight: 15},
}

// DefaultSkills is a map of default skills.
var DefaultSkills = map[string]Skill{
	"Appraise":          {base: 15, weight: 5},
	"Athletics":         {base: 10, weight: 5},
	"Bargain":           {base: 5, weight: 15},
	"Brawl":             {base: 25, weight: 5},
	"Bureaucracy":       {base: 5, weight: 10},
	"Climb":             {base: 40, weight: -40},
	"Command":           {base: 5, weight: 10},
	"Demolition":        {base: 1, weight: 10},
	"Disguise":          {base: 5, weight: 10},
	"Dodge":             {base: -1, weight: 30},
	"Drive: Automobile": {base: 20, weight: -10},
	"Etiquette":         {base: 5, weight: 15},
	"Fast Talk":         {base: 5, weight: 15},
	"Fine Manipulation": {base: 5, weight: 10},
	"First Aid":         {base: 30, weight: 0},
	"Gaming":            {base: 10, weight: 5},
	"Grapple":           {base: 25, weight: 5},
	"Hide":              {base: 10, weight: 10},
	"Insight":           {base: 5, weight: 15},
	"Jump":              {base: 25, weight: 0},
	"Listen":            {base: 25, weight: -5},
	"Navigate":          {base: 10, weight: 10},
	"Perform":           {base: 5, weight: 5},
	"Persuade":          {base: 15, weight: 5},
	"Research":          {base: 25, weight: 5},
	"Ride":              {base: 5, weight: 5},
	"Sense":             {base: 10, weight: 5},
	"Sleight of Hand":   {base: 5, weight: 10},
	"Spot":              {base: 25, weight: -5},
	"Status":            {base: 15, weight: 5},
	"Stealth":           {base: 10, weight: 20},
	"Strategy":          {base: 5, weight: 5},
	"Swim":              {base: 25, weight: 0},
	"Teach":             {base: 10, weight: 5},
	"Throw":             {base: 25, weight: 5},
	"Track":             {base: 10, weight: 10},
}

// ArtSkills is a map of skills.
var ArtSkills = map[string]Skill{
	"Art: Calligraphy": {base: 5, weight: 10},
	"Art: Drawing":     {base: 5, weight: 10},
	"Art: Other":       {base: 5, weight: 10},
	"Art: Painting":    {base: 5, weight: 10},
	"Art: Photography": {base: 5, weight: 10},
	"Art: Sculpture":   {base: 5, weight: 10},
	"Art: Writing":     {base: 5, weight: 10},
}

// AthleticsSkills is a map of skills.
var AthleticsSkills = map[string]Skill{
	"Athletics: Acrobatics":    {base: 10, weight: 10},
	"Athletics: Football":      {base: 10, weight: 10},
	"Athletics: Baseball":      {base: 10, weight: 10},
	"Athletics: Basketball":    {base: 10, weight: 10},
	"Athletics: Bowling":       {base: 10, weight: 10},
	"Athletics: Cricket":       {base: 10, weight: 10},
	"Athletics: Cycling":       {base: 10, weight: 10},
	"Athletics: Golf":          {base: 10, weight: 10},
	"Athletics: Hockey":        {base: 10, weight: 10},
	"Athletics: Rugby":         {base: 10, weight: 10},
	"Athletics: Skating":       {base: 10, weight: 10},
	"Athletics: Skiing":        {base: 10, weight: 10},
	"Athletics: Soccer":        {base: 10, weight: 10},
	"Athletics: Tennis":        {base: 10, weight: 10},
	"Athletics: Track & Field": {base: 10, weight: 10},
}

// ComputerUseSkills is a map of skills.
var ComputerUseSkills = map[string]Skill{
	"Computer Use: Art":         {base: 5, weight: 15},
	"Computer Use: Design":      {base: 5, weight: 15},
	"Computer Use: Gaming":      {base: 5, weight: 15},
	"Computer Use: Hacking":     {base: 5, weight: 25},
	"Computer Use: Magic":       {base: 5, weight: 35},
	"Computer Use: Maintenance": {base: 5, weight: 25},
	"Computer Use: Other":       {base: 5, weight: 15},
	"Computer Use: Programming": {base: 5, weight: 25},
	"Computer Use: Repair":      {base: 5, weight: 25},
}

// RepairSkills is a map of skills.
var RepairSkills = map[string]Skill{
	"Repair: Electrical": {base: 5, weight: 10},
	"Repair: Electronic": {base: 5, weight: 10},
	"Repair: Hydraulic":  {base: 5, weight: 10},
	"Repair: Mechanical": {base: 5, weight: 10},
	"Repair: Plumbing":   {base: 5, weight: 10},
	"Repair: Structural": {base: 5, weight: 10},
}

// CraftSkills is a map of skills.
var CraftSkills = map[string]Skill{
	"Craft: Carpentry":      {base: 5, weight: 10},
	"Craft: Cooking":        {base: 5, weight: 10},
	"Craft: Leatherworking": {base: 5, weight: 10},
	"Craft: Pottery":        {base: 5, weight: 10},
	"Craft: Sewing":         {base: 5, weight: 10},
	"Craft: Woodworking":    {base: 5, weight: 10},
}

// DriveSkills is a map of skills.
var DriveSkills = map[string]Skill{
	"Drive: Automobile":       {base: 20, weight: 20},
	"Drive: Industrial Mover": {base: 20, weight: 10},
	"Drive: Motorcycle":       {base: 20, weight: 10},
	"Drive: Other":            {base: 20, weight: 10},
	"Drive: Tank":             {base: 20, weight: 10},
}

// HeavyMachineSkills is a map of skills.
var HeavyMachineSkills = map[string]Skill{
	"Heavy Machine: Boiler":    {base: 5, weight: 10},
	"Heavy Machine: Bulldozer": {base: 5, weight: 10},
	"Heavy Machine: Crane":     {base: 5, weight: 10},
	"Heavy Machine: Engine":    {base: 5, weight: 10},
	"Heavy Machine: Other":     {base: 5, weight: 10},
	"Heavy Machine: Turbine":   {base: 5, weight: 10},
	"Heavy Machine: Wrecker":   {base: 5, weight: 10},
}

// KnowledgeSkills is a map of skills.
var KnowledgeSkills = map[string]Skill{
	"Knowledge: Accounting":      {base: 10, weight: 10},
	"Knowledge: Anthropology":    {base: 1, weight: 10},
	"Knowledge: Archaeology":     {base: 1, weight: 10},
	"Knowledge: Art History":     {base: 1, weight: 10},
	"Knowledge: Business":        {base: 1, weight: 10},
	"Knowledge: Espionage":       {base: 0, weight: 20},
	"Knowledge: Folklore":        {base: 5, weight: 10},
	"Knowledge: Group":           {base: 0, weight: 10},
	"Knowledge: History":         {base: 20, weight: 10},
	"Knowledge: Law":             {base: 5, weight: 10},
	"Knowledge: Linguistics":     {base: 0, weight: 10},
	"Knowledge: Literature":      {base: 5, weight: 10},
	"Knowledge: Natural History": {base: 10, weight: 10},
	"Knowledge: Occult":          {base: 5, weight: 20},
	"Knowledge: Philosophy":      {base: 1, weight: 10},
	"Knowledge: Politics":        {base: 5, weight: 10},
	"Knowledge: Region":          {base: 0, weight: 10},
	"Knowledge: Religion":        {base: 5, weight: 10},
	"Knowledge: Streetwise":      {base: 5, weight: 20},
}

// MedicineSkills is a map of skills.
var MedicineSkills = map[string]Skill{
	"Medicine: Dermatology":       {base: 5, weight: 20},
	"Medicine: Family Medicine":   {base: 5, weight: 20},
	"Medicine: Immunology":        {base: 5, weight: 20},
	"Medicine: Internal Medicine": {base: 5, weight: 20},
	"Medicine: Neurology":         {base: 5, weight: 30},
	"Medicine: Nuclear Medicine":  {base: 5, weight: 20},
	"Medicine: Oncology":          {base: 5, weight: 20},
	"Medicine: Other":             {base: 5, weight: 20},
	"Medicine: Pathology":         {base: 5, weight: 30},
	"Medicine: Pediatrics":        {base: 5, weight: 20},
	"Medicine: Radiology":         {base: 5, weight: 20},
	"Medicine: Psychotherapy":     {base: 5, weight: 30},
	"Medicine: Psychiatry":        {base: 5, weight: 30},
	"Medicine: Surgery":           {base: 5, weight: 20},
}

// PsychologySkills is a map of skills.
var PsychologySkills = map[string]Skill{
	"Medicine: Psychotherapy": {base: 5, weight: 30},
	"Medicine: Psychiatry":    {base: 5, weight: 30},
	"Knowledge: Psychology":   {base: 5, weight: 30},
}

// PerformSkills is a map of skills.
var PerformSkills = map[string]Skill{
	"Perform: Act":             {base: 5, weight: 15},
	"Perform: Dance":           {base: 5, weight: 15},
	"Perform: Juggle":          {base: 5, weight: 15},
	"Perform: Other":           {base: 5, weight: 15},
	"Perform: Play Instrument": {base: 5, weight: 15},
	"Perform: Recite":          {base: 5, weight: 15},
	"Perform: Sing":            {base: 5, weight: 15},
}

// PilotSkills is a map of skills.
var PilotSkills = map[string]Skill{
	"Pilot: Battleship":      {base: 0, weight: 20},
	"Pilot: Helicopter":      {base: 0, weight: 20},
	"Pilot: Hot Air Balloon": {base: 0, weight: 20},
	"Pilot: Hovercraft":      {base: 0, weight: 20},
	"Pilot: Jet Airliner":    {base: 0, weight: 20},
	"Pilot: Jet Boat":        {base: 0, weight: 20},
	"Pilot: Jet Fighter":     {base: 0, weight: 20},
	"Pilot: Ocean Liner":     {base: 0, weight: 20},
	"Pilot: Propeller Plane": {base: 0, weight: 20},
}

// LanguageSkills is a map of skills.
var LanguageSkills = map[string]Skill{
	"Language: Arabic":   {base: 0, weight: 15},
	"Language: Farsi":    {base: 0, weight: 15},
	"Language: English":  {base: 0, weight: 15},
	"Language: French":   {base: 0, weight: 15},
	"Language: German":   {base: 0, weight: 15},
	"Language: Greek":    {base: 0, weight: 15},
	"Language: Italian":  {base: 0, weight: 15},
	"Language: Japanese": {base: 0, weight: 15},
	"Language: Latin":    {base: 0, weight: 15},
	"Language: Mandarin": {base: 0, weight: 15},
	"Language: Other":    {base: 0, weight: 15},
	"Language: Own":      {base: -1, weight: -1},
	"Language: Spanish":  {base: 0, weight: 15},
	"Language: Swahili":  {base: 0, weight: 15},
	"Language: Swedish":  {base: 0, weight: 15},
	"Language: Enochian": {base: 0, weight: 15},
	"Language: Aramaic":  {base: 0, weight: 15},
	"Language: Hebrew":   {base: 0, weight: 15},
}

// ScienceSkills is a map of skills.
var ScienceSkills = map[string]Skill{
	"Science: Astronomy":   {base: 1, weight: 30},
	"Science: Biology":     {base: 1, weight: 20},
	"Science: Botany":      {base: 1, weight: 20},
	"Science: Chemistry":   {base: 1, weight: 30},
	"Science: Cyptography": {base: 1, weight: 20},
	"Science: Forensics":   {base: 1, weight: 30},
	"Science: Genetics":    {base: 1, weight: 20},
	"Science: Geology":     {base: 1, weight: 20},
	"Science: Mathematics": {base: 10, weight: 30},
	"Science: Meteorology": {base: 1, weight: 20},
	"Science: Pharmacy":    {base: 1, weight: 20},
	"Science: Physics":     {base: 1, weight: 30},
	"Science: Planetology": {base: 1, weight: 20},
	"Science: Psychology":  {base: 5, weight: 20},
	"Science: Thaumaturgy": {base: 0, weight: 30},
	"Science: Zoology":     {base: 5, weight: 20},
}

// TechnologyUseSkills is a map of skills.
var TechnologyUseSkills = map[string]Skill{
	"Technology Use: Communications":      {base: 5, weight: 20},
	"Technology Use: Electronic Security": {base: 5, weight: 20},
	"Technology Use: Electronics":         {base: 5, weight: 20},
	"Technology Use: Other":               {base: 5, weight: 20},
	"Technology Use: Sensor Systems":      {base: 5, weight: 20},
	"Technology Use: Surveillance":        {base: 5, weight: 20},
	"Technology Use: Traps":               {base: 5, weight: 20},
}

// EsotericSkills is a map of the creepy stuff.
var EsotericSkills = map[string]Skill{
	"Cthulhu Mythos": {base: 0, weight: 20},
	"Sorcery":        {base: 0, weight: 20},
}

// TechnicalSkills list of all the possible combat skills.
var TechnicalSkills = joinMaps(
	ComputerUseSkills,
	TechnologyUseSkills,
)

// CombatSkills list of all the possible combat skills.
var CombatSkills = joinMaps(
	ArtillerySkills,
	FirearmSkills,
	HeavyWeaponSkills,
	MeleeWeaponSkills,
	MissileWeaponSkills,
)

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
	PsychologySkills,
	LanguageSkills,
	ScienceSkills,
	TechnologyUseSkills,
	EsotericSkills,
)

// AllSkills is a map of everything.
var AllSkills = joinMaps(
	CombatSkills,
	GeneralSkills,
)
