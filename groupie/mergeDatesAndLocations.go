package groupie

import (
	"sort"
	"strings"
)

func MergeDatesAndLocations(x int) ResponseData3 {
	start := 0
	No := 0
	var relation ResponseData3
	sort.Strings(TheLocations.Index[x].TheData)
	for i := 0; i < len(TheLocations.Index[x].TheData); i++ {
		Dates := ""
		start, Dates = ReturnTheDates(x, No)
		No = start
		Index3 := Index3{
			Location: TheLocations.Index[x].TheData[i],
			TheData:  strings.ReplaceAll(Dates, "*", ""),
		}
		relation.Index3 = append(relation.Index3, Index3)
	}
	return relation
}

func ReturnTheDates(x, start int) (int, string) {
	Dates := ""
	for n := start; n < len(TheDates.Index2[x].TheData); n++ {
		if rune(TheDates.Index2[x].TheData[n][0]) == '*' && n != start {
			return n, Dates[:len(Dates)-2]
		} else {
			Dates += TheDates.Index2[x].TheData[n] +", "
		}
	}
	return 0, Dates[:len(Dates)-2]
}
