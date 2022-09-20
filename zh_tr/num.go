package zh_tr

import (
	"fmt"
	"io"
)

type Uint64 uint64

var nums = [...]string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"}
var secs = [...]string{"", "万", "亿", "万亿", "亿亿"}
var chns = [...]string{"", "拾", "佰", "仟"}

func (num Uint64) String() (str string) {
	if num == 0 { // 为零时直接返回
		return nums[0]
	}

	var pos int       // 当前节数
	var needzero bool // 下一个节是否需要补零
	for num > 0 {
		sec := num % 10000

		if needzero { // 需要时补零
			str = nums[0] + str
		}

		secstr := secString(sec)
		if sec != 0 { // 根据需要添加节权
			str = secstr + secs[pos] + str
		} else {
			str = secstr + str
		}

		needzero = (sec < 1000) && (sec > 0)

		num /= 10000
		pos++
	}

	return
}

func secString(sec Uint64) string {
	var str string
	var pos int
	var zero = true

	for sec > 0 {
		v := sec % 10
		if v == 0 {
			if sec == 0 || !zero {
				zero = true
				str = nums[v] + str
			}
		} else {
			zero = false
			ins := nums[v] + chns[pos]
			str = ins + str
		}

		pos++
		sec /= 10
	}

	return str
}

var pairs = map[rune]struct {
	value  uint64
	isUnit bool
}{
	'十': {10, false},
	'百': {100, false},
	'千': {1000, false},
	'万': {10000, true},
	'亿': {100000000, true},
}

var words = map[rune]uint64{
	'零': 0, '一': 1, '二': 2, '三': 3, '四': 4,
	'五': 5, '六': 6, '七': 7, '八': 8, '九': 9,
}

func (num *Uint64) Scan(state fmt.ScanState, verb rune) error {
	var number, sec uint64

	for {
		// 读取一个字
		v, _, err := state.ReadRune()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if n, ok := words[v]; ok {
			// 是数字
			number = n
		} else if unit, ok := pairs[v]; ok {
			// 是单位
			if unit.isUnit {
				// 是权
				*num += Uint64((sec + number) * unit.value)
				sec = 0
			} else {
				sec += (number * unit.value)
			}

			number = 0

		} else {
			// 其他字符
			break
		}
	}

	*num += Uint64(sec + number)
	return nil
}
