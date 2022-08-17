package misch

import (
	"github.com/Pedraamy/Golang-RL-Chess-AI/state"
	"fmt"
)

func PrintBoard(st *state.State) {
	runes := [8][8]string{}
	for r:= 0; r<8; r++ {
		for c:= 0; c<8; c++ {
			var curr uint64 = 1<<(8*r+c)
			if curr&st.WK != 0 {
				runes[r][c] = "K"
			} else if curr&st.BK != 0 {
				runes[r][c] = "k"
			} else if curr&st.WQ != 0 {
				runes[r][c] = "Q"
			} else if curr&st.BQ != 0 {
				runes[r][c] = "q"
			} else if curr&st.WR != 0 {
				runes[r][c] = "R"
			} else if curr&st.BR != 0 {
				runes[r][c] = "r"
			} else if curr&st.WB != 0 {
				runes[r][c] = "B"
			} else if curr&st.BB != 0 {
				runes[r][c] = "b"
			} else if curr&st.WN != 0 {
				runes[r][c] = "N"
			} else if curr&st.BN != 0 {
				runes[r][c] = "n"
			} else if curr&st.WP != 0 {
				runes[r][c] = "P"
			} else if curr&st.BP != 0 {
				runes[r][c] = "p"
			} else {
				runes[r][c] = "."
			}
		}
	}
	for _, l := range runes {
		fmt.Println(l)
	}
}