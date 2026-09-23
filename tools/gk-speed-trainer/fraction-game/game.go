package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const Goal = 100

func normalize(s string) string {

	s = strings.TrimSpace(s)

	if strings.HasPrefix(s, ".") {
		s = "0" + s
	}

	return s
}

func StartGame() {

	reader := bufio.NewReader(os.Stdin)

	pool := NewPool()

	score := 0

	fmt.Println("==========")
	fmt.Println("连续答对100题即可通关！")
	fmt.Println("答错立即结束！")
	fmt.Println("==========")

	for score < Goal {

		q := pool.Next()

		fmt.Printf("\n(%d/%d)\n", score+1, Goal)
		fmt.Printf("%s = ", q.Fraction)

		input, _ := reader.ReadString('\n')

		input = normalize(input)

		if input == q.Decimal {

			score++

			fmt.Println("✅ 正确！")

			continue
		}

		fmt.Println("❌ 回答错误")
		fmt.Println("正确答案：", q.Decimal)
		fmt.Printf("本次成绩：%d\n", score)

		return
	}

	fmt.Println()
	fmt.Println("🎉🎉🎉 恭喜通关！")
}