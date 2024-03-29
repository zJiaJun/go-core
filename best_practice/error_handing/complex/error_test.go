package complex

import "testing"

func TestHandlingErrorWitchCustomError(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "find by name", args: args{"name"}},
		{name: "find by age", args: args{"age"}},
		{name: "find by sex", args: args{"sex"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HandlingErrorWitchCustomError(tt.args.key)
		})
	}
}
