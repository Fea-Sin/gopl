/*
/*
@Time : 2020/11/20 5:33 下午
@Author : chengqunzhong
@File : gd
@Software: GoLand
*/
/*
 * ReadRune 方法执行UTF-8解码并返回三个值：
 * 解码的rune字符的值，字符UTF-8编码后的长度和一个错误值
 * 我们可预期的错误值只有对应文件结尾的io.EOF
 * 如果输入的是无效UTF-8编码的字符，返回的将是unicode.ReplacementChar表示无效
 * 且编码长度是1
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	fmt.Print("Enter text: ")
	in := bufio.NewReader(os.Stdin)
	_, err := in.ReadString('\n')

	if err != nil {

		for {
			r, n, err := in.ReadRune()
			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
				os.Exit(1)
			}
			if r == unicode.ReplacementChar && n == 1 {
				invalid++
				continue
			}
			counts[r]++
			utflen[n]++
		}

		fmt.Printf("rune\tcount\t")
		for c, n := range counts {
			fmt.Printf("%q\t%d\t", c, n)
		}
		fmt.Print("\nlen\tcount\n")
		for i, n := range utflen {
			if i > 0 {
				fmt.Printf("%d\t%d\t", i, n)
			}
		}
		if invalid > 0 {
			fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
		}
	}
}
