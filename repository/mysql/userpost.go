package mysql

import (
	"bytes"
	"context"
	"database/sql"
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

	query.WriteString("SELECT id, text, tags. image_path, images, last_user_post_comment_id, likes_count, comment_counts, status, created_by, updated_by, create_at, updated_at FROM userposts")
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

	query.WriteString("SELECT COUNT(1) FROM userposts")
	query, params := querySelectParams(ctx, query, request)
	query.WriteString("FROM userposts")
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

	query.WriteString("SELECT id, text, tags. image_path, images, last_user_post_comment_id, likes_count, comment_counts, status, created_by, updated_by, create_at, updated_at FROM userposts")
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

	query.WriteString("SELECT COUNT(1) FROM userposts")
	query, params := querySelectParams(ctx, query, request.UserPostRequest)
	query.WriteString("AND created_by = ? ")
	query.WriteString("FROM userposts")
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

func (r *UserPost) GetActor(ctx context.Context, id int64) (*model.UserResponse, error) {
	var query bytes.Buffer
	var result *model.UserResponse
	var err error

	query.WriteString("SELECT u.id, u.name, u.photo_url, u.`role`, u.rw, reg.name as regency_name, dis.name as district_name, vil.name as village_name")
	query.WriteString("FROM `user` u ")
	query.WriteString(`LEFT JOIN areas reg ON reg.id = u.kabkota_id
	LEFT JOIN areas dis ON dis.id = u.kec_id 
	LEFT JOIN areas vil ON vil.id = u.kel_id`)
	query.WriteString("WHERE u.id = ? ")
	if ctx != nil {
		err = r.conn.GetContext(ctx, result, query.String(), id)
	} else {
		err = r.conn.Get(result, query.String(), id)
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *UserPost) GetIsLikedByUser(ctx context.Context, req *model.IsLikedByUser) (bool, error) {
	var query bytes.Buffer
	var isExists int64
	var err error
	var params []interface{}

	query.WriteString("SELECT COUNT(1) as is_exists FROM likes WHERE WHERE `type` = ? AND user_id = ? AND entity_id = ?")
	params = append(params, req.Type, req.UserID, req.EntityID)
	if ctx != nil {
		err = r.conn.GetContext(ctx, isExists, query.String(), params...)
	} else {
		err = r.conn.Get(isExists, query.String(), params...)
	}

	if err != nil {
		return false, err
	}

	if isExists == 1 {
		return true, nil
	}
	return false, nil
}

func (r *UserPost) GetDetailPost(ctx context.Context, id int64) (*model.PostResponse, error) {
	var query bytes.Buffer
	var result *model.PostResponse
	var err error

	query.WriteString("SELECT id, text, tags. image_path, images, last_user_post_comment_id, likes_count, comment_counts, status, created_by, updated_by, create_at, updated_at FROM userposts")
	query.WriteString("WHERE id = ?")
	if ctx != nil {
		err = r.conn.GetContext(ctx, result, query.String(), id)
	} else {
		err = r.conn.Get(result, query.String(), id)
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *UserPost) InsertPost(ctx context.Context, request *model.CreateNewPostRequestRepository) error {
	var query bytes.Buffer
	var params = make(map[string]interface{})
	var err error
	_, unixTime := helper.GetCurrentTimeUTC()

	query.WriteString("INSERT INTO userposts (`text`, tags, image_path, images, status, created_by, updated_by, created_at, updated_at)")
	query.WriteString("VALUES (:title, :tags, :image_path, :images, :status, :actor, :actor, :created_at, :created_at)")
	params = map[string]interface{}{
		"title":      request.Title,
		"tags":       request.Tags,
		"image_path": request.ImagePathURL,
		"images":     request.Images,
		"status":     request.Status,
		"actor":      request.ActorID,
		"created_at": unixTime,
	}

	if ctx != nil {
		_, err = r.conn.NamedExecContext(ctx, query.String(), params)
	} else {
		_, err = r.conn.NamedExec(query.String(), params)
	}

	if err != nil {
		return err
	}
	return nil
}

func (r *UserPost) UpdateStatusOrTitle(ctx context.Context, request *model.UpdatePostRequest) error {
	var query bytes.Buffer
	var params = make(map[string]interface{})
	var first = true
	var err error
	_, unixTime := helper.GetCurrentTimeUTC()

	query.WriteString("UPDATE userposts")
	if request.Status != nil {
		query.WriteString("status` = :status")
		params["status"] = request.Status
		first = false
	}
	if request.Title != nil {
		if !first {
			query.WriteString(", ")
		}
		query.WriteString("`text` = :title")
		params["title"] = request.Title
		first = false
	}
	if !first {
		query.WriteString(", ")
	}
	query.WriteString(" updated_at = :update")
	params["update"] = unixTime
	query.WriteString(" WHERE id = :id ")
	params["id"] = request.ID

	if ctx != nil {
		_, err = r.conn.NamedExecContext(ctx, query.String(), params)
	} else {
		_, err = r.conn.NamedExec(query.String(), params)
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
