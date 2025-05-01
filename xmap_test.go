package xmap_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/targc/xmap"
)

func TestXMapGet(t *testing.T) {

	type TC struct {
		m        xmap.XMap
		key      string
		expected interface{}
	}

	tcs := []TC{
		{
			m: xmap.XMap{
				"A": map[string]interface{}{
					"B": map[string]interface{}{
						"C": map[string]interface{}{
							"E": "test",
						},
					},
				},
			},
			key:      "A.B.C.E",
			expected: "test",
		},
		{
			m: xmap.XMap{
				"A": map[string]interface{}{
					"B": map[string]interface{}{
						"C": map[string]interface{}{
							"E": "test",
						},
					},
				},
			},
			key:      "A.B.C.E.D",
			expected: nil,
		},
		{
			m:        xmap.XMap{},
			key:      "A.B.C.E.D",
			expected: nil,
		},
		{
			m:        xmap.XMap{},
			key:      "A.B",
			expected: nil,
		},
		{
			m:        xmap.XMap{},
			key:      "A",
			expected: nil,
		},
		{
			m: xmap.XMap{
				"A": map[string]interface{}{
					"B": map[string]interface{}{
						"C": map[string]interface{}{
							"E": "test",
						},
					},
				},
			},
			key: "A",
			expected: map[string]interface{}{
				"B": map[string]interface{}{
					"C": map[string]interface{}{
						"E": "test",
					},
				},
			},
		},
		{
			m: xmap.XMap{
				"A": map[string]interface{}{
					"B": map[string]interface{}{
						"C": map[string]interface{}{
							"E": "test",
						},
					},
				},
			},
			key: "A.B",
			expected: map[string]interface{}{
				"C": map[string]interface{}{
					"E": "test",
				},
			},
		},
		{
			m: xmap.XMap{
				"A": map[string]interface{}{
					"B": map[string]interface{}{
						"C": map[string]interface{}{
							"E": "test",
						},
					},
				},
			},
			key: "A.B",
			expected: map[string]interface{}{
				"C": map[string]interface{}{
					"E": "test",
				},
			},
		},
		{
			m: xmap.XMap{
				"A": map[string]interface{}{
					"B": 123,
				},
			},
			key:      "A.B",
			expected: 123,
		},
	}

	t.Parallel()

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			actual := tc.m.Get(tc.key)

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestXMapGetS(t *testing.T) {

	type TC struct {
		m        xmap.XMap
		key      string
		expected string
	}

	tcs := []TC{
		{
			m: xmap.XMap{
				"A": map[string]interface{}{
					"B": map[string]interface{}{
						"C": map[string]interface{}{
							"E": "test",
						},
					},
				},
			},
			key:      "A.B.C.E",
			expected: "test",
		},
		{
			m: xmap.XMap{
				"A": map[string]interface{}{
					"B": map[string]interface{}{
						"C": map[string]interface{}{
							"E": "test",
						},
					},
				},
			},
			key:      "A.B.C.E.D",
			expected: "",
		},
		{
			m: xmap.XMap{
				"A": map[string]interface{}{
					"B": map[string]interface{}{
						"C": map[string]interface{}{
							"E": "test",
						},
					},
				},
			},
			key:      "A.B",
			expected: "",
		},
	}

	t.Parallel()

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			actual := tc.m.GetS(tc.key)

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
