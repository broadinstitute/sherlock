package errors

import (
	"errors"
	"fmt"
	"github.com/broadinstitute/sherlock/internal/testutils"
	"testing"
)

func Test_convert(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name       string
		args       args
		wantCode   int
		wantStruct ErrorResponse
	}{
		{
			name:     "500",
			args:     args{err: fmt.Errorf("blah %s blah", InternalServerError)},
			wantCode: 500,
			wantStruct: ErrorResponse{
				ToBlame: "server",
				Type:    InternalServerError,
				Message: fmt.Sprintf("blah %s blah", InternalServerError),
			},
		},
		{
			name:     "400",
			args:     args{err: fmt.Errorf("blah %s blah", BadRequest)},
			wantCode: 400,
			wantStruct: ErrorResponse{
				ToBlame: "client",
				Type:    BadRequest,
				Message: fmt.Sprintf("blah %s blah", BadRequest),
			},
		},
		{
			name:     "403",
			args:     args{err: fmt.Errorf("blah %s blah", Forbidden)},
			wantCode: 403,
			wantStruct: ErrorResponse{
				ToBlame: "client",
				Type:    Forbidden,
				Message: fmt.Sprintf("blah %s blah", Forbidden),
			},
		},
		{
			name:     "404",
			args:     args{err: fmt.Errorf("blah %s blah", NotFound)},
			wantCode: 404,
			wantStruct: ErrorResponse{
				ToBlame: "client",
				Type:    NotFound,
				Message: fmt.Sprintf("blah %s blah", NotFound),
			},
		},
		{
			name:     "407",
			args:     args{err: fmt.Errorf("blah %s blah", ProxyAuthenticationRequired)},
			wantCode: 407,
			wantStruct: ErrorResponse{
				ToBlame: "client",
				Type:    ProxyAuthenticationRequired,
				Message: fmt.Sprintf("blah %s blah", ProxyAuthenticationRequired),
			},
		},
		{
			name:     "409",
			args:     args{err: fmt.Errorf("blah %s blah", Conflict)},
			wantCode: 409,
			wantStruct: ErrorResponse{
				ToBlame: "client",
				Type:    Conflict,
				Message: fmt.Sprintf("blah %s blah", Conflict),
			},
		},
		{
			name:     "default to 500",
			args:     args{err: errors.New("blah blah")},
			wantCode: 500,
			wantStruct: ErrorResponse{
				ToBlame: "server",
				Type:    InternalServerError,
				Message: "blah blah",
			},
		},
		{
			name:     "500 overrides other error types",
			args:     args{err: fmt.Errorf("%s %s %s", Conflict, InternalServerError, BadRequest)},
			wantCode: 500,
			wantStruct: ErrorResponse{
				ToBlame: "server",
				Type:    InternalServerError,
				Message: fmt.Sprintf("%s %s %s", Conflict, InternalServerError, BadRequest),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCode, gotStruct := convert(tt.args.err)
			testutils.AssertNoDiff(t, tt.wantCode, gotCode)
			testutils.AssertNoDiff(t, tt.wantStruct, gotStruct)
		})
	}
}
