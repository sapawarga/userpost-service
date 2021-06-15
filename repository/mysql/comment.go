package mysql

import (
	"bytes"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/sapawarga/userpost-service/helper"
	"github.com/sapawarga/userpost-service/model"
)

type Comment struct {
	conn *sqlx.DB
}

func NewComment(conn *sqlx.DB) *Comment {
	return &Comment{
		conn: conn,
	}
}

func (r *Comment) GetLastComment(ctx context.Context, id int64) (*model.CommentResponse, error) {
	var query bytes.Buffer
	var result = &model.CommentResponse{}
	var err error

	query.WriteString("SELECT id, user_post_id, `text` as comment, status, created_by, updated_by, FROM_UNIXTIME(created_at) as created_at, FROM_UNIXTIME(updated_at) as updated_at FROM user_post_comments ")
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

func (r *Comment) GetTotalComments(ctx context.Context, userPostID int64) (*int64, error) {
	var query bytes.Buffer
	var total *int64
	var err error

	query.WriteString("SELECT COUNT(1) as total FROM user_post_comments WHERE user_post_id = ?")

	if ctx != nil {
		err = r.conn.GetContext(ctx, total, query.String(), userPostID)
	} else {
		err = r.conn.Get(total, query.String(), userPostID)
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return total, nil
}

func (r *Comment) GetCommentsByPostID(ctx context.Context, id int64) ([]*model.CommentResponse, error) {
	var query bytes.Buffer
	var result = make([]*model.CommentResponse, 0)
	var err error

	query.WriteString("SELECT id, user_post_id, `text` as comment, status, created_by, updated_by, FROM_UNIXTIME(created_at) as created_at, FROM_UNIXTIME(updated_at) as updated_at FROM user_post_comments ")
	query.WriteString("WHERE user_post_id = ?")

	if ctx != nil {
		err = r.conn.SelectContext(ctx, &result, query.String(), id)
	} else {
		err = r.conn.Select(&result, query.String(), id)
	}

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Comment) Create(ctx context.Context, req *model.CreateCommentRequestRepository) (int64, error) {
	var query bytes.Buffer
	var err error
	var result sql.Result
	_, unixTime := helper.GetCurrentTimeUTC()

	query.WriteString("INSERT INTO user_post_comments (user_post_id, `text`, status, created_by, updated_by, created_at, updated_at) ")
	query.WriteString("VALUES(:user_post_id, :comment, :status, :actor, :actor, :current, :current)")
	params := map[string]interface{}{
		"user_post_id": req.UserPostID,
		"comment":      req.Text,
		"actor":        req.ActorID,
		"current":      unixTime,
		"status":       req.Status,
	}

	if ctx != nil {
		result, err = r.conn.NamedExecContext(ctx, query.String(), params)
	} else {
		result, err = r.conn.NamedExec(query.String(), params)
	}

	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
