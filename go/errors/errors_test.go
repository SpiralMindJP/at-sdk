package errors

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/SpiralMindJP/at-sdk/go/errors/errdetails"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"golang.org/x/exp/slog"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

var errorCodeTests = map[string]struct {
	statusCode int
	grpcCode   codes.Code
	code       ErrorCode
}{
	"NotFound": {
		statusCode: http.StatusNotFound,
		grpcCode:   codes.NotFound,
		code:       NotFound,
	},
	"InvalidArgument": {
		statusCode: http.StatusBadRequest,
		grpcCode:   codes.InvalidArgument,
		code:       InvalidArgument,
	},
	"FailedPrecondition": {
		statusCode: http.StatusBadRequest,
		grpcCode:   codes.FailedPrecondition,
		code:       FailedPrecondition,
	},
	"Aborted": {
		statusCode: http.StatusConflict,
		grpcCode:   codes.Aborted,
		code:       Aborted,
	},
	"AlreadyExists": {
		statusCode: http.StatusConflict,
		grpcCode:   codes.AlreadyExists,
		code:       AlreadyExists,
	},
	"Unauthenticated": {
		statusCode: http.StatusUnauthorized,
		grpcCode:   codes.Unauthenticated,
		code:       Unauthenticated,
	},
	"PermissionDenied": {
		statusCode: http.StatusForbidden,
		grpcCode:   codes.PermissionDenied,
		code:       PermissionDenied,
	},
	"Unknown": {
		statusCode: http.StatusInternalServerError,
		grpcCode:   codes.Unknown,
		code:       Unknown,
	},
}

func TestNewErrorCodeFromGRPCCode(t *testing.T) {
	for name, tt := range errorCodeTests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			code := newErrorCodeFromGRPCCode(tt.grpcCode)
			if code != tt.code {
				t.Errorf("newErrorCodeFromGRPCCode(%v): got %v, want %v", tt.grpcCode, code, tt.code)
			}
		})
	}
}

func TestErrorCodeToHTTPStatus(t *testing.T) {
	for name, tt := range errorCodeTests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			statusCode := tt.code.toHTTPStatus()
			if statusCode != tt.statusCode {
				t.Errorf("ErrorCode(%v).toHTTPStatus(): got %v, want %v", tt.code, statusCode, tt.statusCode)
			}
		})
	}
}

func TestErrorCodeToGRPCCode(t *testing.T) {
	for name, tt := range errorCodeTests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			grpcCode := tt.code.toGRPCCode()
			if grpcCode != tt.grpcCode {
				t.Errorf("ErrorCode(%v).toGRPCCode(): got %v, want %v", tt.code, grpcCode, tt.grpcCode)
			}
		})
	}
}

func TestNew(t *testing.T) {
	const errorMessage = "test error"

	err1 := New(errorMessage)
	if err1 == nil {
		t.Fatal("New() should not return nil")
	}
	if msg := err1.Message; msg != errorMessage {
		t.Errorf("Error.Message: got %v, want %v", msg, errorMessage)
	}
	if code := err1.Code; code != Unknown {
		t.Errorf("Error.Code: got %v, want %v", code, Unknown)
	}
	if code := err1.HTTPStatusCode(); code != http.StatusInternalServerError {
		t.Errorf("Error.HTTPStatusCode(): got %v, want %v", code, http.StatusInternalServerError)
	}

	err2 := Newf("test %s", "error")
	if err2 == nil {
		t.Fatal("New() should not return nil")
	}
	if msg := err2.Message; msg != errorMessage {
		t.Errorf("Error.Message: got %v, want %v", msg, errorMessage)
	}
	if code := err2.Code; code != Unknown {
		t.Errorf("Error.Code: got %v, want %v", code, Unknown)
	}
	if code := err2.HTTPStatusCode(); code != http.StatusInternalServerError {
		t.Errorf("Error.HTTPStatusCode(): got %v, want %v", code, http.StatusInternalServerError)
	}
}

