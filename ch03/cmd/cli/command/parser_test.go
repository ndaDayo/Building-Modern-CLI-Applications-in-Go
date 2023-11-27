package command

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/ndaDayo/Audio-Metadata-CLI/internal/interfaces"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	if strings.Contains(req.URL.String(), "/upload") {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader("")),
		}, nil
	}
	return nil, nil
}

func TestParser(t *testing.T) {
	mc := &MockClient{}
	type fields struct {
		commands []interfaces.Command
	}
	type args struct {
		args []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "upload - failure - does not exist",
			fields: fields{
				commands: []interfaces.Command{
					NewUploadCommand(mc),
				},
			},
			args: args{
				args: []string{"upload", "-filename", "dummy.mp3"},
			},
			wantErr: true,
		},
		{
			name: "upload - success - uploaded",
			fields: fields{
				commands: []interfaces.Command{
					NewUploadCommand(mc),
				},
			},
			args: args{
				args: []string{"upload", "-filename", "testdata/dummy.mp3"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				commands: tt.fields.commands,
			}
			err := p.Parse(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
