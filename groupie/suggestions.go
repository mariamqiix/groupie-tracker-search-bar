package groupie

import (
	"fmt"
	"slices"
	"strings"
)

func suggestions(s string) []string {
	var suggestions []string
	x := 0
	s = strings.ToLower(s)
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), s) {
			suggestions = append(suggestions, artist.Name+" - brand Name")
		}
		if strings.Contains(strings.ToLower(artist.FirstAlbum), s) {
			suggestions = append(suggestions, artist.FirstAlbum+" - First Album")
		}
		suggestions = AppendCreationDate(suggestions, artist.CreationDate, s)
		suggestions = AppendTheMembers(suggestions, artist.Members, s)
		suggestions = AppendTheDates(suggestions, TheDates.Index2[x].TheData, s)
		suggestions = AppendLocation(suggestions, TheLocations.Index, s)
		x++
	}
	return suggestions
}

func AppendTheDates(v, x []string, n string) []string {
	for i := 0; i < len(x); i++ {
		if con(strings.ToLower(x[i]), n){
			v = append(v, x[i]+" - Date")
		}
	}
	return v
}

func AppendTheMembers(v, x []string, n string) []string {
	for j := 0; j < len(x); j++ {
		if strings.Contains(strings.ToLower(x[j]), n) {
			if !slices.Contains(v, x[j]+" - Member") {
				v = append(v, x[j]+" - Member")
			}
		}
	}
	return v
}

func AppendCreationDate(v []string, x interface{}, n string) []string {
	j := fmt.Sprintf("%v", x) // Print the interface data
	if j == n {
		v = append(v, j+" - Creation Date")
	}
	return v
}

func AppendLocation(v []string, x []Index, n string) []string {
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[i].TheData); j++ {
			if strings.Contains(strings.ToLower(x[i].TheData[j]), (n)) == true {
				if !slices.Contains(v, x[i].TheData[j]+" - Location") {
					v = append(v, x[i].TheData[j]+" - Location")
				}
			}
		}
	}
	return v

}
