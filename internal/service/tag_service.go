package service

import (
	"context"
	"github.com/kalougata/bookkeeping/internal/data"
	"github.com/kalougata/bookkeeping/internal/model"
	"github.com/kalougata/bookkeeping/pkg/e"
)

type TagService struct {
	data *data.Data
}

func (ts *TagService) Create(ctx context.Context, req *model.TagInReq) error {
	if exist, err := ts.data.DB.Context(ctx).Table(&model.Tag{}).Where("name = ?", req.Name).Exist(); exist && err == nil {
		return e.ErrBadRequest().WithMsg("标签名已存在~")
	}
	if count, err := ts.data.DB.Context(ctx).Table(&model.Tag{}).Insert(req.ToModel()); err != nil || count <= 0 {
		return e.ErrInternalServer().WithMsg("添加标签失败, 请稍后再试~")
	}

	return nil
}

func NewTagService(data *data.Data) *TagService {
	return &TagService{data}
}