package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/jmoiron/sqlx/reflectx"
)

func TestSqlxReflectx(t *testing.T) {
	type Timestamp struct {
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}
	type JsonField struct {
		Foo string `json:"foo"`
		Bar string `json:"bar"`
	}
	type Post struct {
		ID       int64     `db:"id"`
		UserID   int64     `db:"user_id"`
		Body     string    `db:"body"`
		PostedAt time.Time `db:"posted_at"`
		Deleted  bool      `db:"deleted"`
		NoName   string    `db:"-"`
		JsonBody JsonField `db:"json_body"`
		Timestamp
	}

	p := Post{}
	dbTags := extractDBTags(reflect.TypeOf(p))
	for _, tag := range dbTags {
		fmt.Println(tag)
	}

	fmt.Println("----")

	m := reflectx.NewMapper("db")
	fields := m.TypeMap(reflect.TypeOf(p))
	for tagname, _ := range fields.Names {
		if !strings.Contains(tagname, ".") {
			fmt.Println(tagname)
		}
	}
}

func extractDBTags(t reflect.Type) []string {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	dbTags := make([]string, 0)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		// Get the `db` tag if any
		dbTag, ok := f.Tag.Lookup("db")
		if ok && dbTag != "-" {
			dbTags = append(dbTags, dbTag)
		}

		// If this is an embedded struct, recurse
		if f.Type.Kind() == reflect.Struct {
			dbTags = append(dbTags, extractDBTags(f.Type)...)
			continue
		}
	}

	return dbTags
}
