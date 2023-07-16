/*
 * @Author: liziwei01
 * @Date: 2023-07-13 17:57:13
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-13 19:42:47
 * @Description: file content
 */
package minwindow

func minWindow(s string, t string) string {
	tMap := make(map[byte]int)
	for i := 0; i != len(t); i++ {
		tMap[t[i]]++
	}

	mMap := copyMap(tMap)          // temp map
	lo, hi := 0, 0
	finalLo, finalHi := 0, 0
	winLength := 0
	for hi < len(s) {
		// 如果前面的str对window长度没有帮助，则跳过
		if winLength == 0 && mMap[s[hi]] == 0 {
			lo++
			hi = lo
			continue
		}
		// 如果找全了，则记录这个window并去掉第一位数
		for foundWindow(mMap) {
			if (finalHi - finalLo) == 0 {
				finalLo = lo
				finalHi = hi
			}
			if (hi-lo) < (finalHi-finalLo) {
				finalLo = lo
				finalHi = hi
			}
			mMap[s[lo]]++
			lo++
			winLength--
			for lo <= hi && lo < len(s) && tMap[s[lo]] == 0 {
				lo++
				winLength--
			}
		}
		// 如果找到一个可以用作window的字母，则删去
		mMap[s[hi]]--
		// 走到这里说明winLength不为0，extend window
		hi++
		winLength++
	}

	for foundWindow(mMap) {
		if (finalHi - finalLo) == 0 {
			finalLo = lo
			finalHi = hi
		}
		if (hi-lo) < (finalHi-finalLo) {
			finalLo = lo
			finalHi = hi
		}
		mMap[s[lo]]++
		lo++
		for lo <= hi && lo < len(s) && tMap[s[lo]] == 0 {
			lo++
		}
	}

	return s[finalLo:finalHi]
}

func foundWindow[K comparable, V int](m map[K]V) bool {
	for k := range m {
		if m[k] > 0 {
			return false
		}
	}
	return true
}

func copyMap[K, V comparable](m map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		result[k] = v
	}
	return result
}
