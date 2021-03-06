// Extensions to the go-check unittest framework.
//
// NOTE: see https://github.com/go-check/check/pull/6 for reasons why these
// checkers live here.
package gocheck2

import (
	"bytes"

	. "gopkg.in/check.v1"
)

// -----------------------------------------------------------------------
// IsTrue / IsFalse checker.

type isBoolValueChecker struct {
	*CheckerInfo
	expected bool
}

func (checker *isBoolValueChecker) Check(
	params []interface{},
	names []string) (
	result bool,
	error string) {

	obtained, ok := params[0].(bool)
	if !ok {
		return false, "Argument to " + checker.Name + " must be bool"
	}

	return obtained == checker.expected, ""
}

// The IsTrue checker verifies that the obtained value is true.
//
// For example:
//
//     c.Assert(value, IsTrue)
//
var IsTrue Checker = &isBoolValueChecker{
	&CheckerInfo{Name: "IsTrue", Params: []string{"obtained"}},
	true,
}

// The IsFalse checker verifies that the obtained value is false.
//
// For example:
//
//     c.Assert(value, IsFalse)
//
var IsFalse Checker = &isBoolValueChecker{
	&CheckerInfo{Name: "IsFalse", Params: []string{"obtained"}},
	false,
}

// -----------------------------------------------------------------------
// BytesEquals checker.

type bytesEquals struct{}

func (b *bytesEquals) Check(params []interface{}, names []string) (bool, string) {
	if len(params) != 2 {
		return false, "BytesEqual takes 2 bytestring arguments"
	}
	b1, ok1 := params[0].([]byte)
	b2, ok2 := params[1].([]byte)

	if !(ok1 && ok2) {
		return false, "Arguments to BytesEqual must both be bytestrings"
	}

	if bytes.Equal(b1, b2) {
		return true, ""
	}
	return false, "Byte arrays were different"
}

func (b *bytesEquals) Info() *CheckerInfo {
	return &CheckerInfo{
		Name:   "BytesEquals",
		Params: []string{"bytes_one", "bytes_two"},
	}
}

// BytesEquals checker compares two bytes sequence using bytes.Equal.
//
// For example:
//
//     c.Assert(b, BytesEquals, []byte("bar"))
//
// Main difference between DeepEquals and BytesEquals is that BytesEquals treats
// `nil` as empty byte sequence while DeepEquals doesn't.
//
//     c.Assert(nil, BytesEquals, []byte("")) // succeeds
//     c.Assert(nil, DeepEquals, []byte("")) // fails
var BytesEquals = &bytesEquals{}
