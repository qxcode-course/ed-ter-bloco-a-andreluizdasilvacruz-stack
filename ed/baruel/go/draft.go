package main
import "fmt"

func main() {
	var total, qtd int
	fmt.Scan(&total)
	fmt.Scan(&qtd)

	figs := make([]int, qtd)
	count := make(map[int]int)

	for i := 0; i < qtd; i++ {
		fmt.Scan(&figs[i])
		count[figs[i]]++
	}

	
	repetidas := []int{}
	for _, v := range figs {
		if count[v] > 1 {
			repetidas = append(repetidas, v)
			count[v]-- 
		}
	}

	if len(repetidas) == 0 {
		fmt.Println("N")
	} else {
		for i, v := range repetidas {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(v)
		}
		fmt.Println()
	}

	
	faltando := []int{}
	for i := 1; i <= total; i++ {
		if count[i] == 0 {
			faltando = append(faltando, i)
		}
	}

	if len(faltando) == 0 {
		fmt.Println("N")
	} else {
		for i, v := range faltando {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(v)
		}
		fmt.Println()
    
	}
}

