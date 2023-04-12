package ufc

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: don't panic and instead return an error
func timeStringToSeconds(s string) int {
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		panic(fmt.Sprintf("bad time string %q", s))
	}

	minutes, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(fmt.Sprintf("bad minutes %q", parts[0]))
	}
	seconds, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(fmt.Sprintf("bad seconds %q", parts[1]))
	}

	return (minutes * 60) + seconds
}
