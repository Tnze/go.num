package zh

var nums = [...]string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
var secs = [...]string{"", "万", "亿", "万亿", "亿亿"}
var chns = [...]string{"", "十", "百", "千"}

func String(num uint64) (str string) {
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

func secString(sec uint64) string {
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
