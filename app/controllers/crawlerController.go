package controllers

import (
	"context"
	"fmt"
	"gin-example/utils/db"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/gocolly/colly/extensions"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type book struct {
	BookId       uint              `json:"book_id" gorm:"primaryKey" `
	Class        string            `json:"class"`                                     // 分类
	BookName     string            `json:"book_name"`                                 // 书名
	Author       string            `json:"author"`                                    // 作者
	Intro        string            `json:"intro"`                                     // 简介
	CoverPicture string            `json:"cover_picture"`                             // 封面图
	Chapters     map[string]string `json:"chapters" gorm:"serializer:json;type:json"` // 章节
}

func GetBook(ctx *gin.Context) {
	// 数据库
	database := db.Mongo.Database("books")
	// 集合
	bookCol := database.Collection("book")
	s := book{BookName: "abc", Author: "jade", Chapters: map[string]string{"annd": "string(123)"}}

	one, err := bookCol.InsertOne(context.Background(), s)
	if err != nil {
		log.Fatal(err.Error())
	}

	db.Mysql.Table("books").Create(&s)

	log.Fatalf("Inserted a single document: %s   mysqlID %d", one.InsertedID, s.BookId)

	baseUrl := "https://www.biquge.co"
	b := &book{}
	b.Chapters = make(map[string]string)

	c1 := colly.NewCollector(
		colly.AllowedDomains("www.biquge.co", "www.bqg99.com", "www.cqwsjds.com"),
		//colly.Async(true),
		colly.Debugger(&debug.LogDebugger{}),
	)
	// 随机UserAgent
	extensions.RandomUserAgent(c1)
	// 解决中午乱码
	c1.DetectCharset = true

	c2 := c1.Clone()
	// 异步
	c2.Async = true
	// 限速
	_ = c2.Limit(&colly.LimitRule{
		DomainRegexp: "",
		DomainGlob:   "*.biquge.co/*",
		Delay:        10 * time.Second,
		RandomDelay:  0,
		Parallelism:  1,
	})

	// 采集文章信息
	c1.OnHTML("#wrapper", func(e *colly.HTMLElement) {
		b.Class = e.ChildText(".con_top > a:nth-child(3)")

		b.BookName = e.ChildText("#info h1")
		author := e.ChildText("#info p:nth-child(2)")
		b.Author = strings.TrimSpace(strings.Split(author, "：")[1])
		b.Intro = e.ChildText("#intro p:nth-child(1)")
		b.CoverPicture = e.ChildAttr("#fmimg img[src]", "src")
	})

	// 获取章节列表
	c1.OnHTML("#list > dl", func(e *colly.HTMLElement) {
		e.ForEach("dt:nth-child(1)~dd", func(i int, ele *colly.HTMLElement) {
			href := baseUrl + ele.ChildAttr("a", "href")
			chapter := ele.ChildText("a")
			b.Chapters[href] = chapter

			ctx := colly.NewContext()
			ctx.Put("chapter", chapter)
			c2.Request("GET", href, nil, ctx, nil)
		})
	})

	// 获取文章详情
	c2.OnHTML("#content", func(e *colly.HTMLElement) {
		chapter := e.Request.Ctx.Get("chapter")
		context := e.Text
		context = strings.ReplaceAll(context, "&nbsp;&nbsp;&nbsp;&nbsp;", "")
		//log.Fatalf("%v", context)

		log.Printf("开始保存文章 %s...", chapter)

		dir, err := os.Getwd()
		if err != nil {
		}

		dir = dir + "/storage/books/" + b.BookName

		err = os.MkdirAll(dir, 0766)
		if err != nil {
			log.Fatalf("创建目录[%s]错误：%s", dir, err.Error())
			return
		}

		fileName := dir + "/" + chapter + ".txt"
		// 判断文件是否存在   不存在就创建并下载
		if _, err = os.Stat(fileName); os.IsNotExist(err) {
			f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				log.Fatalln("打开文件错误：", err.Error())
				return
			}

			_, err = f.WriteString(context)
			f.Close()
			if err != nil {
				log.Fatalln("写入文件错误：", err.Error())
				return
			}
			log.Printf("下载文章[%s]end ", fileName)
		} else {
			log.Printf("文章[%s]已经存在 ", fileName)
		}

		//log.Fatalln(e.Text)
	})

	c1.OnRequest(func(r *colly.Request) {
		fmt.Println("c爬取页面：", r.URL)
	})

	err = c1.Visit(fmt.Sprintf("%s/%s/", baseUrl, "28_28615"))
	if err != nil {
		log.Printf("error : %s", err.Error())
		return
	}
	//c1.Wait()
	c2.Wait()

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"book": b,
	})
}
