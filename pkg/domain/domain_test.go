package domain

import (
	"reflect"
	"testing"
)

func getClient() *Client {
	return New(Secret{
		APIKey:    "",
		APISecret: "",
	})
}

func TestClient_GetRecord(t *testing.T) {
	c := getClient()
	type args struct {
		domain string
		type_  string
		name   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantNil bool
		wantErr bool
	}{
		{
			args: args{
				domain: "pdmaker.club",
				type_:  "A",
				name:   "@",
			},
			want:    "16.162.22.101",
			wantErr: false,
		},
		{
			args: args{
				domain: "pdmaker.club",
				type_:  "CNAME",
				name:   "@",
			},
			wantNil: true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.GetRecord(tt.args.domain, tt.args.type_, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantNil && !reflect.DeepEqual(got.Data, tt.want) {
				t.Errorf("GetRecord() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_EnsureRecord(t *testing.T) {
	c := getClient()
	type args struct {
		domain string
		type_  string
		name   string
		value  string
	}
	var tests = []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				domain: "pdmaker.club",
				type_:  "A",
				name:   "test",
				value:  "192.168.1.1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.EnsureRecord(tt.args.domain, tt.args.type_, tt.args.name, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("EnsureRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_EnsureRecordSimplified(t *testing.T) {
	c := getClient()
	type args struct {
		domainAddr string
		type_      string
		value      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				domainAddr: "abc.test.pdmaker.club",
				type_:      "A",
				value:      "192.168.1.1",
			},
			wantErr: false,
		},
		{
			args: args{
				domainAddr: "test.pdmaker.club",
				type_:      "A",
				value:      "192.168.1.1",
			},
			wantErr: false,
		},
		{
			args: args{
				domainAddr: "pdmaker.club",
				type_:      "A",
				value:      "16.162.22.101",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.EnsureRecordSimplified(tt.args.domainAddr, tt.args.type_, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("EnsureRecordSimplified() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
