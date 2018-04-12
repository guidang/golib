package lib

import (
	"github.com/mozillazg/go-pinyin"
	"github.com/skiy/bbcode"
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

	//转 table
	compiler.SetTag("table", func(node *bbcode.BBCodeNode) (*bbcode.HTMLTag, bool) {
		out := bbcode.NewHTMLTag("")
		out.Name = "table"
		return out, true
	})

	//转 tr
	compiler.SetTag("tr", func(node *bbcode.BBCodeNode) (*bbcode.HTMLTag, bool) {
		out := bbcode.NewHTMLTag("")
		out.Name = "tr"

		return out, true
	})
	//转 td
	compiler.SetTag("td", func(node *bbcode.BBCodeNode) (*bbcode.HTMLTag, bool) {
		out := bbcode.NewHTMLTag("")
		out.Name = "td"

		return out, true
	})

	//ul
	compiler.SetTag("list", func(node *bbcode.BBCodeNode) (*bbcode.HTMLTag, bool) {
		out := bbcode.NewHTMLTag("")
		out.Name = "ul"

		return out, true
	})

	//text-align
	compiler.SetTag("align", func(node *bbcode.BBCodeNode) (*bbcode.HTMLTag, bool) {
		out := bbcode.NewHTMLTag("")
		out.Name = "div"
		value := node.GetOpeningTag().Value
		if value != "" {
			out.Attrs["style"] = "text-align: " + value
		}
		return out, true
	})

	//li -> 将 [*] 转为 li
	compiler.SetTag("*", func(node *bbcode.BBCodeNode) (*bbcode.HTMLTag, bool) {
		out := bbcode.NewHTMLTag("")
		out.Name = "li"

		return out, true
	})

	//font -> 清空 font
	compiler.SetTag("font", func(node *bbcode.BBCodeNode) (*bbcode.HTMLTag, bool) {
		out := bbcode.NewHTMLTag("")
		out.Name = ""

		return out, true
	})
	return compiler.Compile(msg)
}
