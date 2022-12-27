package verify

import (
	"regexp"
	"strconv"
)

// IsEmpty 校验是否为空
func IsEmpty(params interface{}) bool {
	switch v := params.(type) {
	case int, int8, int16, int32, int64:
		if v == 0 {
			return false
		}
		return true
	case string:
		if len(v) == 0 || v == "" {
			return false
		}
		return true
	case []int, []int8, []int16, []int32, []int64:
		if v == 0 {
			return false
		}
		return true
	case []string:
		if len(v) == 0 {
			return false
		}
		return true
	case []uint, []uint8, []uint16, []uint32, []uint64:
		if v == 0 {
			return false
		}
		return true
	case float32, float64:
		if v == 0 {
			return false
		}
		return true
	}
	return true
}

// ValidateUserName 校验参数用户名长度
func ValidateUserName(username string) bool {
	if username == "" {
		return false
	}

	if len(username) > 15 || len(username) < 2 {
		return false
	}
	var regex = "^[\u4E00-\u9FA5A-Za-z0-9]+$"
	matches := Matches(regex, username)
	if matches {
		return true
	} else {
		return false
	}
}

// Matches 正则校验
func Matches(regex string, str string) bool {
	compile, err := regexp.Compile(regex)
	if err != nil {
		return false
	}

	matchString := compile.MatchString(str)
	return matchString
}

// MatchStr 匹配字符串
func MatchStr(str string) string {
	//<p\b[^<>]>?
	submatch := regexp.MustCompile("<p>[\\s\\S]*.*?</p>")
	stringSubmatch := submatch.FindStringSubmatch(str)
	stringSub := ""
	for _, i2 := range stringSubmatch {
		stringSub = i2
	}
	return stringSub
}

func VerifyPhoneFormat(Phone string) bool {
	pattern := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(Phone)
}

// VerifyEmailFormat 校验邮箱
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func MatcheNumber(regex string, number int) bool {
	compile, err := regexp.Compile(regex)
	if err != nil {
		return false
	}

	match := compile.MatchString(strconv.Itoa(number))
	return match
}

// IsNumber 校验是否是数字
func IsNumber(number int) bool {
	return MatcheNumber("^-?\\d+$", number)
}

// IsPwdSimple 密码校验：字母+数字
func IsPwdSimple(pwd string) bool {

	if IsEmpty(pwd) {
		return true
	}
	if len(pwd) < 6 {
		return true
	}

	if Matches("(.*)[a-zA-z](.*)", pwd) && Matches("(.*)\\d+(.*)", pwd) {
		return false
	}
	return true
}
