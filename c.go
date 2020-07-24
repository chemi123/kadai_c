package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
  方針: 継承者ごとのスケジュールのリスト(shedules [][]int)から日付ごとに集まれる継承者一覧のリストを作る。
	   そこから集まった継承者が地図を集まった分だけ持つようにし、次に集まった時にそれをさらなる継承者に渡すというのを繰り返す。
	   直感的な実装になりそう。ここでは試さないが、グラフを使った探索もできそう？
*/
func collectMap(schedules [][]int) int {
	return -1
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dataSetNum, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	fmt.Println(dataSetNum)

	for i := 0; i < dataSetNum; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		fmt.Println(n)
		schedules := make([][]int, 0, n)
		maxDay := 0
		for j := 0; j < n; j++ {
			scanner.Scan()
			daysStr := strings.Split(scanner.Text(), " ")
			days := make([]int, 0, len(daysStr))
			for _, dayStr := range daysStr {
				day, _ := strconv.Atoi(dayStr)
				maxDay = max(maxDay, day)
				days = append(days, day)
			}
			schedules = append(schedules, days)
		}
		fmt.Println(schedules)
		fmt.Println(maxDay)
	}
}
