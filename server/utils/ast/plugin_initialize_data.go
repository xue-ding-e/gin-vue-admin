package ast

import (
	"go/ast"
	"io"
)

// PluginDataData 用于生成 initialize/data/data.go 文件
type PluginInitializeData struct {
	Base
	Type             Type   // 类型
	Path             string // 文件路径
	ImportPath       string // 导包路径
	RelativePath     string // 相对路径
	StructName       string // 结构体名称
	PackageName      string // 包名
	FunctionName     string // 函数名
	DataFunctionName string // 数据初始化函数名
}

func (a *PluginInitializeData) Parse(filename string, writer io.Writer) (file *ast.File, err error) {
	if filename == "" {
		if a.RelativePath == "" {
			filename = a.Path
			a.RelativePath = a.Base.RelativePath(a.Path)
			return a.Base.Parse(filename, writer)
		}
		a.Path = a.Base.AbsolutePath(a.RelativePath)
		filename = a.Path
	}
	return a.Base.Parse(filename, writer)
}

// func (a *PluginInitializeData) Parse(filename string, writer io.Writer) (file *ast.File, err error) {
// 	if filename == "" {
// 		filename = a.Path
// 	}

// 	// 如果文件不存在，则创建文件并写入基础代码
// 	if _, err := os.Stat(filename); os.IsNotExist(err) {
// 		err = os.MkdirAll(filepath.Dir(filename), os.ModePerm)
// 		if err != nil {
// 			return nil, err
// 		}
// 		fileContent := fmt.Sprintf(`package %s

// import "github.com/flipped-aurora/gin-vue-admin/server/global"

// func %s() {
//     // 实际的数据初始化逻辑
//     global.GVA_LOG.Info("Initializing plugin data...")
//     // TODO: 添加数据初始化代码
// }
// `, a.PackageName, a.FunctionName)
// 		err = os.WriteFile(filename, []byte(fileContent), 0644)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return a.Base.Parse(filename, writer)
// }

func (a *PluginInitializeData) Injection(file *ast.File) error {
	// 初始模板文件，无需修改
	return nil
}

func (a *PluginInitializeData) Rollback(file *ast.File) error {
	// 初始模板文件，无需回滚
	return nil
}

func (a *PluginInitializeData) Format(filename string, writer io.Writer, file *ast.File) error {
	if filename == "" {
		filename = a.Path
	}
	return a.Base.Format(filename, writer, file)
}
