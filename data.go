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
			"Computer Use: *",
			"CRAFT",
			"Knowledge: Archaeology",
			"Medicine",
			"Science: any",
		},
		optOffset: 8,
		optNum:    2,
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
		optOffset: 8,
		optNum:    2,
	},
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
			 "Language (Own)",
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
		bonus: 20,
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

// Assignments is a map of assignments.
var Assignments = map[string]Assignment{
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
	case "TECHNICAL":
		name, newSkill = randomWeightedSkillChoice(TechnicalSkills)
	case "COMPUTER":
		name, newSkill = randomWeightedSkillChoice(ComputerUseSkills)
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
	"Artillery: Cannon":           {Value: 0, weight: 1},
	"Artillery: Other":            {Value: 0, weight: 1},
	"Artillery: Rocket Launcher":  {Value: 0, weight: 1},
	"Artillery: Turret":           {Value: 0, weight: 1},
	"Artillery: Vehicular Weapon": {Value: 0, weight: 1},
}

// FirearmSkills is a map of skills.
var FirearmSkills = map[string]Skill{
	"Firearm: Assault Rifle":  {Value: 15, weight: 5},
	"Firearm: Esoteric":       {Value: 0, weight: 5},
	"Firearm: Exotic":         {Value: 5, weight: 1},
	"Firearm: Pistol":         {Value: 20, weight: 50},
	"Firearm: Rifle":          {Value: 25, weight: 20},
	"Firearm: Shotgun":        {Value: 30, weight: 20},
	"Firearm: Submachine Gun": {Value: 15, weight: 1},
}

// HeavyWeaponSkills is a map of skills.
var HeavyWeaponSkills = map[string]Skill{
	"Heavy Weapon: Bazooka":           {Value: 0, weight: 1},
	"Heavy Weapon: Grenade Launcher":  {Value: 0, weight: 1},
	"Heavy Weapon: Heavy Machine Gun": {Value: 0, weight: 1},
	"Heavy Weapon: Minigun":           {Value: 0, weight: 1},
	"Heavy Weapon: Other":             {Value: 0, weight: 1},
	"Heavy Weapon: Rocket Launcher":   {Value: 0, weight: 1},
}

// MeleeWeaponSkills is a map of skills.
var MeleeWeaponSkills = map[string]Skill{
	"Melee Weapon: Axe":     {Value: 5, weight: 5},
	"Melee Weapon: Club":    {Value: 5, weight: 10},
	"Melee Weapon: Garrote": {Value: 5, weight: 1},
	"Melee Weapon: Knife":   {Value: 5, weight: 20},
	"Melee Weapon: Other":   {Value: 5, weight: 1},
	"Melee Weapon: Spear":   {Value: 5, weight: 1},
	"Melee Weapon: Staff":   {Value: 5, weight: 5},
	"Melee Weapon: Sword":   {Value: 5, weight: 5},
	"Melee Weapon: Whip":    {Value: 5, weight: 1},
}