func TestConnection(t *testing.T) {
	err := New("test error")

	err = err.WithCode(NotFound)
	if code := err.Code; code != NotFound {
		t.Errorf("Error.Code: got %v, want %v", code, NotFound)
	}

	innerErr := errors.New("inner error")
	err = err.WithError(innerErr)
	if err := err.Err; err != innerErr {
		t.Errorf("Error.Err: got %v, want %v", err, innerErr)
	}
	if err := err.Unwrap(); err != innerErr {
		t.Errorf("Error.Unwrap(): got %v, want %v", err, innerErr)
	}

	err = err.WithSource()
	if err.Source == nil {
		t.Fatal("Error.Source: should not be nil")
	}

	var testAttrs = []slog.Attr{
		slog.Int("test_num", 10),
		slog.String("test_str", "test"),
		slog.Bool("test_bool", true),
	}
	err = err.WithAttrs(testAttrs[:2]...)
	if attrs := err.Attrs; !reflect.DeepEqual(attrs, testAttrs[:2]) {
		t.Errorf("Error.Attrs: got %v, want %v", attrs, testAttrs)
	}
	err = err.AppendAttrs(testAttrs[2:]...)
	if attrs := err.Attrs; !reflect.DeepEqual(attrs, testAttrs) {
		t.Errorf("Error.Attrs: got %v, want %v", attrs, testAttrs)
	}

	var testDetails = []proto.Message{
		&errdetails.BadRequest{
			FieldViolations: []*errdetails.BadRequestFieldViolation{
				{
					Field:       "name",
					Description: "empty",
				},
			},
		},
	}
	err = err.WithDetails(testDetails...)
	if details := err.Details; !reflect.DeepEqual(details, testDetails) {
		t.Errorf("Error.Details: got %v, want %v", details, testDetails)
	}
}

var errorTests = map[string]struct {
	err *Error
	s   string
}{
	"nil error": {
		err: nil,
		s:   "(unknown)",
	},
	"unknown error code": {
		err: New("test error"),
		s:   "test error [unknown]",
	},
	"with error code": {
		err: New("test error").WithCode(NotFound),
		s:   "test error [not-found]",
	},
	"with inner error/unknown error code": {
		err: New("test error").WithError(errors.New("inner error")),
		s:   "test error [unknown]: inner error",
	},
	"with inner error/with error code": {
		err: New("test error").WithError(errors.New("inner error")).WithCode(NotFound),
		s:   "test error [not-found]: inner error",
	},
}

func TestIsError(t *testing.T) {
	err1 := New("test error")
	if !IsError(err1) {
		t.Error("IsError() should return true")
	}
	err2 := fmt.Errorf("test error: %w", New("test inner error"))
	if !IsError(err2) {
		t.Error("IsError() should return true")
	}
	err3 := errors.New("test error")
	if IsError(err3) {
		t.Error("IsError() should return false")
	}
	err4 := fmt.Errorf("test error :%w", errors.New("test inner error"))
	if IsError(err4) {
		t.Error("IsError() should return false")
	}
}

func TestIsKnownError(t *testing.T) {
	err1 := New("test error")
	if IsKnownError(err1) {
		t.Error("IsKnownError() should return false")
	}
	err2 := fmt.Errorf("test error: %w", New("test inner error"))
	if IsKnownError(err2) {
		t.Error("IsKnownError() should return false")
	}

	err3 := NewNotFound("test error")
	if !IsKnownError(err3) {
		t.Error("IsKnownError() should return true")
	}
	err4 := fmt.Errorf("test error: %w", NewNotFound("test inner error"))
	if !IsKnownError(err4) {
		t.Error("IsKnownError() should return true")
	}

	err5 := errors.New("test error")
	if IsKnownError(err5) {
		t.Error("IsKnownError() should return false")
	}
	err6 := fmt.Errorf("test error :%w", errors.New("test inner error"))
	if IsKnownError(err6) {
		t.Error("IsKnownError() should return false")
	}
}

