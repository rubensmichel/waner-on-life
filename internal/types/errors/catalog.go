package errors

import (
	"fmt"
)

var (
	ErrInvalidInput         = newError(KindInvalidInput, "0001", "invalid input")
	ErrEntityNotFound       = newError(KindNotFound, "0002", "entity not found")
	ErrInvalidAppClient     = newError(KindNotFound, "0003", "invalid client")
	ErrMalformedPayload     = newError(KindBusinessRule, "0004", "http request payload is malformed")
	ErrRequestNotFound      = newError(KindBusinessRule, "0005", "request not found")
	ErrUnableRetrieveSecret = newError(KindInternalError, "0006", "Unable to retrieve secret from AWS Secrets Manager")
)

func ErrServiceConnection(serviceName string) Error {
	return newError(KindInternalError, "0010", "unable to establish connection with service "+serviceName)
}

func ErrServiceResponse(serviceName string, err error) Error {
	return newError(KindInternalError, "0011", "error in request to service "+serviceName+": "+err.Error())
}

func ErrServiceUnexpectedStatusCode(serviceName string, statusCode int) Error {
	return newError(KindInternalError, "0012", fmt.Sprintf("error in request to service %s: unexpected status code received: %d", serviceName, statusCode))
}

func ErrServiceMalformedResponse(serviceName string) Error {
	return newError(KindInternalError, "0013", "%s service response body is malformed (failed unmarshaling JSON)")
}

func ErrServiceReadTimeout(serviceName string, timeoutMillis int) Error {
	return newError(KindInternalError, "0014", fmt.Sprintf("unable to receive response from %s service within %v milliseconds", serviceName, timeoutMillis))
}

func ErrUsingRepository(repositoryName string, err error) Error {
	return newError(KindInternalError, "0015", fmt.Sprintf("error when using repository %s: %s", repositoryName, err))
}

func ErrSecretNotFound(secretName string) Error {
	return newError(KindInternalError, "0016", fmt.Sprintf("secret '%s' not found in AWS Secrets Manager", secretName))
}

func ErrUnableToConvertSecret(secretName string) Error {
	return newError(KindInternalError, "0017", fmt.Sprintf("unable to convert string value from secret '%s' to JSON format", secretName))
}

func ErrSecretKeyNotFound(secretKey, secretName string) Error {
	return newError(KindInternalError, "0018", fmt.Sprintf("key '%s' not found in secret '%s'", secretKey, secretName))
}

func ErrPathNotFound(path string) Error {
	return newError(KindNotFound, "0019", fmt.Sprintf("URL path '%s' not found", path))
}

func ErrSPDRestrictionTimelineNotUpdated(actionName string) Error {
	return newError(KindInternalError, "0020", fmt.Sprintf("Restriction was %s, but SPD Restriction Timeline was not updated", actionName))
}
