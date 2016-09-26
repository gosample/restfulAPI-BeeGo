package models

func removeDuplicatesUnordered(elements []string) []string {
    encountered := map[string]bool{}
    // Create a map of all unique elements.
    for v:= range elements {
	encountered[elements[v]] = true
    }
    result := []string{}
    for key, _ := range encountered {
	result = append(result, key)
    }
    return result
}
