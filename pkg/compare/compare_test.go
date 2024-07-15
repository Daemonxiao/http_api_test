package compare

import (
	"testing"
)

func TestContainsMap(t *testing.T) {
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
				"a": map[string]any{
					"b": map[string]any{
						"d": []interface{}{
							map[string]any{
								"h": "1",
								"i": "2",
							},
							map[string]any{
								"j": "1",
								"k": "2",
							},
						},
					},
					"c": map[string]any{
						"l": map[string]any{
							"j": "1",
							"k": "2",
						},
					},
				},
			},
			b: map[string]any{
				"a": map[string]any{
					"b": map[string]any{
						"d": []interface{}{
							map[string]any{
								"h": "1",
							},
						},
					},
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		if got := ContainsMap(tt.a, tt.b); got != tt.want {
			t.Errorf("result %v, want %v", got, tt.want)
		}
	}
}
