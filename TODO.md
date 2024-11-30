## TODO

### 本体

tokenNext 和 clams改造方便增删(下次遇到立刻更改)

C端用户默认api 和默认菜单(下次遇到立刻更改)

common封装登录返回的信息( Login去掉密码验证)  ( 时不知道什么意思 忘了 下次遇到立刻更改)

个人信息默认资料(暂时不知道什么意思 忘了) (下次遇到立刻更改)



各个部分拆分成插件,合并、删除重复的业务,待删除9528角色(测试角色)

### xuedinge 插件

uni 微信一键build上传

uni 退出登录  (注释跳转) 清楚本地个人信息（并刷新) 登录页面微信一键登录

wx支付插件 service G微信小程序支付业务逻辑scan改first

uni 商城

uni 个人地址

uni 前端角色id -> 角色名字 map utils

uni navigator 移动到utils下  navigation文件改名navigator

nick_name 前后端

uni 永久存储

utils 和 model/system 循环导包

go-pay微信插件初始化必须绑定所有所需信息

支付插件



IgnoreRecordNotFoundError

b.TokenNext放到common(待考虑)

后端角色id和名称对应字典(待考虑)

复制两个结构体之间对应的字段

```
import (
    "fmt"
    "github.com/jinzhu/copier"
)

type User struct {
    Name     string
    Age      int
    Location string
}

type Person struct {
    Name string
    Age  int
    City string
}

func main() {
    user := User{
        Name:     "张三",
        Age:      30,
        Location: "北京",
    }

    var person Person
    copier.Copy(&person, &user)

    fmt.Printf("%+v\n", person)
    // 输出: {Name:张三 Age:30 City:}
}
```

## 已完成

### 2024-11-30更新内容

Hasura以及后端常用辅助中间件docker-compose.yaml



### 2024-11-21更新内容

etcd插件初始化链接失败增加panic

去掉UUID替换为雪花ID(注释掉 设置为待用状态而不是一定要使用 有需要再根据文档使用)

登录admin 默认username去掉

中文路径编码

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
