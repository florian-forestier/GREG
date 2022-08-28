package service

import (
	"bufio"
	"fmt"
)

func ParseCache() string {

	buffer := bufio.Writer{}

	final := ""

	for _, key := range cache {
		_, _ = buffer.WriteString(fmt.Sprintf("%s %f\n", key.Name, key.Value))
	}

	Reset()

	return final
}
