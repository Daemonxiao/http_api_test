package compare

import (
	"testing"
)

func TestContianMap(t *testing.T) {
	tests := []struct {
		a    map[string]any
		b    map[string]any
		want bool
	}{
		{
			a: map[string]any{
				"a": "1",
				"b": "2",
			},
			b: map[string]any{
				"a": "1",
			},
			want: true,
		},
		{
			a: map[string]any{
				"a": "2",
				"b": "2",
			},
			b: map[string]any{
				"a": "1",
			},
			want: false,
		},
		{
			a: map[string]any{
				"a": map[string]any{
					"b": map[string]any{
						"c": "1",
						"d": "2",
					},
				},
				"e": []interface{}{
					map[string]any{
						"f": "1",
						"g": "2",
					},
					map[string]any{
						"h": "1",
						"i": "2",
					},
				},
			},
			b: map[string]any{
				"a": map[string]any{
					"b": map[string]any{
						"d": "2",
					},
				},
			},
			want: true,
		},
		{
			a: map[string]any{
				"a": map[string]any{
					"b": map[string]any{
						"c": "1",
						"d": "2",
					},
				},
				"e": []interface{}{
					map[string]any{
						"f": "1",
						"g": "2",
					},
					map[string]any{
						"h": "1",
						"i": "2",
					},
				},
			},
			b: map[string]any{
				"e": []interface{}{
					map[string]any{
						"h": "1",
						"i": "2",
					},
				},
			},
			want: true,
		},
		{
			a: map[string]any{
				"a": map[string]any{
					"b": map[string]any{
						"c": "1",
						"d": "2",
					},
				},
				"e": []interface{}{
					map[string]any{
						"f": "1",
						"g": "2",
					},
					map[string]any{
						"h": "1",
						"i": "2",
					},
				},
			},
			b: map[string]any{
				"e": []interface{}{
					map[string]any{
						"h": "3",
						"i": "2",
					},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		if got := ContainsMap(tt.a, tt.b); got != tt.want {
			t.Errorf("result %v, want %v", got, tt.want)
		}
	}
}
