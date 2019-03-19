package en

import (
	"testing"
)

func TestSpell(t *testing.T) {
	for _, v := range []struct {
		num uint64
		str string
	}{
		{0, "zero"},
		{1, "one"},
		{2, "two"},
		{3, "three"},
		{4, "four"},
		{5, "five"},
		{6, "six"},
		{7, "seven"},
		{8, "eight"},
		{9, "nine"},
		{10, "ten"},
		{11, "eleven"},
		{12, "twelve"},
		{13, "thirteen"},
		{14, "fourteen"},
		{15, "fifteen"},
		{16, "sixteen"},
		{17, "seventeen"},
		{18, "eighteen"},
		{19, "nineteen"},
		{20, "twenty"},
		{30, "thirty"},
		{40, "forty"},
		{50, "fifty"},
		{60, "sixty"},
		{70, "seventy"},
		{80, "eighty"},
		{90, "ninety"},
		{23, "twenty three"},
		{35, "thirty five"},
		{100, "one hundred"},
		{200, "two hundred"},
		{1000, "one thousand"},
		{1001, "one thousand one"},
		{1011, "one thousand eleven"},
		{1000000, "one million"},
		{1000000000, "one billion"},
		{101, "one hundred one"},
		{3261340757, "three billion two hundred sixty one million three hundred forty thousand seven hundred fifty seven"},
	} {
		s := Spell(v.num)
		if s != v.str {
			t.Logf("Spell %d as %q, want %q", v.num, s, v.str)
			t.Fail()
		}
	}
}
