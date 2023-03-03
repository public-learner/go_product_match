package main

type TypeAndValues struct {
	typename string
	values   []string
}
type Match []string
type Variant []string

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getProducts(existing []TypeAndValues, updating []TypeAndValues, _variants []Variant) []Match {
	var result []Match
	updatedExisting := existing
	variants := _variants

	for _, typeAndValues := range updating {
		existingTypes := make([]string, len(updatedExisting))
		for i, typeAndValues := range updatedExisting {
			existingTypes[i] = typeAndValues.typename
		}
		typeName := typeAndValues.typename
		matchingIndex := -1
		for i, v := range existingTypes {
			if v == typeName {
				matchingIndex = i
				break
			}
		}

		if matchingIndex != -1 && typeAndValues.values != nil { // update
			afterValues := typeAndValues.values
			beforeValues := updatedExisting[matchingIndex].values
			difference := make([]string, 0, len(afterValues))
			for _, x := range afterValues {
				if !contains(beforeValues, x) {
					difference = append(difference, x)
				}
			}
			removalValues := make([]string, 0, len(beforeValues))
			for _, removalValue := range beforeValues {
				if !contains(afterValues, removalValue) {
					removalValues = append(removalValues, removalValue)
				}
			}
			if len(removalValues) > 0 {
				for _, removalValue := range removalValues {
					var newVariants []Variant
					for _, variant := range variants {
						if !contains(variant, removalValue) {
							newVariants = append(newVariants, variant)
						}
					}
					variants = newVariants
				}
			}
			if len(difference) > 0 {
				variants = append(variants, difference)
				updatedExisting[matchingIndex] = typeAndValues
			}
		} else if matchingIndex != -1 && typeAndValues.values == nil { // delete
			values := updatedExisting[matchingIndex].values
			for i, variant := range variants {
				var newVariant []string
				for _, variantValue := range variant {
					if !contains(values, variantValue) {
						newVariant = append(newVariant, variantValue)
					}
				}
				variants[i] = newVariant
			}
			updatedExisting = append(updatedExisting[:matchingIndex], updatedExisting[matchingIndex+1:]...)
		} else if typeAndValues.values != nil { // add
			updatedExisting = append(updatedExisting, typeAndValues)
		}
	}
	isVariant := func(match []string) bool {
		if len(variants) == 0 {
			return true
		}
		for _, variant := range variants {
			allMatched := true
			for _, v := range variant {
				if !contains(match, v) {
					allMatched = false
					break
				}
			}
			if allMatched {
				return true
			}
		}
		return false
	}

	var generateMatches func(index int, matchSoFar []string)

	generateMatches = func(index int, matchSoFar []string) {
		if index == len(updatedExisting) {
			if isVariant(matchSoFar) {
				result = append(result, matchSoFar)
			}
			return
		}
		for _, added := range updatedExisting[index].values {
			generateMatches(index+1, append(matchSoFar, added))
		}
	}

	generateMatches(0, make([]string, 0))
	return result
}

func main() {

}
