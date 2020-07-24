package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
  方針1: 継承者ごとのスケジュールのリスト(shedules [][]int)から日付ごとに集まれる継承者一覧のリストを作る。
	    そこから集まった継承者が地図を集まった分だけ持つようにし、次に集まった時にそれをさらなる継承者に渡すというのを繰り返す。
	    直感的な実装になりそう。グラフを使った探索もできそうだったため、これとはまた別にcollectoMap2を実装した。
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

type node struct {
	id   int
	day  int
	left *node
}

/*
  方針2: 縦と横方向で構成されたグラフを使った探索を行う。全ての行(縦方向)の探索が終われば地図を全て集められることになる。
        縦方向は双方向に移動できるが、横方向は左にしか移動できない。小さい日付から始め、探索可能ならその日付を返す。不可なら次に大きい日付を試す。
        全ての日付で探索不可なら-1を返す。

  例: 入力値が[[1, 6], [2, 3], [1, 2], [3, 4, 5]]の場合、以下のようなグラフで表現できる
  id 日付
  0: 1---------6
     |
  1: | 2-3
     | | |
  2: 1-2 |
         |
  3:     3-4-5

  上記のような例の場合、以下の手順を踏む
  1. 1日から探索開始。1日は縦方向に0, 2のid(行)に探索できる。これ以上左にはいけないためここで探索完了。全ての行は探索できていないが1日に探索できるidをキャッシュに保存。
  2. 2日から探索開始。2日は縦方向に1, 2のidに探索できる。idが2から1日の日付に移動できる。ここで1のキャッシュ(idが0, 2)を参照し、2日の探索結果と合わせて0, 1, 2になるがまだ探索できていないため2日の結果をキャッシュに保存。
  3. 3日から探索開始、3日は縦方向に1, 3のidに探索できる。idが1から2日の日付に移動できる。ここで2のキャッシュ(idが0, 1, 2)を参照し、3日の探索結果と合わせて0, 1, 2, 3となり探索完了。3を返す。

  データ構造は左にしか辿れないリンクリストと配列を使う。node01はidが0で日付が1日であることを示す。node06 -> node01と繋がっており、それぞれのnodeを参照することでidと日付が取れる。
    ↓ 0番目は0日に対応した日付がないので空。配列である以上便宜的に要素が存在するだけ。
  [[],[node01, node21], [node12, node22], [node23, node33], [node34], [node35], [node06]]

 時間計算量はO(N*days)になる。測り方は粗い上マシンパワーにも依存するが以下の結果になった。(方針1と方針2で結果の差分がないことは確認している)。
 当然であるがデータセットに対して線形な時間の増え方を見せている。N(<=50)が大きくなれば、二次関数以上の差分が広がる。
 データセット数 速度(方針1) 速度(方針2)
 10000        4.72s      0.86s
 100000       47.12s     8.69s

 データセット数を10000で固定。Nを変更させて計測
 N   速度(方針1) 速度(方針2)
 50  4.72s      0.86s
 100 35.65s     2.91s
 200 283.25s    11.51s
*/
func collectMap2(schedules [][]int, maxDay int) int {
	// 条件は1 <= N <= 50であるため、一人の場合もある。その場合は集まる必要はないため0を返すことにする
	if len(schedules) == 1 {
		return 0
	}

	nodesList := make([][]*node, maxDay+1)
	for id, schedule := range schedules {
		var prevNode *node
		for _, day := range schedule {
			n := &node{id: id, day: day, left: prevNode}
			prevNode = n
			nodesList[day] = append(nodesList[day], n)
		}
	}

	// 日付毎に訪れたidのセット(map[int]bool)を格納したmap。キャッシュ用
	dayIdSet := make(map[int]map[int]bool)
	for day := 1; day <= maxDay; day++ {
		dayIdSet[day] = make(map[int]bool)
		for _, node := range nodesList[day] {
			dayIdSet[day][node.id] = true
			if node.left != nil {
				for id := range dayIdSet[node.left.day] {
					dayIdSet[day][id] = true
				}
			}
		}
		if len(dayIdSet[day]) == len(schedules) {
			return day
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

		// 切り替えて試してみる
		// fmt.Println(collectMap(schedules, maxDay))
		fmt.Println(collectMap2(schedules, maxDay))
	}
}
