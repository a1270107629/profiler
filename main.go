package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	k string
	v int
}
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].v > p[j].v }

func main() {
	if len(os.Args) <= 1 {
		return
	}
	arg := []string{"-T"}
	arg = append(arg, os.Args[1:]...)
	cmd := exec.Command("strace", arg...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		return
	}
	errStr := stderr.String()
	sysSpits := strings.Split(errStr, "\n")
	m := make(map[string]int, 0)
	for i, sysSpit := range sysSpits {
		fmt.Println(sysSpit)
		if i == len(sysSpits)-1 {
			break
		}
		time, err := strconv.Atoi(sysSpit[len(sysSpit)-7 : len(sysSpit)-1])
		if err != nil {
			continue
		}
		i := 0
		for ; i < len(sysSpit); i++ {
			if sysSpit[i] == '(' {
				break
			}
		}
		comm := sysSpit[:i]

		m[comm] += time
	}
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	fmt.Println(p)
}
