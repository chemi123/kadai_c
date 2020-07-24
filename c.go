package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func collectMap(schedules [][]int) int {
	return -1
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
		fmt.Println(schedules)
	}
}
