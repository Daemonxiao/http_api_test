package compare

// ContainsMap
// checks if map1 contains all the key-value pairs of map2 recursively.
func ContainsMap(map1, map2 map[string]interface{}) bool {
	for key, val2 := range map2 {
		val1, ok := map1[key]
		if !ok {
			return false
		}

		switch v2 := val2.(type) {
		case map[string]interface{}:
			v1, ok := val1.(map[string]interface{})
			if !ok {
				return false
			}
			if !ContainsMap(v1, v2) {
				return false
			}
		case []interface{}:
			v1, ok := val1.([]interface{})
			if !ok {
				return false
			}
			if !containsSlice(v1, v2) {
				return false
			}
		default:
			if val1 != val2 {
				return false
			}
		}
	}

	return true
}

func containsSlice(slice1 []interface{}, slice2 []interface{}) bool {
	for i := range slice2 {
		for j := range slice1 {
			isContain := false
			switch v2 := slice2[i].(type) {
			case map[string]interface{}:
				v1, ok := slice1[j].(map[string]interface{})
				if !ok {
					if j == len(slice1)-1 {
						return false
					} else {
						continue
					}
				}
				if !ContainsMap(v1, v2) {
					if j == len(slice1)-1 {
						return false
					} else {
						continue
					}
				} else {
					isContain = true
				}
			case []interface{}:
				v1, ok := slice1[j].([]interface{})
				if !ok {
					if j == len(slice1)-1 {
						return false
					} else {
						continue
					}
				}
				if !containsSlice(v1, v2) {
					if j == len(slice1)-1 {
						return false
					} else {
						continue
					}
				} else {
					isContain = true
				}
			}
			if isContain {
				break
			}
		}
	}
	return true
}
