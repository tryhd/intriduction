package helpers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func FileUpload(c *gin.Context, path string, formFile string) string {

	c.Request.ParseMultipartForm(32 << 20)

	// file, handler, err := c.Request.FormFile("image")
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, err.Error())
	// }
	// defer file.Close()
	// f, err := os.OpenFile(path+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, err.Error())
	// }
	// defer f.Close()
	// io.Copy(f, file)

	// return handler

	file, err := c.FormFile(formFile)

	// The file cannot be received.
	if err != nil {

		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension

	// The file is received, so let's save it
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	if err := c.SaveUploadedFile(file, mydir+path+newFileName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
	return newFileName

}
