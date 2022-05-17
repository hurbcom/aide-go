package aidego

import (
	"math"
	"math/rand"
	"time"

	"github.com/fatih/structs"
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
			_ = structs.New(dest).Field(key).Set(val)
		}
	}
}
