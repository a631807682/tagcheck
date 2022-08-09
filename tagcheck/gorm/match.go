package gorm

import "strings"

type category int8

const (
	// table column definition
	Table category = iota + 1
	// column permission
	Permission
	// gorm internal implementation
	Internal
	// schema relation
	Relation
)

type GTag struct {
	key      string
	category category
	checker  checker
}

// gorm tag full support list, key upper.
var gormTagMap = loadGormTagChecker()

func loadGormTagChecker() (gtm map[string]GTag) {
	anyC := &anyChecker{}
	uintC := &uintChecker{}
	intC := &intChecker{}
	notEmptyC := &notEmptyChecker{}
	emptyOrBoolC := &emptyOrBoolChecker{}

	// https://gorm.io/docs/models.html
	matchTags := []GTag{
		{
			key:      "column",
			category: Table,
			checker:  notEmptyC,
		},
		{
			key:      "type",
			category: Table,
			checker:  notEmptyC,
		},
		{
			key:      "size",
			category: Table,
			checker:  uintC,
		},
		{
			key:      "primaryKey",
			category: Table,
			checker:  emptyOrBoolC,
		},
		{
			key:      "primary_key",
			category: Table,
			checker:  emptyOrBoolC,
		},
		{
			key:      "unique",
			category: Table,
			checker:  emptyOrBoolC,
		},
		{
			key:      "default",
			category: Table,
			checker:  notEmptyC,
		},
		{
			key:      "precision",
			category: Table,
			checker:  uintC,
		},
		{
			key:      "scale",
			category: Table,
			checker:  uintC,
		},
		{
			key:      "not null",
			category: Table,
			checker:  emptyOrBoolC,
		},
		{
			key:      "notnull",
			category: Table,
			checker:  emptyOrBoolC,
		},
		{
			key:      "autoIncrement",
			category: Table,
			checker:  emptyOrBoolC,
		},
		{
			key:      "autoIncrementIncrement",
			category: Table,
			checker:  intC,
		},
		{
			key:      "index",
			category: Table,
			checker:  anyC,
		},
		{
			key:      "uniqueIndex",
			category: Table,
			checker:  anyC,
		},
		{
			key:      "comment",
			category: Table,
			checker:  notEmptyC,
		},
		{
			key:      "serializer",
			category: Internal,
			checker:  notEmptyC,
		},
		{
			key:      "embedded",
			category: Internal,
			checker:  anyC,
		},
		{
			key:      "embeddedPrefix",
			category: Internal,
			checker:  notEmptyC,
		},
		{
			key:      "autoCreateTime",
			category: Internal,
			checker:  anyC,
		},
		{
			key:      "autoUpdateTime",
			category: Internal,
			checker:  anyC,
		},
		{
			key:      "check",
			category: Internal,
			checker:  notEmptyC,
		},
		{
			key:      "<-",
			category: Permission,
			checker:  anyC,
		},
		{
			key:      "->",
			category: Permission,
			checker:  anyC,
		},
		{
			key:      "-",
			category: Permission,
			checker:  anyC,
		},
		{
			key:      "foreignKey",
			category: Relation,
			checker:  notEmptyC,
		},
		{
			key:      "references",
			category: Relation,
			checker:  notEmptyC,
		},
		{
			key:      "polymorphic",
			category: Relation,
			checker:  notEmptyC,
		},
		{
			key:      "many2many",
			category: Relation,
			checker:  notEmptyC,
		},
		{
			key:      "belongsto",
			category: Relation,
			checker:  notEmptyC,
		},
		{
			key:      "polymorphicValue",
			category: Relation,
			checker:  notEmptyC,
		},
		{
			key:      "joinForeignKey",
			category: Relation,
			checker:  notEmptyC,
		},
		{
			key:      "joinReferences",
			category: Relation,
			checker:  notEmptyC,
		},
		{
			key:      "constraint",
			category: Relation,
			checker:  notEmptyC,
		},
	}

	gtm = make(map[string]GTag, len(matchTags))
	for _, gt := range matchTags {
		gtm[strings.ToUpper(gt.key)] = gt
	}
	return
}
