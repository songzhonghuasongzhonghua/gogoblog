package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/songzhonghuasongzhonghua/gogoblog/global"
	"github.com/songzhonghuasongzhonghua/gogoblog/models"
	"github.com/songzhonghuasongzhonghua/gogoblog/utils"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"strings"
)

type ImagesResponse struct {
	FileName  string `json:"file_name"`
	Msg       string `json:"msg"`
	IsSuccess bool   `json:"is_success"`
}

var (
	ImagesWhiteList = []string{
		"jpg",
		"png",
		"gif",
		"webp",
		"svg",
	}
)

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

	var imagesRes []ImagesResponse

	basePath := global.Config.Uploads.Path
	//判断文件保存路径是否存在
	_, err = os.ReadDir(basePath)
	if err != nil {
		//不存在则重新创建
		err = os.Mkdir(basePath, fs.ModePerm)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	for _, file := range filesList {

		//判断文件格式是否为图片格式
		nameList := strings.Split(file.Filename, ".")
		suffixName := strings.ToLower(nameList[len(nameList)-1])
		if !utils.InList(suffixName, ImagesWhiteList) {
			imagesRes = append(imagesRes, ImagesResponse{
				FileName:  file.Filename,
				Msg:       "文件格式不正确",
				IsSuccess: false,
			})

			continue
		}

		//判断文件大小
		filePath := path.Join(basePath, file.Filename)
		fileSize := float64(file.Size) / float64(1024*1024)
		if fileSize >= float64(global.Config.Uploads.Size) {
			imagesRes = append(imagesRes, ImagesResponse{
				FileName:  file.Filename,
				Msg:       fmt.Sprintf("文件大小超过%dMB", global.Config.Uploads.Size),
				IsSuccess: false,
			})
			continue
		}

		//判断文件是否保存过(hash值判断)
		fileObj, err := file.Open()
		fileData, err := io.ReadAll(fileObj)
		if err != nil {
			log.Fatal(err)
		}

		fileHash := utils.MD5(fileData)

		banner := models.BannerModel{}
		err = global.DB.Take(&banner, "hash = ?", fileHash).Error
		//文件已存在
		if err == nil {
			imagesRes = append(imagesRes, ImagesResponse{
				FileName:  banner.Path,
				Msg:       "文件已存在",
				IsSuccess: false,
			})
			continue
		}

		//保存文件至项目
		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			imagesRes = append(imagesRes, ImagesResponse{
				FileName:  file.Filename,
				Msg:       err.Error(),
				IsSuccess: false,
			})
			continue
		}
		imagesRes = append(imagesRes, ImagesResponse{
			FileName:  filePath,
			Msg:       "文件上传成功",
			IsSuccess: true,
		})
		//保存至数据库
		global.DB.Create(&models.BannerModel{
			Hash: fileHash,
			Path: filePath,
			Name: file.Filename,
		})
	}

	utils.Success(c, gin.H{
		"files": imagesRes,
	})

}
