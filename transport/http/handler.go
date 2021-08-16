package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sapawarga/userpost-service/endpoint"
	"github.com/sapawarga/userpost-service/lib/constant"
	"github.com/sapawarga/userpost-service/lib/convert"
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
	r.Handle("/health/live", kithttp.NewServer(endpoint.MakeCheckHealthy(ctx), decodeNoRequest, encodeResponse, opts...)).Methods(constant.HTTP_GET)
	r.Handle("/health/ready", kithttp.NewServer(endpoint.MakeCheckReadiness(ctx, fs), decodeNoRequest, encodeResponse, opts...)).Methods(constant.HTTP_GET)
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
	r.Handle("/user-posts/", processGetList).Methods(constant.HTTP_GET)
	r.Handle("/user-posts/byme", processGetListByMe).Methods(constant.HTTP_GET)
	r.Handle("/user-posts/{id:[0-9]+}", processGetDetail).Methods(constant.HTTP_GET)
	r.Handle("/user-posts/{id:[0-9]+}/comments", processGetComments).Methods(constant.HTTP_GET)
	r.Handle("/user-posts/", processCreatePost).Methods(constant.HTTP_POST)
	r.Handle("/user-posts/{id:[0-9]+}", processUpdate).Methods(constant.HTTP_PUT)
	r.Handle("/user-posts/{id:[0-9]+}/comments", processCreateComment).Methods(constant.HTTP_POST)
	r.Handle("/user-posts/{id:[0-9]+}/like-dislike", processLikeDislike).Methods(constant.HTTP_PUT)

	return r
}

func decodeGetList(ctx context.Context, r *http.Request) (interface{}, error) {
	statusString := r.URL.Query().Get("status")
	limitString := r.URL.Query().Get("limit")
	pageString := r.URL.Query().Get("page")
	districtID, _ := convert.FromStringToInt64(r.URL.Query().Get("kabkota_id"))
	status, _ := convert.FromStringToInt64(statusString)

	if limitString == "" {
		limitString = "10"
	}
	if pageString == "" || pageString == "0" {
		pageString = "1"
	}

	var sortBy, orderBy = "created_at", "DESC"
	if r.URL.Query().Get("sort_by") != "" {
		sortBy = r.URL.Query().Get("sort_by")
	}
	if r.URL.Query().Get("sort_order") != "" {
		orderBy = constant.AscOrDesc[r.URL.Query().Get("sort_order")]
	}

	limit, _ := convert.FromStringToInt64(limitString)
	page, _ := convert.FromStringToInt64(pageString)

	return &endpoint.GetListUserPostRequest{
		ActivityName: convert.SetPointerString(r.URL.Query().Get("text")),
		Username:     convert.SetPointerString(r.URL.Query().Get("name")),
		Category:     convert.SetPointerString(r.URL.Query().Get("tags")),
		Status:       status,
		Page:         page,
		Limit:        limit,
		SortBy:       sortBy,
		OrderBy:      orderBy,
		Search:       convert.SetPointerString(r.URL.Query().Get("search")),
		DistrictID:   districtID,
	}, nil
}

func decodeGetByID(ctx context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	_, id := convert.FromStringToInt64(params["id"])
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
	_, id := convert.FromStringToInt64(params["id"])
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
		if status.Code == constant.STATUS_CREATED {
			w.WriteHeader(http.StatusCreated)
		} else if status.Code == constant.STATUS_UPDATED || status.Code == constant.STATUS_DELETED {
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
	w.WriteHeader(http.StatusUnprocessableEntity)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
