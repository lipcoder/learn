package prose

import "testing"

func TestJoinWithCommas2(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want string
	}{
		{
			name: "empty",
			in:   []string{},
			want: "",
		},
		{
			name: "one item",
			in:   []string{"apple"},
			want: "apple",
		},
		{
			name: "two items",
			in:   []string{"apple", "orange"},
			want: "apple and orange",
		},
		{
			name: "three items",
			in:   []string{"apple", "orange", "pear"},
			want: "apple, orange, and pear",
		},
		{
			name: "many items",
			in:   []string{"a", "b", "c", "d"},
			want: "a, b, c, and d",
		},
	}

	for _, tt := range tests {
		tt := tt // 避免闭包引用同一个变量,这个我还没有太理解
		t.Run(tt.name, func(t *testing.T) {
			got := JoinWithCommas(tt.in)
			if got != tt.want {
				t.Fatalf("JoinWithCommas(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

/*
for _, tt := range tests {
    tt := tt
    t.Run(tt.name, func(t *testing.T) {
        // 未来想并行时，直接打开就行
        // t.Parallel()

        got := Fn(tt.in)
        if got != tt.want { ... }
    })
}
*/
