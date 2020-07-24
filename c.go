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
	// 条件は1 <= N <= 50であるため、一人の場合もある。その場合は集まる必要はないため0を返すことにする
	if len(schedules) == 1 {
		return 0
	}

	// 日付がkey、その日に集まる継承者のリストがvalue。日付をそのままキーとしたかったためサイズはmaxDay+1で指定(keyが0は使わない)
	scheduleMap := make([][]int, maxDay+1)

	// 継承者は0からidとして振り分けることにする(三人いれば上から0, 1, 2)
	for id, schedule := range schedules {
		for _, day := range schedule {
			scheduleMap[day] = append(scheduleMap[day], id)
		}
	}

	// 継承者毎に持っている地図(継承者のid)のmap
	// keyが継承者id, valueが持っている地図のid(setが良いが、golangはsetを提供していないためmap[int]boolで代用)
	successorMap := make([]map[int]bool, len(schedules))
	for i := range successorMap {
		successorMap[i] = make(map[int]bool)
	}

	// 1 <= day <= 30, 1 <= N <= 50の条件で最悪ケースは単純計算だと30 * 50 * 50 * 50 = 3750000となる
	// ただし単純に上記の掛け算にはならないと考えられるため、実測する必要があるがそこまで遅くはない？
	// ただグラフの探索など少し良いやり方はないか考えられそうでもある。
	for day := range scheduleMap {
		for _, id1 := range scheduleMap[day] {
			for _, id2 := range scheduleMap[day] {
				successorMap[id1][id2] = true
				for id3 := range successorMap[id2] {
					if !successorMap[id1][id3] {
						successorMap[id1][id3] = true
					}
				}
				if len(successorMap[id1]) == len(schedules) {
					return day
				}
			}
		}
	}

	return -1
}

/*
  math.Maxはfloat64の引数を取り、float64を返すものしか提供していない。使いにくいので自分でintのmax関数を定義
  https://golang.org/pkg/math/#Max
*/
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