func TestAsError(t *testing.T) {
	err1 := New("test error")
	if err, ok := AsError(err1); !ok {
		t.Error("AsError() should return true")
	} else if err != err1 {
		t.Errorf("AsError(): got %v, want %v", err, err1)
	}

	err2 := fmt.Errorf("test error: %w", err1)
	if err, ok := AsError(err2); !ok {
		t.Error("AsError() should return true")
	} else if err != err1 {
		t.Errorf("AsError(): got %v, want %v", err, err1)
	}

	err3 := errors.New("test error")
	if err, ok := AsError(err3); ok {
		t.Error("AsError() should return false")
	} else if err != nil {
		t.Error("AsError() should return nil error")
	}

	err4 := fmt.Errorf("test error :%w", errors.New("test inner error"))
	if err, ok := AsError(err4); ok {
		t.Error("AsError() should return false")
	} else if err != nil {
		t.Error("AsError() should return nil error")
	}
}

func TestError(t *testing.T) {
	for name, tt := range errorTests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			if s := tt.err.Error(); s != tt.s {
				t.Errorf("Error.Error(): got %v, want %v", s, tt.s)
			}
		})
	}
}

func TestErrorUnwrap(t *testing.T) {
	innerErr := errors.New("inner error")

	err1 := New("test error")
	if err := err1.Unwrap(); err != nil {
		t.Error("Error.Unwrap(): should return nil")
	}

	err2 := New("test error").WithError(innerErr)
	if err := err2.Unwrap(); err != innerErr {
		t.Errorf("Error.Unwrap(): got %v, want %v", err, innerErr)
	}

	err3 := (*Error)(nil)
	if err := err3.Unwrap(); err != nil {
		t.Error("Error(nil).Unwrap(): should return nil")
	}
}

var errorErrorCodeTests = map[string]struct {
	err        *Error
	statusCode int
	grpcCode   codes.Code
}{
	"NotFound": {
		err:        NewNotFound("test error"),
		statusCode: http.StatusNotFound,
		grpcCode:   codes.NotFound,
	},
	"InvalidArgument": {
		err:        NewInvalidArgument("test error"),
		statusCode: http.StatusBadRequest,
		grpcCode:   codes.InvalidArgument,
	},
	"FailedPrecondition": {
		err:        NewFailedPrecondition("test error"),
		statusCode: http.StatusBadRequest,
		grpcCode:   codes.FailedPrecondition,
	},
	"Aborted": {
		err:        NewAborted("test error"),
		statusCode: http.StatusConflict,
		grpcCode:   codes.Aborted,
	},
	"AlreadyExists": {
		err:        NewAlreadyExists("test error"),
		statusCode: http.StatusConflict,
		grpcCode:   codes.AlreadyExists,
	},
	"Unauthenticated": {
		err:        NewUnauthenticated("test error"),
		statusCode: http.StatusUnauthorized,
		grpcCode:   codes.Unauthenticated,
	},
	"PermissionDenied": {
		err:        NewPermissionDenied("test error"),
		statusCode: http.StatusForbidden,
		grpcCode:   codes.PermissionDenied,
	},
	"Unknown": {
		err:        New("test error"),
		statusCode: http.StatusInternalServerError,
		grpcCode:   codes.Unknown,
	},
	"Unknown(nil)": {
		err:        nil,
		statusCode: http.StatusInternalServerError,
		grpcCode:   codes.Unknown,
	},
}

func TestErrorHTTPStatusCode(t *testing.T) {
	for name, tt := range errorErrorCodeTests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			statusCode := tt.err.HTTPStatusCode()
			if statusCode != tt.statusCode {
				t.Errorf("Error(%v).HTTPStatusCode(): got %v, want %v", tt.err, statusCode, tt.statusCode)
			}
		})
	}
}

func TestErrorGRPCCode(t *testing.T) {
	for name, tt := range errorErrorCodeTests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			grpcCode := tt.err.GRPCCode()
			if grpcCode != tt.grpcCode {
				t.Errorf("Error(%v).GRPCCode(): got %v, want %v", tt.err, grpcCode, tt.grpcCode)
			}
		})
	}
}

