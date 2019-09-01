package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ArticleType struct {
	Id       int64
	TypeName string `orm:"column(TypeName)","size(128)"`
	Logo string `orm:"size(128)"`
}

type  ArticleTypeRelation struct {
	Id int64
	TypeId      int64 `orm:"column(TypeId)","size(20)"`
	ArticleId      int64 `orm:"column(ArticleId)","size(20)"`
}

func init() {
	orm.RegisterModelWithPrefix("tb_", new(ArticleType),new(ArticleTypeRelation))
}

// AddArticleType insert a new ArticleType into database and returns
// last inserted Id on success.
func AddArticleType(m *ArticleType) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetArticleTypeById retrieves ArticleType by Id. Returns error if
// Id doesn't exist
func GetArticleTypeById(id int64) (v *ArticleType, err error) {
	o := orm.NewOrm()
	v = &ArticleType{Id: id}
	if err = o.QueryTable(new(ArticleType)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllArticleType retrieves all ArticleType matches certain condition. Returns empty list if
// no records exist
func GetAllArticleType(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ArticleType))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []ArticleType
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateArticleType updates ArticleType by Id and returns error if
// the record to be updated doesn't exist
func UpdateArticleTypeById(m *ArticleType) (err error) {
	o := orm.NewOrm()
	v := ArticleType{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteArticleType deletes ArticleType by Id and returns error if
// the record to be deleted doesn't exist
func DeleteArticleType(id int64) (err error) {
	o := orm.NewOrm()
	v := ArticleType{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ArticleType{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}


//通过文章Id查询所有相关类型
func GetArtcleTypesByArtcleId(id int64) (rs []ArticleType) {
	o := orm.NewOrm()
	_, err := o.Raw("SELECT a.TypeName as TypeName ,a.Logo as Logo FROM tb_article_type a LEFT JOIN tb_article_type_relation b on a.Id=b.TypeId WHERE b.ArticleId=?", id).QueryRows(&rs)
	if err != nil {
		logs.Error("article of",id," type no found ")
	}
	return
}

// 文章作者关联
func ArticleAddType(m []ArticleTypeRelation)(num int,err error) {
	o := orm.NewOrm()
	_, err = o.InsertMulti(100,m)
	return
}