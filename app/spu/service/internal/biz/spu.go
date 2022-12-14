package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Spu struct {
	Id             int64  ` json:"id,omitempty"`
	Title          string ` json:"title,omitempty"`
	Subtitle       string ` json:"subtitle,omitempty"`
	CategoryId     int64  ` json:"category_id,omitempty"`
	RootCategoryId int64  ` json:"root_category_id,omitempty"`
	Price          string ` json:"price,omitempty"`
	Img            string ` json:"img,omitempty"`
	ForThemeImg    string ` json:"for_theme_img,omitempty"`
	Description    string ` json:"description,omitempty"`
	DiscountPrice  string ` json:"discount_price,omitempty"`
	Tags           string ` json:"tags,omitempty"`

	Online int32 `json:"online,omitempty"`
}

type SpuRepo interface {
	CreateSpu(ctx context.Context) (err error)
	GetSpuById(ctx context.Context, id int64) (Spu Spu, err error)
	GetSpuByCategory(ctx context.Context, id int64) (Spus []Spu, err error)
	ListSpuByIds(ctx context.Context, ids []int64) (Spus []Spu, err error)
}
type SpuUsecase struct {
	repo SpuRepo
	log  *log.Helper
}

func NewSpuUsecase(repo SpuRepo, logger log.Logger) *SpuUsecase {
	return &SpuUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *SpuUsecase) ListSpuByIds(ctx context.Context, ids []int64) (Spus []Spu, err error) {
	return uc.repo.ListSpuByIds(ctx, ids)
}
