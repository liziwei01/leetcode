/*
 * @Author: liziwei01
 * @Date: 2023-07-10 03:00:23
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-10 03:38:03
 * @Description: file content
 */
package simplifypath

func simplifyPath(path string) string {
	path = rmSuffixSlash(path)
	path = rmSuccessiveSlash(path)
	path = rmSinglePeriod(path)
	path = rmDoublePeriod(path)
	return path
}

// 去除后缀斜杠
func rmSuffixSlash(path string) string {
	hasSlash := false
	if len(path) <= 1 {
		return "/"
	}
	if path[len(path)-1] == '/' {
		hasSlash = true
		path = path[:len(path)-1]
	}
	if !hasSlash {
		return path
	}
	return rmSuffixSlash(path)
}

// 去除连续斜杠
func rmSuccessiveSlash(path string) string {
	hasSlash := false
	p := ""
	for i := 0; i != len(path); i++ {
		if !hasSlash && path[i] == '/' {
			hasSlash = true
			p += "/"
		} else if hasSlash && path[i] == '/' {
			// pass
		} else {
			hasSlash = false
			p += path[i : i+1]
		}
	}
	return p
}

func rmDoublePeriod(path string) string {
	hasDoublePeriod := false
	upLevelIdx := 0
	if len(path) <= 1 {
		return "/"
	}
	for i := 0; i != len(path)-2; i++ {
		if path[i:i+3] == "/.." {
			if i+4 <= len(path) && path[i+3] != '/' {
				continue
			}
			if i == 0 {
				if len(path) <= 4 {
					return "/"
				}
				return rmDoublePeriod(path[3:])
			}
			path = path[:upLevelIdx] + path[i+3:]
			hasDoublePeriod = true
			break
		}
		if path[i] == '/' {
			upLevelIdx = i
		}
	}

	if !hasDoublePeriod {
		return path
	}
	return rmDoublePeriod(path)
}

func rmSinglePeriod(path string) string {
	hasSinglePeriod := false
	if len(path) <= 1 {
		return "/"
	}
	for i := 0; i != len(path)-2; i++ {
		if path[i:i+3] == "/./" {
			path = path[:i] + path[i+2:]
			hasSinglePeriod = true
			break
		}
	}

	if path[len(path)-2:] == "/." {
		return rmSinglePeriod(path[:len(path)-2])
	}
	if !hasSinglePeriod {
		return path
	}
	return rmSinglePeriod(path)
}
