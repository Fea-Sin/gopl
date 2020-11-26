/*
/*
@Time : 2020/11/26 3:08 下午
@Author : chengqunzhong
@File : gfc
@Software: GoLand
*/
package main

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}

}
