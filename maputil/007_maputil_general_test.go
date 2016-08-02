package maputil

import (
	"strings"
	"testing"
	"time"
)

func TestMapJoin(t *testing.T) {
	input := map[string]interface{}{
		`key1`: `value1`,
		`key2`: true,
		`key3`: 3,
	}

	output := Join(input, `=`, `&`)

	if output == `` {
		t.Error("Output should not be empty")
	}

	if !strings.Contains(output, `key1=value1`) {
		t.Errorf("Output should contain '%s'", `key1=value1`)
	}

	if !strings.Contains(output, `key2=true`) {
		t.Errorf("Output should contain '%s'", `key2=true`)
	}

	if !strings.Contains(output, `key3=3`) {
		t.Errorf("Output should contain '%s'", `key3=3`)
	}
}

func TestMapSplit(t *testing.T) {
	input := `key1=value1&key2=true&key3=3`

	output := Split(input, `=`, `&`)

	if len(output) == 0 {
		t.Error("Output should not be empty")
	}

	if v, ok := output[`key1`]; !ok || v != `value1` {
		t.Errorf("Output should contain key %s => '%s'", `key1`, `value1`)
	}

	if v, ok := output[`key2`]; !ok || v != `true` {
		t.Errorf("Output should contain key %s => '%s'", `key2`, `true`)
	}

	if v, ok := output[`key3`]; !ok || v != `3` {
		t.Errorf("Output should contain key %s => '%s'", `key3`, `3`)
	}
}

type SubtypeTester struct {
	A int
	B int `maputil:"b"`
}

type MyStructTester struct {
	Name                  string
	Subtype1              SubtypeTester
	Active                bool           `maputil:"active"`
	Subtype2              *SubtypeTester `maputil:"subtype2"`
	TimeTest              time.Duration
	IntTest               int32
	Properties            map[string]interface{}
	StrSliceTest          []string
	InterfaceStrSliceTest []string
	nonexported           int
}

func TestStructFromMap(t *testing.T) {
	input := map[string]interface{}{
		`Name`:           `Foo Bar`,
		`active`:         true,
		`should_not_set`: 4,
		`Subtype1`: map[string]interface{}{
			`A`: 1,
			`b`: 2,
		},
		`subtype2`: map[string]interface{}{
			`A`: 3,
			`b`: 4,
		},
		`TimeTest`: 15000000000,
		`IntTest`:  int64(5),
		`Properties`: map[string]interface{}{
			`first`:  1,
			`second`: true,
			`third`:  `three`,
		},
		`StrSliceTest`:          []string{`one`, `two`, `three`},
		`InterfaceStrSliceTest`: []interface{}{`one`, `two`, `three`},
	}

	output := MyStructTester{}

	if err := StructFromMap(input, &output); err == nil {
		if output.Name != `Foo Bar` {
			t.Errorf("output.Name; expected: %s, got: %v", `Foo Bar`, output.Name)
		}

		if !output.Active {
			t.Errorf("output.Active; expected: true, got: false")
		}

		if output.nonexported != 0 {
			t.Errorf("output.nonexported; expected: 0, got: %v", output.nonexported)
		}

		if output.Subtype1.A != 1 {
			t.Errorf("output.Subtype1.A; expected: 1, got: %v", output.Subtype1.A)
		}

		if output.Subtype1.B != 2 {
			t.Errorf("output.Subtype1.B; expected: 2, got: %v", output.Subtype1.B)
		}

		if output.Subtype2 == nil {
			t.Errorf("output.Subtype2; is nil, should be populated with an instance")
		} else {
			if output.Subtype2.A != 3 {
				t.Errorf("output.Subtype2.A; expected: 3, got: %v", output.Subtype2.A)
			}

			if output.Subtype2.B != 4 {
				t.Errorf("output.Subtype2.B; expected: 4, got: %v", output.Subtype2.B)
			}
		}

		if output.TimeTest != time.Duration(15)*time.Second {
			t.Errorf("output.TimeTest; expected: 15s, got: %s", output.TimeTest)
		}

		if output.IntTest != int32(5) {
			t.Errorf("output.IntTest; expected: 5(int32), got: %d(%T)", output.IntTest, output.IntTest)
		}

		if output.Properties == nil {
			t.Errorf("output.Properties; is nil, should be map[string]interface{}")
		} else {
			if v, ok := output.Properties[`first`]; !ok || v != 1 {
				t.Errorf("output.Properties['first']; expected: 1, got: %v(%T)", v, v)
			}

			if v, ok := output.Properties[`second`]; !ok || v != true {
				t.Errorf("output.Properties['second']; expected: true, got: %v(%T)", v, v)
			}

			if v, ok := output.Properties[`third`]; !ok || v != `three` {
				t.Errorf("output.Properties['third']; expected: 'three', got: %v(%T)", v, v)
			}
		}

		if output.StrSliceTest == nil {
			t.Errorf("output.StrSliceTest; is nil, should be []string")
		} else {
			if l := len(output.StrSliceTest); l == 3 {
				if v := output.StrSliceTest[0]; v != `one` {
					t.Errorf("output.StrSliceTest[0]; expected: 'one', got: %s", v)
				}

				if v := output.StrSliceTest[1]; v != `two` {
					t.Errorf("output.StrSliceTest[1]; expected: 'two', got: %s", v)
				}

				if v := output.StrSliceTest[2]; v != `three` {
					t.Errorf("output.StrSliceTest[2]; expected: 'three', got: %s", v)
				}
			} else {
				t.Errorf("output.StrSliceTest; wrong length - expected: 3, got: %d", l)
			}
		}

		if output.InterfaceStrSliceTest == nil {
			t.Errorf("output.InterfaceStrSliceTest; is nil, should be []string")
		} else {
			if l := len(output.InterfaceStrSliceTest); l == 3 {
				if v := output.InterfaceStrSliceTest[0]; v != `one` {
					t.Errorf("output.InterfaceStrSliceTest[0]; expected: 'one', got: %s", v)
				}

				if v := output.InterfaceStrSliceTest[1]; v != `two` {
					t.Errorf("output.InterfaceStrSliceTest[1]; expected: 'two', got: %s", v)
				}

				if v := output.InterfaceStrSliceTest[2]; v != `three` {
					t.Errorf("output.InterfaceStrSliceTest[2]; expected: 'three', got: %s", v)
				}
			} else {
				t.Errorf("output.InterfaceStrSliceTest; wrong length - expected: 3, got: %d", l)
			}
		}
	} else {
		t.Error(err)
	}
}
