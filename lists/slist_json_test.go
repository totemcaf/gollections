package lists

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceList_strings_to_json(t *testing.T) {
	a := Of("an element", "other element")

	str, err := json.Marshal(a)

	assert.Nil(t, err)
	assert.Equal(t, `["an element","other element"]`, string(str))
}

func TestSliceList_ints_to_json(t *testing.T) {
	a := Of(12, 45, 62)

	str, err := json.Marshal(a)

	assert.Nil(t, err)
	assert.Equal(t, `[12,45,62]`, string(str))
}

type sample struct {
	Name string
	Age  int
}

func TestSliceList_structs_to_json(t *testing.T) {
	a := Of(&sample{"Elton", 42}, &sample{"Alison", 25})

	str, err := json.Marshal(a)

	assert.Nil(t, err)
	assert.Equal(t, `[{"Name":"Elton","Age":42},{"Name":"Alison","Age":25}]`, string(str))
}

func TestSliceList_empty_to_json(t *testing.T) {
	a := Of[string]()

	str, err := json.Marshal(a)

	assert.Nil(t, err)

	assert.Equal(t, `[]`, string(str))
}

func TestSliceList_json_to_strings(t *testing.T) {
	jsonStr := `["an element","other element"]`

	var list = Of[string]("")

	err := json.Unmarshal([]byte(jsonStr), &list)

	assert.Nil(t, err)
	assert.Equal(t, Of("an element", "other element"), list)
}

func (s *sliceList[T]) MarshalJSON() ([]byte, error) {
	if s.es == nil {
		return []byte("[]"), nil
	}
	return json.Marshal(s.es)
}

func (s *sliceList[T]) UnmarshalJSON(data []byte) error {
	var elements []T

	if err := json.Unmarshal(data, &elements); err != nil {
		return err
	}

	s.es = elements

	return nil
}
