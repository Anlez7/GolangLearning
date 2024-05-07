package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//创建一个默认的路由引擎
	engine := gin.Default()
	//注册路由
	TestRecFiles(engine)
	engine.Run(":8080")
}

// 接收多个文件
func TestRecFiles(engine *gin.Engine) {
	// 设置内存限制为8M， 默认是32MiB
	engine.MaxMultipartMemory = 8 << 20
	engine.POST("/files", func(context *gin.Context) {
		// 接收图片
		form, _ := context.MultipartForm()
		files := form.File["imgList[]"]
		// 遍历保存
		for _, file := range files {
			_ = context.SaveUploadedFile(file, "./tmp/"+file.Filename)
		}
		context.String(200, "保存成功!")
	})
}
