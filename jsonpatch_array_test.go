package jsonpatch

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var arrayA = `{"containers": ["a", "b"]}`
var arrayB = `{"containers": ["c", "d", "e"]}`

var simpleListA = `{"x": ["a"]}`
var simpleListAB = `{"x": ["a", "b"]}`
var simpleListAC = `{"x": ["a", "c"]}`

func TestListAddAndRemove(t *testing.T) {
	patch, e := CreatePatch([]byte(simpleListA), []byte(simpleListAC))
	assert.NoError(t, e)
	assert.Equal(t, patch, []JsonPatchOperation{
		{Operation: "add", Path: "/x/1", Value: "c"},
	}, "they should be equal")

	patch, e = CreatePatch([]byte(simpleListAB), []byte(simpleListAC))
	assert.NoError(t, e)
	assert.Equal(t, patch, []JsonPatchOperation{
		{Operation: "replace", Path: "/x/1", Value: "c"},
	}, "they should be equal")

}

func TestPatchTwoCompletelyDifferentArrays(t *testing.T) {
	patch, err := CreatePatch([]byte(arrayA), []byte(arrayB))

	assert.NoError(t, err, "patch should not return an error")
	s, _ := json.MarshalIndent(patch, "", "\t")
	t.Log(string(s))

	assert.Equal(t, len(patch), 3, "there should be 3 operations")
	first := patch[0]
	second := patch[1]
	third := patch[2]

	assert.Equal(t, "replace", first.Operation, "first operation should be a replace")
	assert.Equal(t, "replace", second.Operation, "second operation should be a replace")
	assert.Equal(t, "add", third.Operation, "third operation should be an add")
	assert.Equal(t, "/containers/0", first.Path, "the first path should be index 0")
	assert.Equal(t, "/containers/1", second.Path, "the second path should be index 1")
	assert.Equal(t, "/containers/2", third.Path, "the third path should be index 2")
}
