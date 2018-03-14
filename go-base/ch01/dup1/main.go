package main

import (
	"bufio"
	"fmt"
	"os"
)

// 查找重复的行:这是原翻译。或许翻译成：“以行为文本单位，统计文本出现的次数”更合适
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {

		// if input.Text() == "end" {
		// 	break
		// }

		// you can also press hot key "ctrl + d" for the EOF signal
		counts[input.Text()]++

		//fmt.Println(counts)

	}
	//fmt.Println(counts)
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("\n%d\t%s\n", n, line)
		}

	}
	// 加上结果输出：无论有没有重复的line
	fmt.Printf("counts result is %v;\n", counts)

}

/* output:
ChinaDreams:dup1 kangcunhua$ go run main.go
java
c
c++
c++
c++
go
go
ada
basic
pascal
^D
3       c++

2       go
counts result is map[c:1 c++:3 go:2 ada:1 basic:1 pascal:1 java:1];
*/
