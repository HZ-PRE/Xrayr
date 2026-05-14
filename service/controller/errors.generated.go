package controller

import "github.com/HZ-PRE/XrarCore/common/errors"

func newError(values ...interface{}) *errors.Error {
	return errors.New(values...)
}
