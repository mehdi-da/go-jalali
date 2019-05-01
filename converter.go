package jalali

import (
	"time"
)

func GtoJ(t time.Time) (y int, m int, d int) {
	gyear := t.Year()
	gmonth := int(t.Month())
	gday := t.Day()

	_gl := [...]int{0, 31, 60, 91, 121, 152, 182, 213, 244, 274, 305, 335}
	_g := [...]int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}

	ddj := 10
	var gd, jyear, jmonth, jday, gmod int
	if GLeapYear(gyear - 1) {
		ddj = 11
	}
	if GLeapYear(gyear) {
		gd = _gl[gmonth-1] + gday
	} else {
		gd = _g[gmonth-1] + gday
	}
	if gd > 79 {
		jyear = gyear - 621
		gd = gd - 79
		if gd <= 186 {
			gmod = gd % 31
			if gmod == 0 {
				jday = 31
				jmonth = int(gd / 31)
			} else {
				jday = gmod
				jmonth = int(gd/31) + 1
			}
		} else {
			gd = gd - 186
			gmod = gd % 30
			if gmod == 0 {
				jday = 30
				jmonth = int(gd/30) + 6
			} else {
				jday = gmod
				jmonth = int(gd/30) + 7
			}
		}
	} else {
		jyear = gyear - 622
		gd = gd + ddj
		gmod = gd % 30
		if gmod == 0 {
			jday = 30
			jmonth = int(gd/30) + 9
		} else {
			jday = gmod
			jmonth = int(gd/30) + 10
		}
	}
	return jyear, jmonth, jday
}

func GLeapYear(y int) bool {
	if (y%4 == 0) && ((y%100 != 0) || (y%400 == 0)) {
		return true
	} else {
		return false
	}
}

func JLeapYear(y int) bool {
	ary := [...]int{1, 5, 9, 13, 17, 22, 26, 30}
	result := false
	b := y % 33
	for _, v := range ary {
		if b == v {
			result = true
		}
	}
	return result
}
