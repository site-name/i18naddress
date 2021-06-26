package i18naddress

import "regexp"

func stringInSlice(s string, slice *[]string) bool {
	for _, str := range *slice {
		if s == str {
			return true
		}
	}

	return false
}

func regexesToStrings(in []*regexp.Regexp) *[]string {
	res := []string{}
	for _, rg := range in {
		res = append(res, rg.String())
	}

	return &res
}

func filterDuplicate(slice *[]string) *[]string {
	meetMap := make(map[string]bool)
	res := []string{}
	for _, str := range *slice {
		if _, met := meetMap[str]; !met {
			res = append(res, str)
			meetMap[str] = true
		}
	}

	return &res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func filterSlice(slice []string, filter func(s string) bool) []string {
	res := []string{}
	for _, str := range slice {
		if filter(str) {
			res = append(res, str)
		}
	}

	return res
}