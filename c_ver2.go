package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	id   int
	date int
	left *Node
}

/*
  方針: 縦と横方向で構成されたグラフを使った探索を行う。全ての行(縦方向)の探索が終われば地図を全て集められることになる。
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

  データ構造は左にしか辿れないリンクリストと配列を使う。node01はidが0で日付が1日であることを示す。node06 -> node01と繋がっており、node01を参照することでidと日付が取れる。
    ↓ 0番目は0日に対応した日付がないので空。配列である以上便宜的に要素が存在するだけ。
  [[],[node01, node21], [node12, node22], [node23, node33], [node34], [node35], [node06]]
*/
func collectMap2(schedules [][]int) int {
	// 条件は1 <= N <= 50であるため、一人の場合もある。その場合は集まる必要はないため0を返すことにする
	if len(schedules) == 1 {
		return 0
	}

	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dataSetNum, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	for i := 0; i < dataSetNum; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		schedules := make([][]int, 0, n)
		for j := 0; j < n; j++ {
			scanner.Scan()
			daysStr := strings.Split(scanner.Text(), " ")
			days := make([]int, 0, len(daysStr))
			for _, dayStr := range daysStr {
				day, _ := strconv.Atoi(dayStr)
				days = append(days, day)
			}
			schedules = append(schedules, days)
		}
		fmt.Println(collectMap2(schedules))
	}
}
