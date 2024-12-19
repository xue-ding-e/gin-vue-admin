package ast

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

func TestPluginInitializeData_Injection(t *testing.T) {
	type fields struct {
		Type             Type
		Path             string
		ImportPath       string
		StructName       string
		PackageName      string
		FunctionName     string
		DataFunctionName string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "测试插件初始化数据注入",
			fields: fields{
				Type:             TypePluginInitializeData,
				Path:             filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "LocationInfo", "initialize", "data.go"),
				ImportPath:       fmt.Sprintf(`"%s/plugin/LocationInfo/initialize/data"`, global.GVA_CONFIG.AutoCode.Module),
				StructName:       "",     // 不需要结构体名
				PackageName:      "data", // 包名
				FunctionName:     "Data", // 调用的函数名
				DataFunctionName: "Data", // 被调用的数据初始化函数名
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &PluginInitializeData{
				Type:             tt.fields.Type,
				Path:             tt.fields.Path,
				ImportPath:       tt.fields.ImportPath,
				StructName:       tt.fields.StructName,
				PackageName:      tt.fields.PackageName,
				FunctionName:     tt.fields.FunctionName,
				DataFunctionName: tt.fields.DataFunctionName,
			}
			file, err := a.Parse(a.Path, nil)
			if err != nil {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			a.Injection(file)
			err = a.Format(a.Path, nil, file)
			if (err != nil) != tt.wantErr {
				t.Errorf("Injection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPluginInitializeData_Rollback(t *testing.T) {
	type fields struct {
		Type             Type
		Path             string
		ImportPath       string
		StructName       string
		PackageName      string
		FunctionName     string
		DataFunctionName string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "测试插件初始化数据回滚",
			fields: fields{
				Type:             TypePluginInitializeData,
				Path:             filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", "LocationInfo", "initialize", "data.go"),
				ImportPath:       fmt.Sprintf(`"%s/plugin/LocationInfo/initialize/data"`, global.GVA_CONFIG.AutoCode.Module),
				StructName:       "",
				PackageName:      "data",
				FunctionName:     "Data",
				DataFunctionName: "Data",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &PluginInitializeData{
				Type:             tt.fields.Type,
				Path:             tt.fields.Path,
				ImportPath:       tt.fields.ImportPath,
				StructName:       tt.fields.StructName,
				PackageName:      tt.fields.PackageName,
				FunctionName:     tt.fields.FunctionName,
				DataFunctionName: tt.fields.DataFunctionName,
			}
			file, err := a.Parse(a.Path, nil)
			if err != nil {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			a.Rollback(file)
			err = a.Format(a.Path, nil, file)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rollback() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
