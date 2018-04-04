package utils_test

import (
	"testing"
	"github.com/tukangkod/go-gobok/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetFunctionName(t *testing.T) {
	actualResult := utils.GetFunctionName(TestGetFunctionName)
	expectedResult := "github.com/tukangkod/go-gobok/utils_test.TestGetFunctionName"

	assert.Equal(t, expectedResult, actualResult)
}

func BenchmarkGetFunctionName(b *testing.B) {
	for n := 0; n < b.N; n++ {
		utils.GetFunctionName(BenchmarkGetFunctionName)
	}
}
