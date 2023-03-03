package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase(t *testing.T) {
	t.Parallel()
	assertWithTest := assert.New(t)

	var params1 = case1()
	var params2 = case2()
	var params3 = case3()
	var params4 = case4()
	var params5 = case5()
	var params6 = case6()
	var params7 = case7()

	result1 := getProducts(params1.param1, params1.param2, params1.param3)
	result2 := getProducts(params2.param1, params2.param2, params2.param3)
	result3 := getProducts(params3.param1, params3.param2, params3.param3)
	result4 := getProducts(params4.param1, params4.param2, params4.param3)
	result5 := getProducts(params5.param1, params5.param2, params5.param3)
	result6 := getProducts(params6.param1, params6.param2, params6.param3)
	result7 := getProducts(params7.param1, params7.param2, params7.param3)

	assertWithTest.ElementsMatch(result1, case1_result())
	assertWithTest.ElementsMatch(result2, case2_result())
	assertWithTest.ElementsMatch(result3, case3_result())
	assertWithTest.ElementsMatch(result4, case4_result())
	assertWithTest.ElementsMatch(result5, case5_result())
	assertWithTest.ElementsMatch(result6, case6_result())
	assertWithTest.ElementsMatch(result7, case7_result())
}
