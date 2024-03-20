package order

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	OrderService Service
}

func NewHandler(s Service) *handler {
	return &handler{
		OrderService: s,
	}
}

func (h *handler) CreateOrder(ctx *gin.Context) {
	var req CreateOrderRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	res, err := h.OrderService.CreateOrder(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Order has successfully created.",
		"data":    res,
	})
}

func (h *handler) GetOrder(ctx *gin.Context) {
	res, err := h.OrderService.GetOrder()
	if err != nil {
		if err.Error() == "empty" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": "Order still empty.",
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Successfully retrieve orders",
		"data":    res,
	})
}

func (h *handler) DeleteOrder(ctx *gin.Context) {
	var req DeleteOrderRequest

	id, _ := strconv.Atoi(ctx.Param("id"))
	req.Id = uint(id)

	err := h.OrderService.DeleteOrder(&req)

	if err != nil {
		if err.Error() == "not found" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": "Order not found",
			})
			return
		} else if err.Error() == "null" {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Null record affected",
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Successfully delete order",
	})
}

func (h *handler) UpdateOrder(ctx *gin.Context) {
	var req UpdateOrderRequest

	id, _ := strconv.Atoi(ctx.Param("id"))

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	res, err := h.OrderService.UpdateOrder(uint(id), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Null record affected",
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"message": "Successfully delete order",
		"data":    res,
	})
}
