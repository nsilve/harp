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

package cmdutil

import (
	"fmt"

	"github.com/spf13/afero"

	"github.com/elastic/harp/pkg/sdk/types"
	"github.com/elastic/harp/pkg/template/engine"
	"github.com/elastic/harp/pkg/template/files"
)

// Files returns template files.
func Files(fs afero.Fs, basePath string) (engine.Files, error) {
	// Check arguments
	if types.IsNil(fs) {
		return nil, fmt.Errorf("unable to load files without a default filesystem to use")
	}

	// Get appropriate loader
	loader, err := files.Loader(afero.NewReadOnlyFs(fs), basePath)
	if err != nil {
		return nil, fmt.Errorf("unable to process files: %v", err)
	}

	// Crawl and load file content
	fileList, err := loader.Load()
	if err != nil {
		return nil, fmt.Errorf("unable to load files: %v", err)
	}

	// Wrap as template files
	templateFiles := engine.NewFiles(fileList)

	// No error
	return templateFiles, nil
}
