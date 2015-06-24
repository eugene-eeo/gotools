package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func buildIndexes(strs []string) []map[string]int {
	maps := []map[string]int{}
	for _, str := range strs {
		strMap := map[string]int{}
		for idx, item := range strings.Split(str, "\n") {
			strMap[item] = idx
		}
		maps = append(maps, strMap)
	}
	return maps
}

func eachIncludes(ms []map[string]int, s string) bool {
	for _, m := range ms {
		_, ok := m[s]
		if !ok {
			return false
		}
	}
	return true
}

func intersection(maps []map[string]int) map[int]string {
	union := map[int]string{}
	tails := maps[1:]
	for str, idx := range maps[0] {
		if eachIncludes(tails, str) {
			union[idx] = str
		}
	}
	return union
}

func keysOf(m map[int]string) []int {
	buf := []int{}
	for k, _ := range m {
		buf = append(buf, k)
	}
	return buf
}

func main() {
	strings := os.Args[1:]
	if len(strings) == 0 {
		return
	}
	inter := intersection(buildIndexes(strings))
	lines := keysOf(inter)
	sort.Ints(lines)
	for _, item := range lines {
		fmt.Println(inter[item])
	}
}
