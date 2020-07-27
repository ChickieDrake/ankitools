package ankitools

import "testing"

func TestFindCards(t *testing.T) {
	type args struct {
		japanese string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"default", args{japanese: "Sam"}, "this is a response"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindCards(tt.args.japanese); got != tt.want {
				t.Errorf("FindCards() = %v, want %v", got, tt.want)
			}
		})
	}
}
