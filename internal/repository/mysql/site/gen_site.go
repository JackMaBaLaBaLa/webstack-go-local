///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package site

import (
	"fmt"
	"time"

	"github.com/ch3nnn/webstack-go/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *Site {
	return new(Site)
}

func NewQueryBuilder() *siteQueryBuilder {
	return new(siteQueryBuilder)
}

func (t *Site) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type siteQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *siteQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *siteQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&Site{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *siteQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&Site{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *siteQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&Site{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *siteQueryBuilder) First(db *gorm.DB) (*Site, error) {
	ret := &Site{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *siteQueryBuilder) QueryOne(db *gorm.DB) (*Site, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *siteQueryBuilder) QueryAll(db *gorm.DB) ([]*Site, error) {
	var ret []*Site
	err := qb.buildQuery(db).Preload("Category").Find(&ret).Error
	return ret, err
}

func (qb *siteQueryBuilder) Limit(limit int) *siteQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *siteQueryBuilder) Offset(offset int) *siteQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *siteQueryBuilder) WhereId(p mysql.Predicate, value int32) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) WhereIdIn(value []int32) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) WhereIdNotIn(value []int32) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) OrderById(asc bool) *siteQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *siteQueryBuilder) WhereCategoryId(p mysql.Predicate, value int32) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "category_id", p),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) WhereCategoryIdIn(value []int32) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "category_id", "IN"),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) WhereCategoryIdNotIn(value []int32) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "category_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) OrderByCategoryId(asc bool) *siteQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "category_id "+order)
	return qb
}

func (qb *siteQueryBuilder) WhereIsUsed(p mysql.Predicate, value int32) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_used", p),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) WhereTitle(p mysql.Predicate, value string) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "title", p),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) WhereTitleLike(value string) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "title ", "LIKE"),
		"%" + value + "%",
	})
	return qb
}

func (qb *siteQueryBuilder) WhereTitleIn(value []string) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "title", "IN"),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) WhereTitleNotIn(value []string) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "title", "NOT IN"),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) OrderByTitle(asc bool) *siteQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "title "+order)
	return qb
}

func (qb *siteQueryBuilder) WhereThumb(p mysql.Predicate, value string) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "thumb", p),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) WhereThumbIn(value []string) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "thumb", "IN"),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) WhereThumbNotIn(value []string) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "thumb", "NOT IN"),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) OrderByThumb(asc bool) *siteQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "thumb "+order)
	return qb
}

func (qb *siteQueryBuilder) WhereDescription(p mysql.Predicate, value string) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "description", p),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) WhereDescriptionIn(value []string) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "description", "IN"),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) WhereDescriptionNotIn(value []string) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "description", "NOT IN"),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) OrderByDescription(asc bool) *siteQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "description "+order)
	return qb
}

func (qb *siteQueryBuilder) WhereUrl(p mysql.Predicate, value string) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "url", p),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) WhereUrlIn(value []string) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "url", "IN"),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) WhereUrlNotIn(value []string) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "url", "NOT IN"),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) OrderByUrl(asc bool) *siteQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "url "+order)
	return qb
}

func (qb *siteQueryBuilder) WhereCreateTime(p mysql.Predicate, value time.Time) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) WhereCreateTimeIn(value []time.Time) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) WhereCreateTimeNotIn(value []time.Time) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) OrderByCreateTime(asc bool) *siteQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created"+
		"_at "+order)
	return qb
}

func (qb *siteQueryBuilder) WhereUpdateTime(p mysql.Predicate, value time.Time) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) WhereUpdateTimeIn(value []time.Time) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) WhereUpdateTimeNotIn(value []time.Time) *siteQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *siteQueryBuilder) OrderByUpdateTime(asc bool) *siteQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}
