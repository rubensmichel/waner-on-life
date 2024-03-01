package http

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	domainErrors "github.com/rubensmichel/waner-on-life/internal/types/errors"
)

type Request struct {
	DomainName string            `json:"domainName"`
	UrlPath    string            `json:"urlPath"`
	Method     string            `json:"method"`
	Headers    map[string]string `json:"headers"`
	Body       []byte            `json:"body"`
	Params     map[string]string `json:"params"`
	Query      map[string]string `json:"queryParams"`
}

func (req Request) ParseQuery(name string) string {
	return req.Query[name]
}

func (req Request) ParseQueryInt(name string) int {
	intVar, err := strconv.Atoi(req.Query[name])
	if err != nil {
		return 0
	}
	return intVar
}

func (req Request) ParseParamString(name string) string {
	return req.Params[name]
}

func (req Request) ParseParamInt(name string) int {
	intVar, err := strconv.Atoi(req.Params[name])
	if err != nil {
		return 0
	}
	return intVar
}

func (req Request) ParseBody(ctx context.Context, result interface{}) error {
	err := json.Unmarshal(req.Body, result)
	if err != nil {
		var unmarshalErr *json.UnmarshalTypeError
		if errors.As(err, &unmarshalErr) {
			err = errors.Join(err, domainErrors.ErrInvalidInput.WithContext(ctx).WithDetailMessage(unmarshalErr.Field, domainErrors.DetailInvalidDataType))
		} else {
			err = errors.Join(err, domainErrors.ErrInvalidInput.WithContext(ctx).WithDetailDescription(err.Error(), domainErrors.DetailInvalidDataType))
		}
		return err
	}

	return nil
}

func (req Request) String() string {
	return req.Method + " " + req.UrlPath
}
