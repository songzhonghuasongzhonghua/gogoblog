package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/songzhonghuasongzhonghua/gogoblog/global"
	"github.com/songzhonghuasongzhonghua/gogoblog/utils"
	"path"
)

type ImagesResponse struct {
	FileName  string `json:"file_name"`
	Msg       string `json:"msg"`
	IsSuccess bool   `json:"is_success"`
}

func (ImagesApi) ImagesUpload(c *gin.Context) {
	imageFrom, err := c.MultipartForm()
	if err != nil {
		utils.Failed(c, err.Error())
		return
	}
	filesList, ok := imageFrom.File["images"]
	if !ok {
		utils.Failed(c, "不存在的文件")
		return
	}
	//判断文件大小
	//判断文件路径是否相同

	var imagesRes []ImagesResponse

	for _, file := range filesList {

		filePath := path.Join("uploads", file.Filename)
		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			//append(imagesRes, ImagesResponse{
			//	FileName: file.Filename,
			//	Msg:      fmt.Printf("文件大小超过%dMB", global.Config.Uploads.Size),
			//})
			continue
		}
		fileSize := float64(file.Size) / float64(1024*1024)
		if fileSize >= float64(global.Config.Uploads.Size) {
			imagesRes = append(imagesRes, ImagesResponse{
				FileName:  file.Filename,
				Msg:       fmt.Sprintf("文件大小超过%dMB", global.Config.Uploads.Size),
				IsSuccess: false,
			})
			continue
		}

		imagesRes = append(imagesRes, ImagesResponse{
			FileName:  file.Filename,
			Msg:       "文件上传成功",
			IsSuccess: true,
		})

	}

	utils.Success(c, gin.H{
		"files": imagesRes,
	})

}
