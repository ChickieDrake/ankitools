package convert_test

import (
	"github.com/ChickieDrake/ankitools/convert"
	"reflect"
	"testing"
)

func TestToRequestMessage(t *testing.T) {
	type args struct {
		action convert.Action
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Happy path", args{action: "fake"}, `{"action":"fake","version":6}`, false},
		{"Contains unescaped character", args{action: `"`}, `{"action":"\"","version":6}`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			converter := convert.New()
			got, err := converter.ToRequestMessage(tt.args.action)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToRequestMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToRequestMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessageToDeckList(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name      string
		args      args
		wantDecks []string
		wantErr   bool
	}{
		{
			name:      "Happy Path 1",
			args:      args{message: `{"result": ["Hello"],"error": null}`},
			wantDecks: []string{"Hello"},
			wantErr:   false,
		},
		{
			name:      "Happy Path 2",
			args:      args{message: `{"result": ["Hello", "Hi"],"error": null}`},
			wantDecks: []string{"Hello", "Hi"},
			wantErr:   false,
		},
		{
			name:      "Error from JSON",
			args:      args{message: `{"result": null,"error": "unsupported action"}`},
			wantDecks: nil,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			converter := convert.New()
			gotDecks, err := converter.ToDeckList(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToDeckList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDecks, tt.wantDecks) {
				t.Errorf("ToDeckList() gotDecks = %v, want %v", gotDecks, tt.wantDecks)
			}
		})
	}
}
