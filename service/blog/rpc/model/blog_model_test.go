package model

import (
	"testing"

	"github.com/go-xorm/builder"
	"github.com/he2121/go-blog/common/sql-helper"
	"github.com/stretchr/testify/assert"
)

func TestStructToConds(t *testing.T) {
	a := int64(1)
	where := WhereBlog{IDs: []int64{1, 2, 3}, FolderID: &a}
	conds, err := sql_helper.WrapWhere(where)
	assert.Nil(t, err)
	sqlStr, args, err := builder.Select("id").From("table").Where(builder.And(conds...)).ToSQL()
	assert.True(t, sqlStr == "select id from table where id in (?,?,?) and folder_id = ?")
	assert.True(t, len(args) == 2)
}
