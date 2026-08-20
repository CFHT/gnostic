// Copyright 2020 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package generator

import (
	"strings"
	"unicode"

	"google.golang.org/protobuf/reflect/protoreflect"
)

// contains returns true if an array contains a specified string.
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// appendUnique appends a string, to a string slice, if the string is not already in the slice
func appendUnique(s []string, e string) []string {
	if !contains(s, e) {
		return append(s, e)
	}
	return s
}

// singular produces the singular form of a collection name.
func singular(plural string) string {
	if strings.HasSuffix(plural, "ves") {
		return strings.TrimSuffix(plural, "ves") + "f"
	}
	if strings.HasSuffix(plural, "ies") {
		return strings.TrimSuffix(plural, "ies") + "y"
	}
	if strings.HasSuffix(plural, "s") {
		return strings.TrimSuffix(plural, "s")
	}
	return plural
}

func getValueKind(message protoreflect.MessageDescriptor) string {
	valueField := getValueField(message)
	return valueField.Kind().String()
}

func getValueField(message protoreflect.MessageDescriptor) protoreflect.FieldDescriptor {
	fields := message.Fields()
	return fields.ByName("value")
}

// Copied in from generate-gnostic/generate-compiler.go
// Returns a "snake case" form of a camel-cased string.
func camelCaseToSnakeCase(input string) string {
	out := ""
	for index, runeValue := range input {
		//fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
		if runeValue >= 'A' && runeValue <= 'Z' {
			if index > 0 {
				out += "_"
			}
			out += string(runeValue - 'A' + 'a')
		} else {
			out += string(runeValue)
		}
	}
	return out
}

// Copied in from generate-gnostic/generate-compiler.go
func snakeCaseToCamelCase(input string) string {
	out := ""

	words := strings.Split(input, "_")

	for i, word := range words {
		if (i > 0) && len(word) > 0 {
			w := []rune(word)
			w[0] = unicode.ToUpper(w[0])
			out += string(w)
		} else {
			out += word
		}
	}

	return out
}
