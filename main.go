package main

import (
	"github.com/Pedraamy/Golang-RL-Chess-AI/state"
	"fmt"
)


func main() {
	var curr1, curr2 uint64
	curr1 = 1<<63
	curr2 = 1<<55
	curr2 |= curr1
	curr2 |= 1<<58
	curr2 += 4
	curr2 += 8
	res := state.GetPositionsFromBoard(curr2)
	fmt.Println(res)
}
