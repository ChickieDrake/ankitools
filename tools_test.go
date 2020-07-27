package ankitools

import "testing"

func TestDeckNamesAndIds(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeckNamesAndIds()
			if (err != nil) != tt.wantErr {
				t.Errorf("DeckNamesAndIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeckNamesAndIds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createActionString(t *testing.T) {
	type args struct {
		action string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createActionString(tt.args.action)
			if (err != nil) != tt.wantErr {
				t.Errorf("createActionString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("createActionString() = %v, want %v", got, tt.want)
			}
		})
	}
}
