// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package category

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsErrCategoryNotExist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CategoryErrorReason_ERR_CATEGORY_NOT_EXIST.String() && e.Code == 404
}

func ErrorErrCategoryNotExist(format string, args ...interface{}) *errors.Error {
	return errors.New(404, CategoryErrorReason_ERR_CATEGORY_NOT_EXIST.String(), fmt.Sprintf(format, args...))
}

func IsErrCategoryInvalidTitle(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CategoryErrorReason_ERR_CATEGORY_INVALID_TITLE.String() && e.Code == 400
}

func ErrorErrCategoryInvalidTitle(format string, args ...interface{}) *errors.Error {
	return errors.New(400, CategoryErrorReason_ERR_CATEGORY_INVALID_TITLE.String(), fmt.Sprintf(format, args...))
}

func IsErrCategoryInvalidId(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CategoryErrorReason_ERR_CATEGORY_INVALID_ID.String() && e.Code == 400
}

func ErrorErrCategoryInvalidId(format string, args ...interface{}) *errors.Error {
	return errors.New(400, CategoryErrorReason_ERR_CATEGORY_INVALID_ID.String(), fmt.Sprintf(format, args...))
}

func IsErrCategoryPreExisting(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CategoryErrorReason_ERR_CATEGORY_PRE_EXISTING.String() && e.Code == 400
}

func ErrorErrCategoryPreExisting(format string, args ...interface{}) *errors.Error {
	return errors.New(400, CategoryErrorReason_ERR_CATEGORY_PRE_EXISTING.String(), fmt.Sprintf(format, args...))
}

func IsErrCategoryForeignKeyConstraint(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CategoryErrorReason_ERR_CATEGORY_FOREIGN_KEY_CONSTRAINT.String() && e.Code == 400
}

func ErrorErrCategoryForeignKeyConstraint(format string, args ...interface{}) *errors.Error {
	return errors.New(400, CategoryErrorReason_ERR_CATEGORY_FOREIGN_KEY_CONSTRAINT.String(), fmt.Sprintf(format, args...))
}
