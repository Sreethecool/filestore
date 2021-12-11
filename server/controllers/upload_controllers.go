package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Sreethecool/filestore/server/types"
	"github.com/labstack/echo/v4"
)

func Upload(c echo.Context) error {
	resp := types.Response{}
	errors := map[string]string{}
	form, err := c.MultipartForm()
	if err != nil {
		resp.Message = "Invalid Input File data"
		return c.JSON(http.StatusBadRequest, resp)
	}
	files := form.File["files"]
	for index, file := range files {
		src, err := file.Open()
		if err != nil {
			errors[file.Filename] = fmt.Sprintf("Cant read input file %d", index)
			continue
		}
		defer src.Close()

		path := "upload/"
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err := os.Mkdir(path, os.ModePerm)
			if err != nil {
				fmt.Println("Cant create upload folder")
				path = "./"
			}
		}

		dst, err := os.OpenFile(path+file.Filename, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Cant Open/create File", index, err.Error())
			resp.Message = "Error in Uploading"
			errors[file.Filename] = fmt.Sprintf("Cant create file file %d", index)
			continue
		}

		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			fmt.Println("Cant write to File", index, err.Error())
			resp.Message = "Error in Uploading"
			errors[file.Filename] = fmt.Sprintf("Cant create file file %d", index)
			continue
		}

	}

	if len(errors) > 0 {
		resp.Message = fmt.Sprintf("Error in Uploading %d files out of %d files", len(errors), len(files))
		resp.Data = errors
		return c.JSON(http.StatusAccepted, resp)
	}

	resp.Message = fmt.Sprintf("<p>Uploaded successfully %d files with field.</p>", len(files))
	return c.JSON(http.StatusOK, resp)
}
