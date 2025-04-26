package core

import (
	"flag"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"os"
	"path/filepath"

	"github.com/flipped-aurora/gin-vue-admin/server/core/internal"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

/*
打印日志不要使用全局global.GVA_LOG因为此时还没加载
*/

// Viper 配置
func Viper() *viper.Viper {
	config := getConfigPath()

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
		panic(fmt.Errorf("fatal error unmarshal config: %w", err))
	}

	//mergeConfigPath := "生产环境所需/config"
	//loadAdditionalConfig(mergeConfigPath, &global.GVA_CONFIG)
	// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}

// 递归读取指定路径下的所有*.yaml文件
func loadAdditionalConfig(mergeConfigPath string, config *config.Server) {
	// 定义支持的配置文件扩展名
	supportedExts := []string{".yaml", ".yml"}

	// 递归遍历目录
	err := filepath.Walk(mergeConfigPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 获取文件扩展名
		ext := filepath.Ext(path)

		// 检查是否是支持的配置文件类型
		isSupported := false
		for _, supportedExt := range supportedExts {
			if ext == supportedExt {
				isSupported = true
				break
			}
		}

		// 只处理支持的配置文件类型
		if !info.IsDir() && isSupported {
			v := viper.New()
			v.SetConfigFile(path)
			v.SetConfigType(ext[1:]) // 去掉点号，只保留扩展名部分
			err := v.ReadInConfig()
			if err != nil {
				panic(fmt.Errorf("读取配置文件 %s 失败: %s", path, err))
			}
			v.WatchConfig()

			v.OnConfigChange(func(e fsnotify.Event) {
				fmt.Println("config file changed:", e.Name)
				if err = v.Unmarshal(config); err != nil {
					fmt.Println(err)
				}
			})

			if err = v.Unmarshal(config); err != nil {
				panic(fmt.Errorf("fatal error unmarshal config: %w", err))
			}
			fmt.Printf("成功加载配置文件: %s\n", path)
		}
		return nil
	})

	if err != nil {
		panic(fmt.Errorf("遍历配置文件目录失败: %s\n", err))
	}
}

// getConfigPath 获取配置文件路径, 优先级: 命令行 > 环境变量 > 默认值
func getConfigPath() (config string) {
	// `-c` flag parse
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()
	if config != "" { // 命令行参数不为空 将值赋值于config
		fmt.Printf("您正在使用命令行的 '-c' 参数传递的值, config 的路径为 %s\n", config)
		return
	}
	if env := os.Getenv(internal.ConfigEnv); env != "" { // 判断环境变量 GVA_CONFIG
		config = env
		fmt.Printf("您正在使用 %s 环境变量, config 的路径为 %s\n", internal.ConfigEnv, config)
		return
	}

	switch gin.Mode() { // 根据 gin 模式文件名
	case gin.DebugMode:
		config = internal.ConfigDebugFile
	case gin.ReleaseMode:
		config = internal.ConfigReleaseFile
	case gin.TestMode:
		config = internal.ConfigTestFile
	}
	fmt.Printf("您正在使用 gin 的 %s 模式运行, config 的路径为 %s\n", gin.Mode(), config)

	_, err := os.Stat(config)
	if err != nil || os.IsNotExist(err) {
		config = internal.ConfigDefaultFile
		fmt.Printf("配置文件路径不存在, 使用默认配置文件路径: %s\n", config)
	}

	return
}
