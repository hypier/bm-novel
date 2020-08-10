package draft

import (
	"strconv"
	"unicode"
)

type position struct {
	begin int
	end   int
}

func (p *position) isNull() bool {
	if p.begin == -1 || p.end == -1 {
		return true
	}

	return false
}

type positions struct {
	head position
	tail position
}

func (ps *positions) getSpitePos() (int, error) {
	if ps.head.begin > 3 {
		return ps.head.begin, nil
	}

	if !ps.tail.isNull() {
		return ps.tail.end, nil
	}

	return 0, ErrNoPosition
}

func null() *position {
	return &position{-1, -1}
}

func compare(p1, p2 position) int {

	if p1.begin > p2.begin {
		return 1
	} else if p1.begin < p2.begin {
		return -1
	} else {
		return 0
	}
}

func cNumberToInt(s string) (int, bool) {
	// 1.全数字转换
	if numberInt, err := strconv.Atoi(s); err == nil {
		return numberInt, true
	}

	cNum := map[rune]int{'零': 0, '一': 1, '二': 2, '两': 2, '三': 3, '四': 4, '五': 5, '六': 6, '七': 7, '八': 8, '九': 9, '十': 10, '百': 100, '千': 1000, '万': 10000, '亿': 100000000, '壹': 1, '贰': 2, '叁': 3, '肆': 4, '伍': 5, '陆': 6, '柒': 7, '捌': 8, '玖': 9, '拾': 10, '佰': 100, '仟': 1000}
	total, temp := 0, 0

	for i, c := range s {
		// 2.判断是否是全中文
		if !unicode.Is(unicode.Han, c) {
			return 0, false
		}

		val := cNum[c]

		// 3.判断是否是单位
		if val >= 10 {
			// 判断首位是否是单位
			if i == 0 {
				total = val
			} else {
				total += temp * val
			}

			temp = 0
		} else {
			temp = val
		}
	}

	// 未尾数字处理
	if temp > 0 {
		total += temp
	}

	return total, true
}

func counter(value, step int) func() int {
	num := value
	return func() int {
		num += step
		return num
	}
}
