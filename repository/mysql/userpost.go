package mysql

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sapawarga/userpost-service/lib/convert"
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

	query.WriteString(`SELECT up.id, up.text, up.tags, up.image_path, up.images, up.last_user_post_comment_id, up.likes_count, up.comments_count, up.status, 
		up.created_by, up.updated_by, up.created_at, up.updated_at FROM user_posts up
		LEFT JOIN user u ON u.id = up.created_by`)
	query, params := querySelectParams(ctx, query, request)
	if request.Limit != nil && request.Offset != nil {
		query.WriteString(" LIMIT ?, ?")
		params = append(params, request.Offset, request.Limit)
	}
	if request.OrderBy != nil && request.SortBy != nil {
		query.WriteString(fmt.Sprintf(" ORDER BY %s %s", *request.SortBy, *request.OrderBy))
	}

	if ctx != nil {
		err = r.conn.SelectContext(ctx, &result, query.String(), params...)
	} else {
		err = r.conn.Select(&result, query.String(), params...)
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

	query.WriteString("SELECT COUNT(1) FROM user_posts up LEFT JOIN user u ON u.id = up.created_by")
	query, params := querySelectParams(ctx, query, request)
	if ctx != nil {
		err = r.conn.GetContext(ctx, &total, query.String(), params...)
	} else {
		err = r.conn.Get(&total, query.String(), params...)
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

	query.WriteString(`SELECT id, text, tags, image_path, images, last_user_post_comment_id, likes_count, comments_count, status, 
	created_by, updated_by, created_at, updated_at FROM user_posts WHERE created_by = ?`)
	query, params := querySelectParams(ctx, query, request.UserPostRequest)
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
		err = r.conn.SelectContext(ctx, &result, query.String(), params...)
	} else {
		err = r.conn.Select(&result, query.String(), params...)
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

	query.WriteString("SELECT COUNT(1) FROM user_posts ")
	query.WriteString(" WHERE created_by = ? ")
	query, params := querySelectParams(ctx, query, request.UserPostRequest)
	params = append(params, request.ActorID)

	if ctx != nil {
		err = r.conn.GetContext(ctx, &total, query.String(), params...)
	} else {
		err = r.conn.Get(&total, query.String(), params...)
	}

	if err != nil {
		return nil, err
	}

	return total, nil
}

func (r *UserPost) GetActor(ctx context.Context, id int64) (*model.UserResponse, error) {
	var query bytes.Buffer
	var result = &model.UserResponse{}
	var err error

	query.WriteString("SELECT u.id, u.name, u.photo_url, u.`role`, u.rw, reg.name as regency_name, dis.name as district_name, vil.name as village_name")
	query.WriteString(` FROM user u `)
	query.WriteString(`LEFT JOIN areas reg ON reg.id = u.kabkota_id
	LEFT JOIN areas dis ON dis.id = u.kec_id 
	LEFT JOIN areas vil ON vil.id = u.kel_id`)
	query.WriteString(" WHERE u.id = ? ")
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

func (r *UserPost) GetDetailPost(ctx context.Context, id int64) (*model.PostResponse, error) {
	var query bytes.Buffer
	var result = &model.PostResponse{}
	var err error

	query.WriteString("SELECT id, text, tags, image_path, images, last_user_post_comment_id, likes_count, comments_count, status, created_by, updated_by, created_at, updated_at FROM user_posts")
	query.WriteString(" WHERE id = ?")
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

func (r *UserPost) CheckIsExistLikeOnPostBy(ctx context.Context, request *model.AddOrRemoveLikeOnPostRequest) (bool, error) {
	var query bytes.Buffer
	var total *int64
	var err error
	var params []interface{}

	query.WriteString("SELECT 1	FROM likes WHERE `type` = ? AND user_id  = ? AND entity_id = ?")
	params = append(params, request.TypeEntity, request.ActorID, request.UserPostID)

	if ctx != nil {
		err = r.conn.GetContext(ctx, total, query.String(), params...)
	} else {
		err = r.conn.Get(total, query.String(), params...)
	}

	if total == nil || err == sql.ErrNoRows {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *UserPost) InsertPost(ctx context.Context, request *model.CreateNewPostRequestRepository) error {
	var query bytes.Buffer
	var err error
	_, unixTime := convert.GetCurrentTimeUTC()

	query.WriteString("INSERT INTO user_posts (`text`, tags, image_path, images, status, created_by, updated_by, created_at, updated_at)")
	query.WriteString("VALUES (:title, :tags, :image_path, :images, :status, :actor, :actor, :created_at, :created_at)")
	params := map[string]interface{}{
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

func (r *UserPost) AddLikeOnPost(ctx context.Context, request *model.AddOrRemoveLikeOnPostRequest) error {
	var query bytes.Buffer
	var err error
	_, unixTime := convert.GetCurrentTimeUTC()

	query.WriteString("INSERT INTO likes (`type`, user_id, entity_id, created_at, updated_at) ")
	query.WriteString("VALUES(:type_entity, :user_id, :entity_id, :current, :current)")
	params := map[string]interface{}{
		"type_entity": request.TypeEntity,
		"user_id":     request.ActorID,
		"entity_id":   request.UserPostID,
		"current":     unixTime,
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

func (r *UserPost) UpdateDetailOfUserPost(ctx context.Context, request *model.UpdatePostRequest) error {
	var query bytes.Buffer
	var params = make(map[string]interface{})
	var fields []string
	var err error
	_, unixTime := convert.GetCurrentTimeUTC()

	query.WriteString("UPDATE user_posts SET ")
	if request.Status != nil {
		fields = append(fields, "status")
		params["status"] = request.Status
	}
	if request.Title != nil {
		fields = append(fields, "text")
		params["text"] = request.Title
	}
	if request.LastCommentID != nil {
		fields = append(fields, "last_user_post_comment_id")
		query.WriteString(" comments_count = comments_count + 1 , ")
		params["last_user_post_comment_id"] = request.LastCommentID
	}
	fields = append(fields, "updated_at")
	params["updated_at"] = unixTime
	query.WriteString(updateQuery(ctx, fields...) + " WHERE id = :id ")
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

func (r *UserPost) RemoveLikeOnPost(ctx context.Context, request *model.AddOrRemoveLikeOnPostRequest) error {
	var query bytes.Buffer
	var params []interface{}
	var err error

	query.WriteString("DELETE FROM likes WHERE `type` = ? AND user_id  = ? AND entity_id = ? ")
	params = append(params, request.TypeEntity, request.ActorID, request.UserPostID)

	if ctx != nil {
		_, err = r.conn.ExecContext(ctx, query.String(), params...)
	} else {
		_, err = r.conn.Exec(query.String(), params...)
	}

	if err != nil {
		return err
	}
	return nil
}
