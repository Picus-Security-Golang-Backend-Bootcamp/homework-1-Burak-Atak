package contains

import "strings"

func Contains(list []string, element string) (map[string]int, bool) {
	//var books []string
	books := map[string]int{}
	for i, v := range list {
		if strings.Contains(v, element) {
			//books = append(books, v)
			books[v] = i + 1
		}
	}

	if len(books) == 0 {
		return books, false
	} else {
		return books, true
	}

}
