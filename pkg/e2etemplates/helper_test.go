package e2etemplates

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func diff(a, b string) (int, string) {
	aSplit := strings.Split(a, "\n")
	bSplit := strings.Split(b, "\n")

	length := len(aSplit)
	if len(bSplit) < length {
		length = len(bSplit)
	}

	for i := 0; i < length; i++ {
		if aSplit[i] != bSplit[i] {
			return i + 1, fmt.Sprintf("a: %q b: %q", aSplit[i], bSplit[i])
		}
	}

	if len(aSplit) > len(bSplit) {
		return len(bSplit) + 1, fmt.Sprintf("a: %q b: EOF", aSplit[len(bSplit)])
	}
	if len(aSplit) < len(bSplit) {
		return len(aSplit) + 1, fmt.Sprintf("a: EOF b: %q", bSplit[len(aSplit)])
	}

	return 0, ""
}

func Test_diff(t *testing.T) {
	testCases := []struct {
		name               string
		a                  string
		b                  string
		expectedLine       int
		expectedDifference string
		errorMatcher       func(err error) bool
	}{
		{
			name: "case 1",
			a: `x: 1
			    y: 2
			`,
			b: `x: 1
			    y: 2
			`,
			expectedLine:       0,
			expectedDifference: "",
		},
		{
			name: "case 2",
			a: `x: 1
			    y: 2`,
			b: `x: 4
			    y: 2`,
			expectedLine:       1,
			expectedDifference: `a: "x: 1" b: "x: 4"`,
		},
		{
			name: "case 3",
			a: `x: 1
			    y: 2`,
			b: `x: 1
			    y: 5`,
			expectedLine:       2,
			expectedDifference: `a: "\t\t\t    y: 2" b: "\t\t\t    y: 5"`,
		},
		{
			name: "case 4",
			a: `x: 1
			    y: 2`,
			b: `x: 1
			    y: 2
			    z`,
			expectedLine:       3,
			expectedDifference: `a: EOF b: "\t\t\t    z"`,
		},
		{
			name: "case 5",
			a: `x: 1
			    y: 2
			    z`,
			b: `x: 1
			    y: 2`,
			expectedLine:       3,
			expectedDifference: `a: "\t\t\t    z" b: EOF`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			line, difference := diff(tc.a, tc.b)

			if !reflect.DeepEqual(line, tc.expectedLine) {
				t.Fatalf("line == %d, want %d", line, tc.expectedLine)
			}

			if !reflect.DeepEqual(difference, tc.expectedDifference) {
				t.Fatalf("difference == %s, want %s", difference, tc.expectedDifference)
			}
		})
	}
}
