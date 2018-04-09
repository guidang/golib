package lib

import (
	"github.com/frustra/bbcode"
	"github.com/mozillazg/go-pinyin"
	"strings"
)

/**
将拼音首字母提出来
*/
func ToPinyinSlug(str string) string {
	py := pinyin.NewArgs()
	py.Style = pinyin.FirstLetter

	var slug string
	slug = pinyin.Slug(str, py)
	slug = strings.Replace(slug, "-", "", -1)

	return slug
}

/**
bbcode 转 html
*/
func BBCodeToHtml(msg string) string {
	compiler := bbcode.NewCompiler(true, true)
	return compiler.Compile(msg)
}
