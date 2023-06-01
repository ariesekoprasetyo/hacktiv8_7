package controllers

import (
	"github.com/ariesekoprasetyo/hacktiv8_7/orders"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ApiResponse struct {
	Code          int         `json:"code"`
	Status        string      `json:"status"`
	Message       string      `json:"message"`
	MessageDetail string      `json:"message_detail"`
	Data          interface{} `json:"data"`
}

func OrderPost(c *gin.Context) {
	var err error
	var orderReq orders.Orders
	if err = c.ShouldBindJSON(&orderReq); err != nil {
		response := ApiResponse{
			Code:          http.StatusBadRequest,
			Status:        "Gagal Bind Json",
			Message:       "Gagal Menambahkan Order",
			MessageDetail: err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if err = orders.CreateOrder(orderReq); err != nil {
		response := ApiResponse{
			Code:          http.StatusBadRequest,
			Status:        "Gagal",
			Message:       "Gagal Menambahkan Order",
			MessageDetail: err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	c.JSON(http.StatusOK, ApiResponse{
		Code:          http.StatusOK,
		Status:        "Berhasil",
		Message:       "Berhasil Menambahkan Order",
		MessageDetail: "",
		Data:          orderReq,
	})
}

func OrderGetAllData(c *gin.Context) {
	result, err := orders.GetAllData()
	if err != nil {
		response := ApiResponse{
			Code:          http.StatusBadRequest,
			Status:        "Gagal",
			Message:       "Gagal Get All Data Order",
			MessageDetail: err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	c.JSON(http.StatusOK, ApiResponse{
		Code:          http.StatusOK,
		Status:        "Berhasil",
		Message:       "Berhasil Get All Data Order",
		MessageDetail: "",
		Data:          result,
	})
}

func OrderGetDataById(c *gin.Context) {
	id := c.Param("id")
	num, _ := strconv.ParseUint(id, 10, 64)
	result, err := orders.GetDataById(uint(num))
	if err != nil {
		response := ApiResponse{
			Code:          http.StatusBadRequest,
			Status:        "Gagal",
			Message:       "Gagal Get Data Order By Id",
			MessageDetail: err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	c.JSON(http.StatusOK, ApiResponse{
		Code:          http.StatusOK,
		Status:        "Berhasil",
		Message:       "Berhasil Get Data By Id",
		MessageDetail: "",
		Data:          result,
	})
}

func OrderUpdate(c *gin.Context) {
	var orderUpdate orders.Update
	if err := c.ShouldBindJSON(&orderUpdate); err != nil {
		response := ApiResponse{
			Code:          http.StatusBadRequest,
			Status:        "Gagal Bind Json",
			Message:       "Gagal Update Order",
			MessageDetail: err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	id := c.Param("id")
	num, _ := strconv.ParseUint(id, 10, 64)
	result, err := orders.UpdateOrder(uint(num), orderUpdate)
	if err != nil {
		response := ApiResponse{
			Code:          http.StatusBadRequest,
			Status:        "Gagal",
			Message:       "Gagal Menambahkan Order",
			MessageDetail: err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	c.JSON(http.StatusOK, ApiResponse{
		Code:          http.StatusOK,
		Status:        "Berhasil",
		Message:       "Berhasil Update Order",
		MessageDetail: "",
		Data:          result,
	})
}

func OrderDelete(c *gin.Context) {
	id := c.Param("id")
	num, _ := strconv.ParseUint(id, 10, 64)
	err := orders.DeleteOrder(uint(num))
	if err != nil {
		response := ApiResponse{
			Code:          http.StatusBadRequest,
			Status:        "Gagal",
			Message:       "Gagal Menghapus Order",
			MessageDetail: err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	c.JSON(http.StatusOK, ApiResponse{
		Code:          http.StatusOK,
		Status:        "Berhasil",
		Message:       "Berhasil Mehapus Order",
		MessageDetail: "",
		Data:          nil,
	})
}
