package repositories

import (
	"context"
	"fmt"
	"github.com/david-drvar/xws2021-nistagram/agent_application/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/agent_application/util"
	"github.com/david-drvar/xws2021-nistagram/agent_application/util/images"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ProductRepository interface {
	CreateProduct(context.Context, persistence.Product) error
	SaveProductPhoto(ctx context.Context, product persistence.Product) error
	GetAllProductsByAgentId(ctx context.Context, id string) ([]persistence.Product, error)
	GetAllProducts(ctx context.Context) ([]persistence.Product, error)
}

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepo(db *gorm.DB) (*productRepository, error) {
	if db == nil {
		panic("UserRepository not created, gorm.DB is nil")
	}

	return &productRepository{DB: db}, nil
}

func (repository *productRepository) CreateProduct(ctx context.Context, product persistence.Product) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	product.Id = uuid.New().String()
	resultUser := repository.DB.Create(&product)
	if resultUser.Error != nil {
		return resultUser.Error
	}

	if product.Photo != "" {
		err := repository.SaveProductPhoto(ctx, product)
		if err != nil {
			return err
		}
	}

	return nil
}

func (repository *productRepository) GetAllProductsByAgentId(ctx context.Context, id string) ([]persistence.Product, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllProductsByAgentId")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var products []persistence.Product
	resultUser := repository.DB.Where("agent_id = ?", id).Find(&products)
	if resultUser.Error != nil {
		return nil, resultUser.Error
	}

	return products, nil
}

func (repository *productRepository) GetAllProducts(ctx context.Context) ([]persistence.Product, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllProducts")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var products []persistence.Product
	resultUser := repository.DB.Find(&products)
	if resultUser.Error != nil {
		return nil, resultUser.Error
	}

	return products, nil
}

func (repository *productRepository) SaveProductPhoto(ctx context.Context, product persistence.Product) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "SaveUserProfilePhoto")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	mimeType, err := images.GetImageType(product.Photo)
	if err != nil {
		return err
	}

	t := time.Now()
	formatted := fmt.Sprintf("%s%d%02d%02d%02d%02d%02d%02d", product.Id, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
	name := formatted + "." + mimeType

	err = images.SaveImage(name, product.Photo)
	if err != nil {
		return err
	}

	product.Photo = util.GetContentLocation(name)
	db := repository.DB.Model(&product).Where("id = ?", product.Id).Updates(product)
	if db.Error != nil {
		return db.Error
	}

	return nil
}
