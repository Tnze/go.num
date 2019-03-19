package en

import "strings"

//Spell a number in English
func Spell(num uint64) string {
	return spell(num, 0)
}

var numbers = [20]string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
	"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var ty = [10]string{"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
var sections = [4]string{"thousand", "million", "billion"}

func spell(num uint64, section int) string {
	if num == 0 {
		return numbers[0]
	}
	var sb strings.Builder
	var needSec bool
	if num >= 1000 {
		sb.WriteString(" " + spell(num/1000, section+1))
		num %= 1000
	}
	if num >= 100 {
		sb.WriteString(" " + numbers[num/100] + " " + "hundred")
		num %= 100
		needSec = true
	}
	if num >= 20 {
		sb.WriteString(" " + ty[num/10])
		num %= 10
		needSec = true
	}
	if num > 0 {
		sb.WriteString(" " + numbers[num])
		needSec = true
	}
	if needSec && section > 0 {
		sb.WriteString(" " + sections[section-1])
	}
	return sb.String()[1:]
}
