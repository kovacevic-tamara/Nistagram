package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/content_service/model/domain"
	"github.com/david-drvar/xws2021-nistagram/content_service/services"
	"gorm.io/gorm"
)

type CampaignGrpcController struct {
	service 	*services.CampaignService
	jwtManager  *common.JWTManager
}

func NewCampaignController(db *gorm.DB, jwtManager *common.JWTManager) (*CampaignGrpcController, error) {
	service, err := services.NewCampaignService(db)
	if err != nil {
		return nil, err
	}

	return &CampaignGrpcController{
		service,
		jwtManager,
	}, nil
}

func (controller *CampaignGrpcController) GetCampaign(ctx context.Context, in *protopb.RequestId) (*protopb.Campaign, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCampaign")
	defer span.Finish()
	// claims, _ := controller.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	campaign, err := controller.service.GetCampaign(ctx, in.Id)
	if err != nil { return &protopb.Campaign{}, err }

	return campaign.ConvertToGrpc(), nil
}

func (controller *CampaignGrpcController) GetCampaigns(ctx context.Context, in *protopb.EmptyRequestContent) (*protopb.CampaignArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCampaigns")
	defer span.Finish()
	claims, _ := controller.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	campaigns, err := controller.service.GetCampaigns(ctx, claims.UserId)
	if err != nil { return &protopb.CampaignArray{}, err }

	responseCampaigns := []*protopb.Campaign{}
	for _, campaign := range campaigns{
		responseCampaigns = append(responseCampaigns, campaign.ConvertToGrpc())
	}

	return &protopb.CampaignArray{
		Campaigns: responseCampaigns,
	}, nil
}

func (controller *CampaignGrpcController) CreateCampaign(ctx context.Context, in *protopb.Campaign) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateCampaign")
	defer span.Finish()
	// claims, _ := controller.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	campaign := domain.Campaign{}
	campaign = campaign.ConvertFromGrpc(in)

	err := controller.service.CreateCampaign(ctx, campaign)
	if err != nil { return &protopb.EmptyResponseContent{}, err }

	return &protopb.EmptyResponseContent{}, nil
}

func (controller *CampaignGrpcController) UpdateCampaign(ctx context.Context, in *protopb.Campaign) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateCampaign")
	defer span.Finish()
	// claims, _ := controller.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var campaign domain.Campaign
	campaign = campaign.ConvertFromGrpc(in)

	err := controller.service.UpdateCampaign(ctx, campaign)
	if err != nil { return &protopb.EmptyResponseContent{}, err }

	return &protopb.EmptyResponseContent{}, nil
}

func (controller *CampaignGrpcController) DeleteCampaign(ctx context.Context, in *protopb.RequestId) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteCampaign")
	defer span.Finish()
	// claims, _ := controller.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := controller.service.DeleteCampaign(ctx, in.Id)
	if err != nil { return &protopb.EmptyResponseContent{}, err }

	return &protopb.EmptyResponseContent{}, nil
}

func (controller *CampaignGrpcController) GetCampaignStats(ctx context.Context, in *protopb.RequestId) (*protopb.CampaignStats, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetCampaignStats")
	defer span.Finish()
	claims, _ := controller.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	stats, err := controller.service.GetCampaignStatistics(ctx, claims.UserId, in.Id)
	if err != nil { return &protopb.CampaignStats{}, err }

	return stats.ConvertToGrpc(), nil
}