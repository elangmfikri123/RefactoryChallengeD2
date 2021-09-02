package main

import "fmt"

func removeDuplicate(kata []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range kata {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {
	kata := []string{" aku", "kamu", "kita", "kita", "akan", "milik", "dia", "milik", "akan"}
	fmt.Println(kata)
	removeDuplicateValuesSlice := removeDuplicate(kata)
	fmt.Println(removeDuplicateValuesSlice)
}
