package main

import (
	"fmt"
)

var (
	a     [30000]byte
	prog  = "++++++++++[>++++++++++<-]>++++.+."
	p, pc int
)

func loop(inc int) {
	for i := inc; i != 0; pc += inc {
		fmt.Printf("\nloop:inc->%d,pc->%d,i->%d\n", inc, pc, i)
		switch prog[pc+inc] {
		case '[':
			i++
			fmt.Printf("\n[%d\n", i)
		case ']':
			i--
			fmt.Printf("\n%d]  %s\n", i, prog[pc+inc])
		}
	}
}

func main() {
	for {
		switch prog[pc] {
		case '>':
			p++
		case '<':
			p--
		case '+':
			a[p]++
		case '-':
			a[p]--
		case '.':
			fmt.Print(string(a[p]))
			fmt.Printf("\n %d     %d \n", p, pc)
		case '[':
			if a[p] == 0 {
				fmt.Printf("\n[%d\n", p)
				loop(1)
			}
		case ']':
			if a[p] != 0 {
				fmt.Printf("\n pc:%d  %d]\n", pc, p)
				loop(-1)
			}
		default:
			fmt.Printf("error")
		}
		pc++
		if pc == len(prog) {
			return
		}
	}
}
