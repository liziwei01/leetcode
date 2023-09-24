package leihuo

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)
// sample input
// 2
// 草果,10001
// 桑白皮,10002
// 2
// [2023-07-24 17:24:04][game][ItemMgr.py:482][ItemLog]:[OpenBox]操作类型:1,player_id=[300011][剑灵],抽到道具ID=[10004,1;10005,1]
// [2023-07-24 17:24:05][game][ItemMgr.py:482][ItemLog]:[OpenBox]操作类型:2,player_id=[330011][游戏人生],抽到道具ID=[10004,1;10002,58;10005,27;10001,98;10003,59]
func main() {
	N := 0
	fmt.Scan(&N)
	itemNameIDin := ""
	itemIDNameMap := make(map[string]string)
	itemIDNameSlice := make([]string, 0)
	for i := 0; i != N; i++ {
		fmt.Scan(&itemNameIDin)
		itemID, itemName := splitIDName(itemNameIDin)
		itemIDNameMap[itemID] = itemName
		itemIDNameSlice = append(itemIDNameSlice, itemNameIDin)
	}

	M := 1
	fmt.Scan(&M)
	whateverDate := ""
	boxLog := ""
	itemIDDropMap := make(map[string]int)
	wrongItemID := make([]string, 0)
	dropSum := 0
	for i := 0; i != M; i++ {
		fmt.Scan(&whateverDate)
		fmt.Scan(&boxLog)
		opType, items := splitOpTypeItems(boxLog)
		if opType == "1" {
			dropSum += 1
		} else {
			dropSum += 10
		}
		for _, itemIDandNum := range items {
			tmpItemID, tmpItemNum := splitIDNum(itemIDandNum)
			itemIDDropMap[tmpItemID] += tmpItemNum
			if _, ok := itemIDNameMap[tmpItemID]; !ok {
				wrongItemID = append(wrongItemID, tmpItemID)
			}
		}
	}

	if len(wrongItemID) != 0 {
		reportWrong(wrongItemID)
		return
	}

	for _, itemNameIDin := range itemIDNameSlice {
		itemID, itemName := splitIDName(itemNameIDin)
		totalDrop := itemIDDropMap[itemID]
		possibility := float64(totalDrop) / float64(dropSum)
		strPossibility := strconv.FormatFloat(possibility, 'f', 4, 64)
		fmt.Printf("%s:%d,%s\n", itemName, totalDrop, strPossibility[:5])
	}

}

func splitIDName(s string) (string, string) {
	itemNameIDSlice := strings.Split(s, ",")
	return itemNameIDSlice[1], itemNameIDSlice[0]
}

func splitOpTypeItems(s string) (string, []string) {
	tmpSlice := strings.Split(s, "=")
	strIncludeOpType := tmpSlice[0]
	sliceIncludeOpType := strings.Split(strIncludeOpType, ":")[5]
	strOpType := strings.Split(sliceIncludeOpType, ",")[0]
	itemsStr := tmpSlice[2]
	itemsStr = strings.TrimLeft(itemsStr, "[")
	itemsStr = strings.TrimRight(itemsStr, "]")
	items := strings.Split(itemsStr, ";")
	return strOpType, items
}

func splitIDNum(s string) (string, int) {
	itemIDNumSlice := strings.Split(s, ",")
	tmpItemNum, _ := strconv.Atoi(itemIDNumSlice[1])
	return itemIDNumSlice[0], tmpItemNum
}

func reportWrong(wrongItemID []string) {
	sort.Strings(wrongItemID)
	realWrongItemID := make([]string, 0)
	for i := 0; i != len(wrongItemID); i++ {
		if i != 0 && wrongItemID[i] == wrongItemID[i-1] {
			continue
		}
		realWrongItemID = append(realWrongItemID, wrongItemID[i])
	}
	fmt.Print("Find Wrong ItemId:")
	for i := 0; i != len(realWrongItemID); i++ {
		if i == len(realWrongItemID)-1 {
			fmt.Printf("%s", realWrongItemID[i])
		} else {
			fmt.Printf("%s,", realWrongItemID[i])
		}
	}
	return
}
