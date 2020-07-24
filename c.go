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
func collectMap(schedules [][]int, maxDay int) int {
	// 日付がkey、その日に集まる継承者のリストがvalue。日付をそのままキーとしたかったためサイズはmaxDay+1で指定(keyが0は使わない)
	scheduleMap := make([][]int, maxDay+1)

	// 継承者は0からidとして振り分けることにする(三人いれば上から0, 1, 2)
	for id, schedule := range schedules {
		for _, day := range schedule {
			scheduleMap[day] = append(scheduleMap[day], id)
		}
	}

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

	for i := 0; i < dataSetNum; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
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
		fmt.Println(collectMap(schedules, maxDay))
	}
}
