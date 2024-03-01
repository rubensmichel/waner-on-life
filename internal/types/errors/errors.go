package errors

import (
	"context"
	"errors"
	"fmt"

	"github.com/rubensmichel/waner-on-life/internal/infra/telemetry/request"
)

type kindType int

const (
	KindInvalidInput kindType = iota
	KindBusinessRule
	KindInternalError
	KindNotFound
	KindConflict
	KindNotImplemented

	applicationErrorPrefix = "FPS-"
)

type Error struct {
	errorType
}

type errorType struct {
	Kind        kindType `json:"-"`
	RequestID   string   `json:"id,omitempty"`
	Description string   `json:"description"`
	Code        string   `json:"code,omitempty"`
	Details     []Detail `json:"error_details,omitempty"`
}

type detailMessage string

const (
	DetailInvalidDataType          detailMessage = "INVALID_DATA_TYPE"
	DetailInvalidLength            detailMessage = "INVALID_LENGTH"
	DetailInvalidFormat            detailMessage = "INVALID_FORMAT"
	DetailInvalidValue             detailMessage = "INVALID_VALUE"
	DetailRequiredAttributeMissing detailMessage = "REQUIRED_ATTRIBUTE_MISSING"
)

type Detail struct {
	Attribute   string          `json:"attribute,omitempty"`
	Description string          `json:"description,omitempty"`
	Messages    []detailMessage `json:"messages"`
}

func newError(kind kindType, code, description string) Error {
	if code != "" {
		code = applicationErrorPrefix + code
	}
	return Error{
		errorType: errorType{
			Code:        code,
			Kind:        kind,
			Description: description,
		},
	}
}

func (e Error) Error() string {
	return e.Description
}

func (e Error) Is(target error) bool {
	var other Error
	if errors.As(target, &other) {
		return e.Kind == other.Kind && e.Code == other.Code && e.Description == other.Description
	}
	return false
}

func (e Error) WithContext(ctx context.Context) Error {
	req := request.FromContext(ctx)
	if req != nil {
		e.RequestID = req.ID
	}
	return e
}

func (e Error) WithDetail(detail Detail) Error {
	e.Details = append(e.Details, detail)
	return e
}

func (e Error) WithDetailMessage(attribute string, m detailMessage) Error {
	for i, d := range e.Details {
		if d.Attribute == attribute {
			e.Details[i].AddMessage(m)
			return e
		}
	}

	e.Details = append(e.Details, *NewDetail().SetAttribute(attribute).AddMessage(m))
	return e
}

func (e Error) WithDetailDescription(description string, m detailMessage) Error {
	e.Details = append(e.Details, *NewDetail().SetDescription(description).AddMessage(m))
	return e
}

func (e Error) Log() errorType {
	return e.errorType
}

func (e Error) Equals(other Error) bool {
	if !e.Is(other) || e.RequestID != other.RequestID || len(e.Details) != len(other.Details) {
		return false
	}

	for i := range other.Details {
		if e.Details[i].Attribute != other.Details[i].Attribute ||
			len(e.Details[i].Messages) != len(other.Details[i].Messages) {
			return false
		}
		for j := range other.Details[i].Messages {
			if e.Details[i].Messages[j] != other.Details[i].Messages[j] {
				return false
			}
		}
	}
	return true
}

func NewDetail() *Detail {
	return &Detail{}
}

func (d *Detail) SetAttribute(attribute string) *Detail {
	d.Attribute = attribute
	return d
}

func (d *Detail) SetDescription(description string) *Detail {
	d.Description = description
	return d
}

func (d *Detail) AddMessage(m detailMessage) *Detail {
	d.Messages = append(d.Messages, m)
	return d
}

func FormatMessage(parent *Error, customMessageParameters ...any) *Error {
	newAppErr := *parent
	newAppErr.Description = fmt.Sprintf(parent.Description, customMessageParameters...)
	return &newAppErr
}
