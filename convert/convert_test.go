package convert

import (
	"errors"
	"testing"
)

func TestConvertToRequestMessage(t *testing.T) {
	type args struct {
		action string
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
			got, err := ConvertToRequestMessage(tt.args.action)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToRequestMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertToRequestMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToRequestMessage_err(t *testing.T) {
	// for testing (easier than mocking)
	marshal_function = marshal_function_returns_error

	got, err := ConvertToRequestMessage("anystring")

	// assert
	if got != "" {
		t.Errorf("convert: Expected converted value to be empty if json.Marshal returned an error, received: %s", got)
	}
	if err == nil {
		t.Error("convert: Expected an error if the json.Marshal returned an error")
	}
	expected_error_message := "Test message"
	if err.Error() != expected_error_message {
		t.Errorf("convert: Expected a different error message, received: %s", err.Error())
	}

}

func marshal_function_returns_error(input interface{}) (output []byte, err error) {
	err = errors.New("Test message")
	return
}
