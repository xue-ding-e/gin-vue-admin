## 已完成

###  2024-12-27

NickName -> Nickname

### 2024-12-26

web  api请求加入baseurl

待删除9528角色(测试角色)

### 2024-11-30更新内容

昵称统一字段为nickname 前后端和数据库

uni navigator 移动到utils下  navigation文件改名navigator

后端wx 创建用户增加appid

修改CommonUser 的Authorities 外键 gorm标签指定名称  以及替换SysAuthority 中的Users为CommonUser

wx支付插件 service G微信小程序支付业务逻辑scan改first

utils 和 model/system 循环导包

Hasura以及后端常用辅助中间件docker-compose.yaml

Business 拼写问题 全局替换


### 2024-11-21更新内容

etcd插件初始化链接失败增加panic

去掉UUID替换为雪花ID(注释掉 设置为待用状态而不是一定要使用 有需要再根据文档使用)

登录admin 默认username去掉

中文路径编码

### 2024-11-21更新内容

登录admin 默认username去掉

```
item.path = encodeURIComponent(item.path);  // 路径中文编码
url

const formatRouter = (routes, routeMap, parent) => {
  routes && routes.forEach(item => {
    item.parent = parent
    item.meta.btns = item.btns
    item.meta.hidden = item.hidden
    item.path = encodeURIComponent(item.path);  // 路径中文编码
    if (item.meta.defaultMenu === true) {
      if (!parent) {
        item = { ...item, path: `/${item.path}` }
        notLayoutRouterArr.push(item)
      }
    }
    routeMap[item.name] = item
    if (item.children && item.children.length > 0) {
      formatRouter(item.children, routeMap, item)
    }
  })
}
```


## TODO

### 本体

gin-mode生产环境

#### 急!!!!!

web原生添加Pinia 持久化

添加GVA_MODEL2(去掉ID)

web增加xInputCents组件

```
 不显示 record not found 错误
 
return &gorm.Config{
		Logger: logger.New(NewWriter(general, log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  general.LogLevel(),
			IgnoreRecordNotFoundError: true, // 不显示 record not found 错误
			Colorful:                  true,
		}),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
```

web wallet 金额;

getUrl增加的base64图片的判断

所有json的userName改为username

自动化代码,web template , form增加default 值(closeDialog等直接使用默认空值) 

tokenNext 和 clams改造方便增删(下次遇到立刻更改)  遇到了 例如authority等信息

C端用户默认api 和默认菜单(下次遇到立刻更改)

common封装登录返回的信息( Login去掉密码验证)  ( 时不知道什么意思 忘了 下次遇到立刻更改)

个人信息默认资料(暂时不知道什么意思 忘了  个人信息例如性别 昵称 国家这些 有必要的相关内容改为指针nil) (下次遇到立刻更改)

原生组件整理(现在有很多错误)

自动化代码注册,增加权限

#### 不急

后端生成gorm的时候调换顺序,将gorm的注释comment放到最前面

menu底部横向滑动栏,固定位置 以及其他横向滑动栏检查 (或者始终将滑动栏固定到顶部或者底部fiexd等样式)

web端增加minio  sts策略

前端自动化代码增加一键展开所有(树形结构自己重构 , 重构为前端构建而不是后端构建)增加去掉分页

python 重置postgresql游标改为golang

插件导入数据增加.sql的形式 , 增加插件自动初始化权限(菜单 , api) 的代码

server dockerfile  gva原生更新到xuedinge分支

build状态删除console.log

各个部分拆分成插件,合并、删除重复的业务