func TestGRPCStatus(t *testing.T) {
	const errorMessage = "test error"
	var testDetails = []proto.Message{
		&errdetails.BadRequest{
			FieldViolations: []*errdetails.BadRequestFieldViolation{
				{
					Field:       "id",
					Description: "zero",
				},
				{
					Field:       "name",
					Description: "empty",
				},
			},
		},
	}

	testErr := New(errorMessage).
		WithCode(InvalidArgument).
		WithDetails(testDetails...)

	st := testErr.ToGRPCStatus()
	if st == nil {
		t.Fatal("Error.ToGRPCStatus() should not return nil")
	}
	if code := st.Code(); code != codes.InvalidArgument {
		t.Errorf("status.Status.Code(): got %v, want %v", st, codes.InvalidArgument)
	}
	if msg := st.Message(); msg != errorMessage {
		t.Errorf("status.Status.Message(): got %v, want %v", msg, errorMessage)
	}

	sp := st.Proto()
	b, err := proto.Marshal(sp)
	if err != nil {
		t.Fatal(err)
	}

	testSt := &spb.Status{}
	err = proto.Unmarshal(b, testSt)
	if err != nil {
		t.Fatal(err)
	}

	diffOpts := []cmp.Option{
		cmpopts.IgnoreUnexported(errdetails.BadRequest{}, errdetails.BadRequestFieldViolation{}),
	}

	testErr2 := FromGRPCStatus(status.FromProto(testSt))
	if diff := cmp.Diff(testErr.Details, testErr2.Details, diffOpts...); diff != "" {
		t.Errorf("Error.Details, differs: (-want +got)\n%s", diff)
	}

	testErr3 := FromGRPCError(status.FromProto(testSt).Err())
	if diff := cmp.Diff(testErr.Details, testErr3.Details, diffOpts...); diff != "" {
		t.Errorf("Error.Details, differs: (-want +got)\n%s", diff)
	}
}

func TestMarshalJSON(t *testing.T) {
	const testMessage = "test message"
	const testCode = InvalidArgument
	var testDetails = []proto.Message{
		&errdetails.ErrorInfo{
			Reason: "test/error",
		},
		&errdetails.BadRequest{
			FieldViolations: []*errdetails.BadRequestFieldViolation{
				{Field: "id", Description: "zero"},
				{Field: "name", Description: "empty"},
			},
		},
		&errdetails.PreconditionFailure{
			Violations: []*errdetails.PreconditionFailureViolation{
				{Type: "foo", Subject: "FOO", Description: "foo foo"},
				{Type: "bar", Subject: "BAR", Description: "bar bar"},
			},
		},
		&errdetails.QuotaFailure{
			Violations: []*errdetails.QuotaFailureViolation{
				{Subject: "FOO", Description: "foo foo"},
				{Subject: "BAR", Description: "bar bar"},
			},
		},
	}

	var json = []byte(fmt.Sprintf(`{"code":"%s","message":"%s","details":[`+
		`{"type":"error-info","reason":"test/error"},`+
		`{"type":"bad-request","fieldViolations":[{"field":"id","description":"zero"},{"field":"name","description":"empty"}]},`+
		`{"type":"precondition-failure","violations":[{"type":"foo","subject":"FOO","description":"foo foo"},{"type":"bar","subject":"BAR","description":"bar bar"}]},`+
		`{"type":"quota-failure","violations":[{"subject":"FOO","description":"foo foo"},{"subject":"BAR","description":"bar bar"}]}`+
		`]}`, testCode, testMessage))

	b, err := New(testMessage).WithCode(testCode).WithDetails(testDetails...).MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	want := string(json)
	got := string(b)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Error.MarshalJSON(), differs: (-want +got)\n%s", diff)
	}
}

