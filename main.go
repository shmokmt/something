package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	somethingList := []string{
		"サムス",
		"寒川神社（さむかわじんじゃ）",
		"三浪津駅（さむなんじんえき）",
		"サム・ワーシントン",
		"沙村広明",
		"サムライスピリッツ",
		"サムイ島",
		"サムギョプサル",
		"佐村河内守",
		"侍",
		"侍ジャパン",
		"サムライジャック",
		"サムゲタン",
		"サムネイル",
		"サムソン",
		"SAMSUNG",
		"サムシングエルス",
		"サムシングフォー",
		"サムシングブルー",
		"サム・スミス",
		"サムライハート",
		"SUM関数",
		"SAML Single Sign-On",
		"サムソナイト",
		"thumb（親指）",
		"サムイル要塞",
		"作務衣",
		"サム山",
		"サムシャープ広場",
		"サムタビシ大聖堂",
		"サムタブロ修道院",
		"サムター要塞",
		"サムルノリ",
		"サム・アルトマン",
	}
	rand.Seed(time.Now().UnixNano())
	max := len(somethingList)
	fmt.Printf("寒すぎて、%sになった\n", somethingList[rand.Intn(max)])
}
