// Helper functions for tlfgen.

package tlfgen

import (
	"encoding/binary"
	"encoding/hex"
	"hash/fnv"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"

	randomdata "github.com/Pallinder/go-randomdata"
	log "github.com/sirupsen/logrus"
)

// RAND is the randomizer.
var RAND = rand.New(rand.NewSource(time.Now().UnixNano()))

var dataDir = setDataDir()

func setDataDir() string {
	dir := os.Getenv("GOPATH") + "/src/github.com/gruevyhat/tlfgen/assets/"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}
	return dir
}

func readJSON(filename string) []byte {
	raw, _ := ioutil.ReadFile(filename)
	return raw
}

func arrayContains(arr []string, s string) bool {
	for _, a := range arr {
		if a == s {
			return true
		}
	}
	return false
}

func arrayRemoveInt(a []int, i int) []int {
	a = append(a[:i], a[i+1:]...)
	return a
}

func arrayRemoveString(a []string, i int) []string {
	a = append(a[:i], a[i+1:]...)
	return a
}

func arraySum(a []int) int {
	sum := 0
	for _, i := range a {
		sum += i
	}
	return sum
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func setSeed(charHash string) (string, error) {
	if charHash == "" {
		defaultSeed := time.Now().UTC().UnixNano()
		charHash = strconv.FormatInt(defaultSeed, 16)
	}
	h, err := hex.DecodeString(charHash)
	if err != nil {
		return charHash, err
	}
	seed := binary.BigEndian.Uint64(h)
	RAND.Seed(int64(seed))
	log.Info("Set new seed:", seed)
	return charHash, nil
}

func randomName(gender string) string {
	randomdata.CustomRand(RAND)
	var name string
	switch gender {
	case "Male":
		name = randomdata.FullName(randomdata.Male)
	case "Female":
		name = randomdata.FullName(randomdata.Female)
	}
	return name
}

func sampleWithoutReplacement(choices []string, n int) []string {
	samples := []string{}
	idxs := RAND.Perm(len(choices))
	for i := 0; i < n; i++ {
		samples = append(samples, choices[idxs[i]])
	}
	return samples
}

func randomChoice(choices []string) string {
	n := len(choices)
	if n > 0 {
		r := RAND.Intn(n)
		return choices[r]
	}
	return ""
}

func randomInt(max int) int {
	// Returns an int in [1..max].
	return RAND.Intn(max) + 1
}

func weightedRandomChoice(choices []string, weights []int) string {
	sum := 0
	for _, w := range weights {
		sum += w
	}
	if sum > 0 {
		r := randomInt(sum)
		total := 0
		for i, w := range weights {
			total += w
			if r <= total {
				return choices[i]
			}
		}
	} else {
		log.Error("Skill weights sum to zero!")
	}
	return ""
}

// Die represents a single die of the form <code>D6+<pips>.
type Die struct {
	sides int
	code  int
}

func (d Die) roll() (n int) {
	n = 0
	for i := 0; i < d.code; i++ {
		n += randomInt(d.sides)
	}
	return n
}

func (d Die) toStr() (dieStr string) {
	dieStr = strconv.Itoa(d.code) + "D" + strconv.Itoa(d.sides)
	return dieStr
}
