package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

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

		if value, ok := m[comm]; ok {
			m[comm] = value + time
		} else {
			m[comm] = time
		}
	}

	fmt.Println(m)
}
