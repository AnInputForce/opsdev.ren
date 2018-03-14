package main

import (
	"bufio"
	"fmt"
	"os"
)

// 查找重复的行:这是原翻译。或许翻译成：“以行为文本单位，统计文本出现的次数”更合适
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countlines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countlines(f, counts)
			//fmt.Println(f.Name())
			//fmt.Println(counts)
			f.Close()

		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

func countlines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		//println(input.Text())
	}
	//fmt.Println("------------countlinse------")
	//fmt.Println(f.Name())
	//fmt.Println(counts)
}

/* output:
ChinaDreams:dup2 kangcunhua$ go build
ChinaDreams:dup2 kangcunhua$ ./dup2 ranking2015.txt ranking2016.txt ranking2017.txt
7       c++
6       go
6       python
2       delphi
10      java
4       c
6       ada
4       basic
4       pascal
ChinaDreams:dup2 kangcunhua$ ./dup2
java
java
c
c
c++
c++
c++
go
go
ada
basic
pascal
pascal
pascal
basic
basic^D
2      java
2       c
3       c++
2       go
3       basic
3       pascal
*/
