package jsonpatch

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var arrayA = `{"containers": ["a", "b"]}`
var arrayB = `{"containers": ["c", "d", "e"]}`

func TestPatchTwoCompletelyDifferentArrays(t *testing.T) {
	patch, err := CreatePatch([]byte(arrayA), []byte(arrayB))

	assert.NoError(t, err, "patch should not return an error")
	s, _ := json.MarshalIndent(patch, "", "\t")
	t.Log(string(s))

	assert.Equal(t, len(patch), 5, "there should be 5 operations")
	first := patch[0]
	second := patch[1]

	assert.Equal(t, "remove", first.Operation, "first operation should be a remove")
	assert.Equal(t, "remove", second.Operation, "second operation should be a remove")
	assert.Equal(t, "/containers/1", first.Path, "the first path should be index 1")
	assert.Equal(t, "/containers/0", second.Path, "the second path should be index 0")
}
