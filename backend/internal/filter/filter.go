package filter

import (
	"regexp"
	"unicode"
)

// Filter 过滤器
type Filter struct{}

// New 创建过滤器
func New() *Filter {
	return &Filter{}
}

// ByLength 按长度过滤
func (f *Filter) ByLength(password string, minLen, maxLen int) bool {
	len := len(password)
	return len >= minLen && len <= maxLen
}

// ByCharType 按字符类型过滤
func (f *Filter) ByCharType(password string, requireUpper, requireLower, requireDigit, requireSymbol bool) bool {
	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSymbol := false

	for _, ch := range password {
		if unicode.IsUpper(ch) {
			hasUpper = true
		} else if unicode.IsLower(ch) {
			hasLower = true
		} else if unicode.IsDigit(ch) {
			hasDigit = true
		} else if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
			hasSymbol = true
		}
	}

	if requireUpper && !hasUpper {
		return false
	}
	if requireLower && !hasLower {
		return false
	}
	if requireDigit && !hasDigit {
		return false
	}
	if requireSymbol && !hasSymbol {
		return false
	}

	return true
}

// ByRegex 按正则表达式过滤
func (f *Filter) ByRegex(password string, pattern string) bool {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}
	return re.MatchString(password)
}

// ByCommonWeakness 过滤常见弱点
func (f *Filter) ByCommonWeakness(password string) bool {
	// 过滤连续重复字符
	for i := 0; i < len(password)-2; i++ {
		if password[i] == password[i+1] && password[i+1] == password[i+2] {
			return false
		}
	}

	// 过滤连续序列（如123、abc）
	if isSequential(password) {
		return false
	}

	return true
}

// isSequential 检查是否为连续序列
func isSequential(s string) bool {
	if len(s) < 3 {
		return false
	}

	for i := 0; i < len(s)-2; i++ {
		if s[i]+1 == s[i+1] && s[i+1]+1 == s[i+2] {
			return true
		}
	}

	return false
}

// CalcEntropy 计算密码熵值
func (f *Filter) CalcEntropy(password string) float64 {
	charsetSize := 0

	hasLower := false
	hasUpper := false
	hasDigit := false
	hasSymbol := false

	for _, ch := range password {
		if unicode.IsLower(ch) {
			hasLower = true
		} else if unicode.IsUpper(ch) {
			hasUpper = true
		} else if unicode.IsDigit(ch) {
			hasDigit = true
		} else {
			hasSymbol = true
		}
	}

	if hasLower {
		charsetSize += 26
	}
	if hasUpper {
		charsetSize += 26
	}
	if hasDigit {
		charsetSize += 10
	}
	if hasSymbol {
		charsetSize += 32
	}

	if charsetSize == 0 {
		return 0
	}

	// Entropy = log2(charset_size) * password_length
	var entropy float64
	for i := 0; i < len(password); i++ {
		entropy += 1.0
	}
	entropy *= float64(charsetSize)

	// 简化计算
	return float64(charsetSize) * float64(len(password)) / 10
}
