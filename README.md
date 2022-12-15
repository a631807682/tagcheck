# tagcheck

Go static analysis, detecting gorm tag bugs.

[![go report card](https://goreportcard.com/badge/github.com/a631807682/tagcheck "go report card")](https://goreportcard.com/report/github.com/a631807682/tagcheck)
[![test status](https://github.com/a631807682/tagcheck/workflows/tests/badge.svg?branch=main "test status")](https://github.com/a631807682/tagcheck/actions)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/a631807682/tagcheck)

## Install

> go install github.com/a631807682/tagcheck/cmd/tagcheck

## Usage

- Check [models file](./tagcheck/testdata/src/gormtag) struct tag

  > tc -tests ./...

- Warning message

  ```log
  tagcheck/testdata/src/gormtag/models.go:61:36: not support Gorm option "amount_off"
  tagcheck/testdata/src/gormtag/models.go:62:36: not support Gorm option "percent_off"
  tagcheck/testdata/src/gormtag/models.go:84:18: duplicate Gorm option "column"
  tagcheck/testdata/src/gormtag/models.go:85:18: not support Gorm option "column" value "" can not be empty
  tagcheck/testdata/src/gormtag/models.go:87:18: not support Gorm option "size" value "a" not an uint
  tagcheck/testdata/src/gormtag/models.go:88:18: not support Gorm option "size" value "-1" not an uint
  tagcheck/testdata/src/gormtag/models.go:89:18: not support Gorm option "autoIncrement" value "fals" not empty or bool
  tagcheck/testdata/src/gormtag/models.go:90:18: not support Gorm option "default" value "" can not be empty
  tagcheck/testdata/src/gormtag/models.go:92:19: not support Gorm option "foreignKey" value "" can not be empty
  tagcheck/testdata/src/gormtag/models.go:98:20: not support Gorm option "amount_off"
  ```

## TODO

- [x] Command Line
- [ ] Auto Fix
- [ ] Visual Studio Plugin
