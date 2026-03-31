package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func getMen(vet []int) []int {
	var resp []int
	for _, v := range vet {
		if v > 0 {
			resp = append(resp, v)
		}
	}
	if resp == nil {
		return []int{}
	}
	return resp
}

func getCalmWomen(vet []int) []int {
	var resp []int
	for _, v := range vet {
		if v < 0 && v > -10 {
			resp = append(resp, v)
		}
	}
	if resp == nil {
		return []int{}
	}
	return resp
}

func sortVet(vet []int) []int {
	resp := make([]int, len(vet))
	copy(resp, vet)
	sort.Ints(resp)
	return resp
}

func sortStress(vet []int) []int {
	resp := make([]int, len(vet))
	copy(resp, vet)
	sort.SliceStable(resp, func(i, j int) bool {
		return abs(resp[i]) < abs(resp[j])
	})
	return resp
}

func reverse(vet []int) []int {
	resp := make([]int, len(vet))
	for i := 0; i < len(vet); i++ {
		resp[i] = vet[len(vet)-1-i]
	}
	if resp == nil {
		return []int{}
	}
	return resp
}

func unique(vet []int) []int {
	var resp []int
	vistos := make(map[int]bool)
	for _, v := range vet {
		if !vistos[v] {
			resp = append(resp, v)
			vistos[v] = true
		}
	}
	if resp == nil {
		return []int{}
	}
	return resp
}

func repeated(vet []int) []int {
	var resp []int
	vistos := make(map[int]bool)
	for _, v := range vet {
		if vistos[v] {
			resp = append(resp, v)
		} else {
			vistos[v] = true
		}
	}
	if resp == nil {
		return []int{}
	}
	return resp
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Split(line, " ")

		fmt.Println("$" + line)

		if args[0] == "end" {
			break
		}

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return []int{}
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}