package main

type testcase struct {
	param1 []TypeAndValues
	param2 []TypeAndValues
	param3 []Variant
}

func case1() testcase {
	var case1 testcase
	case1 = testcase{
		[]TypeAndValues{}, []TypeAndValues{
			{
				typename: "Style",
				values:   []string{"Classic", "Modern"},
			},
		}, []Variant{},
	}

	return case1
}
func case1_result() []Match {
	return []Match{
		[]string{"Classic"},
		[]string{"Modern"},
	}
}
func case2() testcase {
	var case2 testcase
	case2 = testcase{
		[]TypeAndValues{
			{
				typename: "Style",
				values:   []string{"Classic", "Modern"},
			},
		}, []TypeAndValues{
			{
				typename: "Size",
				values:   []string{"Medium", "Large", "Small"},
			},
		}, []Variant{
			[]string{"Classic"},
			[]string{"Modern"},
		},
	}

	return case2
}
func case2_result() []Match {
	return []Match{
		[]string{"Classic", "Large"},
		[]string{"Classic", "Small"},
		[]string{"Classic", "Medium"},
		[]string{"Modern", "Large"},
		[]string{"Modern", "Medium"},
		[]string{"Modern", "Small"},
	}
}
func case3() testcase {
	var case3 testcase
	case3 = testcase{
		[]TypeAndValues{
			{
				typename: "Style",
				values:   []string{"Classic", "Modern"},
			},
			{
				typename: "Size",
				values:   []string{"Medium", "Large", "Small"},
			},
		}, []TypeAndValues{
			{
				typename: "Color",
				values:   []string{"Red", "Blue"},
			},
		}, []Variant{
			[]string{"Classic", "Medium"},
			[]string{"Classic", "Small"},
			[]string{"Modern", "Large"},
			[]string{"Modern", "Small"},
		},
	}

	return case3
}
func case3_result() []Match {
	return []Match{
		[]string{"Classic", "Medium", "Red"},
		[]string{"Classic", "Medium", "Blue"},
		[]string{"Classic", "Small", "Red"},
		[]string{"Classic", "Small", "Blue"},
		[]string{"Modern", "Large", "Red"},
		[]string{"Modern", "Large", "Blue"},
		[]string{"Modern", "Small", "Red"},
		[]string{"Modern", "Small", "Blue"},
	}
}
func case4() testcase {
	var case4 testcase
	case4 = testcase{
		[]TypeAndValues{
			{
				typename: "Style",
				values:   []string{"Classic", "Modern"},
			},
			{
				typename: "Size",
				values:   []string{"Medium", "Large", "Small"},
			},
		}, []TypeAndValues{
			{
				typename: "Size",
				values:   []string{"Large", "Small"},
			},
		}, []Variant{
			[]string{"Classic", "Large"},
			[]string{"Classic", "Medium"},
			[]string{"Classic", "Small"},
			[]string{"Modern", "Large"},
			[]string{"Modern", "Medium"},
			[]string{"Modern", "Small"},
		},
	}

	return case4
}
func case4_result() []Match {
	return []Match{
		[]string{"Classic", "Large"},
		[]string{"Classic", "Small"},
		[]string{"Modern", "Large"},
		[]string{"Modern", "Small"},
	}
}
func case5() testcase {
	var case5 testcase
	case5 = testcase{
		[]TypeAndValues{
			{
				typename: "Style",
				values:   []string{"Classic", "Modern"},
			},
			{
				typename: "Size",
				values:   []string{"Medium", "Large", "Small"},
			},
			{
				typename: "Color",
				values:   []string{"Red", "Blue"},
			},
		}, []TypeAndValues{
			{
				typename: "Size",
				values:   []string{"Large", "Small", "ExtraLarge"},
			},
		}, []Variant{
			[]string{"Classic", "Large", "Blue"},
			[]string{"Classic", "Medium", "Red"},
			[]string{"Classic", "Medium", "Blue"},
			[]string{"Modern", "Large", "Red"},
			[]string{"Modern", "Medium", "Blue"},
			[]string{"Modern", "Small", "Red"},
			[]string{"Modern", "Small", "Blue"},
		},
	}

	return case5
}
func case5_result() []Match {
	return []Match{
		[]string{"Classic", "Large", "Blue"},
		[]string{"Classic", "ExtraLarge", "Red"},
		[]string{"Classic", "ExtraLarge", "Blue"},
		[]string{"Modern", "Large", "Red"},
		[]string{"Modern", "Small", "Red"},
		[]string{"Modern", "Small", "Blue"},
		[]string{"Modern", "ExtraLarge", "Red"},
		[]string{"Modern", "ExtraLarge", "Blue"},
	}
}
func case6() testcase {
	var case6 testcase
	case6 = testcase{
		[]TypeAndValues{
			{
				typename: "Style",
				values:   []string{"Classic", "Modern"},
			},
			{
				typename: "Size",
				values:   []string{"Medium", "Large", "Small"},
			},
			{
				typename: "Color",
				values:   []string{"Red", "Blue"},
			},
		}, []TypeAndValues{
			{
				typename: "Color",
			},
		}, []Variant{
			[]string{"Classic", "Large", "Blue"},
			[]string{"Classic", "Medium", "Red"},
			[]string{"Classic", "Medium", "Blue"},
			[]string{"Modern", "Large", "Red"},
			[]string{"Modern", "Small", "Red"},
			[]string{"Modern", "Small", "Blue"},
		},
	}

	return case6
}
func case6_result() []Match {
	return []Match{
		[]string{"Classic", "Large"},
		[]string{"Classic", "Medium"},
		[]string{"Modern", "Large"},
		[]string{"Modern", "Small"},
	}
}
func case7() testcase {
	var case7 testcase
	case7 = testcase{
		[]TypeAndValues{
			{
				typename: "Style",
				values:   []string{"Classic", "Modern"},
			},
			{
				typename: "Size",
				values:   []string{"Medium", "Large", "Small"},
			},
			{
				typename: "Color",
				values:   []string{"Red", "Blue"},
			},
		}, []TypeAndValues{
			{
				typename: "Size",
				values:   []string{"Large", "Medium", "Small"},
			},
		}, []Variant{
			[]string{"Classic", "Large", "Blue"},
			[]string{"Classic", "Medium", "Red"},
			[]string{"Classic", "Medium", "Blue"},
			[]string{"Modern", "Large", "Red"},
			[]string{"Modern", "Medium", "Blue"},
			[]string{"Modern", "Small", "Red"},
			[]string{"Modern", "Small", "Blue"},
		},
	}

	return case7
}
func case7_result() []Match {
	return []Match{
		[]string{"Classic", "Large", "Blue"},
		[]string{"Classic", "Medium", "Red"},
		[]string{"Classic", "Medium", "Blue"},
		[]string{"Modern", "Large", "Red"},
		[]string{"Modern", "Medium", "Blue"},
		[]string{"Modern", "Small", "Red"},
		[]string{"Modern", "Small", "Blue"},
	}
}
