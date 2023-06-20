// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package errors

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsAuthFailure(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_AUTH_FAILURE.String() && e.Code == 403
}

func ErrorAuthFailure(format string, args ...interface{}) *errors.Error {
	return errors.New(403, ErrorReason_AUTH_FAILURE.String(), fmt.Sprintf(format, args...))
}

func IsInvalidPayloadType(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_INVALID_PAYLOAD_TYPE.String() && e.Code == 500
}

func ErrorInvalidPayloadType(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_INVALID_PAYLOAD_TYPE.String(), fmt.Sprintf(format, args...))
}

func IsMessageIdGenerateFailure(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_MESSAGE_ID_GENERATE_FAILURE.String() && e.Code == 500
}

func ErrorMessageIdGenerateFailure(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_MESSAGE_ID_GENERATE_FAILURE.String(), fmt.Sprintf(format, args...))
}

func IsUploadFileTooLarge(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UPLOAD_FILE_TOO_LARGE.String() && e.Code == 500
}

func ErrorUploadFileTooLarge(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_UPLOAD_FILE_TOO_LARGE.String(), fmt.Sprintf(format, args...))
}
