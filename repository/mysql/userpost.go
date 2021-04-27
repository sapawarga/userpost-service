package mysql

import (
	"bytes"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sapawarga/userpost-service/helper"
	"github.com/sapawarga/userpost-service/model"
)

type UserPost struct {
	conn *sqlx.DB
}

func NewUserPost(conn *sqlx.DB) *UserPost {
	return &UserPost{
		conn: conn,
	}
}

func (r *UserPost) GetListPost(ctx context.Context, request *model.UserPostRequest) ([]*model.PostResponse, error) {
	var query bytes.Buffer
	var result = make([]*model.PostResponse, 0)
	var err error

	query.WriteString("SELECT id, text, tags. image_path, images, last_user_post_comment_id, likes_count, comment_counts, status, created_by, updated_by, create_at, updated_at")
	query, params := querySelectParams(ctx, query, request)
	if request.Limit != nil && request.Offset != nil {
		query.WriteString("LIMIT ?, ?")
		params = append(params, request.Offset, request.Limit)
	}
	if request.OrderBy != nil && request.SortBy != nil {
		query.WriteString(" ORDER BY ? ?")
		params = append(params, request.OrderBy, request.SortBy)
	}
	if ctx != nil {
		err = r.conn.SelectContext(ctx, result, query.String(), params...)
	} else {
		err = r.conn.Select(result, query.String(), params...)
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *UserPost) GetMetadataPost(ctx context.Context, request *model.UserPostRequest) (*int64, error) {
	var query bytes.Buffer
	var total *int64
	var err error

	query.WriteString("SELECT COUNT(1)")
	query, params := querySelectParams(ctx, query, request)

	if ctx != nil {
		err = r.conn.GetContext(ctx, total, query.String(), params...)
	} else {
		err = r.conn.Get(total, query.String(), params...)
	}

	if err != nil {
		return nil, err
	}

	return total, nil
}

func (r *UserPost) GetListPostByMe(ctx context.Context, request *model.UserPostByMeRequest) ([]*model.PostResponse, error) {
	var query bytes.Buffer
	var result = make([]*model.PostResponse, 0)
	var err error

	query.WriteString("SELECT id, text, tags. image_path, images, last_user_post_comment_id, likes_count, comment_counts, status, created_by, updated_by, create_at, updated_at")
	query, params := querySelectParams(ctx, query, request.UserPostRequest)
	query.WriteString("AND created_by = ? ")
	params = append(params, request.ActorID)
	if request.Limit != nil && request.Offset != nil {
		query.WriteString("LIMIT ?, ?")
		params = append(params, request.Offset, request.Limit)
	}
	if request.OrderBy != nil && request.SortBy != nil {
		query.WriteString(" ORDER BY ? ?")
		params = append(params, request.OrderBy, request.SortBy)
	}

	if ctx != nil {
		err = r.conn.SelectContext(ctx, result, query.String(), params...)
	} else {
		err = r.conn.Select(result, query.String(), params...)
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *UserPost) GetMetadataPostByMe(ctx context.Context, request *model.UserPostByMeRequest) (*int64, error) {
	var query bytes.Buffer
	var total *int64
	var err error

	query.WriteString("SELECT COUNT(1)")
	query, params := querySelectParams(ctx, query, request.UserPostRequest)
	query.WriteString("AND created_by = ? ")
	params = append(params, request.ActorID)

	if ctx != nil {
		err = r.conn.GetContext(ctx, total, query.String(), params...)
	} else {
		err = r.conn.Get(total, query.String(), params...)
	}

	if err != nil {
		return nil, err
	}

	return total, nil
}

// func (r *UserPost) GetActor(ctx context.Context, id int64) (*model.UserResponse, error) {
// 	var query bytes.Buffer
// 	var result *model.UserResponse

// }

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
		first = false
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
