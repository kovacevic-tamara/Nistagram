package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/grpc_common"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type PostGrpcController struct {
	service    *services.PostService
	jwtManager *common.JWTManager
}

func NewPostController(db *gorm.DB, jwtManager *common.JWTManager) (*PostGrpcController, error) {
	service, err := services.NewPostService(db)
	if err != nil {
		return nil, err
	}

	return &PostGrpcController{
		service,
		jwtManager,
	}, nil
}

func (c *PostGrpcController) CreatePost(ctx context.Context, in *protopb.Post) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePost")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	}  else if claims.UserId == "" {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.InvalidArgument, "no user id provided")
	}  else if claims.UserId != in.UserId {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot create post for another user")
	}

	var post *domain.Post
	post = post.ConvertFromGrpc(in)

	for _, media := range post.Media{
		for _, tag := range media.Tags {
			following, err := grpc_common.CheckFollowInteraction(ctx, tag.UserId, post.UserId)
			if err != nil {
				return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot tag selected users")
			}

			isPublic, err := grpc_common.CheckIfPublicProfile(ctx, in.Id)
			if err != nil {
				return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
			}

			isBlocked, err := grpc_common.CheckIfBlocked(ctx, in.Id, claims.UserId)
			if err != nil {
				return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
			}

			// If used is blocked or his profile is private and did not approve your request
			if isBlocked || (!isPublic && !following.IsApprovedRequest ) {
				return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot tag selected users")
			}

			username, err := grpc_common.GetUsernameById(ctx, tag.UserId)
			if username == "" || err != nil {
				return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot tag selected users")
			}
		}
	}

	err = c.service.CreatePost(ctx, post)
	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	}

	return &protopb.EmptyResponseContent{}, nil
}

func (c *PostGrpcController) GetAllPosts(ctx context.Context, in *protopb.EmptyRequestContent) (*protopb.ReducedPostArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllPosts")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.ReducedPostArray{}, status.Errorf(codes.Unknown, err.Error())
	}else if claims.UserId == ""{
		return &protopb.ReducedPostArray{}, status.Errorf(codes.InvalidArgument, "no user id is provided")
	}

	userIds, err := grpc_common.GetHomepageUsers(ctx, claims.UserId)
	if err != nil {
		return &protopb.ReducedPostArray{}, status.Errorf(codes.Unknown, err.Error())
	}

	posts, err := c.service.GetAllPosts(ctx, userIds)
	if err != nil {
		return &protopb.ReducedPostArray{}, status.Errorf(codes.Unknown, err.Error())
	}

	responsePosts := []*protopb.ReducedPost{}
	for _, post := range posts {
		responsePosts = append(responsePosts, post.ConvertToGrpc())
	}

	return &protopb.ReducedPostArray{
		Posts: responsePosts,
	}, nil
}

func (c *PostGrpcController) GetPostsForUser(ctx context.Context, in *protopb.RequestId) (*protopb.ReducedPostArray, error){
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostsForUser")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.ReducedPostArray{}, status.Errorf(codes.Unknown, err.Error())
	}else if claims.UserId == "" || in.Id == "" {
		return &protopb.ReducedPostArray{}, status.Errorf(codes.InvalidArgument, "no user id is provided")
	}

	if in.Id != claims.UserId{
		followConnection, err := grpc_common.CheckFollowInteraction(ctx, in.Id, claims.UserId)
		if err != nil {
			return &protopb.ReducedPostArray{}, status.Errorf(codes.Unknown, err.Error())
		}

		isPublic, err := grpc_common.CheckIfPublicProfile(ctx, in.Id)
		if err != nil {
			return &protopb.ReducedPostArray{}, status.Errorf(codes.Unknown, err.Error())
		}

		isBlocked, err := grpc_common.CheckIfBlocked(ctx, in.Id, claims.UserId)
		if err != nil {
			return &protopb.ReducedPostArray{}, status.Errorf(codes.Unknown, err.Error())
		}

		// If used is blocked or his profile is private and did not approve your request
		if isBlocked || (!isPublic && !followConnection.IsApprovedRequest) {
			return &protopb.ReducedPostArray{}, nil
		}
	}

	posts, err := c.service.GetPostsForUser(ctx, in.Id)
	if err != nil{
		return &protopb.ReducedPostArray{}, status.Errorf(codes.Unknown, err.Error())
	}

	responsePosts := domain.ConvertMultipleReducedPostsToGrpc(posts)

	return &protopb.ReducedPostArray{ Posts: responsePosts }, nil
}

