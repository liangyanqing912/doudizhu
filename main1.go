package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var cards []int
	list := os.Args
	for i := 1; i < len(list); i++ {
		a, err := strconv.Atoi(list[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		cards = append(cards, a)
	}
	A := CardType(cards)
	fmt.Println(A)

}

//冒泡对牌进行排序
func sort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

//判断牌的类型
func CardType(cards []int) string {
	var q string
	if len(cards) == 0 {
		q = "错误"
		return q
	} else if singleCard(cards) {
		q = "单排"
		return q
	} else if alongCard(cards) {
		q = "顺子"
		return q
	} else if apairCard(cards) {
		q = "对子"
		return q
	} else if alongApairCard(cards) {
		q = "连对"
		return q
	} else if threeWithoutCard(cards) {
		q = "三不带"
		return q
	} else if threeBeltsOneCard(cards) {
		q = "三带一"
		return q
	} else if fourBeltsTwoCard(cards) {
		q = "三带二"
		return q
	} else if aircraftCard(cards) {
		q = "飞机"
		return q
	} else if bombCard(cards) {
		q = "炸弹"
		return q
	} else if towJokerCard(cards) {
		q = "王炸"
		return q
	} else {
		return "错误"
	}
}

//单牌
func singleCard(cards []int) (bl bool) {
	if len(cards) == 1 {
		bl = true
	}
	return
}

//对子
func apairCard(cards []int) (bl bool) {
	if len(cards) == 2 {
		if cards[0] == cards[1] {
			bl = true
		}
	}
	return
}

//连对
func alongApairCard(cards []int) (bl bool) {
	bl = true
	if len(cards) < 6 || len(cards) > 24 || len(cards)%2 != 0 {
		bl = false
		return
	}
	cards = sort(cards)
	for _, v := range cards {
		if v == 15 || v == 16 || v == 17 {
			bl = false
			return

		}
	}
	for i := 0; i < len(cards); i = i + 2 {
		if cards[i] != cards[i+1] {
			bl = false
			break
		}
		if i < len(cards)-2 {
			if cards[i+2]-cards[i] != 1 {
				bl = false
				break
			}
		}
	}
	return
}

//三不带
func threeWithoutCard(cards []int) (bl bool) {
	if len(cards) == 3 {
		if cards[0] == cards[1] && cards[0] == cards[2] {
			bl = true
		}
	}
	return
}

//三带一
func threeBeltsOneCard(cards []int) (bl bool) {
	if len(cards) == 4 {
		cards = sort(cards)
		a := cards[0]
		b := cards[1]
		c := cards[2]
		d := cards[3]
		if a == b && a == c && a != d {
			bl = true
		} else if d == c && d == b && d != a {
			bl = true
		}
	}
	return
}

//顺子
func alongCard(cards []int) (bl bool) {
	bl = true
	if len(cards) < 5 || len(cards) > 12 {
		bl = false
		return
	}
	cards = sort(cards)
	for i := 0; i < len(cards)-1; i++ {
		a := cards[i]
		b := cards[i+1]
		//大小王，2不能加入连牌
		if a == 15 || a == 16 || a == 17 || b == 15 || b == 16 || b == 17 {
			bl = false
			break
		} else {
			if b-a != 1 {
				bl = false
				break
			}
		}
	}
	return
}

//炸弹
func bombCard(cards []int) (bl bool) {
	if len(cards) == 4 {
		if cards[0] == cards[1] && cards[0] == cards[2] && cards[0] == cards[3] {
			bl = true
		}
	}
	return
}

//四带二
func fourBeltsTwoCard(cards []int) (bl bool) {
	if len(cards) == 6 {
		cards = sort(cards)
		for i := 0; i < 3; i++ {
			a := cards[i]
			b := cards[i+1]
			c := cards[i+2]
			d := cards[i+3]
			if a == b && a == c && a == d {
				bl = true
				break
			}
		}
	}
	return
}

//是否是飞机牌
func aircraftCard(cards []int) (bl bool) {
	bl = true
	if len(cards) < 6 {
		bl = false
		return
	}
	cards = sort(cards)
	bl = aircraftBeltsCard(cards) || aircraftWithOutCard(cards)
	return
}

//飞机不带
func aircraftWithOutCard(cards []int) (bl bool) {
	bl = true
	for _, v := range cards {
		if v == 16 || v == 17 {
			bl = false
			return
		}
	}
	for i := 0; i < len(cards); i = i + 3 {
		if cards[i] != cards[i+1] || cards[i] != cards[i+2] || cards[i+1] != cards[i+2] {
			bl = false
			break
		}
		if i < len(cards)-3 {
			if cards[i+3]-cards[i] != 1 {
				bl = false
				break
			}
		}
	}
	return
}

//飞机带
func aircraftBeltsCard(cards []int) (bl bool) {
	bl = true
	var grade = make(map[int]int)
	for i := 0; i < len(cards); i++ {
		grade[cards[i]] += 1
	}
	fmt.Printf("每张牌的命中次数%v\n", grade)
	var hitThree []int
	var hitTwo []int
	var hitOne []int
	for index, v := range grade {
		if v == 3 {
			hitThree = append(hitThree, index)
		} else if v == 2 {
			hitTwo = append(hitTwo, index)
		} else if v == 1 {
			hitOne = append(hitOne, index)
		}
	}
	for i := 0; i < len(hitThree)-1; i++ {
		if hitThree[i+1]-hitThree[i] != 1 {
			bl = false
			break
		}
	}
	if len(hitTwo) == 0 {
		if len(hitOne) != len(hitThree) {
			bl = false
		}
	} else {
		if len(hitTwo) != len(hitThree) {
			bl = false
		}
	}
	return
}

//王炸
func towJokerCard(cards []int) (bl bool) {
	if len(cards) == 2 {
		if cards[0] == 99 || cards[0] == 100 {
			if cards[0]+cards[1] == 199 {
				bl = true
			}
		}

	}
	return
}
