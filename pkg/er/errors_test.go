package er

import (
	"reflect"
	"testing"
)

func TestAPPError_WithArgs(t *testing.T) {
	type fields struct {
		Status int
		Code   int
		Msg    string
		Data   interface{}
	}
	type args struct {
		a []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *APPError
	}{
		{
			name:   "test 123 456 then 123 test 456",
			fields: fields{Msg: "%s test %s"},
			args:   args{a: []interface{}{"123", "456"}},
			want:   &APPError{Msg: "123 test 456"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &APPError{
				Status: tt.fields.Status,
				Code:   tt.fields.Code,
				Msg:    tt.fields.Msg,
				Data:   tt.fields.Data,
			}
			if got := e.WithArgs(tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPPError_Error(t *testing.T) {
	type fields struct {
		Status int
		Code   int
		Msg    string
		Data   interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "error to string then success",
			fields: fields{Msg: "message"},
			want:   "message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &APPError{
				Status: tt.fields.Status,
				Code:   tt.fields.Code,
				Msg:    tt.fields.Msg,
				Data:   tt.fields.Data,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPPError_WithMsg(t *testing.T) {
	type fields struct {
		Status int
		Code   int
		Msg    string
		Data   interface{}
	}
	type args struct {
		msg string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *APPError
	}{
		{
			name:   "append message then success",
			fields: fields{},
			args:   args{msg: "message"},
			want:   &APPError{Status: 0, Code: 0, Msg: "message", Data: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &APPError{
				Status: tt.fields.Status,
				Code:   tt.fields.Code,
				Msg:    tt.fields.Msg,
				Data:   tt.fields.Data,
			}
			if got := e.WithMsg(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPPError_WithData(t *testing.T) {
	type fields struct {
		Status int
		Code   int
		Msg    string
		Data   interface{}
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *APPError
	}{
		{
			name:   "append data then success",
			fields: fields{},
			args:   args{data: "data"},
			want:   &APPError{Status: 0, Code: 0, Msg: "", Data: "data"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &APPError{
				Status: tt.fields.Status,
				Code:   tt.fields.Code,
				Msg:    tt.fields.Msg,
				Data:   tt.fields.Data,
			}
			if got := e.WithData(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithData() = %v, want %v", got, tt.want)
			}
		})
	}
}
