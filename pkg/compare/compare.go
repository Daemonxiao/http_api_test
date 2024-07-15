package compare

func ContainsMap(map1, map2 map[string]interface{}) bool {
	for key, val2 := range map2 {
		val1, ok := map1[key]
		if !ok {
			return false // map1中缺少map2中的键
		}

		// 比较值的类型
		switch v2 := val2.(type) {
		case map[string]interface{}:
			v1, ok := val1.(map[string]interface{})
			if !ok {
				return false // 类型不匹配
			}
			if !ContainsMap(v1, v2) {
				return false // 递归比较结果不包含
			}
		case []interface{}:
			v1, ok := val1.([]interface{})
			if !ok {
				return false // 类型不匹配
			}
			if !containsSlice(v1, v2) {
				return false // 数组元素不包含
			}
		default:
			if val1 != val2 {
				return false // 值不相等
			}
		}
	}

	return true
}

// containsSlice 函数比较两个切片
func containsSlice(slice1 []interface{}, slice2 []interface{}) bool {
	for i := range slice2 {
		for j := range slice1 {
			isContain := false
			switch v2 := slice2[i].(type) {
			case map[string]interface{}:
				v1, ok := slice1[j].(map[string]interface{})
				if !ok {
					if j == len(slice1)-1 {
						return false // 类型不匹配
					} else {
						continue
					}
				}
				if !ContainsMap(v1, v2) {
					if j == len(slice1)-1 {
						return false // 递归比较结果不包含
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
						return false // 类型不匹配
					} else {
						continue
					}
				}
				if !containsSlice(v1, v2) {
					if j == len(slice1)-1 {
						return false // 递归比较结果不包含
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
