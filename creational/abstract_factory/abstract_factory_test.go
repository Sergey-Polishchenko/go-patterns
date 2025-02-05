package abstractfactory

import (
	"reflect"
	"testing"
)

func TestGUIFactory(t *testing.T) {
	tests := []struct {
		name       string
		content    string
		wantHeader Header
		wantBody   Body
		wantErr    bool
	}{
		{
			name:       "Valid content for GUIHeader",
			content:    "GUI Header",
			wantHeader: &GUIHeader{content: "GUI Header"},
			wantBody:   nil,
			wantErr:    false,
		},
		{
			name:       "Valid content for GUIBody",
			content:    "GUI Body",
			wantHeader: nil,
			wantBody:   &GUIBody{content: "GUI Body"},
			wantErr:    false,
		},
		{
			name:       "Empty content for GUIHeader",
			content:    "",
			wantHeader: nil,
			wantBody:   nil,
			wantErr:    true,
		},
		{
			name:       "Empty content for GUIBody",
			content:    "",
			wantHeader: nil,
			wantBody:   nil,
			wantErr:    true,
		},
	}

	factory := NewGUIFactory()

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				if tt.wantHeader != nil {
					header, err := factory.CreateHeader(tt.content)
					if (err != nil) != tt.wantErr {
						t.Errorf(
							"GUIFactory.CreateHeader() error = %v, wantErr %v",
							err,
							tt.wantErr,
						)
						return
					}
					if !tt.wantErr && !reflect.DeepEqual(header, tt.wantHeader) {
						t.Errorf("GUIFactory.CreateHeader() = %v, want %v", header, tt.wantHeader)
					}
				}

				if tt.wantBody != nil {
					body, err := factory.CreateBody(tt.content)
					if (err != nil) != tt.wantErr {
						t.Errorf("GUIFactory.CreateBody() error = %v, wantErr %v", err, tt.wantErr)
						return
					}
					if !tt.wantErr && !reflect.DeepEqual(body, tt.wantBody) {
						t.Errorf("GUIFactory.CreateBody() = %v, want %v", body, tt.wantBody)
					}
				}
			},
		)
	}
}

func TestTUIFactory(t *testing.T) {
	tests := []struct {
		name       string
		content    string
		wantHeader Header
		wantBody   Body
		wantErr    bool
	}{
		{
			name:       "Valid content for TUIHeader",
			content:    "TUI Header",
			wantHeader: &TUIHeader{content: "TUI Header"},
			wantBody:   nil,
			wantErr:    false,
		},
		{
			name:       "Valid content for TUIBody",
			content:    "TUI Body",
			wantHeader: nil,
			wantBody:   &TUIBody{content: "TUI Body"},
			wantErr:    false,
		},
		{
			name:       "Empty content for TUIHeader",
			content:    "",
			wantHeader: nil,
			wantBody:   nil,
			wantErr:    true,
		},
		{
			name:       "Empty content for TUIBody",
			content:    "",
			wantHeader: nil,
			wantBody:   nil,
			wantErr:    true,
		},
	}

	factory := NewTUIFactory()

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				if tt.wantHeader != nil {
					header, err := factory.CreateHeader(tt.content)
					if (err != nil) != tt.wantErr {
						t.Errorf(
							"TUIFactory.CreateHeader() error = %v, wantErr %v",
							err,
							tt.wantErr,
						)
						return
					}
					if !tt.wantErr && !reflect.DeepEqual(header, tt.wantHeader) {
						t.Errorf("TUIFactory.CreateHeader() = %v, want %v", header, tt.wantHeader)
					}
				}

				if tt.wantBody != nil {
					body, err := factory.CreateBody(tt.content)
					if (err != nil) != tt.wantErr {
						t.Errorf("TUIFactory.CreateBody() error = %v, wantErr %v", err, tt.wantErr)
						return
					}
					if !tt.wantErr && !reflect.DeepEqual(body, tt.wantBody) {
						t.Errorf("TUIFactory.CreateBody() = %v, want %v", body, tt.wantBody)
					}
				}
			},
		)
	}
}
