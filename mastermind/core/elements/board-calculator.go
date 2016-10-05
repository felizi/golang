package elements

import (
	"../validade"
)

func calcHits(userCodes, boardCodes []int) (misses, hitsChecked, hitsOutOfOrderChecked []int) {
	hitsChecked, hitsOutOfOrderChecked = calcHitsAndHitsOutOfOrder(userCodes, boardCodes)
	misses = calcMisses(boardCodes, hitsChecked, hitsOutOfOrderChecked)
	return misses, hitsChecked, hitsOutOfOrderChecked
}

func calcHitsAndHitsOutOfOrder(userCodes, boardCodes []int) (hitsCheck, hitsOutOfOrderCheck []int){
	for i := range userCodes {
		for x := range boardCodes {
			if i == x && userCodes[i] == boardCodes[x] {
				hitsCheck = append(hitsCheck, x)
			}
		}
	}
	for i := range userCodes {
		for x := range boardCodes {
			if userCodes[i] == boardCodes[x] && !validate.Contains(hitsCheck, x){
				hitsOutOfOrderCheck = append(hitsOutOfOrderCheck, x)
			}
		}
	}
	return hitsCheck, hitsOutOfOrderCheck
}

func calcMisses(sliceRange, hitsCheck, hitsOutOfOrderCheck []int) (misses []int) {
	for x := range sliceRange {
		if !validate.Contains(hitsCheck, x) && !validate.Contains(hitsOutOfOrderCheck, x) {
			misses = append(misses, x)
		}
	}
	return misses
}

func calcRemainingMoves(b *Board) int {
	return b.Rows - b.checks
}