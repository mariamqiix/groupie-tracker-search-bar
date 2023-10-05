package groupie

import (
	"fmt"
	"slices"
	"strings"
)

func Search(s string) ([]Artist, string) {
	TheArtists := artists
	var artists []Artist
	x := 0
	n, cond := RetutrnCond(s)
	n = strings.ToLower(n)
	Inlocation := searchByLocation2(TheLocations.Index, n)
	for _, artist := range TheArtists {
		if cond == "dates" {
			if checkTheDates(TheDates.Index2[x].TheData, n) {
				artists = append(artists, artist)
			}
		} else if cond == "Member" {
			if chceckarr(artist.Members, n) {
				artists = append(artists, artist)
			}
		} else if cond == "Creation" {
			if chceckInter(artist.CreationDate, n) {
				artists = append(artists, artist)
			}
		} else if cond == "Location" {
			if slices.Contains(Inlocation, x) {
				artists = append(artists, artist)
			}
		} else if cond == "brand" {
			if strings.Contains(strings.ToLower(artist.Name), n) {
				artists = append(artists, artist)
			}
		} else if cond == "Album" {
			if strings.Contains(strings.ToLower(artist.FirstAlbum), n){
				artists = append(artists, artist)
			}
		} else {
			if strings.Contains(strings.ToLower(artist.Name), n) || chceckarr(artist.Members, n) || chceckInter(artist.CreationDate, n) || slices.Contains(Inlocation, x) || strings.Contains(strings.ToLower(artist.FirstAlbum), n) || checkTheDates(TheDates.Index2[x].TheData, n) {
				artists = append(artists, artist)
			}
		}
		x++
	}
	return artists, ""
}

func con(s, n string) bool {
	for i := 0; i < len(s); i++ {
		if len(n) > 0 && s[i] == n[0] {
			i++
			if i+len(n) < len(s) {
				for d := 1; d < len(n); d++ {
					if s[i] == n[d] {
						if i == len(n)-1 {
							return true
						}
						i++
					}
				}
			}
		}
	}
	return false
}

func chceckarr(x []string, n string) bool {
	if slices.Contains(x, n) {
		return true
	}
	for i := 0; i < len(x); i++ {
		if strings.Contains(strings.ToLower(x[i]), n) {
			return true
		}
	}
	return false
}

func checkTheDates(x []string, n string) bool {
	for i := 0; i < len(x); i++ {
		if strings.Contains(strings.ToLower(x[i]), n) {
			return true
		}
	}
	return false
}

func chceckInter(x interface{}, n string) bool {
	j := fmt.Sprintf("%v", x) // Print the interface data
	if j == n {
		return true
	}
	return false
}

func searchByLocation2(x []Index, n string) []int {
	var indx []int
	for i := 0; i < len(x); i++ {
		flag2 := false
		for j := 0; j < len(x[i].TheData); j++ {
			if strings.Contains(strings.ToLower(x[i].TheData[j]), n) && !flag2 {
				indx = append(indx, i)
				flag2 = true
			}
		}
	}
	return indx

}

func RetutrnCond(s string) (string, string) {
	cond := ""
	if strings.Contains(s, " - Date") {
		cond = "dates"
		s = s[: len(s)-7]
	} else if strings.Contains(s, " - Member") {
		cond = "Member"
		s = s[: len(s)-9]
	} else if strings.Contains(s, " - Creation Date") {
		cond = "Creation"
		s = s[: len(s)-16]
	} else if strings.Contains(s, " - Location") {
		cond = "Location"
		s = s[: len(s)-11]
	} else if strings.Contains(s, " - brand Name") {
		cond = "brand"
		s = s[: len(s)-14]
	} else if strings.Contains(s, " - First Album") {
		cond = "Album"
		s = s[: len(s)-14]
	} else {
		cond = "All"
	}
	return s, cond
}
