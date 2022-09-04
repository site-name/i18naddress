package i18naddress

import (
	"embed"
	"regexp"

	"github.com/samber/lo"
)

//go:embed google-i18n-address/i18naddress/data
var assets embed.FS

const assetsPrefix = "google-i18n-address/i18naddress/data"

func RegexesToStrings(in []*regexp.Regexp) []string {
	return lo.Map(in, func(exp *regexp.Regexp, _ int) string {
		if exp != nil {
			return exp.String()
		}

		return ""
	})
}
