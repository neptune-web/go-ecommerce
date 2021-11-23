package data

import (
	"context"
	"mall-go/internal/biz"
	"mall-go/internal/data/model/banner"

	"github.com/go-kratos/kratos/v2/log"
)

type bannerRepo struct {
	data *Data
	log  *log.Helper
}

func NewBannerRepo(data *Data, logger log.Logger) biz.BannerRepo {
	return &bannerRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *bannerRepo) GetBannerById(ctx context.Context, id int64) (b biz.Banner, err error) {
	po, err := r.data.db.Banner.Query().Where(banner.ID(id)).WithItem().First(ctx)
	if err != nil {
		return 
	}
	var items []biz.BannerItem
	for _, v := range po.Edges.Item {

		item := biz.BannerItem{
			ID:       v.ID,
			Name:     v.Name,
			Img:      v.Img,
			Keyword:  v.Keyword,
			Type:     v.Type,
			BannerID: v.BannerID,
		}
		items = append(items, item)
	}
	return biz.Banner{
		Id:          po.ID,
		Name:        po.Name,
		Title:       po.Title,
		Description: po.Description,
		Img:         po.Description,
		Items:       items,
	}, nil
}
