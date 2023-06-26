package gin

import (
	"log"
	"strconv"
)

// removeFastByIndex does not perform bounds-checking. It expects a valid index as input. This means that negative
// values or indices that are greater or equal to the initial len(s) will cause Go to panic
func removeFastByIndex(albums []album, index int) []album {
	// return a slice containing the structs before and after the index
	return append(albums[:index], albums[index+1:]...)
}

func stringToInt(iString string) (int, error) {
	iInt, err := strconv.Atoi(iString)
	if err != nil {
		log.Println("Error during conversion")
		return 0, err
	} else {
		return iInt, nil
	}
}
