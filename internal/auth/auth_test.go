package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type args struct {
		headers http.Header
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "Validate base case", args: args{http.Header{"Authorization": []string{"ApiKey SuperCoolAuthValue"}}}, want: "SuperCoolAuthValue", wantErr: false},
		{name: "Validate empty header case", args: args{http.Header{}}, want: "", wantErr: true},
		{name: "Validate malformed header", args: args{http.Header{"Authorization": []string{"NotApiKey SuperCoolAuthValue"}}}, want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.args.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}