// MissileWeaponSkills is a map of skills.
var MissileWeaponSkills = map[string]Skill{
	"Missile Weapon: Blowgun":        {Value: 5, weight: 1},
	"Missile Weapon: Boomerang":      {Value: 5, weight: 1},
	"Missile Weapon: Bow":            {Value: 5, weight: 20},
	"Missile Weapon: Crossbow":       {Value: 5, weight: 1},
	"Missile Weapon: Dart":           {Value: 5, weight: 1},
	"Missile Weapon: Javelin":        {Value: 5, weight: 1},
	"Missile Weapon: Other":          {Value: 5, weight: 1},
	"Missile Weapon: Shuriken":       {Value: 5, weight: 1},
	"Missile Weapon: Sling":          {Value: 5, weight: 1},
	"Missile Weapon: Spear":          {Value: 5, weight: 1},
	"Missile Weapon: Throwing Axe":   {Value: 5, weight: 1},
	"Missile Weapon: Throwing Knife": {Value: 5, weight: 10},
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

// DefaultSkills is a map of default skills.
var DefaultSkills = map[string]Skill{
	"Appraise":          {Value: 15, weight: 1},
	"Art":               {Value: 5, weight: 1},
	"Athletics":         {Value: 10, weight: 1},
	"Bargain":           {Value: 5, weight: 1},
	"Brawl":             {Value: 25, weight: 1},
	"Bureaucracy":       {Value: 5, weight: 1},
	"Climb":             {Value: 40, weight: 1},
	"Command":           {Value: 5, weight: 1},
	"Computer Use":      {Value: 5, weight: 1},
	"Craft":             {Value: 5, weight: 1},
	"Demolition":        {Value: 1, weight: 1},
	"Disguise":          {Value: 5, weight: 1},
	"Dodge":             {Value: -1, weight: 10},
	"Drive":             {Value: 20, weight: 1},
	"Etiquette":         {Value: 5, weight: 1},
	"Fast Talk":         {Value: 5, weight: 1},
	"Fine Manipulation": {Value: 5, weight: 1},
	"First Aid":         {Value: 30, weight: 1},
	"Gaming":            {Value: 10, weight: 1},
	"Geology":           {Value: 0, weight: 1},
	"Grapple":           {Value: 25, weight: 1},
	"Heavy Machine":     {Value: 5, weight: 1},
	"Hide":              {Value: 10, weight: 1},
	"Insight":           {Value: 5, weight: 1},
	"Jump":              {Value: 25, weight: 1},
	"Language: Own":     {Value: -1, weight: 1},
	"Listen":            {Value: 25, weight: 1},
	"Medicine":          {Value: 5, weight: 1},
	"Navigate":          {Value: 10, weight: 1},
	"Perform":           {Value: 5, weight: 1},
	"Persuade":          {Value: 15, weight: 1},
	"Research":          {Value: 25, weight: 10},
	"Ride":              {Value: 5, weight: 1},
	"Science":           {Value: 1, weight: 1},
	"Sense":             {Value: 10, weight: 1},
	"Sleight of Hand":   {Value: 5, weight: 1},
	"Spot":              {Value: 25, weight: 10},
	"Status":            {Value: 15, weight: 1},
	"Stealth":           {Value: 10, weight: 1},
	"Strategy":          {Value: 5, weight: 1},
	"Swim":              {Value: 25, weight: 1},
	"Teach":             {Value: 10, weight: 1},
	"Technology Use":    {Value: 5, weight: 1},
	"Throw":             {Value: 25, weight: 1},
	"Track":             {Value: 10, weight: 1},
}

// ArtSkills is a map of skills.
var ArtSkills = map[string]Skill{
	"Art: Calligraphy": {Value: 5, weight: 1},
	"Art: Drawing":     {Value: 5, weight: 1},
	"Art: Other":       {Value: 5, weight: 1},
	"Art: Painting":    {Value: 5, weight: 1},
	"Art: Photography": {Value: 5, weight: 1},
	"Art: Sculpture":   {Value: 5, weight: 1},
	"Art: Writing":     {Value: 5, weight: 1},
}

// AthleticsSkills is a map of skills.
var AthleticsSkills = map[string]Skill{
	"Athletics: Acrobatics":                     {Value: 10, weight: 1},
	"Athletics: American and Canadian Football": {Value: 10, weight: 1},
	"Athletics: Baseball":                       {Value: 10, weight: 1},
	"Athletics: Basketball":                     {Value: 10, weight: 1},
	"Athletics: Bowling":                        {Value: 10, weight: 1},
	"Athletics: Cricket":                        {Value: 10, weight: 1},
	"Athletics: Cycling":                        {Value: 10, weight: 1},
	"Athletics: Golf":                           {Value: 10, weight: 1},
	"Athletics: Hockey":                         {Value: 10, weight: 1},
	"Athletics: Rugby":                          {Value: 10, weight: 1},
	"Athletics: Skating":                        {Value: 10, weight: 1},
	"Athletics: Skiing":                         {Value: 10, weight: 1},
	"Athletics: Soccer":                         {Value: 10, weight: 1},
	"Athletics: Tennis":                         {Value: 10, weight: 1},
	"Athletics: Track & Field":                  {Value: 10, weight: 1},
}

// ComputerUseSkills is a map of skills.
var ComputerUseSkills = map[string]Skill{
	"Computer Use: Art":         {Value: 5, weight: 1},
	"Computer Use: Design":      {Value: 5, weight: 1},
	"Computer Use: Gaming":      {Value: 5, weight: 1},
	"Computer Use: Hacking":     {Value: 5, weight: 1},
	"Computer Use: Magic":       {Value: 5, weight: 1},
	"Computer Use: Maintenance": {Value: 5, weight: 1},
	"Computer Use: Other":       {Value: 5, weight: 1},
	"Computer Use: Programming": {Value: 5, weight: 1},
	"Computer Use: Repair":      {Value: 5, weight: 1},
}

// CraftSkills is a map of skills.
var CraftSkills = map[string]Skill{
	"Craft: Carpentry":      {Value: 5, weight: 1},
	"Craft: Cooking":        {Value: 5, weight: 1},
	"Craft: Leatherworking": {Value: 5, weight: 1},
	"Craft: Pottery":        {Value: 5, weight: 1},
	"Craft: Sewing":         {Value: 5, weight: 1},
	"Craft: Woodworking":    {Value: 5, weight: 1},
}

// DriveSkills is a map of skills.
var DriveSkills = map[string]Skill{
	"Drive: Automobile":       {Value: 20, weight: 1},
	"Drive: Industrial Mover": {Value: 20, weight: 1},
	"Drive: Motorcycle":       {Value: 20, weight: 1},
	"Drive: Other":            {Value: 20, weight: 1},
	"Drive: Tank":             {Value: 20, weight: 1},
}

// HeavyMachineSkills is a map of skills.
var HeavyMachineSkills = map[string]Skill{
	"Heavy Machine: Boiler":    {Value: 5, weight: 1},
	"Heavy Machine: Bulldozer": {Value: 5, weight: 1},
	"Heavy Machine: Crane":     {Value: 5, weight: 1},
	"Heavy Machine: Engine":    {Value: 5, weight: 1},
	"Heavy Machine: Other":     {Value: 5, weight: 1},
	"Heavy Machine: Turbine":   {Value: 5, weight: 1},
	"Heavy Machine: Wrecker":   {Value: 5, weight: 1},
}

// KnowledgeSkills is a map of skills.
var KnowledgeSkills = map[string]Skill{
	"Knowledge: Accounting":      {Value: 10, weight: 1},
	"Knowledge: Anthropology":    {Value: 1, weight: 1},
	"Knowledge: Archaeology":     {Value: 1, weight: 1},
	"Knowledge: Art History":     {Value: 1, weight: 1},
	"Knowledge: Business":        {Value: 1, weight: 1},
	"Knowledge: Espionage":       {Value: 0, weight: 1},
	"Knowledge: Folklore":        {Value: 5, weight: 1},
	"Knowledge: Group":           {Value: 0, weight: 1},
	"Knowledge: History":         {Value: 20, weight: 1},
	"Knowledge: Law":             {Value: 5, weight: 1},
	"Knowledge: Linguistics":     {Value: 0, weight: 1},
	"Knowledge: Literature":      {Value: 5, weight: 1},
	"Knowledge: Natural History": {Value: 10, weight: 1},
	"Knowledge: Occult":          {Value: 5, weight: 1},
	"Knowledge: Philosophy":      {Value: 1, weight: 1},
	"Knowledge: Politics":        {Value: 5, weight: 1},
	"Knowledge: Region":          {Value: 0, weight: 1},
	"Knowledge: Religion":        {Value: 5, weight: 1},
	"Knowledge: Streetwise":      {Value: 5, weight: 1},
}

// MedicineSkills is a map of skills.
var MedicineSkills = map[string]Skill{
	"Medicine: Dermatology":       {Value: 5, weight: 1},
	"Medicine: Family Medicine":   {Value: 5, weight: 1},
	"Medicine: Immunology":        {Value: 5, weight: 1},
	"Medicine: Internal Medicine": {Value: 5, weight: 1},
	"Medicine: Neurology":         {Value: 5, weight: 1},
	"Medicine: Nuclear Medicine":  {Value: 5, weight: 1},
	"Medicine: Oncology":          {Value: 5, weight: 1},
	"Medicine: Other":             {Value: 5, weight: 1},
	"Medicine: Pathology":         {Value: 5, weight: 1},
	"Medicine: Pediatrics":        {Value: 5, weight: 1},
	"Medicine: Radiology":         {Value: 5, weight: 1},
	"Medicine: Surgery":           {Value: 5, weight: 1},
}

// PerformSkills is a map of skills.
var PerformSkills = map[string]Skill{
	"Perform: Act":             {Value: 5, weight: 1},
	"Perform: Dance":           {Value: 5, weight: 1},
	"Perform: Juggle":          {Value: 5, weight: 1},
	"Perform: Other":           {Value: 5, weight: 1},
	"Perform: Play Instrument": {Value: 5, weight: 1},
	"Perform: Recite":          {Value: 5, weight: 1},
	"Perform: Sing":            {Value: 5, weight: 1},
}

// PilotSkills is a map of skills.
var PilotSkills = map[string]Skill{
	"Pilot: Battleship":      {Value: 0, weight: 1},
	"Pilot: Helicopter":      {Value: 0, weight: 1},
	"Pilot: Hot Air Balloon": {Value: 0, weight: 1},
	"Pilot: Hovercraft":      {Value: 0, weight: 1},
	"Pilot: Jet Airliner":    {Value: 0, weight: 1},
	"Pilot: Jet Boat":        {Value: 0, weight: 1},
	"Pilot: Jet Fighter":     {Value: 0, weight: 1},
	"Pilot: Ocean Liner":     {Value: 0, weight: 1},
	"Pilot: Propeller Plane": {Value: 0, weight: 1},
}

// LanguageSkills is a map of skills.
var LanguageSkills = map[string]Skill{
	"Language: (various non-native)": {Value: 0, weight: 1},
	"Language: Own":                  {Value: -1, weight: 1},
}

// ScienceSkills is a map of skills.
var ScienceSkills = map[string]Skill{
	"Science: Astronomy":   {Value: 1, weight: 1},
	"Science: Biology":     {Value: 1, weight: 1},
	"Science: Botany":      {Value: 1, weight: 1},
	"Science: Chemistry":   {Value: 1, weight: 1},
	"Science: Cyptography": {Value: 1, weight: 1},
	"Science: Forensics":   {Value: 1, weight: 1},
	"Science: Genetics":    {Value: 1, weight: 1},
	"Science: Geology":     {Value: 1, weight: 1},
	"Science: Mathematics": {Value: 10, weight: 1},
	"Science: Meteorology": {Value: 1, weight: 1},
	"Science: Pharmacy":    {Value: 1, weight: 1},
	"Science: Physics":     {Value: 1, weight: 1},
	"Science: Planetology": {Value: 1, weight: 1},
	"Science: Psychology":  {Value: 5, weight: 1},
	"Science: Thaumaturgy": {Value: 0, weight: 1},
	"Science: Zoology":     {Value: 5, weight: 1},
}

// TechnologyUseSkills is a map of skills.
var TechnologyUseSkills = map[string]Skill{
	"Technology Use: Communications":      {Value: 5, weight: 1},
	"Technology Use: Electronic Security": {Value: 5, weight: 1},
	"Technology Use: Electronics":         {Value: 5, weight: 1},
	"Technology Use: Other":               {Value: 5, weight: 1},
	"Technology Use: Sensor Systems":      {Value: 5, weight: 1},
	"Technology Use: Surveillance":        {Value: 5, weight: 1},
	"Technology Use: Traps":               {Value: 5, weight: 1},
}

// EsotericSkills is a map of the creepy stuff.
var EsotericSkills = map[string]Skill{
	"Cthulhu Mythos": {Value: 0, weight: 1},
	"Sorcery":        {Value: 0, weight: 1},
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
