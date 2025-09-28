package id1143

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name     string
		text1    string
		text2    string
		expected int
	}{
		{
			name:     "basic case",
			text1:    "abcde",
			text2:    "ace",
			expected: 3, // "ace"
		},
		{
			name:     "another basic case",
			text1:    "abc",
			text2:    "abc",
			expected: 3, // "abc"
		},
		{
			name:     "no common subsequence",
			text1:    "abc",
			text2:    "def",
			expected: 0,
		},
		{
			name:     "empty first string",
			text1:    "",
			text2:    "abc",
			expected: 0,
		},
		{
			name:     "empty second string",
			text1:    "abc",
			text2:    "",
			expected: 0,
		},
		{
			name:     "both strings empty",
			text1:    "",
			text2:    "",
			expected: 0,
		},
		{
			name:     "single character match",
			text1:    "a",
			text2:    "a",
			expected: 1,
		},
		{
			name:     "single character no match",
			text1:    "a",
			text2:    "b",
			expected: 0,
		},
		{
			name:     "longer strings with multiple matches",
			text1:    "abcdefghij",
			text2:    "acegi",
			expected: 5, // "acegi"
		},
		{
			name:     "leetcode example",
			text1:    "abcba",
			text2:    "abcbcba",
			expected: 5, // "abcba"
		},
		{
			name:     "characters in different order",
			text1:    "abcd",
			text2:    "dcba",
			expected: 1, // "a", "b", "c", or "d"
		},
		{
			name:     "repeated characters",
			text1:    "aaa",
			text2:    "aaa",
			expected: 3,
		},
		{
			name:     "partial match",
			text1:    "abcdef",
			text2:    "acdf",
			expected: 4, // "acdf"
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestCommonSubsequence(tt.text1, tt.text2)
			if result != tt.expected {
				t.Errorf("longestCommonSubsequence(%q, %q) = %d, expected %d",
					tt.text1, tt.text2, result, tt.expected)
			}
		})
	}
}

func BenchmarkLongestCommonSubsequence(b *testing.B) {
	text1 := "abcdefghijklmnopqrstuvwxyz"
	text2 := "acegikmoqsuwy"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestCommonSubsequence(text1, text2)
	}
}
