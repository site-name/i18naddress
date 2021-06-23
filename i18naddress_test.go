package i18naddress

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLoadValidationData(t *testing.T) {
	data, err := LoadValidationData("us")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(data))
}

func TestLoadCountryData(t *testing.T) {
	db, data, err := loadCountryData("zz")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(db)
	fmt.Println(data)
}

func TestFilterDuplicate(t *testing.T) {
	type testCase struct {
		in     []string
		expect []string
	}
	cases := []testCase{
		{[]string{"1", "2", "3", "4", "5"}, []string{"1", "2", "3", "4", "5"}},
		{[]string{"1", "2", "3", "4", "5", "5"}, []string{"1", "2", "3", "4", "5"}},
		{[]string{"1", "2", "3", "4", "4"}, []string{"1", "2", "3", "4"}},
	}

	for i, ca := range cases {
		res := filterDuplicate(&ca.in)
		if !reflect.DeepEqual(*res, ca.expect) {
			t.Fatalf("Test case #%d failed: got %v, expect: %v", i+1, *res, ca.expect)
		} else {
			fmt.Println(*res)
		}
	}
}

func TestGetValidationRules(t *testing.T) {
	pr := &Params{}
	rules, err := GetValidationRules(pr)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(rules)
}
