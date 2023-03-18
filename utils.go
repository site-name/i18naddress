package i18naddress

import (
	"embed"
	"regexp"

	"github.com/samber/lo"
)

//go:embed data
var assets embed.FS

const assetsPrefix = "data"

func RegexesToStrings(in []*regexp.Regexp) []string {
	return lo.Map(in, func(exp *regexp.Regexp, _ int) string {
		if exp != nil {
			return exp.String()
		}

		return ""
	})
}
