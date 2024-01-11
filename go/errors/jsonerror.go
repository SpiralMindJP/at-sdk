package errors

import (
	"encoding/json"

	"github.com/SpiralMindJP/at-sdk/go/errors/errdetails"
	"google.golang.org/protobuf/proto"
)

type safeArray[T any] []T

var _ json.Marshaler = safeArray[int](nil)

func (a safeArray[T]) MarshalJSON() ([]byte, error) {
	if len(a) == 0 {
		var empty [0]T
		return json.Marshal(empty)
	}

	return json.Marshal([]T(a))
}

type detailType string

const (
	errorInfoType           detailType = "error-info"
	badRequestType          detailType = "bad-request"
	preconditionFailureType detailType = "precondition-failure"
	quotaFailureType        detailType = "quota-failure"
)

type jsonErrorInfo struct {
	Type   detailType `json:"type"`
	Reason string     `json:"reason"`
}

func newJSONErrorInfo(info *errdetails.ErrorInfo) *jsonErrorInfo {
	return &jsonErrorInfo{
		Type:   errorInfoType,
		Reason: info.GetReason(),
	}
}

type jsonBadRequestFieldViolation struct {
	Field       string `json:"field"`
	Description string `json:"description"`
}

func newJSONBadRequestFieldViolation(fv *errdetails.BadRequestFieldViolation) *jsonBadRequestFieldViolation {
	return &jsonBadRequestFieldViolation{
		Field:       fv.GetField(),
		Description: fv.GetDescription(),
	}
}

func newJSONBadRequestFieldViolations(violations []*errdetails.BadRequestFieldViolation) []*jsonBadRequestFieldViolation {
	if len(violations) == 0 {
		return nil
	}

	results := make([]*jsonBadRequestFieldViolation, len(violations))
	for i := range violations {
		results[i] = newJSONBadRequestFieldViolation(violations[i])
	}

	return results
}

type jsonBadRequest struct {
	Type            detailType                               `json:"type"`
	FieldViolations safeArray[*jsonBadRequestFieldViolation] `json:"fieldViolations"`
}

func newJSONBadRequest(br *errdetails.BadRequest) *jsonBadRequest {
	return &jsonBadRequest{
		Type:            badRequestType,
		FieldViolations: newJSONBadRequestFieldViolations(br.GetFieldViolations()),
	}
}

type jsonPreconditionViolation struct {
	Type        string `json:"type"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
}

func newJSONPreconditionViolation(fv *errdetails.PreconditionFailureViolation) *jsonPreconditionViolation {
	return &jsonPreconditionViolation{
		Type:        fv.GetType(),
		Subject:     fv.GetSubject(),
		Description: fv.GetDescription(),
	}
}

func newJSONPreconditionViolations(violations []*errdetails.PreconditionFailureViolation) []*jsonPreconditionViolation {
	if len(violations) == 0 {
		return nil
	}

	results := make([]*jsonPreconditionViolation, len(violations))
	for i := range violations {
		results[i] = newJSONPreconditionViolation(violations[i])
	}

	return results
}

type jsonPreconditionFailure struct {
	Type       detailType                            `json:"type"`
	Violations safeArray[*jsonPreconditionViolation] `json:"violations"`
}

func newJSONPreconditionFailure(br *errdetails.PreconditionFailure) *jsonPreconditionFailure {
	return &jsonPreconditionFailure{
		Type:       preconditionFailureType,
		Violations: newJSONPreconditionViolations(br.GetViolations()),
	}
}

type jsonQuotaViolation struct {
	Subject     string `json:"subject"`
	Description string `json:"description"`
}

func newJSONQuotaViolation(fv *errdetails.QuotaFailureViolation) *jsonQuotaViolation {
	return &jsonQuotaViolation{
		Subject:     fv.GetSubject(),
		Description: fv.GetDescription(),
	}
}

func newJSONQuotaViolations(violations []*errdetails.QuotaFailureViolation) []*jsonQuotaViolation {
	if len(violations) == 0 {
		return nil
	}

	results := make([]*jsonQuotaViolation, len(violations))
	for i := range violations {
		results[i] = newJSONQuotaViolation(violations[i])
	}

	return results
}

type jsonQuotaFailure struct {
	Type       detailType                     `json:"type"`
	Violations safeArray[*jsonQuotaViolation] `json:"violations"`
}

func newJSONQuotaFailure(br *errdetails.QuotaFailure) *jsonQuotaFailure {
	return &jsonQuotaFailure{
		Type:       quotaFailureType,
		Violations: newJSONQuotaViolations(br.GetViolations()),
	}
}

func newJSONDetails(details []proto.Message) []any {
	if len(details) == 0 {
		return nil
	}

	results := make([]any, 0, len(details))
	for _, req := range details {
		switch v := req.(type) {
		case *errdetails.ErrorInfo:
			results = append(results, newJSONErrorInfo(v))
		case *errdetails.BadRequest:
			results = append(results, newJSONBadRequest(v))
		case *errdetails.PreconditionFailure:
			results = append(results, newJSONPreconditionFailure(v))
		case *errdetails.QuotaFailure:
			results = append(results, newJSONQuotaFailure(v))
		}
	}

	return results
}

type jsonError struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Details safeArray[any] `json:"details"`
}

func newJSONError(err *Error) *jsonError {
	return &jsonError{
		Code:    string(err.Code),
		Message: err.Message,
		Details: newJSONDetails(err.Details),
	}
}
