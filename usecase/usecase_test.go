package usecase_test

import (
	"context"
	"fmt"
	"os"
	"reflect"

	"github.com/sapawarga/userpost-service/mocks"
	"github.com/sapawarga/userpost-service/mocks/testcases"
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
		data := testcases.GetListUserPostData[idx]
		mockPostRepo.EXPECT().GetListPost(ctx, gomock.Any()).Return(data.ResponseGetList.Result, data.ResponseGetList.Error).Times(1)
		mockPostRepo.EXPECT().GetMetadataPost(ctx, gomock.Any()).Return(data.ResponseMetadata.Result, data.ResponseMetadata.Error).Times(1)
		mockCommentsRepo.EXPECT().GetLastComment(ctx, data.GetLastCommentParams).Return(data.ResponseGetLastComment.Result, data.ResponseGetLastComment.Error).Times(len(data.ResponseGetList.Result))
		mockPostRepo.EXPECT().GetActor(ctx, data.GetActorParams).Return(data.ResponseGetActor.Result, data.ResponseGetActor.Error).Times(len(data.ResponseGetList.Result) * 3)
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
		ctx := context.Background()
		data := testcases.GetDetailUserPostData[idx]
		mockPostRepo.EXPECT().GetDetailPost(ctx, data.GetUserPostParams).Return(data.ResponseGetDetailUserPost.Result, data.ResponseGetDetailUserPost.Error).Times(1)
		mockCommentsRepo.EXPECT().GetLastComment(ctx, data.GetLastCommentParams).Return(data.ResponseGetLastComment.Result, data.ResponseGetLastComment.Error).Times(1)
		mockCommentsRepo.EXPECT().GetTotalComments(ctx, data.GetTotalCommentsParams).Return(data.ResponseGetTotalComment.Result, data.ResponseGetTotalComment.Error).Times(1)
		mockPostRepo.EXPECT().GetActor(ctx, data.GetActorParams).Return(data.ResponseGetActor.Result, data.ResponseGetActor.Error).Times(3)
		resp, err := userPost.GetDetailPost(ctx, data.UsecaseParams)
		if err != nil {
			Expect(err).NotTo(BeNil())
			Expect(resp).To(BeNil())
		} else {
			Expect(err).To(BeNil())
			Expect(resp).NotTo(BeNil())
		}
	}

	var unitTestLogic = map[string]map[string]interface{}{
		"GetListUserPost":   {"func": GetListUserPostLogic, "test_case_count": len(testcases.GetListUserPostData), "desc": testcases.ListUserPostDescription()},
		"GetDetailUserPost": {"func": GetDetailUserPostLogic, "test_case_count": len(testcases.GetDetailUserPostData), "desc": testcases.ListUserPostDetailDescription()},
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
