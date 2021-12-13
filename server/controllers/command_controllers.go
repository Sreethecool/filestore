package controllers

import (
	"bytes"
	"html/template"
	"net/http"
	"strings"

	"github.com/Sreethecool/filestore/server/command"
	"github.com/Sreethecool/filestore/server/models"
	"github.com/labstack/echo/v4"
)

func List(c echo.Context) error {
	resp := models.Response{}

	cmd := "ls upload/"
	res, err := command.Execute(cmd)
	if err != nil {
		resp.Message = "Unable to get the list"
		c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Message = "List of files success"
	resp.Data = res
	return c.JSON(http.StatusOK, resp)
}

func Delete(c echo.Context) error {
	var req models.DeleteRequest
	var resp models.Response
	if err := c.Bind(&req); err != nil {
		resp.Message = "Invalid Data"
		return c.JSON(http.StatusBadRequest, resp)
	}

	cmd := "rm upload/" + req.Filename

	_, err := command.Execute(cmd)
	if err != nil {
		resp.Message = "Unable to delete the file"
		c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Message = "File deleted successfully"
	return c.JSON(http.StatusOK, resp)
}

func Execute(c echo.Context) error {
	var req models.ExecuteRequest
	var resp models.Response
	if err := c.Bind(&req); err != nil {
		resp.Message = "Invalid Data"
		return c.JSON(http.StatusBadRequest, resp)
	}

	temp := ""
	found := false
	cmd := strings.ToLower(req.Cmd)
	if temp, found = models.CmdTemplate[cmd]; !found {
		resp.Message = "Invalid Request"
		return c.JSON(http.StatusBadRequest, resp)
	}
	t := template.Must(template.New("command").Parse(temp))
	param := map[string]string{
		"folder": "upload/",
		"count":  "10",
		"order":  "head",
	}

	for i := 0; i < len(req.Args); i++ {
		if req.Args[i] == "--limit" || req.Args[i] == "-n" && i < len(req.Args)-1 {
			param["count"] = req.Args[i+1]
			i++
		} else if req.Args[i] == "--order" && i < len(req.Args)-1 {
			param["order"] = req.Args[i+1]
			i++
		}
	}
	var out bytes.Buffer
	err := t.Execute(&out, param)
	if err != nil {
		resp.Message = "Unable to process request"
		c.JSON(http.StatusInternalServerError, resp)
	}

	res, err := command.Execute(out.String())
	if err != nil {
		resp.Message = "Unable to process request"
		c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Message = "request processed"
	resp.Data = res
	return c.JSON(http.StatusOK, resp)
}
