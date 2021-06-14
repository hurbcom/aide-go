package v4

import (
	"math"
	"math/rand"
	"regexp"
	"time"

	"github.com/fatih/structs"
)

var (
	regexpCommaAlphaNum *regexp.Regexp = regexp.MustCompile(
		`[^A-Za-z0-9,]`)
)

func Round(value float64, precision int) float64 {
	exponential := math.Pow10(precision)
	return math.Ceil(value*exponential) / exponential
}

func RandomInt(bottom, top int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(top-bottom) + bottom
}

// Fill merges data from struct instance to another
// By @titpetric suggested in https://scene-si.org/2016/06/01/golang-tips-and-tricks
func Fill(dest interface{}, src interface{}) {
	mSrc := structs.Map(src)
	mDest := structs.Map(dest)
	for key, val := range mSrc {
		if _, ok := mDest[key]; ok {
			structs.New(dest).Field(key).Set(val)
		}
	}
}
