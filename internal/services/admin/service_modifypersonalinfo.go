package admin

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/admin"
)

type ModifyData struct {
	Nickname string // 昵称
	Mobile   string // 手机号
}

func (s *service) ModifyPersonalInfo(ctx core.Context, id int32, modifyData *ModifyData) (err error) {
	data := map[string]interface{}{
		"nickname":     modifyData.Nickname,
		"mobile":       modifyData.Mobile,
		"updated_user": ctx.SessionUserInfo().UserName,
	}

	qb := admin.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	return
}
