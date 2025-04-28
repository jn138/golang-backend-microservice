package mysql

import (
	"errors"
	"fmt"
	"golang-backend-microservice/model"
	"strconv"

	sq "github.com/Masterminds/squirrel"
)

type column struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func BuildQueryColumns(data map[string]interface{}) []column {
	col := make([]column, len(data))
	i := 0
	for key, val := range data {
		col[i] = column{Name: key, Value: fmt.Sprintf("%v", val)}
		i++
	}
	return col
}

func BuildSelectQuery(req *model.MySqlReqArgs) (string, []interface{}, error) {
	query := sq.Select("*").From(req.Table)

	if len(req.Where) > 0 {
		query = query.Where(req.Where)
	}
	if len(req.WhereGreater) > 0 {
		query = query.Where(sq.Gt(req.WhereGreater))
	}
	if len(req.WhereLess) > 0 {
		query = query.Where(sq.Lt(req.WhereLess))
	}
	if len(req.WhereNot) > 0 {
		query = query.Where(sq.NotEq(req.WhereNot))
	}

	if req.Limit != "" {
		limit, err := strconv.ParseUint(req.Limit, 10, 0)
		if err != nil {
			return "", nil, errors.New("unable to convert limit")
		}
		query = query.Limit(limit)
	}

	queryString, args, err := query.ToSql()
	if err != nil {
		return "", nil, err
	}
	return queryString, args, nil
}

func BuildInsertQuery(req *model.MySqlReqArgs) (string, []interface{}, error) {
	query := sq.Insert(req.Table)

	data := BuildQueryColumns(req.Data)
	cols := make([]string, len(data))
	vals := make([]interface{}, len(data))
	for i, d := range data {
		cols[i] = d.Name
		vals[i] = d.Value
	}
	query = query.Columns(cols...)
	query = query.Values(vals...)

	queryString, args, err := query.ToSql()
	if err != nil {
		return "", nil, err
	}
	return queryString, args, nil
}

func BuildUpdateQuery(req *model.MySqlReqArgs) (string, []interface{}, error) {
	query := sq.Update(req.Table)

	if len(req.Where) == 0 && len(req.WhereGreater) == 0 &&
		len(req.WhereLess) == 0 && len(req.WhereNot) == 0 {
		err := errors.New("where cannot be empty")
		return "", nil, err
	}
	if len(req.Where) > 0 {
		query = query.Where(req.Where)
	}
	if len(req.WhereGreater) > 0 {
		query = query.Where(sq.Gt(req.WhereGreater))
	}
	if len(req.WhereLess) > 0 {
		query = query.Where(sq.Lt(req.WhereLess))
	}
	if len(req.WhereNot) > 0 {
		query = query.Where(sq.NotEq(req.WhereNot))
	}

	data := BuildQueryColumns(req.Data)
	for _, d := range data {
		query = query.Set(d.Name, d.Value)
	}

	queryString, args, err := query.ToSql()
	if err != nil {
		return "", nil, err
	}
	return queryString, args, nil
}

func BuildDeleteQuery(req *model.MySqlReqArgs) (string, []interface{}, error) {
	query := sq.Delete(req.Table)

	if len(req.Where) == 0 && len(req.WhereGreater) == 0 &&
		len(req.WhereLess) == 0 && len(req.WhereNot) == 0 {
		err := errors.New("where cannot be empty")
		return "", nil, err
	}
	if len(req.Where) > 0 {
		query = query.Where(req.Where)
	}
	if len(req.WhereGreater) > 0 {
		query = query.Where(sq.Gt(req.WhereGreater))
	}
	if len(req.WhereLess) > 0 {
		query = query.Where(sq.Lt(req.WhereLess))
	}
	if len(req.WhereNot) > 0 {
		query = query.Where(sq.NotEq(req.WhereNot))
	}

	queryString, args, err := query.ToSql()
	if err != nil {
		return "", nil, err
	}
	return queryString, args, nil
}
