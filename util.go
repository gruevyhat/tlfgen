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
	"strings"
	"time"
)

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
		if a == s || strings.HasPrefix(a, s) || strings.HasSuffix(s, a) {
			return true
		}
	}
	return false
}

func arrayRemove(s string, a []string) []string {
	for i, x := range a {
		if x == "" || x == s {
			a = append(a[:i], a[i+1:]...)
		}
	}
	return a
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
	rand.Seed(int64(seed))
	log.Info("Set new seed:", seed)
	return charHash, nil
}

func sampleWithoutReplacement(choices []string, n int) []string {
	samples := []string{}
	idxs := rand.Perm(len(choices))
	for i := 0; i < n; i++ {
		samples = append(samples, choices[idxs[i]])
	}
	return samples
}

func randomChoice(choices []string) string {
	n := len(choices)
	if n > 0 {
		r := rand.Intn(n)
		return choices[r]
	}
	return ""
}

func randomInt(max int) int {
	// Returns an int in [1..max].
	return rand.Intn(max) + 1
}

func weightedRandomChoice(choices []string, weights []int) string {
	sum := 0
	for _, w := range weights {
		sum += w
	}
	r := randomInt(sum)
	total := 0
	for i, w := range weights {
		total += w
		if r <= total {
			return choices[i]
		}
	}
	return choices[0]
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
