package utils_test

import (
	"testing"
	"github.com/tukangkod/go-gobok/utils"
	"github.com/stretchr/testify/assert"
)

// single tag
func TestTagMapSingleTag(t *testing.T) {
	actualResult := utils.TagMap("a=b")
	expectedResult := map[string]string(map[string]string{"a": "b"})

	assert.Equal(t, expectedResult, actualResult)
}

// multiple tag
func TestTagMapMultipleTag(t *testing.T) {
	actualResult := utils.TagMap("a=b&c=d&e=f&g=h&i=j&k=l")
	expectedResult := map[string]string(map[string]string{"a": "b", "c": "d", "e": "f", "g": "h", "i": "j", "k": "l"})

	assert.Equal(t, expectedResult, actualResult)
}

// empty string return empty map
func TestTagMapEmptyTag(t *testing.T) {
	actualResult := utils.TagMap("")
	expectedResult := map[string]string(map[string]string{})

	assert.Equal(t, expectedResult, actualResult)
}

// string without equal (=) will be ignored
func TestTagMapSingleWrongFormatTag(t *testing.T) {
	actualResult := utils.TagMap("abc")
	expectedResult := map[string]string(map[string]string{})

	assert.Equal(t, expectedResult, actualResult)
}

// string only wrong tag will be ignored
func TestTagMapWrongFormatTagTwo(t *testing.T) {
	actualResult := utils.TagMap("abc&def=ghi")
	expectedResult := map[string]string(map[string]string{"def": "ghi"})

	assert.Equal(t, expectedResult, actualResult)
}

// wrong string tag will be ignored
func TestTagMapExtraEqualTag(t *testing.T) {
	actualResult := utils.TagMap("abc=def=ghi")
	expectedResult := map[string]string(map[string]string{})

	assert.Equal(t, expectedResult, actualResult)
}
func TestTagMapExtraEqualTagTwo(t *testing.T) {
	actualResult := utils.TagMap("a==b")
	expectedResult := map[string]string(map[string]string{})

	assert.Equal(t, expectedResult, actualResult)
}

func TestTagMapMultipleAmpTag(t *testing.T) {
	actualResult := utils.TagMap("a&b&c&d")
	expectedResult := map[string]string(map[string]string{})

	assert.Equal(t, expectedResult, actualResult)
}

func TestTagMapExtraAmpTag(t *testing.T) {
	actualResult := utils.TagMap("abc=def&&ghi=jkl")
	expectedResult := map[string]string(map[string]string{"abc": "def", "ghi": "jkl"})

	assert.Equal(t, expectedResult, actualResult)
}

func TestTagMapMultipleAmpAndEqualOnlyTag(t *testing.T) {
	actualResult := utils.TagMap("&&&===")
	expectedResult := map[string]string(map[string]string{})

	assert.Equal(t, expectedResult, actualResult)
}

// MarshalMsg
// Empty
func TestMarshalMsgEmptyTag(t *testing.T) {
	msg := map[string]string(map[string]string{})

	actualResult := utils.MarshalMsg(msg)
	expectedResult := "{}"

	assert.Equal(t, expectedResult, actualResult)
}

// Single
func TestMarshalMsgSingleTag(t *testing.T) {
	msg := map[string]string(map[string]string{"a": "b"})

	actualResult := utils.MarshalMsg(msg)
	expectedResult := "{\"a\":\"b\"}"

	assert.Equal(t, expectedResult, actualResult)
}

// Multiple
func TestMarshalMsgMultipleTag(t *testing.T) {
	msg := map[string]string(map[string]string{"a": "b", "c": "d", "e": "f", "g": "h", "i": "j", "k": "l"})

	actualResult := utils.MarshalMsg(msg)
	expectedResult := "{\"a\":\"b\",\"c\":\"d\",\"e\":\"f\",\"g\":\"h\",\"i\":\"j\",\"k\":\"l\"}"

	assert.Equal(t, expectedResult, actualResult)
}
