package match

import (
	"fmt"
	"math"

	"github.com/antzucaro/matchr"
)

// Match test for matching strings
func Match() {
	
		string1 := "Family"
		string2 := "gsnuly222"
		
		fmt.Println(MatchStr(string1, string2))
}

// MatchStr export
func MatchStr(s1, s2 string) float64 {
	r := matchr.Jaro(s1, s2)
	roundMatch := math.Round(r * 100)/100
	return roundMatch
}