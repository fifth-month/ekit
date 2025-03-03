// Copyright 2021 ecodeclub
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package slice

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymmetricDiffSet(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		dst  []int
		want []int
	}{
		{
			src:  []int{1, 2, 4, 3},
			dst:  []int{4, 5, 6, 1},
			want: []int{2, 3, 5, 6},
			name: "normal test",
		},
		{
			src:  []int{1, 1, 2, 3, 4},
			dst:  []int{4, 5, 6, 1, 7, 6},
			want: []int{3, 6, 7, 5, 2},
			name: "deduplicate",
		},
		{
			src:  []int{},
			dst:  []int{1},
			want: []int{1},
			name: "src length is 0",
		},
		{
			src:  []int{1, 3, 5},
			dst:  []int{2, 4},
			want: []int{1, 3, 2, 4, 5},
			name: "not exist same ele",
		},
		{
			name: "both nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := SymmetricDiffSet[int](tt.src, tt.dst)
			assert.ElementsMatch(t, tt.want, res)
		})
	}
}

func TestSymmetricDiffSetFunc(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		dst  []int
		want []int
	}{
		{
			src:  []int{1, 2, 3, 4},
			dst:  []int{4, 5, 6, 1},
			want: []int{2, 3, 5, 6},
			name: "normal test",
		},
		{
			src:  []int{1, 1, 2, 3, 4},
			dst:  []int{4, 5, 6, 1, 7, 6},
			want: []int{3, 6, 7, 5, 2},
			name: "deduplicate",
		},
		{
			src:  []int{},
			dst:  []int{1},
			want: []int{1},
			name: "src length is 0",
		},
		{
			src:  []int{1, 3, 5},
			dst:  []int{2, 4},
			want: []int{1, 3, 2, 4, 5},
			name: "not exist same ele",
		},
		{
			name: "both nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := SymmetricDiffSetFunc[int](tt.src, tt.dst, func(src, dst int) bool {
				return src == dst
			})
			assert.ElementsMatch(t, tt.want, res)
		})
	}
}

func ExampleSymmetricDiffSet() {
	res := SymmetricDiffSet[int]([]int{1, 3, 4, 2}, []int{2, 5, 7, 3})
	sort.Ints(res)
	fmt.Println(res)
	// Output:
	// [1 4 5 7]
}

func ExampleSymmetricDiffSetFunc() {
	res := SymmetricDiffSetFunc[int]([]int{1, 3, 4, 2}, []int{2, 5, 7, 3}, func(src, dst int) bool {
		return src == dst
	})
	sort.Ints(res)
	fmt.Println(res)
	// Output:
	// [1 4 5 7]
}