func (c *PostGrpcController) GetPostById(ctx context.Context, id string) (*protopb.Post, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostById")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.Post{}, status.Errorf(codes.Unknown, err.Error())
	}else if claims.UserId == ""{
		return &protopb.Post{}, status.Errorf(codes.Unknown, err.Error())
	}else if id == "" {
		return &protopb.Post{}, status.Errorf(codes.InvalidArgument, "cannot retrieve non-existing posts")
	}

	post, err := c.service.GetPostById(ctx, id)
	if err != nil {
		return &protopb.Post{}, status.Errorf(codes.Unknown, err.Error())
	}

	if post.UserId != claims.UserId{
		following, err := grpc_common.CheckFollowInteraction(ctx, post.UserId, claims.UserId)
		if err != nil { return &protopb.Post{}, status.Errorf(codes.Unknown, err.Error()) }

		isPublic, err := grpc_common.CheckIfPublicProfile(ctx, post.UserId)
		if err != nil { return &protopb.Post{}, status.Errorf(codes.Unknown, err.Error()) }

		isBlocked, err := grpc_common.CheckIfBlocked(ctx, post.UserId, claims.UserId)
		if err != nil {
			return &protopb.Post{}, status.Errorf(codes.Unknown, err.Error())
		}

		if (!following.IsApprovedRequest && !isPublic) || isBlocked {
			return &protopb.Post{}, status.Errorf(codes.PermissionDenied, "cannot retrieve this post")
		}
	}

	grpcPost := post.ConvertToGrpc()
	return grpcPost, nil
}

func (c *PostGrpcController) RemovePost(ctx context.Context, id string) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemovePost")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	}else if claims.UserId == ""{
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot remove other people's posts")
	}else if id == "" {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "cannot remove non-existing posts")
	}

	err = c.service.RemovePost(ctx, id, claims.UserId)
	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	}

	return &protopb.EmptyResponseContent{}, nil
}

func (c *PostGrpcController) SearchContentByLocation(ctx context.Context, in *protopb.SearchLocationRequest) (*protopb.ReducedPostArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	location := in.Location

	posts, err := c.service.SearchContentByLocation(ctx, location)
	if err != nil {
		return &protopb.ReducedPostArray{
			Posts: []*protopb.ReducedPost{},
		}, status.Errorf(codes.Unknown, err.Error())
	}

	responsePosts := []*protopb.ReducedPost{}
	for _, post := range posts {
		responsePosts = append(responsePosts, post.ConvertToGrpc())
	}

	return &protopb.ReducedPostArray{
		Posts: responsePosts,
	}, nil
}

func (c *PostGrpcController) GetPostsByHashtag(ctx context.Context, in *protopb.Hashtag) (*protopb.ReducedPostArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostsByHashtag")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	posts, err := c.service.GetPostsByHashtag(ctx, in.Text)
	if err != nil {
		return &protopb.ReducedPostArray{
			Posts: []*protopb.ReducedPost{},
		}, status.Errorf(codes.Unknown, err.Error())
	}

	responsePosts := []*protopb.ReducedPost{}
	for _, post := range posts {
		responsePosts = append(responsePosts, post.ConvertToGrpc())
	}

	return &protopb.ReducedPostArray{
		Posts: responsePosts,
	}, nil
}
