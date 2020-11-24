package main

import (
	"bytes"
	"fmt"

	"github.com/spf13/afero"

	"goshop/front-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//初始化配置文件
func InitConfig() {
	buf := &utils.Config{}
	UnmarshalYaml(fmt.Sprintf("./conf/%s.app.yaml", gin.Mode()), buf)
	utils.C = buf

	//app解析
	baseInfo := &utils.Base{}
	UnmarshalYaml("./conf/app.yaml", baseInfo)
	utils.C.Base = baseInfo
}

func UnmarshalYaml(fileName string, data interface{}) {
	v := viper.New()
	v.SetConfigType("yaml")

	b, err := afero.ReadFile(afero.NewOsFs(), fileName)
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败, err: %v", err))
	}
	r := bytes.NewBuffer(b)

	if utils.FileExists("./conf/local.app.yaml") {
		localData, err := afero.ReadFile(afero.NewOsFs(), "./conf/local.app.yaml")
		if err != nil {
			panic(fmt.Errorf("读取local配置文件失败, err: %v", err))
		}
		r.Write([]byte("\n"))
		r.Write(localData)
	}

	if err := v.ReadConfig(r); err != nil {
		panic(fmt.Errorf("读取配置文件失败, err: %v", err))
	}

	if err := v.Unmarshal(data); err != nil {
		panic(fmt.Sprintf("解析配置文件出错, err: %v", err))
	}

}
