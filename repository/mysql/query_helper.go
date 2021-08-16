package mysql

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/sapawarga/userpost-service/lib/convert"
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
	if params.Search != nil {
		query.WriteString(" WHERE (user.name LIKE ? OR up.text LIKE ? )")
		queryParams = append(queryParams, "%"+convert.GetStringFromPointer(params.Search)+"%", "%"+convert.GetStringFromPointer(params.Search)+"%")
	}

	if params.ActivityName != nil {
		query.WriteString(andWhere(ctx, query, "up.text", "LIKE"))
		queryParams = append(queryParams, "%"+convert.GetStringFromPointer(params.ActivityName)+"%")
	}

	if params.Category != nil {
		query.WriteString(andWhere(ctx, query, "LOWER(up.tags)", "LIKE"))
		queryParams = append(queryParams, "%"+strings.ToLower(convert.GetStringFromPointer(params.Category))+"%")
	}

	if params.Status != nil {
		query.WriteString(andWhere(ctx, query, "up.status", "="))
		queryParams = append(queryParams, convert.GetInt64FromPointer(params.Status))
	}

	if params.Username != nil {
		query.WriteString(andWhere(ctx, query, "user.name", "LIKE"))
		queryParams = append(queryParams, "%"+convert.GetStringFromPointer(params.Username)+"%")
	}

	if params.RegencyID != nil {
		query.WriteString(andWhere(ctx, query, "user.kabkota_id", "="))
		queryParams = append(queryParams, convert.GetInt64FromPointer(params.RegencyID))
	}

	return query, queryParams
}

func andWhere(ctx context.Context, query bytes.Buffer, field string, action string) string {
	var newQuery string
	if strings.Contains(query.String(), "WHERE") {
		newQuery = fmt.Sprintf(" AND %s %s ? ", field, action)
	} else {
		newQuery = fmt.Sprintf(" WHERE %s %s ? ", field, action)
	}
	return newQuery
}

func updateQuery(ctx context.Context, fields ...string) string {
	var query bytes.Buffer
	query.WriteString(fmt.Sprintf(" %s = :%s ", fields[0], fields[0]))
	fields = fields[1:]
	for _, field := range fields {
		query.WriteString(fmt.Sprintf(" , %s = :%s ", field, field))
	}
	return query.String()
}
