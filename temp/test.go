package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  // 匿名字段
	school string
	loan   float32
}

type Employee struct {
	Human   // 匿名字段
	company string
	money   float32
}

// Human 实现 SayHi 方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human 实现 Sing 方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

// Employee 重载 Human 的 SayHi 方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

// Interface Men 被 Human, Student 和 Employee 实现
// 因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	// 定义 Men 类型的变量i
	var i Men

	// i 能存储 Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	// i 也能存储 Employee
	i = tom
	fmt.Println("This is tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	// 定义了 slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	// 这三个都是不同类型的元素，但是他们实现了 interface 同一个接口
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}
}

func sortPeople(names []string, heights []int) []string {
	res := make([]string, 0)
	return res
}

// func arithmeticTriplets(nums []int, diff int) int {
// 	count := 0
// 	for i := len(nums) - 1; i >= 0; i-- {
// 		// count = 0
// 		for j := i - 1; j >= 0; j-- {
// 			if nums[i]-nums[j] == diff {
// 				fmt.Println(nums[i], nums[j], "ok")
// 				count++
// 			} else {
// 				fmt.Println(nums[i], nums[j])
// 			}
// 		}
// 	}

// 	return count
// }

// result := make([]rune, len(s))
// var strs = []rune{}

// for _, v := range s {
// 	strs = append(strs, v)
// }

// for i := 0; i < len(strs); i++ {
// 	indx := indices[i]
// 	result[indx] = strs[i]

// }
// fmt.Println(string(result))

// return string(result)

// for _, v := range s {
// 	sliceStr = append(sliceStr, v)
// 	fmt.Printf("%q ", string(sliceStr))
// 	fmt.Println(sliceStr[1])
// }