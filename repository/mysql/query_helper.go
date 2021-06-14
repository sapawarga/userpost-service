package mysql

import (
	"bytes"
	"context"
	"fmt"

	"github.com/sapawarga/userpost-service/helper"
	"github.com/sapawarga/userpost-service/model"
)

func (r *UserPost) CheckHealthReadiness(ctx context.Context) error {
	var err error
	if ctx != nil {
		err = r.conn.PingContext(ctx)
	} else {
		err = r.conn.Ping()
	}

	if err != nil {
		return err
	}

	return nil
}

func querySelectParams(ctx context.Context, query bytes.Buffer, params *model.UserPostRequest) (newQuery bytes.Buffer, queryParams []interface{}) {
	var first = true
	if params.ActivityName != nil {
		qBuffer := isFirstQuery(ctx, first, helper.SELECT_QUERY)
		query.WriteString(qBuffer.String())
		query.WriteString(fmt.Sprintf(`text LIKE LOWER(%s) `, "'%"+helper.GetStringFromPointer(params.ActivityName)+"%'"))
		first = false
	}

	if params.Category != nil {
		qBuffer := isFirstQuery(ctx, first, helper.SELECT_QUERY)
		query.WriteString(qBuffer.String())
		query.WriteString(" tags = ? ")
		queryParams = append(queryParams, params.Category)
		first = false
	}

	if params.Status != nil {
		qBuffer := isFirstQuery(ctx, first, helper.SELECT_QUERY)
		query.WriteString(qBuffer.String())
		query.WriteString(" status = ?")
		queryParams = append(queryParams, params.Status)
	}

	return query, queryParams
}

func isFirstQuery(ctx context.Context, isFirst bool, queryType string) bytes.Buffer {
	var query bytes.Buffer
	if queryType == helper.SELECT_QUERY {
		if isFirst {
			query.WriteString(" WHERE ")
		} else {
			query.WriteString(" AND ")
		}
	} else if queryType == helper.UPDATE_QUERY {
		if !isFirst {
			query.WriteString(" , ")
		}
	}

	return query
}
