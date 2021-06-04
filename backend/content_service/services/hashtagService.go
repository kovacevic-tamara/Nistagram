package services

import (
	"context"
	"errors"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/repositories"
	"gorm.io/gorm"
)

type HashtagService struct {
	hashtagRepository repositories.HashtagRepository
}

func NewHashtagService(db *gorm.DB) (*HashtagService, error) {
	hashtagRepository, err := repositories.NewHashtagRepo(db)
	if err != nil {
		return nil, err
	}

	return &HashtagService{
		hashtagRepository,
	}, err
}

func (service HashtagService) CreateHashtag(ctx context.Context, text string) (*domain.Hashtag, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateHashtag")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if text == "" {
		return nil, errors.New("cannot create empty hashtag")
	}

	hashtag, err := service.hashtagRepository.CreateHashtag(ctx, text)
	if err != nil {
		return nil, errors.New("hashtag does not exist")
	}

	return &domain.Hashtag{Id: hashtag.Id, Text: hashtag.Text}, nil
}
