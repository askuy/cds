package logic

import (
	"cds/galaxy/internal/svc"
	"cds/galaxy/internal/types"
	"cds/tools/clickhousex"
	"context"
	"sort"

	"github.com/tal-tech/go-zero/core/logx"
)

type ListDatabasesLogic struct {
	ctx context.Context
	logx.Logger
}

func NewListDatabasesLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListDatabasesLogic {
	return ListDatabasesLogic{
		ctx:    ctx,
		Logger: logx.WithContext(ctx),
	}
	// TODO need set model here from svc
}

func (l *ListDatabasesLogic) ListDatabases(req string) (*types.StringList, error) {
	cli := clickhousex.TakeClientx(req)
	dbs, e := clickhousex.ShowDatabases(cli)
	if e != nil {
		return nil, e
	}
	sort.Strings(dbs)
	return &types.StringList{StringList: dbs}, nil
}
