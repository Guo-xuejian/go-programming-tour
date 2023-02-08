package word

import (
	"strings"
	"unicode"
)

// ToUpper 实现单词转全大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower 实现单词转全小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// UnderscoreToUpperCamelCase 实现下划线转大写驼峰
func UnderscoreToUpperCamelCase(s string) string {
	// 替换所有的下划线为空格，利用 ToTitle 将首字母全部大写，最后替换空格即可
	s = strings.Replace(s, "_", " ", -1)
	s = strings.ToTitle(s)
	return strings.Replace(s, " ", "", -1)
}

// UnderscoreToLowerCamelCase 实现下划线转小写驼峰
func UnderscoreToLowerCamelCase(s string) string {
	// 核转大写驼峰类似，转换完成后堆首字母进行处理即可
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CamelCaseToUnderscore 实现驼峰转下划线
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 { // 首字母大小写都行
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) { // 大写字母先加入一个下划线，然后添加其小写形式字母
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
