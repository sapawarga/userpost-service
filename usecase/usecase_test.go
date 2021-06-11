package usecase_test

import (
	"context"
	"fmt"
	"os"
	"reflect"

	"github.com/sapawarga/userpost-service/helper"
	mocks "github.com/sapawarga/userpost-service/mocks"
	"github.com/sapawarga/userpost-service/mocks/testcases"
	"github.com/sapawarga/userpost-service/model"
	"github.com/sapawarga/userpost-service/usecase"

	kitlog "github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Usecase", func() {
	var (
		mockPostRepo     *mocks.MockPostI
		mockCommentsRepo *mocks.MockCommentI
		userPost         usecase.UsecaseI
	)

	BeforeEach(func() {
		logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
		mockSvc := gomock.NewController(GinkgoT())
		mockSvc.Finish()
		mockPostRepo = mocks.NewMockPostI(mockSvc)
		mockCommentsRepo = mocks.NewMockCommentI(mockSvc)
		userPost = usecase.NewPost(mockPostRepo, mockCommentsRepo, logger)
	})

	// DECLARE UNIT TEST FUNCTION

	var GetListUserPostLogic = func(idx int) {
		ctx := context.Background()
		actor := map[string]interface{}{
			"id": int64(1),
		}
		actorData := &model.ActorFromContext{Data: actor}
		ctx = context.WithValue(ctx, helper.ACTORKEY, actorData)
		data := testcases.GetListUserPostData[idx]
		mockPostRepo.EXPECT().GetListPost(ctx, gomock.Any()).Return(data.ResponseGetList.Result, data.ResponseGetList.Error).Times(1)
		mockPostRepo.EXPECT().GetMetadataPost(ctx, gomock.Any()).Return(data.ResponseMetadata.Result, data.ResponseMetadata.Error).Times(1)
		mockCommentsRepo.EXPECT().GetLastComment(ctx, data.GetLastCommentParams).Return(data.ResponseGetLastComment.Result, data.ResponseGetLastComment.Error).Times(len(data.ResponseGetList.Result))
		mockPostRepo.EXPECT().GetActor(ctx, data.GetActorParams).Return(data.ResponseGetActor.Result, data.ResponseGetActor.Error).Times(len(data.ResponseGetList.Result) * 3)
		mockPostRepo.EXPECT().CheckIsExistLikeOnPostBy(ctx, gomock.Any()).Return(data.CheckIsLikedResponse.Result, data.CheckIsLikedResponse.Error).Times(len(data.ResponseGetList.Result))
		resp, err := userPost.GetListPost(ctx, &data.UsecaseParams)
		if err != nil {
			Expect(err).NotTo(BeNil())
			Expect(resp).To(BeNil())
		} else {
			Expect(err).To(BeNil())
			Expect(resp.Metadata.Page).To(Equal(data.ResponseUsecase.Result.Metadata.Page))
			Expect(resp.Metadata.Total).To(Equal(data.ResponseUsecase.Result.Metadata.Total))
			Expect(resp).NotTo(BeNil())
		}
	}

	var GetDetailUserPostLogic = func(idx int) {
		ctx := context.TODO()
		actor := map[string]interface{}{
			"id": int64(1),
		}
		actorData := &model.ActorFromContext{Data: actor}
		ctx = context.WithValue(ctx, helper.ACTORKEY, actorData)
		data := testcases.GetDetailUserPostData[idx]
		mockPostRepo.EXPECT().GetDetailPost(ctx, data.GetUserPostParams).Return(data.ResponseGetDetailUserPost.Result, data.ResponseGetDetailUserPost.Error).Times(1)
		mockCommentsRepo.EXPECT().GetLastComment(ctx, data.GetLastCommentParams).Return(data.ResponseGetLastComment.Result, data.ResponseGetLastComment.Error).Times(1)
		mockCommentsRepo.EXPECT().GetTotalComments(ctx, data.GetTotalCommentsParams).Return(data.ResponseGetTotalComment.Result, data.ResponseGetTotalComment.Error).Times(1)
		mockPostRepo.EXPECT().GetActor(ctx, data.GetActorParams).Return(data.ResponseGetActor.Result, data.ResponseGetActor.Error).Times(3)
		mockPostRepo.EXPECT().CheckIsExistLikeOnPostBy(ctx, data.IsLikedRequest).Return(data.CheckIsLikedResponse.Result, data.CheckIsLikedResponse.Error).Times(1)
		resp, err := userPost.GetDetailPost(ctx, data.UsecaseParams)
		if err != nil {
			Expect(err).NotTo(BeNil())
			Expect(resp).To(BeNil())
		} else {
			Expect(err).To(BeNil())
			Expect(resp).NotTo(BeNil())
		}
	}

	var CreateNewPostLogic = func(idx int) {
		ctx := context.TODO()
		actor := map[string]interface{}{
			"id": int64(1),
		}
		actorData := &model.ActorFromContext{Data: actor}
		ctx = context.WithValue(ctx, helper.ACTORKEY, actorData)
		data := testcases.CreateNewUserPostData[idx]
		mockPostRepo.EXPECT().InsertPost(ctx, data.RepositoryRequest).Return(data.MockRepository).Times(1)
		err := userPost.CreateNewPost(ctx, data.UsecaseRequest)
		if err != nil {
			Expect(err).NotTo(BeNil())
		} else {
			Expect(err).To(BeNil())
		}
	}

	var UpdateUserPostLogic = func(idx int) {
		ctx := context.TODO()
		data := testcases.UpdateUserPostDetailData[idx]
		mockPostRepo.EXPECT().GetDetailPost(ctx, data.GetDetailParam).Return(nil, data.ResponseGetDetailRepo.Error).Times(1)
		mockPostRepo.EXPECT().UpdateStatusOrTitle(ctx, data.UpdateUserPostParam).Return(data.ResponseUpdateUserPostRepo).Times(1)
		if err := userPost.UpdateTitleOrStatus(ctx, data.UsecaseRequest); err != nil {
			Expect(err).NotTo(BeNil())
		} else {
			Expect(err).To(BeNil())
		}
	}

	var GetCommentsLogic = func(idx int) {
		ctx := context.TODO()
		data := testcases.GetCommentsData[idx]
		mockCommentsRepo.EXPECT().GetCommentsByPostID(ctx, data.GetCommentsByIDRequestRepository).Return(data.ResponseGetComments.Result, data.ResponseGetComments.Error).Times(1)
		mockPostRepo.EXPECT().GetActor(ctx, data.GetActorParams).Return(data.ResponseGetActor.Result, data.ResponseGetActor.Error).Times(len(testcases.GetCommentsData) * 3)
		resp, err := userPost.GetCommentsByPostID(ctx, data.UsecaseParams)
		if err != nil {
			Expect(err).NotTo(BeNil())
			Expect(resp).To(BeNil())
		} else {
			fmt.Println(resp)
			Expect(resp).To(Equal(data.ResponseUsecase.Result))
		}
	}

	var CreateCommentLogic = func(idx int) {
		ctx := context.TODO()
		actor := map[string]interface{}{
			"id": int64(1),
		}
		actorData := &model.ActorFromContext{Data: actor}
		ctx = context.WithValue(ctx, helper.ACTORKEY, actorData)
		data := testcases.CreateCommentOnAPostData[idx]
		mockCommentsRepo.EXPECT().Create(ctx, data.RepositoryRequest).Return(data.MockRepository).Times(1)
		if err := userPost.CreateCommentOnPost(ctx, data.UsecaseRequest); err != nil {
			Expect(err).ToNot(BeNil())
		} else {
			Expect(err).To(BeNil())
		}
	}

	var GetListPostByMeLogic = func(idx int) {
		ctx := context.Background()
		actor := &model.ActorFromContext{Data: map[string]interface{}{
			"id": int64(1),
		}}
		ctx = context.WithValue(ctx, helper.ACTORKEY, actor)
		data := testcases.GetListUserPostByMeData[idx]
		mockPostRepo.EXPECT().GetListPostByMe(ctx, gomock.Any()).Return(data.ResponseGetList.Result, data.ResponseGetList.Error).Times(1)
		mockPostRepo.EXPECT().GetMetadataPostByMe(ctx, gomock.Any()).Return(data.ResponseMetadata.Result, data.ResponseMetadata.Error).Times(1)
		mockCommentsRepo.EXPECT().GetLastComment(ctx, data.GetLastCommentParams).Return(data.ResponseGetLastComment.Result, data.ResponseGetLastComment.Error).Times(len(data.ResponseGetList.Result))
		mockCommentsRepo.EXPECT().GetTotalComments(ctx, data.GetTotalCommentsParams).Return(data.ResponseGetTotalComment.Result, data.ResponseGetTotalComment.Error).Times(len(data.ResponseGetList.Result))
		mockPostRepo.EXPECT().GetActor(ctx, data.GetActorParams).Return(data.ResponseGetActor.Result, data.ResponseGetActor.Error).Times(len(data.ResponseGetList.Result) * 3)
		mockPostRepo.EXPECT().CheckIsExistLikeOnPostBy(ctx, gomock.Any()).Return(data.CheckIsLikedResponse.Result, data.CheckIsLikedResponse.Error).Times(len(data.ResponseGetList.Result))
		resp, err := userPost.GetListPostByMe(ctx, data.UsecaseParams)
		if err != nil {
			Expect(err).NotTo(BeNil())
			Expect(resp).To(BeNil())
		} else {
			Expect(err).To(BeNil())
			Expect(resp.Metadata.Page).To(Equal(data.ResponseUsecase.Result.Metadata.Page))
			Expect(resp.Metadata.Total).To(Equal(data.ResponseUsecase.Result.Metadata.Total))
			Expect(resp).NotTo(BeNil())
		}
	}

	var LikeOrDislikePostLogic = func(idx int) {
		ctx := context.Background()
		actor := &model.ActorFromContext{Data: map[string]interface{}{
			"id": int64(1),
		}}
		ctx = context.WithValue(ctx, helper.ACTORKEY, actor)
		data := testcases.LikeOrDislikePostData[idx]
		mockPostRepo.EXPECT().CheckIsExistLikeOnPostBy(ctx, data.CheckIsLikedRequest).Return(data.MockCheckIsLiked.Result, data.MockCheckIsLiked.Error).Times(1)
		mockPostRepo.EXPECT().AddLikeOnPost(ctx, data.AddLikeOnPostRequest).Return(data.MockAddLikeOnPost).Times(1)
		mockPostRepo.EXPECT().RemoveLikeOnPost(ctx, data.RemoveLikeOnPostRequest).Return(data.MockRemoveLikeOnPost).Times(1)
		err := userPost.LikeOrDislikePost(ctx, data.UsecaseRequest)
		if err != nil {
			Expect(err).To(Equal(data.MockUsecase))
		} else {
			Expect(err).To(BeNil())
		}
	}

	var CheckReadinessLogic = func(idx int) {
		ctx := context.Background()
		data := testcases.CheckReadinessData[idx]
		mockPostRepo.EXPECT().CheckHealthReadiness(ctx).Return(data.MockCheckReadiness).Times(1)
		if err := userPost.CheckHealthReadiness(ctx); err != nil {
			Expect(err).NotTo(BeNil())
		} else {
			Expect(err).To(BeNil())
		}
	}

	var unitTestLogic = map[string]map[string]interface{}{
		"GetListUserPost":     {"func": GetListUserPostLogic, "test_case_count": len(testcases.GetListUserPostData), "desc": testcases.ListUserPostDescription()},
		"GetDetailUserPost":   {"func": GetDetailUserPostLogic, "test_case_count": len(testcases.GetDetailUserPostData), "desc": testcases.ListUserPostDetailDescription()},
		"CreateNewPost":       {"func": CreateNewPostLogic, "test_case_count": len(testcases.CreateNewUserPostData), "desc": testcases.CreateNewUserPostDescription()},
		"UpdateUserPost":      {"func": UpdateUserPostLogic, "test_case_count": len(testcases.UpdateUserPostDetailData), "desc": testcases.UpdateUserPostDetailDescription()},
		"GetCommentsByID":     {"func": GetCommentsLogic, "test_case_count": len(testcases.GetCommentsData), "desc": testcases.ListGetCommentsDescription()},
		"CreateComment":       {"func": CreateCommentLogic, "test_case_count": len(testcases.CreateCommentOnAPostData), "desc": testcases.CreateCommentDescription()},
		"GetListUserPostByMe": {"func": GetListPostByMeLogic, "test_case_count": len(testcases.GetListUserPostByMeData), "desc": testcases.ListUserPostByMeDescription()},
		"LikeOrDislikeOnPost": {"func": LikeOrDislikePostLogic, "test_case_count": len(testcases.LikeOrDislikePostData), "desc": testcases.LikeOrDislikePostDescription()},
		"CheckReadiness":      {"func": CheckReadinessLogic, "test_case_count": len(testcases.CheckReadinessData), "desc": testcases.CheckReadinessDescription()},
	}

	for _, val := range unitTestLogic {
		s := reflect.ValueOf(val["desc"])
		var arr []TableEntry
		for i := 0; i < val["test_case_count"].(int); i++ {
			fmt.Println(s.Index(i).String())
			arr = append(arr, Entry(s.Index(i).String(), i))
		}
		DescribeTable("Function ", val["func"], arr...)
	}
})
