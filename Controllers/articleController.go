package Controllers

import (
	"GinDemo/Models"
	"GinDemo/Services"
	"GinDemo/Util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加文章
func AddArticle(ctx *gin.Context)  {
	var article Models.Article
	var contentx []Models.TbContent
	err:=ctx.ShouldBindJSON(&article)
	fmt.Println(len([]rune(article.Content)))
	article.Content=string([]rune(article.Content)[:len(article.Content)-1])
	//fmt.Println(article.Content)
	//获取当前时间，并赋值给结构体
	article.CreatedTime=Util.InitTime()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	con:=article.Content
	if len([]rune(con))>15374 {
		article.Content=string([]rune(article.Content)[:15374])
		art,dberr := Services.InsertArticle(article)
		contentx=Services.Interception(con,art.Id)
		conerr:=Services.InsertContent(contentx)
		if dberr!=nil||conerr!=nil {
			ctx.JSON(http.StatusAccepted,gin.H{
				"Code":0,
				"Message":"数据插入错误",
				"ArtData":dberr,
				"ContentData":conerr,
			})
		}else {
			ctx.JSON(http.StatusOK,gin.H{
				"Code":1,
				"Message":"数据插入成功",
				"Result":"Success",
			})
		}
	}else {
		art,dberr := Services.InsertArticle(article)
		if dberr!=nil {
			ctx.JSON(http.StatusAccepted,gin.H{
				"Code":0,
				"Message":"数据插入错误",
				"Data":dberr,
			})
		}else {
			ctx.JSON(http.StatusOK,gin.H{
				"Code":1,
				"Message":"数据插入成功",
				"Result":"Success",
				"Data":art,
			})
		}
	}

}

//查找全部文章
func FindAllArticle(ctx *gin.Context)  {
	result,err :=Services.FindAllArticle()
	fmt.Println("数据获取",result)
	if err != nil||len(result)==0 {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"Code":0,
			"Message":"数据获取错误",
		})
	}else {
		ctx.JSON(http.StatusOK,gin.H{
			"Code":1,
			"Message":"数据获取成功",
			"Data":result,
		})
	}
}

//查找前6篇文章
func FindArticle6(ctx *gin.Context)  {
	result,err :=Services.FindArticle6()
	if err != nil||len(result)==0 {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"Code":0,
			"Message":"数据获取错误",
		})
	}else {
		ctx.JSON(http.StatusOK,gin.H{
			"Code":1,
			"Message":"数据获取成功",
			"Data":result,
		})
	}
}
//根据id查找文章
func FindArticleById(ctx *gin.Context)  {
	articleid := ctx.Query("articleid")
	result,err :=Services.FindArticleByIdMapper(articleid)
	if err != nil||len(result)==0 {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"Code":0,
			"Message":"数据获取错误",
		})
	}else {
		ctx.JSON(http.StatusOK,gin.H{
			"Code":1,
			"Message":"数据获取成功",
			"Data":result,
		})
	}
}

//查找文章大体分类
func FindGeneralsort(ctx *gin.Context)  {
	result,err :=Services.FindAllGeneral()
	if err != nil||len(result)==0 {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"Code":0,
			"Message":"数据获取错误",
		})
	}else {
		ctx.JSON(http.StatusOK,gin.H{
			"Code":1,
			"Message":"数据获取成功",
			"Data":result,
		})
	}
}
//查找文章分类
func FindArticlesort(ctx *gin.Context)  {
	generalsort := ctx.Query("generalsort")
	fmt.Println(generalsort)
	result,err :=Services.FindAllArticleSort(generalsort)
	if err != nil||len(result)==0 {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"Code":0,
			"Message":"数据获取错误",
		})
	}else {
		ctx.JSON(http.StatusOK,gin.H{
			"Code":1,
			"Message":"数据获取成功",
			"Data":result,
		})
	}
}
//查找文章每个大体分类
func FindArticlexController(ctx *gin.Context)  {
	id:=ctx.Query("generalid")
	generalid,err :=strconv.Atoi(id)
	fmt.Println(id)
	result:=Services.FindArticlexServices(generalid)
	if err != nil||len(result)==0 {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"Code":0,
			"Message":"数据获取错误",
		})
	}else {
		ctx.JSON(http.StatusOK,gin.H{
			"Code":1,
			"Message":"数据获取成功",
			"Data":result,
		})
	}
}
//查找文章每个大体分类
func FindArticleBySortController(ctx *gin.Context)  {
	id:=ctx.Query("articlesort")
	articlesort,err :=strconv.Atoi(id)
	fmt.Println(id)
	result:=Services.FindArticleBySortService(articlesort)
	if err != nil||len(result)==0 {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest,gin.H{
			"Code":0,
			"Message":"数据获取错误",
		})
	}else {
		ctx.JSON(http.StatusOK,gin.H{
			"Code":1,
			"Message":"数据获取成功",
			"Data":result,
		})
	}
}