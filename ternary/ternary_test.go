package ternary_test

import (
	"testing"

	"github.com/mylxsw/go-utils/assert"
	"github.com/mylxsw/go-utils/ternary"
)

func TestIfElse(t *testing.T) {

	assert.Equal(t, "positive", ternary.IfElse(true, "positive", "negative"))
	assert.Equal(t, "negative", ternary.IfElse(false, "positive", "negative"))

	assert.Equal(t, "positive", ternary.IfElseLazy(true, func() string { return "positive" }, func() string { return "negative" }))
	assert.Equal(t, "negative", ternary.IfElseLazy(false, func() string { return "positive" }, func() string { return "negative" }))
}

func TestIf(t *testing.T) {
	assert.Equal(t, "positive", ternary.If(true, "positive", ""))
	assert.Equal(t, "", ternary.If(false, "positive", ""))

	assert.Equal(t, "positive", ternary.IfLazy(true, func() string { return "positive" }, func() string { return "" }))
	assert.Equal(t, "", ternary.IfLazy(false, func() string { return "positive" }, func() string { return "" }))
}
