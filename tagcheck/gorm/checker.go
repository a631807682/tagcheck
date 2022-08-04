package gorm

import (
	"errors"
	"strconv"
)

type checker interface {
	check(value string) error
}

// match everything
type anyChecker struct{}

func (anyChecker) check(value string) error {
	return nil
}

// not empty
type notEmptyChecker struct{}

func (notEmptyChecker) check(value string) error {
	if value == "" {
		return errors.New("can not be empty")
	}
	return nil
}

// match uint
type uintChecker struct{}

func (uintChecker) check(value string) error {
	num, err := strconv.Atoi(value)
	if err != nil || num < 0 {
		return errors.New("not an uint")
	}

	return nil
}

// match int
type intChecker struct{}

func (intChecker) check(value string) error {
	_, err := strconv.Atoi(value)
	if err != nil {
		return errors.New("not an int")
	}

	return nil
}

type emptyOrBoolChecker struct{}

func (emptyOrBoolChecker) check(value string) error {
	if value == "" {
		return nil
	}

	booltypes := []string{"false", "true", "1", "0"}
	var isBool bool
	for _, bt := range booltypes {
		if bt == value {
			isBool = true
			break
		}
	}

	if !isBool {
		return errors.New("not empty or bool")
	}
	return nil
}
