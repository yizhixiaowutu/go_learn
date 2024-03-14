package main

import "fmt"

func main() {
	// 键盘录入
	// 方式一 fmt.Scanln
	// var age int
	// fmt.Println("录取年龄：")

	// fmt.Scanln(&age)

	// var name string
	// fmt.Println("录取姓名：")
	// fmt.Scanln(&name)

	// var score float32
	// fmt.Println("录取分数：")
	// fmt.Scanln(&score)

	// fmt.Printf("录取年龄：%v, 录取姓名: %v, 录取分数 %v \n", age, name, score)

	// 方式二 fmt.Scanf
	var age int
	var name string
	var score float32

	fmt.Println("依次录入年龄、姓名、分数，以空格分隔。")
	fmt.Scanf("%d %s %f", &age, &name, &score)

	fmt.Printf("录取年龄：%v, 录取姓名: %v, 录取分数 %v \n", age, name, score)
}
