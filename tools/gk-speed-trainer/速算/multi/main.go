package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	Target      = 100
	RecentLimit = 20
)

type Stats struct {
	SingleSingle int
	DoubleSingle int
	DoubleDouble int
	DoubleTriple int
}

func main() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	reader := bufio.NewReader(os.Stdin)

	recent := make([]string, 0, RecentLimit)

	var stats Stats

	count := 0

	totalDuration := time.Duration(0)

	fmt.Println("==========================================")
	fmt.Println("      公考资料分析乘法速算训练")
	fmt.Println("==========================================")
	fmt.Println("连续答对100题即可通关")
	fmt.Println("输入错误立即结束")
	fmt.Println("==========================================")

	for count < Target {

		a, b, typ := generateQuestion(r, count, recent)

		key := fmt.Sprintf("%d*%d", a, b)
		recent = pushRecent(recent, key)

		switch typ {
		case 1:
			stats.SingleSingle++
		case 2:
			stats.DoubleSingle++
		case 3:
			stats.DoubleDouble++
		case 4:
			stats.DoubleTriple++
		}

		answer := a * b

		fmt.Printf("\n[%3d/%3d] %d × %d = ", count+1, Target, a, b)

		start := time.Now()

		for {

			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("读取失败")
				return
			}

			input = strings.TrimSpace(input)

			userAnswer, err := strconv.Atoi(input)

			if err != nil {

				fmt.Print("请输入数字：")
				continue
			}

			duration := time.Since(start)

			totalDuration += duration

			if userAnswer != answer {

				fmt.Println("\n❌ 回答错误！")
				fmt.Printf("正确答案：%d\n", answer)

				printResult(
					false,
					count,
					totalDuration,
					stats,
				)

				return
			}

			count++

			avg := totalDuration / time.Duration(count)

			fmt.Printf(
				"✅ %.2fs  平均 %.2fs\n",
				duration.Seconds(),
				avg.Seconds(),
			)

			if count%10 == 0 && count != Target {

				fmt.Printf(
					"🔥 连击 %d！继续保持！\n",
					count,
				)
			}

			break
		}
	}

	printResult(
		true,
		count,
		totalDuration,
		stats,
	)
}

func generateQuestion(r *rand.Rand, correct int, recent []string) (int, int, int) {

	for {

		var a, b int
		var typ int

		p := r.Intn(100)

		// 根据训练进度动态调整概率
		switch {

		// 1~20题
		case correct < 20:

			switch {
			case p < 80:
				a = r.Intn(7) + 3
				b = r.Intn(7) + 3
				typ = 1

			default:
				a = r.Intn(90) + 10
				b = r.Intn(7) + 3
				typ = 2
			}

		// 21~40题
		case correct < 40:

			switch {
			case p < 60:
				a = r.Intn(7) + 3
				b = r.Intn(7) + 3
				typ = 1

			case p < 90:
				a = r.Intn(90) + 10
				b = r.Intn(7) + 3
				typ = 2

			default:
				a = r.Intn(90) + 10
				b = r.Intn(90) + 10
				typ = 3
			}

		// 41~60题
		case correct < 60:

			switch {
			case p < 40:
				a = r.Intn(7) + 3
				b = r.Intn(7) + 3
				typ = 1

			case p < 80:
				a = r.Intn(90) + 10
				b = r.Intn(7) + 3
				typ = 2

			default:
				a = r.Intn(90) + 10
				b = r.Intn(90) + 10
				typ = 3
			}

		// 61~80题
		case correct < 80:

			switch {
			case p < 20:
				a = r.Intn(7) + 3
				b = r.Intn(7) + 3
				typ = 1

			case p < 60:
				a = r.Intn(90) + 10
				b = r.Intn(7) + 3
				typ = 2

			case p < 90:
				a = r.Intn(90) + 10
				b = r.Intn(90) + 10
				typ = 3

			default:
				a = r.Intn(90) + 10
				b = r.Intn(900) + 100
				typ = 4
			}

		// 81~100题
		default:

			switch {
			case p < 10:
				a = r.Intn(7) + 3
				b = r.Intn(7) + 3
				typ = 1

			case p < 40:
				a = r.Intn(90) + 10
				b = r.Intn(7) + 3
				typ = 2

			case p < 80:
				a = r.Intn(90) + 10
				b = r.Intn(90) + 10
				typ = 3

			default:
				a = r.Intn(90) + 10
				b = r.Intn(900) + 100
				typ = 4
			}
		}

		// 随机交换左右乘数
		if r.Intn(2) == 0 {
			a, b = b, a
		}

		key := fmt.Sprintf("%d*%d", a, b)

		if !contains(recent, key) {
			return a, b, typ
		}
	}
}

func contains(list []string, key string) bool {
	for _, v := range list {
		if v == key {
			return true
		}
	}
	return false
}

func pushRecent(list []string, key string) []string {
	list = append(list, key)

	if len(list) > RecentLimit {
		list = list[1:]
	}

	return list
}

func printResult(
	pass bool,
	count int,
	total time.Duration,
	stats Stats,
) {

	fmt.Println("\n==========================================")

	if pass {

		fmt.Println("🎉🎉🎉 恭喜通关！")
		fmt.Printf("连续答对 %d 题！\n", Target)

	} else {

		fmt.Printf("训练结束！连续答对：%d/%d\n", count, Target)
	}

	fmt.Println("------------------------------------------")

	if count > 0 {

		avg := total / time.Duration(count)

		fmt.Printf("总耗时：%.2f 秒\n", total.Seconds())
		fmt.Printf("平均每题：%.2f 秒\n", avg.Seconds())

		fmt.Println("------------------------------------------")

		fmt.Println("题型统计：")
		fmt.Printf("3~9 × 3~9         : %d\n", stats.SingleSingle)
		fmt.Printf("10~99 × 3~9       : %d\n", stats.DoubleSingle)
		fmt.Printf("10~99 × 10~99     : %d\n", stats.DoubleDouble)
		fmt.Printf("10~99 × 100~999   : %d\n", stats.DoubleTriple)

		fmt.Println("------------------------------------------")

		sec := avg.Seconds()

		switch {

		case sec < 1.5:
			fmt.Println("★★★★★ 速算大师")
			fmt.Println("你的乘法速算已经达到优秀水平，可以直接用于考公资料分析。")

		case sec < 2.5:
			fmt.Println("★★★★☆ 熟练")
			fmt.Println("速度很好，再继续提升两位数乘法即可。")

		case sec < 3.5:
			fmt.Println("★★★☆☆ 良好")
			fmt.Println("已经具备不错基础，建议加强两位数×两位数训练。")

		case sec < 5:
			fmt.Println("★★☆☆☆ 入门")
			fmt.Println("建议继续练习个位数和两位数乘法，提高熟练度。")

		default:
			fmt.Println("★☆☆☆☆ 新手")
			fmt.Println("建议每天坚持练习100题，重点提升口算速度。")
		}
	}

	fmt.Println("==========================================")
}