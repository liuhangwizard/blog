package VBRouter

import (
	. "vb-server/controller"
	. "vb-server/middleware"
	. "vb-server/server"
)

func InitRouter() {

	//controller
	adminController := NewVBControllerAdmin()
	infoController := NewVBControllerInfo()
	blogController := NewVBControllerBlog()
	blogRecordController := NewVBControllerBlogRecord()
	tagController := NewVBControllerTag()
	tagRecordController := NewVBControllerTagRecord()
	catagoryController := NewVBControllerCatagory()
	imageController:=NewVBControllerImage()

	//mapping
	//admin
	Main.Use(VBMiddlewareLog)
	Main.POST("/logout", VBMiddlewareAuth, adminController.Logout)
	Main.POST("/login", adminController.Login)
	Main.GET("/ping", adminController.Ping)


	//permission
	permission := Main.Group("/info")
	permission.Use(VBMiddlewareAuth)
	{
		permission.POST("/user", infoController.GetUserInfo)
	}

	//blog
	blog := Main.Group("/blog")
	{
		blog.PUT("/status", VBMiddlewareAuth, blogController.UpdateStatus)
		blog.PUT("/change", VBMiddlewareAuth, blogController.Update)
		blog.POST("/new", VBMiddlewareAuth, blogController.New)
		blog.POST("/queryparam", VBMiddlewareAuth, blogController.QueryParam)

		blog.POST("/info", VBMiddlewareAuth, blogController.ManageInfo)

		//client article
		blog.POST("/artilceinfo", blogController.ClientArticleInfo)
		blog.POST("/artilcepropose", blogController.ClientArticlePropose)

		//client main
		blog.POST("/maincloud", blogController.ClientMainCloud)
		blog.POST("/mainrecent", blogController.ClientMainRecent)
	}

	//blog record
	blogRecord := Main.Group("/recordblog")
	{
		blogRecord.POST("/add",blogRecordController.RecordNew)
		blogRecord.POST("/info", VBMiddlewareAuth, blogRecordController.RecordInfo)
		blogRecord.POST("/queryparam", VBMiddlewareAuth, blogRecordController.QueryParam)
		blogRecord.POST("/maincount", blogRecordController.ClientMainCount)
		blogRecord.GET("/chartblog", VBMiddlewareAuth, blogRecordController.RecordChartBlog)
	}

	//catagory
	catagory := Main.Group("/catagory")
	{
		catagory.POST("/all", catagoryController.GetCatagoryList)
		catagory.POST("/new", VBMiddlewareAuth, catagoryController.New)
		catagory.PUT("/remove", VBMiddlewareAuth, catagoryController.Delete)
	}

	//tag
	tag := Main.Group("/tag")
	{
		tag.POST("/all", tagController.GetTagList)
		tag.POST("/blog", tagController.GetTagBlogList)
		tag.POST("/new", VBMiddlewareAuth, tagController.New)
		tag.PUT("/remove", VBMiddlewareAuth, tagController.Delete)
	}

	//tag record
	tagRecord := Main.Group("/recordtag")
	{
		tagRecord.GET("/list", VBMiddlewareAuth, tagRecordController.RecordTagList)
	}

	//image
	imageOSS:=Main.Group("/image")
	imageOSS.Use(VBMiddlewareAuth)
	{
		imageOSS.POST("/url",imageController.UploadSign)
		imageOSS.POST("/type",imageController.UpdateType)
		imageOSS.POST("",imageController.QueryParam)
		imageOSS.PUT("",imageController.New)


	}
}
