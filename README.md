# Modeltools
#### GO语言连接Mysql生成对应的model，包括对应字段类型、注释等。生成基础的结构体，不局限于某一个ORM。
 
  **源码码地址---------**
  ##### github：[https://github.com/longzongqin/modeltools](https://github.com/longzongqin/modeltools)
  ##### 码云：[https://gitee.com/longzongqin/modeltools](https://gitee.com/longzongqin/modeltools)
 
 

 **生成示例---------**

```go 
  package models

  // 管理员表
  type AdminInfo struct {
  	Id int `json:"id"` 
  	Name string `json:"name"` // 姓名
  	Username string `json:"username"` // 用户名 
  	Password string `json:"password"` // 密码
  	RoleInfoId int `json:"role_info_id"` // 角色ID
  	Status int8 `json:"status"` // -1删除，0正常，1禁用
  }
```

**参数配置--------conf.go**

```go 
  package conf
  
  // model保存路径
  const ModelPath = "./models/"
  // 是否覆盖已存在model
  const ModelReplace = true
  // 数据库驱动
  const DriverName = "mysql"
  
  type DbConf struct {
  	Host   string
  	Port   string
  	User   string
  	Pwd    string
  	DbName string
  }
  // 数据库链接配置
  var MasterDbConfig DbConf = DbConf{
  	Host:   "127.0.0.1",
  	Port:   "3306",
  	User:   "root",
  	Pwd:    "long",
  	DbName: "mvideo",
  }
```

**生成model--------**
```go
package main

import (
	"modeltools/dbtools"
	"modeltools/generate"
)


func main() {
	//初始化数据库
	dbtools.Init()
	//generate.Genertate() //生成所有表信息
	generate.Genertate("admin_info","video_info") //生成指定表信息，可变参数可传入多个表名
}


```