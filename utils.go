package i18naddress

import (
	"regexp"

	"github.com/samber/lo"
)

const assetsPrefix = "data"

func RegexesToStrings(in []*regexp.Regexp) []string {
	return lo.Map(in, func(exp *regexp.Regexp, _ int) string {
		if exp != nil {
			return exp.String()
		}

		return ""
	})
}
