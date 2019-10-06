package zh

var nums = [...]string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
var secs = [...]string{"", "万", "亿", "万亿", "亿亿"}
var chns = [...]string{"", "十", "百", "千"}

func String(num int64) string {

	if num == 0 {
		return "零"
	}
	if num < 0 {
		return "负" + String(-num)
	}

	var str string
	var pos int
	var zero bool
	for num > 0 {
		sec := num % 10000
		if zero {
			str = nums[0] + str
		}

		secstr := secString(sec)
		if sec != 0 {
			secstr += secs[pos]
		}
		str = secstr + str

		zero = (sec < 1000) && (sec > 0)

		num /= 10000
		pos++
	}

	return str
}

func secString(sec int64) string {
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
