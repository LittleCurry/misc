package misc

import "regexp"

func MonitorPhoneNum() []string {
	return []string{"17740808015", "18001612148"}
}

func CheckPhoneNum(phone string) bool {

	// 移动
	rgx1 := regexp.MustCompile("^((13[4-9])|(147)|(15[0-2,7-9])|(178)|(18[2-4,7-8]))\\d{8}|(1705)\\d{7}$")
	// 联通
	rgx2 := regexp.MustCompile("^((13[0-2])|(145)|(15[5-6])|(176)|(18[5,6]))\\d{8}|(1709)\\d{7}$")
	// 电信
	rgx3 := regexp.MustCompile("^((133)|(153)|(177)|(18[0,1,9]))\\d{8}$")
	if !rgx1.MatchString(phone) && !rgx2.MatchString(phone) && !rgx3.MatchString(phone) {
		return false
	}
	return true
}
