package errors

import (
	"encoding/json"
	"testing"

	"github.com/SpiralMindJP/at-sdk/go/errors/errdetails"
	"github.com/google/go-cmp/cmp"
)

var safeArrayTests = map[string]struct {
	array safeArray[string]
	json  string
}{
	"some items": {
		array: []string{"foo", "bar"},
		json:  `["foo","bar"]`,
	},
	"empty": {
		array: []string{},
		json:  `[]`,
	},
	"nil": {
		array: nil,
		json:  `[]`,
	},
}

func TestSafeArray(t *testing.T) {
	for name, tt := range safeArrayTests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			b, err := json.Marshal(tt.array)
			if err != nil {
				t.Fatal(err)
			}
			got := string(b)
			if got != tt.json {
				t.Errorf("json.Marshal(%v): got %v, want %v", tt.array, got, tt.json)
			}
		})
	}
}

func TestNewJSONError(t *testing.T) {
	const testMessage = "test error"
	const testCode = InvalidArgument

	want := &jsonError{
		Code:    string(testCode),
		Message: testMessage,
		Details: []any{
			&jsonErrorInfo{
				Type:   errorInfoType,
				Reason: "test/error",
			},
			&jsonBadRequest{
				Type: badRequestType,
				FieldViolations: []*jsonBadRequestFieldViolation{
					{Field: "id", Description: "zero"},
					{Field: "name", Description: "empty"},
				},
			},
			&jsonPreconditionFailure{
				Type: preconditionFailureType,
				Violations: []*jsonPreconditionViolation{
					{Type: "foo", Subject: "FOO", Description: "foo foo"},
					{Type: "bar", Subject: "BAR", Description: "bar bar"},
				},
			},
			&jsonQuotaFailure{
				Type: quotaFailureType,
				Violations: []*jsonQuotaViolation{
					{Subject: "FOO", Description: "foo foo"},
					{Subject: "BAR", Description: "bar bar"},
				},
			},
		},
	}

	got := newJSONError(New(testMessage).
		WithCode(testCode).
		WithDetails(
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
		))

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("newJSONError(), differs: (-want +got)\n%s", diff)
	}
}
