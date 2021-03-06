// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package transpiler

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubpath(t *testing.T) {
	testCases := map[string][]struct {
		root       string
		path       string
		resultPath string
		isSubpath  bool
	}{
		"linux": {
			{"/", "a", "/a", true},
			{"/a", "b", "/a/b", true},
			{"/a", "b/c", "/a/b/c", true},

			{"/a/b", "/a/c", "/a/c", false},

			{"/a/b", "/a/b/../c", "/a/c", false},
			{"/a/b", "../c", "/a/c", false},
			{"/a", "/a/b/c", "/a/b/c", true},
			{"/a", "/A/b/c", "/A/b/c", false},
		},
		"darwin": {
			{"/", "a", "/a", true},
			{"/a", "b", "/a/b", true},
			{"/a", "b/c", "/a/b/c", true},
			{"/a/b", "/a/c", "/a/c", false},
			{"/a/b", "/a/b/../c", "/a/c", false},
			{"/a/b", "../c", "/a/c", false},
			{"/a", "/a/b/c", "/a/b/c", true},
			{"/a", "/A/b/c", "/a/b/c", true},
		},
		// (Windows issue) See issue:
		//"windows": {
		//	{"/", "a", "\\a", true},
		//	{"/a", "b", "\\a\\b", true},
		//	{"/a", "b/c", "\\a\\b\\c", true},
		//	//
		//	{"/a/b", "/a/c", "\\a\\c", false},
		//	{"/a/b", "/a/b/../c", "\\a\\c", false},
		//	{"/a/b", "../c", "\\a\\c", false},
		//	{"/a", "/a/b/c", "\\a\\b\\c", true},
		//	{"/a", "/A/b/c", "\\a\\b\\c", true},
		//},
	}

	osSpecificTests, found := testCases[runtime.GOOS]
	if !found {
		return
	}

	for _, test := range osSpecificTests {
		t.Run(fmt.Sprintf("[%s] root:'%s path: %s'", runtime.GOOS, test.root, test.path), func(t *testing.T) {
			newPath, result := joinPaths(test.root, test.path)
			assert.Equal(t, test.resultPath, newPath)
			assert.Equal(t, test.isSubpath, result)
		})
	}
}
