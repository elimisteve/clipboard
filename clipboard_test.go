// Copyright 2013 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clipboard_test

import (
	"testing"

	. "."
)

func TestCopyAndPaste(t *testing.T) {
	expected := []byte("日本語")

	err := WriteAll(expected)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	if string(actual) != string(expected) {
		t.Errorf("want %s, got %s", expected, actual)
	}
}

func TestMultiCopyAndPaste(t *testing.T) {
	expected1 := []byte("French: éèêëàùœç")
	expected2 := []byte("Weird UTF-8: 💩☃")

	err := WriteAll(expected1)
	if err != nil {
		t.Fatal(err)
	}

	actual1, err := ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if string(actual1) != string(expected1) {
		t.Errorf("want %s, got %s", expected1, actual1)
	}

	err = WriteAll(expected2)
	if err != nil {
		t.Fatal(err)
	}

	actual2, err := ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if string(actual2) != string(expected2) {
		t.Errorf("want %s, got %s", expected2, actual2)
	}
}

func BenchmarkReadAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadAll()
	}
}

func BenchmarkWriteAll(b *testing.B) {
	text := []byte("いろはにほへと")
	for i := 0; i < b.N; i++ {
		WriteAll(text)
	}
}
