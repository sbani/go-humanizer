package collection

import (
	"strconv"
	"strings"
)

// Oxford returns a string from a string collection by adding ", " as a seperator and the word "and" for the last seperator.
// A limit will print all elements to the limit and elements after the limit will be grouped with a "and n more".
// No limit is represented with with a limit <= 0
func Oxford(collection []string, limit int) string {
	len := len(collection)

	switch len {
	case 0:
		return ""
	case 1:
		return collection[0]
	case 2:
		return formatOnlyTwo(collection)
	}

	if limit > 0 {
		return formatCommaSeparatedWithLimit(collection, limit, len)
	}

	return formatCommaSeparated(collection, len)
}

// Format a collection of two strings
func formatOnlyTwo(collection []string) string {
	return collection[0] + " and " + collection[1]
}

// Format with limit
func formatCommaSeparatedWithLimit(collection []string, limit int, count int) string {
	display := strings.Join(collection[:limit], ", ")
	moreCount := count - limit
	return display + ", and " + strconv.Itoa(moreCount) + " more"
}

// Format without limit
func formatCommaSeparated(collection []string, count int) string {
	display := strings.Join(collection[:(count-1)], ", ")
	return display + ", and " + collection[(count-1)]
}
