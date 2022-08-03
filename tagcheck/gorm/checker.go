package gorm

type checker interface {
	check(value string) error
}

// match everything
type anyChecker struct{}

func (anyChecker) check(value string) error {
	return nil
}

// match uint
type uintChecker struct{}

func (uintChecker) check(value string) error {
	return nil
}
