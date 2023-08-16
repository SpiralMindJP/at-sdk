package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"runtime"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

type ErrorCode string

const (
	// Unknown は、エラーの理由が不明であることを表します。
	Unknown ErrorCode = ""

	// NotFound は、エンティティが存在しないことを表します。
	NotFound ErrorCode = "not-found"

	// InvalidArgument は、クライアントが無効な引数を指定したことを表わします。
	InvalidArgument ErrorCode = "invalid-argument"

	// FailedPrecondition は、システムがオペレーションの実行に必要な状態ではないため、オペレーションが拒否されたことを表わします。
	FailedPrecondition ErrorCode = "failed-precondition"

	// Aborted はオペレーションが実行時の問題のために中止されたことを表わします。。
	Aborted ErrorCode = "aborted"

	// AlreadyExists は、作成しようとしたエンティティが既に存在していることを表わします。
	AlreadyExists ErrorCode = "already-exists"

	// Unauthenticated は、操作に対する有効な認証資格情報がないことを表わします。
	Unauthenticated ErrorCode = "unauthenticated"

	// PermissionDenied は、操作に対する権限がないことを表します。
	PermissionDenied ErrorCode = "permission-denied"
)

func newErrorCodeFromGRPCCode(code codes.Code) ErrorCode {
	switch code {
	case codes.NotFound:
		return NotFound
	case codes.InvalidArgument:
		return InvalidArgument
	case codes.FailedPrecondition:
		return FailedPrecondition
	case codes.Aborted:
		return Aborted
	case codes.AlreadyExists:
		return AlreadyExists
	case codes.Unauthenticated:
		return Unauthenticated
	case codes.PermissionDenied:
		return PermissionDenied
	}

	return Unknown
}

func (code ErrorCode) toHTTPStatus() int {
	switch code {
	case NotFound:
		return http.StatusNotFound
	case InvalidArgument:
		return http.StatusBadRequest
	case FailedPrecondition:
		return http.StatusBadRequest
	case Aborted:
		return http.StatusConflict
	case AlreadyExists:
		return http.StatusConflict
	case Unauthenticated:
		return http.StatusUnauthorized
	case PermissionDenied:
		return http.StatusForbidden
	}

	return http.StatusInternalServerError
}

func (code ErrorCode) toGRPCCode() codes.Code {
	switch code {
	case NotFound:
		return codes.NotFound
	case InvalidArgument:
		return codes.InvalidArgument
	case FailedPrecondition:
		return codes.FailedPrecondition
	case Aborted:
		return codes.Aborted
	case AlreadyExists:
		return codes.AlreadyExists
	case Unauthenticated:
		return codes.Unauthenticated
	case PermissionDenied:
		return codes.PermissionDenied
	}

	return codes.Unknown
}

// Error は、HTTP のエラーを表します。
type Error struct {
	// エラーの理由。
	Code ErrorCode

	// エラーメッセージ。
	Message string

	// 内部エラー。
	Err error

	// ログ出力用の属性。
	Attrs []slog.Attr

	// エラーの詳細。
	Details []proto.Message

	// プログラムカウンター。
	pc uintptr
}

var (
	_ json.Marshaler = (*Error)(nil)
)

func newError(code ErrorCode, skip int, msg string) *Error {
	var pcs [1]uintptr

	runtime.Callers(
		skip+2,
		pcs[:],
	)

	return &Error{
		Code:    code,
		Message: msg,
		pc:      pcs[0],
	}
}

func newErrorf(code ErrorCode, skip int, format string, a ...any) *Error {
	return newError(code, skip+1, fmt.Sprintf(format, a...))
}

func New(text string) *Error {
	return newError(Unknown, 1, text)
}

func Newf(format string, a ...any) *Error {
	return newErrorf(Unknown, 1, format, a...)
}

func FromGRPCStatus(s *status.Status) *Error {
	err := newError(newErrorCodeFromGRPCCode(s.Code()), 1, s.Message())

	ps := s.Proto()
	if len(ps.Details) > 0 {
		details := make([]proto.Message, 0, len(ps.Details))
		for _, d := range ps.Details {
			detail, err := d.UnmarshalNew()
			if err == nil {
				details = append(details, detail)
			}
		}
		err.Details = details
	}

	return err
}

func FromGRPCError(err error) *Error {
	return FromGRPCStatus(status.Convert(err))
}

func IsError(err error) bool {
	for err != nil {
		_, ok := err.(*Error)
		if ok {
			return true
		}
		err = errors.Unwrap(err)
	}
	return false
}

func IsKnownError(err error) bool {
	for err != nil {
		atErr, ok := err.(*Error)
		if ok && atErr.Code != Unknown {
			return true
		}
		err = errors.Unwrap(err)
	}
	return false
}

