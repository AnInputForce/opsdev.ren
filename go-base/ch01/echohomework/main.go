package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("练习1：修订echo，打印os.Args[0],即：执行命令本身的名字.\n----------------------------------------------------")
	fmt.Printf("执行命令本身的名字：%s\n", strings.Join(os.Args[0:1], " "))
	fmt.Println("练习2：修订echo，打印每个参数的索引和值，每个一行.\n----------------------------------------------------")
	for t, arg := range os.Args[1:] {
		fmt.Printf("第%d个索引，值为：%s\n", t+1, arg)
	}
	fmt.Println("练习3：比较echo三个版本的优化时间。\n----------------------------------------------------")
	echo1()
	echo2()
	echo3()

}
func echo1() {
	var s, sep string
	t1 := time.Now() // get current time
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	//fmt.Println(s)
	elapsed := time.Since(t1)
	fmt.Printf("echo1 elapsed: %s\n", elapsed)
}
func echo2() {
	s, sep := "", ""
	t1 := time.Now() // get current time

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	//fmt.Println(s)

	elapsed := time.Since(t1)
	fmt.Printf("echo2 elapsed: %s\n", elapsed)
}
func echo3() {

	t1 := time.Now() // get current time

	strings.Join(os.Args[1:], " ")
	//fmt.Println(strings.Join(os.Args[1:], " "))

	elapsed := time.Since(t1)
	fmt.Printf("echo3 elapsed: %s\n", elapsed)
}

/* output:
ChinaDreams:echohomework kangcunhua$ ./echohomework Happy families are all alike, every unhappy family is unhappy in its own way.
练习1：修订echo，打印os.Args[0],即：执行命令本身的名字.
-------------------------------------
执行命令本身的名字：./echohomework
练习2：修订echo，打印每个参数的索引和值，每个一行.
-------------------------------------
第1个索引，值为：Happy
第2个索引，值为：families
第3个索引，值为：are
第4个索引，值为：all
第5个索引，值为：alike,
第6个索引，值为：every
第7个索引，值为：unhappy
第8个索引，值为：family
第9个索引，值为：is
第10个索引，值为：unhappy
第11个索引，值为：in
第12个索引，值为：its
第13个索引，值为：own
第14个索引，值为：way.
练习3：比较echo三个版本的优化时间。
-------------------------------------
echo1 elapsed: 3.752µs
echo2 elapsed: 2.216µs
echo3 elapsed: 1.553µs
*/
