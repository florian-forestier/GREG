package service

import(
	"strconv"
)

func ParseCache() string {
	
	final := ""
	
	for _, key := range cache {
		final = final + key.Name + " " + strconv.FormatFloat(key.Value, 'f', -1, 32) + "\n"
	}
	
	Reset()
	
	return final
}