func AsError(err error) (*Error, bool) {
	for err != nil {
		atErr, ok := err.(*Error)
		if ok {
			return atErr, true
		}
		err = errors.Unwrap(err)
	}
	return nil, false
}

func (err *Error) Error() string {
	if err == nil {
		return "(unknown)"
	}

	code := string(err.Code)
	if code == "" {
		code = "unknown"
	}

	if err.Err == nil {
		return fmt.Sprintf("%s [%s]", err.Message, code)
	}

	return fmt.Sprintf("%s [%s]: %v", err.Message, code, err.Err)
}

func (err *Error) Unwrap() error {
	if err == nil {
		return nil
	}

	return err.Err
}

func (err *Error) HTTPStatusCode() int {
	if err == nil {
		return http.StatusInternalServerError
	}

	return err.Code.toHTTPStatus()
}

func (err *Error) GRPCCode() codes.Code {
	if err == nil {
		return codes.Unknown
	}

	return err.Code.toGRPCCode()
}

func (err *Error) ToGRPCStatus() *status.Status {
	s := status.New(err.GRPCCode(), err.Message)

	if len(err.Details) > 0 {
		details := make([]*anypb.Any, len(err.Details))
		for _, detail := range err.Details {
			any, err := anypb.New(detail)
			if err != nil {
				continue
			}
			details = append(details, any)
		}

		sp := s.Proto()
		sp.Details = details

		s = status.FromProto(sp)
	}

	return s
}

func (err *Error) WithCode(code ErrorCode) *Error {
	err.Code = code
	return err
}

func (err *Error) WithError(inner error) *Error {
	err.Err = inner
	return err
}

func (err *Error) WithAttrs(attrs ...slog.Attr) *Error {
	err.Attrs = attrs
	return err
}

func (err *Error) AppendAttrs(attrs ...slog.Attr) *Error {
	err.Attrs = append(err.Attrs, attrs...)
	return err
}

func (err *Error) WithDetails(details ...proto.Message) *Error {
	err.Details = details
	return err
}

func (err *Error) PC() uintptr {
	return err.pc
}

func (err *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(newJSONError(err))
}

func NewNotFound(text string) *Error {
	return newError(NotFound, 1, text)
}

func NewNotFoundf(format string, a ...any) *Error {
	return newErrorf(NotFound, 1, format, a...)
}

func NewInvalidArgument(text string) *Error {
	return newError(InvalidArgument, 1, text)
}

func NewInvalidArgumentf(format string, a ...any) *Error {
	return newErrorf(InvalidArgument, 1, format, a...)
}

func NewFailedPrecondition(text string) *Error {
	return newError(FailedPrecondition, 1, text)
}

func NewFailedPreconditionf(format string, a ...any) *Error {
	return newErrorf(FailedPrecondition, 1, format, a...)
}

func NewAborted(text string) *Error {
	return newError(Aborted, 1, text)
}

func NewAbortedf(format string, a ...any) *Error {
	return newErrorf(Aborted, 1, format, a...)
}

func NewAlreadyExists(text string) *Error {
	return newError(AlreadyExists, 1, text)
}

func NewAlreadyExistsf(format string, a ...any) *Error {
	return newErrorf(AlreadyExists, 1, format, a...)
}

func NewUnauthenticated(text string) *Error {
	return newError(Unauthenticated, 1, text)
}

func NewUnauthenticatedf(format string, a ...any) *Error {
	return newErrorf(Unauthenticated, 1, format, a...)
}

func NewPermissionDenied(text string) *Error {
	return newError(PermissionDenied, 1, text)
}

func NewPermissionDeniedf(format string, a ...any) *Error {
	return newErrorf(PermissionDenied, 1, format, a...)
}

func IsErrorCode(err error, code ErrorCode) bool {
	var _err *Error
	if errors.As(err, &_err) {
		return _err.Code == code
	}

	return false
}

func IsNotFound(err error) bool {
	return IsErrorCode(err, NotFound)
}

func IsInvalidArgument(err error) bool {
	return IsErrorCode(err, InvalidArgument)
}

func IsFailedPrecondition(err error) bool {
	return IsErrorCode(err, FailedPrecondition)
}

func IsAborted(err error) bool {
	return IsErrorCode(err, Aborted)
}

func IsAlreadyExists(err error) bool {
	return IsErrorCode(err, AlreadyExists)
}

func IsUnauthenticated(err error) bool {
	return IsErrorCode(err, Unauthenticated)
}

func IsPermissionDenied(err error) bool {
	return IsErrorCode(err, PermissionDenied)
}
