// Copyright 2013 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package clipboard read/write on clipboard
package clipboard

// ReadAll read data from clipboard
func ReadAll() ([]byte, error) {
	return readAll()
}

// WriteAll write string to clipboard
func WriteAll(data []byte) error {
	return writeAll(data)
}
