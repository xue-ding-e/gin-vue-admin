package initialize

import (
	"net/http"
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/docs"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrPermission
	}

	return f, nil
}

// 初始化总路由
func Routers() *fiber.App {
	app := fiber.New()
	app.Use(gin.Recovery())
	if global.GVA_CONFIG.System.Mode == "debug" {
		app.Use(logger.New())
	}
	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example
	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面3行注释
	// Router.Static("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/assets", "./dist/assets")   // dist里面的静态资源
	// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	// 静态文件处理，如原先的 Router.StaticFS
	// Fiber 推荐使用 app.Static(prefix, rootDir, fiber.Static{}) 这样的方法
	// 如果你需要只读或者禁止目录浏览，可以设置  fiber.Static{Browse: false} 等
	app.Static(global.GVA_CONFIG.Local.StorePath, global.GVA_CONFIG.Local.StorePath)

	// 跨域，如需跨域可以打开下面的注释
	//app.Use(middleware.Cors())        // 直接放行全部跨域请求
	//app.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	// global.GVA_LOG.Info("use middleware cors")
	// 如需 swagger，这里仅示例保留配置。Fiber 官方有 swagger 插件可替代
	docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
	app.Get(global.GVA_CONFIG.System.RouterPrefix+"/swagger/*", func(c *fiber.Ctx) error {
		// 此处需要替换为 fiber swagger 的实际用法
		// TODO swagger替换
		// 这里仅示例直接返回提示
		return c.SendString("请使用 Fiber 对应的 swagger 实现。")
	})
	global.GVA_LOG.Info("register swagger handler")

	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := app.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup := app.Group(global.GVA_CONFIG.System.RouterPrefix)

	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())

	{
		// 健康监测
		PublicGroup.Get("/health", func(c *fiber.Ctx) error {
			return c.Status(http.StatusOK).JSON("ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}

	{
		systemRouter.InitApiRouter(PrivateGroup, PublicGroup)       // 注册功能api路由
		systemRouter.InitJwtRouter(PrivateGroup)                    // jwt相关路由
		systemRouter.InitUserRouter(PrivateGroup)                   // 注册用户路由
		systemRouter.InitMenuRouter(PrivateGroup)                   // 注册menu路由
		systemRouter.InitSystemRouter(PrivateGroup)                 // system相关路由
		systemRouter.InitCasbinRouter(PrivateGroup)                 // 权限相关路由
		systemRouter.InitAutoCodeRouter(PrivateGroup, PublicGroup)  // 创建自动化代码
		systemRouter.InitAuthorityRouter(PrivateGroup)              // 注册角色路由
		systemRouter.InitSysDictionaryRouter(PrivateGroup)          // 字典管理
		systemRouter.InitAutoCodeHistoryRouter(PrivateGroup)        // 自动化代码历史
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)     // 操作记录
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)    // 字典详情管理
		systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)     // 按钮权限管理
		systemRouter.InitSysExportTemplateRouter(PrivateGroup)      // 导出模板
		systemRouter.InitSysParamsRouter(PrivateGroup, PublicGroup) // 参数管理
		exampleRouter.InitCustomerRouter(PrivateGroup)              // 客户路由
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由

	}

	//插件路由安装
	InstallPlugin(PrivateGroup, PublicGroup, Router)

	// 注册业务路由
	initBizRouter(PrivateGroup, PublicGroup)

	global.GVA_ROUTERS = app.Stack()

	global.GVA_LOG.Info("router register success")
	return app
}
