// Copyright 2020 Tetrate
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"fmt"

	cmdinit "github.com/tetratelabs/getenvoy/pkg/cmd/extension/init"
)

// CategoryLanguageTuple represents a combination of extension category and
// programming language.
type CategoryLanguageTuple struct {
	Category, Language string
}

func (t CategoryLanguageTuple) String() string {
	return fmt.Sprintf("category=%s, language=%s", t.Category, t.Language)
}

// GetCategoryLanguageCombinations returns all combinations of a supported
// extension category with a supported programming language.
func GetCategoryLanguageCombinations() []CategoryLanguageTuple {
	tuples := make([]CategoryLanguageTuple, 0)
	for _, lang := range cmdinit.SupportedLanguages {
		for _, category := range cmdinit.SupportedCategories[lang.Value] {
			tuples = append(tuples, CategoryLanguageTuple{category.Value, lang.Value})
		}
	}
	return tuples
}
