package main

import (
	"testing"
)

func TestGroupWords(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[string][]string
	}{
		{
			name:  "basic anagram groups",
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			name:  "mixed case handling",
			input: []string{"ПяТак", "пЯткА", "ТяпКа", "ЛистОк", "СЛИТОК", "стоЛИК"},
			expected: map[string][]string{
				"ПяТак":  {"ПяТак", "пЯткА", "ТяпКа"},
				"ЛистОк": {"ЛистОк", "СЛИТОК", "стоЛИК"},
			},
		},
		{
			name:     "ignore single words",
			input:    []string{"один", "два", "три"},
			expected: map[string][]string{},
		},
		{
			name:  "duplicates handling",
			input: []string{"пятак", "пятак", "пятка", "пятка", "тяпка"},
			expected: map[string][]string{
				"пятак": {"пятак", "пятка", "тяпка"},
			},
		},
		{
			name:  "different first word",
			input: []string{"пятка", "пятак", "тяпка", "столик", "листок", "слиток"},
			expected: map[string][]string{
				"пятка":  {"пятак", "пятка", "тяпка"},
				"столик": {"листок", "слиток", "столик"},
			},
		},
		{
			name:  "english words",
			input: []string{"listen", "silent", "enlist", "triangle", "integral", "altering", "alerting"},
			expected: map[string][]string{
				"listen":   {"enlist", "listen", "silent"},
				"triangle": {"alerting", "altering", "integral", "triangle"},
			},
		},
		{
			name:     "empty input",
			input:    []string{},
			expected: map[string][]string{},
		},
		{
			name:  "different length words",
			input: []string{"кот", "ток", "окт", "к", "отк", "то"},
			expected: map[string][]string{
				"кот": {"кот", "окт", "отк", "ток"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := groupWords(tt.input)

			if len(result) != len(tt.expected) {
				t.Errorf("expected %d groups, got %d", len(tt.expected), len(result))
			}

			for key, expectedGroup := range tt.expected {
				resultGroup, ok := result[key]
				if !ok {
					t.Errorf("missing expected key: %s", key)
					continue
				}

				if len(resultGroup) != len(expectedGroup) {
					t.Errorf("for key %s: expected %d words, got %d", key, len(expectedGroup), len(resultGroup))
					continue
				}

				for i, word := range expectedGroup {
					if resultGroup[i] != word {
						t.Errorf("for key %s: at index %d expected %s, got %s", key, i, word, resultGroup[i])
					}
				}
			}

			for key := range result {
				if _, ok := tt.expected[key]; !ok {
					t.Errorf("unexpected group with key: %s", key)
				}
			}
		})
	}
}
