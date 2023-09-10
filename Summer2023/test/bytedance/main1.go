package bytedance

import (
	"fmt"
)

func main1() {
	s := ""
	t := ""
	fmt.Scan(&s, &t)
	// plen := prefixLen(s, t)
	// slen := suffixLen(s, t)

	slen := len(s)
	tlen := len(t)
	mlen := min(slen, tlen)

	prefixLen := prefixLength(s, t, slen, tlen, mlen)
	suffixLen := suffixLength(s, t, slen, tlen, mlen)

	prefixLenShorter := false
	if prefixLen < suffixLen {
		prefixLenShorter = true
	}
	similarity := 0

	// 如果s没有碰头 abc...qwe abc....qwe 3*3
	if prefixLen+suffixLen < slen-1 {
		// 长度相同 无所谓

		// 长度不同 给小的+1
		if prefixLenShorter {
			similarity = (prefixLen + 1) * suffixLen
		} else {
			similarity = prefixLen * (suffixLen + 1)
		}
		fmt.Printf("%d\n", similarity)
		return
	}

	// 如果s差一碰头 abc.qwe abc..qwe 3*3
	if prefixLen+suffixLen == slen-1 {
		// 长度相同 abc.qwe abc.qwe
		if slen == tlen {
			similarity = slen * slen
			fmt.Printf("%d\n", similarity)
			return
		}

		// 长度不同 s长awcc.ccwer awccwer 4*5 => awccwccwer awccwer 5*6
		if slen > tlen {
			news1 := s[:prefixLen] + t[prefixLen:prefixLen+1] + s[slen-suffixLen:]
			prefixLen1 := prefixLength(news1, t, slen, tlen, mlen)
			suffixLen1 := suffixLength(news1, t, slen, tlen, mlen)
			news2 := s[:prefixLen] + t[slen-suffixLen-1:slen-suffixLen] + s[slen-suffixLen:]
			prefixLen2 := prefixLength(news2, t, slen, tlen, mlen)
			suffixLen2 := suffixLength(news2, t, slen, tlen, mlen)
			if prefixLen1*suffixLen1 > prefixLen2*suffixLen2 {
				similarity = prefixLen1 * suffixLen1
			} else {
				similarity = prefixLen2 * suffixLen2
			}
			fmt.Printf("%d\n", similarity)
			return
		} else {
			// 长度不同 s短 abc.qwer abc..qwer => abcdqwer abcd.qwer
			if prefixLenShorter {
				similarity = (prefixLen + 1) * suffixLen
			} else {
				similarity = prefixLen * (suffixLen + 1)
			}
			fmt.Printf("%d\n", similarity)
			return
		}

	}

	// 如果s正好碰头 abcqwe abc.qwe 3*3
	if prefixLen+suffixLen == slen {
		// 长度相同 不可能

		// 长度不同s长 abczzt abczt 4*2 => abcczt abczt 3*3
		if slen > tlen {
			// // 取长补短
			// if prefixLenShorter {
			//     similarity = (prefixLen+1) * (suffixLen-1)
			// } else if prefixLen == suffixLen {
			//     // aczbztz acztz 3*3 => acztztz acztz 5*3
			//     similarity = prefixLen * suffixLen
			// }
			news1 := s[:prefixLen] + t[prefixLen:prefixLen+1] + s[slen-suffixLen:]
			prefixLen1 := prefixLength(news1, t, slen, tlen, mlen)
			suffixLen1 := suffixLength(news1, t, slen, tlen, mlen)
			news2 := s[:prefixLen] + t[slen-suffixLen-1:slen-suffixLen] + s[slen-suffixLen:]
			prefixLen2 := prefixLength(news2, t, slen, tlen, mlen)
			suffixLen2 := suffixLength(news2, t, slen, tlen, mlen)
			if prefixLen1*suffixLen1 > prefixLen2*suffixLen2 {
				similarity = prefixLen1 * suffixLen1
			} else {
				similarity = prefixLen2 * suffixLen2
			}
			fmt.Printf("%d\n", similarity)
			return
		} else {

			// 长度不同s短 abcqwert abc..qwert 3*5 => abcdwert abcd.wert 4*4

			if prefixLenShorter {
				similarity = (prefixLen + 1) * (suffixLen - 1)
			} else if prefixLen == suffixLen {
				similarity = prefixLen * suffixLen
			} else {
				similarity = (prefixLen - 1) * (suffixLen + 1)
			}
			fmt.Printf("%d\n", similarity)
			return
		}
	}

	// 如果s重叠 abcdddefg abcddd......dddefg
	if prefixLen+suffixLen > slen {
		// 长度相同 abcdddefg abcdddefg
		if slen == tlen {
			similarity = slen * slen
			fmt.Printf("%d\n", similarity)
			return
		}

		// 长度不同 abcddddefg abcdddefg
		similarity = prefixLen * suffixLen
		fmt.Printf("%d\n", similarity)
		return
	}

}

func prefixLength(s, t string, slen, tlen, mlen int) int {
	for i := 0; i != mlen; i++ {
		if s[i] == t[i] {
			continue
		} else {
			return i
		}
	}
	return mlen
}

func suffixLength(s, t string, slen, tlen, mlen int) int {
	for i := 0; i != mlen; i++ {
		if s[slen-1-i] == t[tlen-1-i] {
			continue
		} else {
			return i
		}
	}
	return mlen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
