package convert_test

import (
	"github.com/ChickieDrake/ankitools/convert"
	"github.com/ChickieDrake/ankitools/types"
	"reflect"
	"testing"
)

func TestConverter_ToRequestMessage(t *testing.T) {

	var noteUpdate = &types.NoteUpdate{
		NoteID: 1483959289817,
		Fields: &map[string]string{
			"Front": "new front content",
			"Back":  "new back content",
		},
	}

	type args struct {
		action convert.Action
		params *convert.Params
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Happy path", args{action: "fake"}, `{"action":"fake","version":6}`, false},
		{"Contains unescaped character", args{action: `"`}, `{"action":"\"","version":6}`, false},
		{"Contains params", args{action: `fake`, params: &convert.Params{Query: "deck:current"}}, `{"action":"fake","version":6,"params":{"query":"deck:current"}}`, false},
		{
			"All params",
			args{action: `fake`, params: &convert.Params{Query: "deck:current", Notes: []int{1483959289817, 1483959291695}, Tag: "testing", Note: noteUpdate}},
			`{"action":"fake","version":6,"params":{"query":"deck:current","notes":[1483959289817,1483959291695],"tags":"testing","note":{"id":1483959289817,"fields":{"Back":"new back content","Front":"new front content"}}}}`,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			converter := convert.New()
			got, err := converter.ToRequestMessage(tt.args.action, tt.args.params)
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

func TestConverter_ToDeckNameList(t *testing.T) {
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
			gotDecks, err := converter.ToDeckNameList(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToDeckNameList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDecks, tt.wantDecks) {
				t.Errorf("ToDeckNameList() gotDecks = %v, want %v", gotDecks, tt.wantDecks)
			}
		})
	}
}

func TestConverter_ToNoteIDList(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name:    "Happy Path 1",
			args:    args{message: `{"result": [1483959289817, 1483959291695],"error": null}`},
			want:    []int{1483959289817, 1483959291695},
			wantErr: false,
		},
		{
			name:    "Happy Path 2",
			args:    args{message: `{"result": [148395928234, 14839345695],"error": null}`},
			want:    []int{148395928234, 14839345695},
			wantErr: false,
		},
		{
			name:    "Empty List",
			args:    args{message: `{"result": [],"error": null}`},
			want:    []int{},
			wantErr: false,
		},
		{
			name:    "Error from JSON",
			args:    args{message: `{"result": null,"error": "unsupported action"}`},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := convert.New()
			got, err := c.ToNoteIDList(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToNoteIDList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToNoteIDList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConverter_ToNoteList(t *testing.T) {
	expectedFieldsMap := &map[string]types.Field{
		"Front": {Value: "front content", Order: 0},
		"Back":  {Value: "back content", Order: 1},
	}

	type args struct {
		message string
	}
	tests := []struct {
		name    string
		args    args
		want    []*types.NoteInfo
		wantErr bool
	}{
		{
			name: "Happy Path 1",
			args: args{message: `{
					"result":[        
						{
            				"noteId":1502298033753,
            				"modelName": "Basic",
            				"tags":["tag","another_tag"],
							"fields": {
								"Front": {"value": "front content", "order": 0},
								"Back": {"value": "back content", "order": 1}
							},
							"cards": [ 1234566, 3456778 ]
        				}
					],
					"error": null
				}`},
			want: []*types.NoteInfo{{
				NoteID:    1502298033753,
				ModelName: "Basic",
				Tags: &[]string{
					"tag",
					"another_tag",
				},
				Fields:  expectedFieldsMap,
				CardIDs: &[]int{1234566, 3456778},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := convert.New()
			got, err := c.ToNoteList(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToNoteList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			//assert.True(reflect.DeepEqual(got.Fields, tt.want.Fields))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToNoteList() got = %v, want %v", got, tt.want)
			}
		})
	}
}
