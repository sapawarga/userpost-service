package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sapawarga/userpost-service/endpoint"
	"github.com/sapawarga/userpost-service/helper"
	"github.com/sapawarga/userpost-service/usecase"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type err interface {
	error() error
}

func MakeHandlerHealthy(ctx context.Context, fs usecase.UsecaseI, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	r := mux.NewRouter()
	r.Handle("/health/live", kithttp.NewServer(endpoint.MakeCheckHealthy(ctx), decodeNoRequest, encodeResponse, opts...)).Methods(helper.HTTP_GET)
	r.Handle("/health/ready", kithttp.NewServer(endpoint.MakeCheckReadiness(ctx, fs), decodeNoRequest, encodeResponse, opts...)).Methods(helper.HTTP_GET)
	return r
}

func MakeHTTPHandler(ctx context.Context, fs usecase.UsecaseI, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	processGetList := kithttp.NewServer(endpoint.MakeGetListUserPost(ctx, fs), decodeGetList, encodeResponse, opts...)
	processGetListByMe := kithttp.NewServer(endpoint.MakeGetListUserPostByMe(ctx, fs), decodeGetList, encodeResponse, opts...)
	processGetDetail := kithttp.NewServer(endpoint.MakeGetDetailUserPost(ctx, fs), decodeGetByID, encodeResponse, opts...)
	processGetComments := kithttp.NewServer(endpoint.MakeGetCommentsByID(ctx, fs), decodeGetByID, encodeResponse, opts...)
	processCreatePost := kithttp.NewServer(endpoint.MakeCreateNewPost(ctx, fs), decodeCreatePost, encodeResponse, opts...)
	processCreateComment := kithttp.NewServer(endpoint.MakeCreateComment(ctx, fs), decodeCreateComment, encodeResponse, opts...)
	processLikeDislike := kithttp.NewServer(endpoint.MakeLikeOrDislikePost(ctx, fs), decodeGetByID, encodeResponse, opts...)
	processUpdate := kithttp.NewServer(endpoint.MakeUpdateStatusOrTitle(ctx, fs), decodeCreateComment, encodeResponse, opts...)

	r := mux.NewRouter()

	// TODO: handle token middleware
	r.Handle("/userpost/", processGetList).Methods(helper.HTTP_GET)
	r.Handle("/userpost/byme", processGetListByMe).Methods(helper.HTTP_GET)
	r.Handle("/userpost/{id}", processGetDetail).Methods(helper.HTTP_GET)
	r.Handle("/userpost/{id}/comments", processGetComments).Methods(helper.HTTP_GET)
	r.Handle("/userpost/", processCreatePost).Methods(helper.HTTP_POST)
	r.Handle("/userpost/{id}/comments", processCreateComment).Methods(helper.HTTP_POST)
	r.Handle("/userpost/{id}/like-dislike", processLikeDislike).Methods(helper.HTTP_PUT)
	r.Handle("/userpost/{id}", processUpdate).Methods(helper.HTTP_PUT)
	return r
}

func decodeGetList(ctx context.Context, r *http.Request) (interface{}, error) {
	statusString := r.URL.Query().Get("status")
	limitString := r.URL.Query().Get("limit")
	pageString := r.URL.Query().Get("page")

	if limitString == "" {
		limitString = "10"
	}
	if pageString == "" || pageString == "0" {
		pageString = "1"
	}

	status, _ := helper.ConvertFromStringToInt64(statusString)
	limit, _ := helper.ConvertFromStringToInt64(limitString)
	page, _ := helper.ConvertFromStringToInt64(pageString)

	return &endpoint.GetListUserPostRequest{
		ActivityName: helper.SetPointerString(r.URL.Query().Get("activity_name")),
		Username:     helper.SetPointerString(r.URL.Query().Get("username")),
		Category:     helper.SetPointerString(r.URL.Query().Get("category")),
		Status:       status,
		Page:         page,
		Limit:        limit,
		SortBy:       helper.SetPointerString(r.URL.Query().Get("sort_by")),
		OrderBy:      helper.SetPointerString(r.URL.Query().Get("order_by")),
	}, nil
}

func decodeGetByID(ctx context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	_, id := helper.ConvertFromStringToInt64(params["id"])
	return &endpoint.GetByID{
		ID: id,
	}, nil
}

func decodeCreatePost(ctx context.Context, r *http.Request) (interface{}, error) {
	reqBody := &endpoint.CreateNewPostRequest{}
	if err := json.NewDecoder(r.Body).Decode(reqBody); err != nil {
		return nil, err
	}
	return reqBody, nil
}

func decodeCreateComment(ctx context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	_, id := helper.ConvertFromStringToInt64(params["id"])
	reqBody := &endpoint.CreateCommentRequest{}
	if err := json.NewDecoder(r.Body).Decode(reqBody); err != nil {
		return nil, err
	}
	reqBody.UserPostID = id
	return reqBody, nil
}

func decodeNoRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(err); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	status, ok := response.(*endpoint.StatusResponse)
	if ok {
		if status.Code == helper.STATUS_CREATED {
			w.WriteHeader(http.StatusCreated)
		} else if status.Code == helper.STATUS_UPDATED || status.Code == helper.STATUS_DELETED {
			w.WriteHeader(http.StatusNoContent)
			_ = json.NewEncoder(w).Encode(nil)
			return nil
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}

	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
