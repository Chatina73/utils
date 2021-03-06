package utils

import (
	"fmt"
	"syscall"
	"testing"

	"github.com/akfaew/test"
)

func TestSlash(t *testing.T) {
	tests := []struct {
		input string
		ret1  string
		ret2  string
	}{
		{"", "", ""},
		{"test", "test", ""},
		{"test/", "test", ""},
		{"test///", "test", "//"},
		{"test/case", "test", "case"},
		{"test/case/a/b/c", "test", "case/a/b/c"},
		{"/case/a/b/c", "", "case/a/b/c"},
	}

	for _, tc := range tests {
		ret1, ret2 := Slash(tc.input)
		test.EqualStr(t, ret1, tc.ret1)
		test.EqualStr(t, ret2, tc.ret2)
	}
}

func TestHasElem(t *testing.T) {
	tests := []struct {
		arr  interface{}
		elem interface{}
		ret  bool
	}{
		{"", "", false},
		{5, "", false},
		{[]int{}, 2, false},
		{[]int{1}, 1, true},
		{[]int{1}, 2, false},
		{[]int{1, 2}, 2, true},
		{[]int{1, 2, 3}, 2, true},
		{[]string{}, 2, false},
		{[]string{}, "", false},
		{[]string{"a"}, "", false},
		{[]string{"a"}, "a", true},
		{[]string{"a", "b"}, "b", true},
		{[]string{"a", "b", "c"}, "b", true},
	}

	for _, tc := range tests {
		ret := HasElem(tc.arr, tc.elem)
		test.True(t, ret == tc.ret)
	}
}

func TestErrors(t *testing.T) {
	errs := []error{
		fmt.Errorf("First error: server error"),
		fmt.Errorf("Second error: %s", syscall.ENOPKG.Error()),
		fmt.Errorf("Third error: %s", syscall.ENOTCONN.Error()),
	}
	test.FixtureExtra(t, "Errors", Errors(errs).Error())
	test.NoError(t, Errors([]error{}))

	var el ErrorList
	test.NoError(t, el.Error())
	el.Append(errs[0])
	test.FixtureExtra(t, "One", el.Error().Error())
	el.Append(errs[1])
	el.Append(errs[2])
	test.FixtureExtra(t, "Three", el.Error().Error())
}

func TestSum(t *testing.T) {
	for _, s := range []string{"a", "", "://!%$"} {
		t.Run(Sum(s), func(t *testing.T) {
			test.Fixture(t, fmt.Sprintf("%s -> %s\n", s, Sum(s)))
		})
	}
}
