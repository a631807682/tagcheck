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
var gormTagMap = map[string]GTag{}

func initGormTagChecker() {
	anyC := &anyChecker{}
	uintC := uintChecker{}
	// boolC := boolChecker{}

	// https://gorm.io/docs/models.html
	matchTags := []GTag{
		{
			key:      "column",
			category: Table,
			checker:  anyC,
		},
		{
			key:      "type",
			category: Table,
			checker:  anyC,
		},
		{
			key:      "size",
			category: Table,
			checker:  uintC,
		},
		{
			key:      "primaryKey",
			category: Table,
			checker:  anyC,
		},
		{
			key:      "primary_key",
			category: Table,
			checker:  anyC,
		},
		{
			key:      "unique",
			category: Table,
			checker:  anyC,
		},
		{
			key:      "default",
			category: Table,
			checker:  anyC,
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
			checker:  anyC,
		},
		{
			key:      "autoIncrement",
			category: Table,
			checker:  anyC,
		},
		{
			key:      "autoIncrementIncrement",
			category: Table,
			checker:  anyC,
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
			checker:  anyC,
		},
		{
			key:      "serializer",
			category: Internal,
			checker:  anyC,
		},
		{
			key:      "embedded",
			category: Internal,
			checker:  anyC,
		},
		{
			key:      "embeddedPrefix",
			category: Internal,
			checker:  anyC,
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
			checker:  anyC,
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
			checker:  anyC,
		},
		{
			key:      "references",
			category: Relation,
			checker:  anyC,
		},
		{
			key:      "polymorphic",
			category: Relation,
			checker:  anyC,
		},
		{
			key:      "many2many",
			category: Relation,
			checker:  anyC,
		},
		{
			key:      "belongsto",
			category: Relation,
			checker:  anyC,
		},
		{
			key:      "polymorphicValue",
			category: Relation,
			checker:  anyC,
		},
		{
			key:      "joinForeignKey",
			category: Relation,
			checker:  anyC,
		},
		{
			key:      "joinReferences",
			category: Relation,
			checker:  anyC,
		},
		{
			key:      "constraint",
			category: Relation,
			checker:  anyC,
		},
	}

	for _, gt := range matchTags {
		gormTagMap[strings.ToUpper(gt.key)] = gt
	}
}