var newSomthingTests = map[string]struct {
	newFunc  func(text string) *Error
	newfFunc func(format string, a ...any) *Error
	code     ErrorCode
}{
	"NotFound": {
		newFunc:  NewNotFound,
		newfFunc: NewNotFoundf,
		code:     NotFound,
	},
	"InvalidArgument": {
		newFunc:  NewInvalidArgument,
		newfFunc: NewInvalidArgumentf,
		code:     InvalidArgument,
	},
	"FailedPrecondition": {
		newFunc:  NewFailedPrecondition,
		newfFunc: NewFailedPreconditionf,
		code:     FailedPrecondition,
	},
	"Aborted": {
		newFunc:  NewAborted,
		newfFunc: NewAbortedf,
		code:     Aborted,
	},
	"AlreadyExists": {
		newFunc:  NewAlreadyExists,
		newfFunc: NewAlreadyExistsf,
		code:     AlreadyExists,
	},
	"Unauthenticated": {
		newFunc:  NewUnauthenticated,
		newfFunc: NewUnauthenticatedf,
		code:     Unauthenticated,
	},
	"PermissionDenied": {
		newFunc:  NewPermissionDenied,
		newfFunc: NewPermissionDeniedf,
		code:     PermissionDenied,
	},
}

func TestNewSomthing(t *testing.T) {
	const testMessage = "test message"
	for name, tt := range newSomthingTests {
		name := name
		tt := tt
		t.Run(name, func(t *testing.T) {
			err1 := tt.newFunc(testMessage)
			if err1.Message != testMessage {
				t.Errorf("New%s().Message: got %v, want %v", name, err1.Message, testMessage)
			}
			if err1.Code != tt.code {
				t.Errorf("New%s().Code: got %v, want %v", name, err1.Code, tt.code)
			}

			err2 := tt.newfFunc("test %s", "message")
			if err2.Message != testMessage {
				t.Errorf("New%s().Message: got %v, want %v", name, err2.Message, testMessage)
			}
			if err2.Code != tt.code {
				t.Errorf("New%s().Code: got %v, want %v", name, err2.Code, tt.code)
			}
		})
	}
}

var isErrorCodeTests = map[string]struct {
	err  error
	code ErrorCode
	is   bool
}{
	"err is NotFound": {
		err:  NewNotFound("test error"),
		code: NotFound,
		is:   true,
	},
	"err is not NotFound": {
		err:  NewNotFound("test error"),
		code: Unknown,
		is:   false,
	},
	"err is InvalidArgument": {
		err:  NewInvalidArgument("test error"),
		code: InvalidArgument,
		is:   true,
	},
	"err is not InvalidArgument": {
		err:  NewInvalidArgument("test error"),
		code: NotFound,
		is:   false,
	},
	"err is not Error": {
		err:  errors.New("not Error"),
		code: Unknown,
		is:   false,
	},
}

func TestIsErrorCode(t *testing.T) {
	for name, tt := range isErrorCodeTests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			is := IsErrorCode(tt.err, tt.code)
			if is != tt.is {
				t.Errorf("IsErrorCode(): got %v, want %v", is, tt.is)
			}
		})
	}
}

var isSomthingTests = map[string]struct {
	err    error
	isFunc func(err error) bool
}{
	"IsNotFound": {
		err:    NewNotFound("test error"),
		isFunc: IsNotFound,
	},
	"IsInvalidArgument": {
		err:    NewInvalidArgument("test error"),
		isFunc: IsInvalidArgument,
	},
	"IsFailedPrecondition": {
		err:    NewFailedPrecondition("test error"),
		isFunc: IsFailedPrecondition,
	},
	"IsAborted": {
		err:    NewAborted("test error"),
		isFunc: IsAborted,
	},
	"IsAlreadyExists": {
		err:    NewAlreadyExists("test error"),
		isFunc: IsAlreadyExists,
	},
	"IsUnauthenticated": {
		err:    NewUnauthenticated("test error"),
		isFunc: IsUnauthenticated,
	},
	"IsPermissionDenied": {
		err:    NewPermissionDenied("test error"),
		isFunc: IsPermissionDenied,
	},
}

func TestIsSomthing(t *testing.T) {
	unknownErr := New("unknwon error")

	for name, tt := range isSomthingTests {
		name := name
		tt := tt
		t.Run(name, func(t *testing.T) {
			if !tt.isFunc(tt.err) {
				t.Errorf("%s(%v): should return true", name, tt.err)
			}
			if tt.isFunc(unknownErr) {
				t.Errorf("%s(%v): should not return true", name, tt.err)
			}
		})
	}
}
