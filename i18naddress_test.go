package i18naddress

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestLoadValidationData(t *testing.T) {
	reader, err := LoadValidationData("vn")
	if err != nil {
		t.Fatal(err)
	}

	buf := bytes.Buffer{}
	_, err = io.Copy(&buf, reader)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(buf.String())
}

func TestLoadCountryData(t *testing.T) {
	db, data, err := loadCountryData("vn")
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
		res := filterDuplicate(ca.in)
		if !reflect.DeepEqual(res, ca.expect) {
			t.Fatalf("Test case #%d failed: got %v, expect: %v", i+1, res, ca.expect)
		} else {
			fmt.Println(res)
		}
	}
}

func TestGetValidationRules(t *testing.T) {
	pr := &Params{
		CountryCode: "vn",
	}
	rules, err := GetValidationRules(pr)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(rules)
}

func TestAssets(t *testing.T) {
	data, err := assets.ReadFile(assetsPrefix + "/ac.json")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(data))
}
