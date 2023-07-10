/*
 * @Author: liziwei01
 * @Date: 2023-07-09 13:01:34
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-09 14:13:43
 * @Description: file content
 */
package fulljustify

type line struct {
	words     []string
	spaces    []int // appended spaces after each word
	n         int   // num of words
	tempWidth int   // temp width
}

func fullJustify(words []string, maxWidth int) []string {
	lines := typesetting(words, maxWidth)
	return addSpaces(lines, maxWidth)
}

func typesetting(words []string, maxWidth int) []line {
	tempWidth := 0
	lines := make([]line, 0)
	lineIdx := 0
	for i := 0; i != len(words); i++ {
		if tempWidth == 0 {
			oneLine := line{
				words:  []string{words[i]},
				spaces: make([]int, 0),
			}
			oneLine.spaces = append(oneLine.spaces, 1)
			lines = append(lines, oneLine)
			tempWidth += len(words[i])
			lines[lineIdx].n = 1
			lines[lineIdx].tempWidth = tempWidth
		} else if tempWidth+1+len(words[i]) <= maxWidth {
			lines[lineIdx].words = append(lines[lineIdx].words, words[i])
			// 每个单词后面应至少有一个空格
			lines[lineIdx].spaces = append(lines[lineIdx].spaces, 1)
			lines[lineIdx].n++
			tempWidth = tempWidth + 1 + len(words[i])
			continue
		} else {
			// 左右对齐最右边应当无空格
			lines[lineIdx].spaces[lines[lineIdx].n-1] = 0
			lines[lineIdx].tempWidth = tempWidth
			tempWidth = 0
			lineIdx++
			i--
		}
	}

	return lines
}

func addSpaces(lines []line, maxWidth int) []string {
	article := make([]string, len(lines))
	lastLineIdx := len(lines) - 1
	for i := 0; i != lastLineIdx; i++ {
		if lines[i].n == 1 {
			article[i] = lines[i].AlignLeft(maxWidth)
		} else {
			article[i] = lines[i].AlignLeftandRight(maxWidth)
		}
	}
	article[lastLineIdx] = lines[lastLineIdx].AlignLeft(maxWidth)
	return article
}

// 左右对齐
func (l *line) AlignLeftandRight(maxWidth int) string {
	// 单词数减一为每行初始单词间的空格数
	for l.tempWidth < maxWidth {
		for i := 0; i != l.n-1; i++ {
			if l.tempWidth < maxWidth {
				l.spaces[i]++
				l.tempWidth++
			} else {
				break
			}
		}
	}

	return l.String()
}

// 左对齐
func (l *line) AlignLeft(maxWidth int) string {
	// 补满空格
	for l.tempWidth < maxWidth {
		l.spaces[l.n-1]++
		l.tempWidth++
	}

	return l.String()[0:maxWidth]
}

func (l *line) String() string {
	str := ""
	for i := 0; i != l.n; i++ {
		str = str + l.words[i] + spaces(l.spaces[i])
	}
	return str
}

func spaces(n int) string {
	str := ""
	for i := 0; i != n; i++ {
		str = str + " "
	}

	return str
}
