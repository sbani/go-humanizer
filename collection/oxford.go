package collection

import (
    "strings"
    "strconv"
)

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
    
    if limit != -1 {
        return formatCommaSeparatedWithLimit(collection, limit, len)
    }
    
    return formatCommaSeparated(collection, len)
}

func formatOnlyTwo(collection []string) string {
    return collection[0] + " and " + collection[1]
}

func formatCommaSeparatedWithLimit(collection []string, limit int, count int) string {
    display := strings.Join(collection[:limit], ", ")
    moreCount := count - limit
    return display + " and " + strconv.Itoa(moreCount) + " more"
}

func formatCommaSeparated(collection []string, count int) string {
    display := strings.Join(collection[:(count - 1)], ", ")
    return display + " and " + collection[(count - 1)]
}
