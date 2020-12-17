package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	somethingList := []string{
		"サムス",
		"サムライスピリッツ",
		"サムイ島",
		"サムギョプサル",
		"佐村河内守",
		"侍",
		"サムゲタン",
		"サムネイル",
		"サムソン",
		"SAMSUNG",
		"サムシングエルス",
		"サムシングフォー",
		"サムシングブルー",
		"サム・スミス",
		"SUM関数",
	}
	rand.Seed(time.Now().UnixNano())
	max := len(somethingList)
	fmt.Printf("寒すぎて、%sになった\n", somethingList[rand.Intn(max)])
}
