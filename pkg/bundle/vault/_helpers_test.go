// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package vault

import (
	"regexp"
	"testing"

	fuzz "github.com/google/gofuzz"
)

func Test_matchPathRule_Fuzz(t *testing.T) {
	// Making sure the function never panics
	for i := 0; i < 50; i++ {
		f := fuzz.New()

		// Prepare arguments
		var (
			value   string
			regexps []*regexp.Regexp
		)

		f.Fuzz(&value)
		f.Fuzz(&regexps)

		// Execute
		matchPathRule(value, regexps)
	}
}

func Test_collect_Fuzz(t *testing.T) {
	// Making sure the function never panics
	for i := 0; i < 50; i++ {
		f := fuzz.New()

		// Prepare arguments
		var (
			values        []string
			regexps       []*regexp.Regexp
			appendIfMatch bool
		)

		f.Fuzz(&values)
		f.Fuzz(&regexps)
		f.Fuzz(&appendIfMatch)

		// Execute
		collect(values, regexps, appendIfMatch)
	}
}
