package maputil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiffuseOneTierScalar(t *testing.T) {
	assert := require.New(t)

	var err error

	input := make(map[string]interface{})
	output := make(map[string]interface{})

	input["id"] = "test"
	input["enabled"] = true
	input["float"] = 2.7

	output, err = DiffuseMap(input, ".")
	assert.NoError(err)

	v, ok := output["id"]
	assert.True(ok)
	assert.Equal("test", v)

	v, ok = output["enabled"]
	assert.True(ok)
	assert.Equal(true, v)

	v, ok = output["float"]
	assert.True(ok)
	assert.Equal(2.7, v)
}

func TestDiffuseOneTierComplex(t *testing.T) {
	assert := require.New(t)

	var err error

	input := make(map[string]interface{})
	output := make(map[string]interface{})

	input["array"] = []string{"first", "third", "fifth"}
	input["numary"] = []int{9, 7, 3}
	input["things"] = map[string]int{"one": 1, "two": 2, "three": 3}

	output, err = DiffuseMap(input, ".")
	assert.NoError(err)

	//  test string array
	_, ok := output["array"]
	assert.True(ok)

	for i, v := range output["array"].([]string) {
		assert.Equal(v, input["array"].([]string)[i])
	}

	//  test int array
	_, ok = output["numary"]
	assert.True(ok)

	for i, v := range output["numary"].([]int) {
		assert.Equal(v, input["numary"].([]int)[i])
	}

	//  test string-int map
	_, ok = output["things"]
	assert.True(ok)

	for k, v := range output["things"].(map[string]int) {
		inputValue, ok := input["things"].(map[string]int)[k]
		assert.True(ok)
		assert.Equal(v, inputValue)
	}
}

func TestDiffuseMultiTierScalar(t *testing.T) {
	assert := require.New(t)

	var err error

	input := make(map[string]interface{})
	output := make(map[string]interface{})

	input["items.0"] = 54
	input["items.1"] = 77
	input["items.2"] = 82

	output, err = DiffuseMap(input, ".")
	assert.NoError(err)

	i_items, ok := output["items"]
	assert.True(ok)

	items := i_items.([]interface{})

	for i, v := range []int{54, 77, 82} {
		assert.True(i < len(items))
		assert.Equal(v, items[i].(int))
	}
}

func TestDiffuseMultiTierComplex(t *testing.T) {
	assert := require.New(t)

	var err error

	input := make(map[string]interface{})
	output := make(map[string]interface{})

	input["items.0.name"] = "First"
	input["items.0.age"] = 54
	input["items.1.name"] = "Second"
	input["items.1.age"] = 77
	input["items.2.name"] = "Third"
	input["items.2.age"] = 82

	output, err = DiffuseMap(input, ".")
	assert.NoError(err)

	i_items, ok := output["items"]
	assert.True(ok)

	items := i_items.([]interface{})
	assert.Len(items, 3)

	for item_id, obj := range items {
		for k, v := range obj.(map[string]interface{}) {
			inValue, ok := input[fmt.Sprintf("items.%d.%s", item_id, k)]
			assert.True(ok)
			assert.Equal(v, inValue)
		}
	}
}

func TestDiffuseMultiTierMixed(t *testing.T) {
	assert := require.New(t)

	var err error

	input := make(map[string]interface{})
	output := make(map[string]interface{})

	input["items.0.tags"] = []string{"base", "other"}
	input["items.1.tags"] = []string{"thing", "still-other", "more-other"}
	input["items.2.tags"] = []string{"last"}

	output, err = DiffuseMap(input, ".")
	assert.NoError(err)

	i_items, ok := output["items"]
	assert.True(ok)

	items := i_items.([]interface{})

	assert.Len(items, 3)

	for item_id, obj := range items {
		for k, v := range obj.(map[string]interface{}) {
			vAry := v.([]string)

			inValue, ok := input[fmt.Sprintf("items.%d.%s", item_id, k)]
			assert.True(ok)

			inValueAry := inValue.([]string)

			for i, vAryV := range vAry {
				assert.Equal(inValueAry[i], vAryV)
			}
		}
	}
}
