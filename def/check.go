package def

// 判断是否为字母
func IsLetter(c int32) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

// 判断是否为数字
func IsDigit(c int32) bool {
	return c >= '0' && c <= '9'
}

// 检查用户名是否合法
func CheckUsername(str string) bool {
	for _, v := range str {
		if !IsLetter(v) {
			return false
		}
	}
	return true
}

// 检查密码
func CheckPassword(str string) bool {
	if len(str) < 6 {
		return false
	}
	for _, v := range str {
		if !(IsLetter(v) || IsDigit(v)) {
			return false
		}
	}
	return true
}
