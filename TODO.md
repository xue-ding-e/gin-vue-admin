## TODO

各个部分拆分成插件,合并、删除重复的业务,待删除9528角色(测试角色)
去掉UUID替换为雪花ID(注释掉 设置为待用状态而不是一定要使用 有需要再根据文档使用)

C端用户默认api 和默认菜单

go-pay微信插件初始化必须绑定所有所需信息

etcd插件初始化信息(链接信息panic等) , 获取锁等信息将fmt.Errorf 为 global.GVA_LOG.中的内容

common封装登录返回的信息( Login去掉密码验证)

支付插件

个人信息默认资料



IgnoreRecordNotFoundError

xuedinge 插件

登录admin 默认username去掉

b.TokenNext放到common

前端随机生成字符串库

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