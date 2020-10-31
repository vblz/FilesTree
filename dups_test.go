package main

import (
	"github.com/github.com/vblz/FilesTree/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test_findDups_Empty(t *testing.T) {
	res := findDuplicates(nil)
	assert.Empty(t, res)
}

func Test_findDups_NoDups(t *testing.T) {
	var testData = map[string]os.FileInfo{
		"a":           store.FileInfo{NameField: "a field name", SizeField: 184, ModeField: os.ModeDir},
		"/full/path":  store.FileInfo{NameField: "test name", SizeField: 22000},
		"/full/path2": store.FileInfo{NameField: "test name 2 ", SizeField: 11},
	}

	res := findDuplicates(testData)
	assert.Empty(t, res)
}

func Test_findDups(t *testing.T) {
	var firstArray = []string{
		"/first_pair",
		"/first_pair/path2",
	}
	var secondArray = []string{
		"/sec_pair/path",
		"secondpair",
		"secondpair/2",
	}

	var testData = map[string]os.FileInfo{
		"a": store.FileInfo{SizeField: 184, ModeField: os.ModeDir},
		"b": store.FileInfo{SizeField: 184, ModeField: os.ModeDir},
		"c": store.FileInfo{SizeField: 15, ModeField: os.ModeDir},

		"/zero/path":  store.FileInfo{SizeField: 0},
		"/zero/path2": store.FileInfo{SizeField: 0},

		firstArray[0]: store.FileInfo{SizeField: 11},
		firstArray[1]: store.FileInfo{SizeField: 11},

		secondArray[0]: store.FileInfo{SizeField: 5155874},
		secondArray[1]: store.FileInfo{SizeField: 5155874},
		secondArray[2]: store.FileInfo{SizeField: 5155874},
	}

	res := findDuplicates(testData)
	require.Len(t, res, 2)

	firstResArray := res[0]
	secondResArray := res[1]
	if len(firstResArray) != 2 {
		firstResArray = res[1]
		secondResArray = res[0]
	}

	assert.EqualValues(t, firstArray, firstResArray)
	assert.EqualValues(t, secondArray, secondResArray)
}
