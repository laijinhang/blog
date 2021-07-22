package route

import (
	"blog/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func init() {
	router := httprouter.New()
	blogRoute(router)
	userRoute(router)
	http.ListenAndServe(":8080", router)
}

// 博客
func blogRoute(router *httprouter.Router) {
	var archiverController controller.Archiver
	// 文章
	router.GET("/index/archives/:id", archiverController.Get)
	// 搜索
	router.GET("/index/search/:keyword", archiverController.Search)
	// 分页
	router.GET("/index/page/:page", archiverController.GetList)
	// 文章点击排名
	router.GET("/index/ranking", archiverController.Ranking)
	// 新增文章

	// 编辑文章

	// 删除文章
	// 保存为草稿
	// 上传图片

}

// 用户
func userRoute(router *httprouter.Router) {
	var userController controller.User
	// 用户信息
	router.GET("/index/user", userController.Info)
	// 登录
	router.POST("/index/user", userController.Login)
	// 注销
	router.DELETE("/index/user", userController.Logout)
	// 更新
	router.PUT("/index/user", userController.Update)
}
