## 已完成

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

tokenNext 和 clams改造方便增删(下次遇到立刻更改)  遇到了 例如authority等信息

C端用户默认api 和默认菜单(下次遇到立刻更改)

common封装登录返回的信息( Login去掉密码验证)  ( 时不知道什么意思 忘了 下次遇到立刻更改)

个人信息默认资料(暂时不知道什么意思 忘了  个人信息例如性别 昵称 国家这些 有必要的相关内容改为指针nil) (下次遇到立刻更改)

build状态删除console.log

各个部分拆分成插件,合并、删除重复的业务,待删除9528角色(测试角色)