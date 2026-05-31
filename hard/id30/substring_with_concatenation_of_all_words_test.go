package id30

import (
	"slices"
	"sort"
	"testing"
)

func Test_findSubstring(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s     string
		words []string
		want  []int
	}{
		{
			"example 1",
			"barfoothefoobarman",
			[]string{"foo", "bar"},
			[]int{0, 9},
		},
		{
			"example 2",
			"wordgoodgoodgoodbestword",
			[]string{"word", "good", "best", "word"},
			[]int{},
		},
		{
			"example 3",
			"barfoofoobarthefoobarman",
			[]string{"bar", "foo", "the"},
			[]int{6, 9, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findSubstring(tt.s, tt.words)
			if !EqualIntSliceUnordered(got, tt.want) {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func EqualIntSliceUnordered(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	sa := slices.Clone(a)
	sb := slices.Clone(b)

	sort.Ints(sa)
	sort.Ints(sb)

	return slices.Equal(sa, sb)
}
