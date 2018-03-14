package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 查找重复的行:这是原翻译。或许翻译成：“以行为文本单位，统计文本出现的次数”更合适
func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

/* output:
ChinaDreams:dup3 kangcunhua$ ./dup3 ranking2015.txt ranking2016.txt ranking2017.txt
10      java
4       c
7       c++
6       ada
4       basic
6       python
6       go
4       pascal
2       delphi
*/
