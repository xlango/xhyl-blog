package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strings"
	"xhylblog/dbconn"

	"github.com/astaxie/beego/orm"
)

var MongoClient *dbconn.MongoClient

type Article struct {
	Id      int64
	Title   string `orm:"size(255)"`
	Content string `orm:"size(255)"`
	Image   string `orm:"size(255)"`
}

//文章主要内容
type ArticleContent struct {
	_id bson.ObjectId `bson:"_id,omitempty"` //mongo objectId  	a := bson.NewObjectId()
	Id int64
	Paragraphs []Paragraph //段落
}
//段落内容
type Paragraph struct {
	Title string
	Content string
	Images   []string
}

//添加文章
type ArticleDetailModel struct {
	Article *Article //存入mysql
	Paragraphs []Paragraph //存入mongo
}

func init() {
	orm.RegisterModelWithPrefix("tb_", new(Article))
	MongoClient=&dbconn.MongoClient{
		Database:"xhyl",
		Collection:"article_paragraph",
	}
}

// AddArticle insert a new Article into database and returns
// last inserted Id on success.
func AddArticle(m *Article) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetArticleById retrieves Article by Id. Returns error if
// Id doesn't exist
func GetArticleById(id int64) (v *Article, err error) {
	o := orm.NewOrm()
	v = &Article{Id: id}
	if err = o.QueryTable(new(Article)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllArticle retrieves all Article matches certain condition. Returns empty list if
// no records exist
func GetAllArticle(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Article))
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

	var l []Article
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				paragraphsMongo := MongoClient.FindOne(map[string]interface{}{"id": v.Id})
				var paragraphs []Paragraph
				if paragraphsMongo!=nil {
					data,_:=bson.Marshal(paragraphsMongo)
					articleContent:=&ArticleContent{}
					bson.Unmarshal(data, articleContent)
					paragraphs=articleContent.Paragraphs
				}
				articleDetail:=&ArticleDetailModel{
					Article:&v,
					Paragraphs:paragraphs,
				}
				ml = append(ml, articleDetail)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				paragraphsMongo := MongoClient.FindOne(map[string]interface{}{"id": v.Id})
				var paragraphs []Paragraph
				if paragraphsMongo!=nil {
					data,_:=bson.Marshal(paragraphsMongo)
					articleContent:=&ArticleContent{}
					bson.Unmarshal(data, articleContent)
					paragraphs=articleContent.Paragraphs
				}
				articleDetail:=&ArticleDetailModel{
					Article:&v,
					Paragraphs:paragraphs,
				}
				ml = append(ml, articleDetail)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateArticle updates Article by Id and returns error if
// the record to be updated doesn't exist
func UpdateArticleById(m *Article) (err error) {
	o := orm.NewOrm()
	v := Article{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteArticle deletes Article by Id and returns error if
// the record to be deleted doesn't exist
func DeleteArticle(id int64) (err error) {
	o := orm.NewOrm()
	v := Article{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Article{Id: id}); err == nil {
			logs.Info("(mysql)Number of article records deleted in database:", num)
		}
		del := MongoClient.Delete(map[string]interface{}{"id": id})
		if !del{
			logs.Error("[error] delete article from mongo fail!")
		}
	}
	return
}

/**
将博客主要内容存入mongo
 */
func AddArticleContentToMongo(m *ArticleDetailModel) (err error) {
	if id, err := AddArticle(m.Article);err==nil{
		bson.NewObjectId()

		content:=&ArticleContent{
			Id:id,
			Paragraphs:m.Paragraphs,
		}
		insert := MongoClient.Insert(content)
		if !insert {
			logs.Error("[error] insert article to mongo fail!")
		}
	}
	return
}

// UpdateArticle updates Article by Id and returns error if
// the record to be updated doesn't exist
func UpdateArticleDetailById(m *ArticleDetailModel) (err error) {
	o := orm.NewOrm()
	v := Article{Id: m.Article.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m.Article); err == nil {
			logs.Info("(mysql)Number of article records updated in database:", num)
		}
		update := MongoClient.Update(map[string]interface{}{"id": m.Article.Id}, map[string]interface{}{"paragraph": m.Paragraphs})
		if !update{
			logs.Error("[error] update article to mongo fail!")
		}
	}
	return
}