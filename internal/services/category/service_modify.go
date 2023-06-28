package category

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/category"
)

type UpdateCategoryData struct {
	Name string // 菜单名称
	Icon string // 图标
}

func (s *service) Modify(ctx core.Context, id int32, categoryData *UpdateCategoryData) (err error) {
	data := map[string]interface{}{
		"title": categoryData.Name,
		"icon":  categoryData.Icon,
	}

	qb := category.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	return
}
