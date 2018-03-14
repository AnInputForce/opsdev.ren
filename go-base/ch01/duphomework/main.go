package main

import (
	"bufio"
	"fmt"
	"os"
)

// 查找重复的行:这是原翻译。或许翻译成：“以行为文本单位，统计文本出现的次数”更合适
// 练习1.4：修订dup2，出现重复的行时，打印文件名
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
		if counts[input.Text()] > 1 {
			fmt.Printf("在文件%s发现数据%s\t重复了%d次；\n", f.Name(), input.Text(), counts[input.Text()])
		}
	}
}

/* output:
ChinaDreams:duphomework kangcunhua$ go buildChinaDreams:duphomework kangcunhua$ ./duphomeworkranking2010.txt ranking2011.txt ranking2012.txt
在文件ranking2011.txt发现数据go 重复了2次；
在文件ranking2011.txt发现数据python     重复了2次；
在文件ranking2012.txt发现数据java       重复了2次；
在文件ranking2012.txt发现数据java       重复了3次；
在文件ranking2012.txt发现数据c  重复了2次；
在文件ranking2012.txt发现数据c  重复了3次；
在文件ranking2012.txt发现数据c++        重复了2次；
在文件ranking2012.txt发现数据c++        重复了3次；
在文件ranking2012.txt发现数据c++        重复了4次；
在文件ranking2012.txt发现数据go 重复了3次；
在文件ranking2012.txt发现数据go 重复了4次；
在文件ranking2012.txt发现数据ada        重复了2次；
在文件ranking2012.txt发现数据basic      重复了2次；
在文件ranking2012.txt发现数据pascal     重复了2次；
在文件ranking2012.txt发现数据pascal     重复了3次；
在文件ranking2012.txt发现数据pascal     重复了4次；
在文件ranking2012.txt发现数据basic      重复了3次；
在文件ranking2012.txt发现数据basic      重复了4次；
3       java
3       c
4       c++
4       go
2       ada
4       basic
4       pascal
2       python
*/
