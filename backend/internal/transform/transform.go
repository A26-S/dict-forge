package transform

import (
	"strings"
	"unicode"
)

// Transformer 变形处理器
type Transformer struct{}

// New 创建变形处理器
func New() *Transformer {
	return &Transformer{}
}

// Capitalize 首字母大写
func (t *Transformer) Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

// ToUpper 全大写
func (t *Transformer) ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower 全小写
func (t *Transformer) ToLower(s string) string {
	return strings.ToLower(s)
}

// Reverse 字符串反转
func (t *Transformer) Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Leet 1337 变形
func (t *Transformer) Leet(s string) string {
	replacer := strings.NewReplacer(
		"a", "@",
		"e", "3",
		"i", "!",
		"o", "0",
		"s", "$",
		"t", "7",
		"l", "1",
	)
	return replacer.Replace(s)
}

// NumberReplace 数字替换
func (t *Transformer) NumberReplace(s string) []string {
	replacements := map[string][]string{
		"0": {"O", "零"},
		"1": {"I", "L", "l", "一"},
		"3": {"E"},
		"5": {"S"},
		"8": {"B"},
	}

	var results []string
	results = append(results, s)

	for old, news := range replacements {
		for _, new := range news {
			results = append(results, strings.ReplaceAll(s, old, new))
		}
	}

	return results
}

// MixedCase 混合大小写
func (t *Transformer) MixedCase(s string) []string {
	var results []string
	runes := []rune(s)

	// 生成所有可能的大小写组合（限制为2^n不超过1000)
	if len(runes) > 10 {
		return []string{s} // 太长则跳过
	}

	for i := 0; i < (1 << uint(len(runes))); i++ {
		combination := make([]rune, len(runes))
		for j, r := range runes {
			if (i & (1 << uint(j))) != 0 {
				combination[j] = unicode.ToUpper(r)
			} else {
				combination[j] = unicode.ToLower(r)
			}
		}
		results = append(results, string(combination))
	}

	return results
}
