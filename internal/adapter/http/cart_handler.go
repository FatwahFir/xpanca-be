package httpadapter

import (
	"net/http"
	"strconv"

	"github.com/FatwahFir/xpanca-be/internal/dto"
	"github.com/FatwahFir/xpanca-be/internal/usecase"
	"github.com/FatwahFir/xpanca-be/pkg/response"
	"github.com/gin-gonic/gin"
)

type CartHandler struct{ UC *usecase.CartUsecase }

func NewCartHandler(r gin.IRouter, uc *usecase.CartUsecase) {
	h := &CartHandler{UC: uc}
	grp := r.Group("/cart") // <- r di sini sudah authGroup
	grp.GET("", h.Get)
	grp.POST("/add", h.Add)
	grp.POST("/:pid/inc", h.Inc)
	grp.POST("/:pid/dec", h.Dec)
	grp.DELETE("/:pid", h.Remove)
}

func getPID(ID string) (uint, error) {
	pid64, err := strconv.ParseUint(ID, 10, 0)
	if err != nil {
		return 0, err
	}
	pid := uint(pid64)
	return pid, nil
}

func getUserID(c *gin.Context) uint {
	v, _ := c.Get("user_id")
	if id, ok := v.(uint); ok {
		return id
	}

	if id64, ok := v.(uint64); ok {
		return uint(id64)
	}

	if s, ok := v.(string); ok {
		if u, err := strconv.ParseUint(s, 10, 64); err == nil {
			return uint(u)
		}
	}

	return 0
}

func (h *CartHandler) Get(c *gin.Context) {
	uid := getUserID(c)
	cart, err := h.UC.Get(c, uid)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, response.CodeNotFound, "cart not found", err.Error())
		return
	}

	var items []dto.CartItemResponse
	var subtotal int64
	for _, item := range cart.Items {
		var thumb *string
		if len(item.Product.Images) > 0 {
			thumb = &item.Product.Images[0].URL
		}
		line := int64(item.Qty) * item.Product.Price
		items = append(items, dto.CartItemResponse{
			ProductID: item.ProductID,
			Name:      item.Product.Name,
			Category:  item.Product.Category,
			Price:     item.Product.Price,
			Thumbnail: thumb,
			Qty:       item.Qty,
			LineTotal: line,
		})
		subtotal += line
	}
	resp := dto.CartResponse{
		Items:    items,
		Subtotal: subtotal,
		Count:    len(items),
	}

	response.OK(c, resp)
}

func (h *CartHandler) Add(c *gin.Context) {
	var req dto.CartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Err(c, http.StatusBadRequest, response.CodeBadRequest, "invalid payload", err.Error())
		return
	}
	if req.Qty == 0 {
		req.Qty = 1
	}
	uid := getUserID(c)
	if err := h.UC.Add(c, uid, req.ProductID, req.Qty); err != nil {
		response.Err(c, http.StatusInternalServerError, response.CodeServerError, "Internal server error", err.Error())
		return
	}
	response.OK(c, gin.H{"message": "added"})
}

func (h *CartHandler) Inc(c *gin.Context) {
	pid, err := getPID(c.Param("pid"))
	if err != nil {
		response.Err(c, http.StatusBadRequest, response.CodeBadRequest, "invalid product id", err.Error())
		return
	}
	uid := getUserID(c)
	if err := h.UC.Inc(c, uid, pid); err != nil {
		response.Err(c, http.StatusInternalServerError, response.CodeNotFound, "No data found", err.Error())
		return
	}
	response.OK(c, gin.H{"message": "increased"})
}

func (h *CartHandler) Dec(c *gin.Context) {
	pid, err := getPID(c.Param("pid"))
	if err != nil {
		response.Err(c, http.StatusBadRequest, response.CodeBadRequest, "invalid product id", err.Error())
		return
	}
	uid := getUserID(c)
	if err := h.UC.Dec(c, uid, pid); err != nil {
		response.Err(c, http.StatusInternalServerError, response.CodeBadRequest, "Invalid payload", err.Error())
		return
	}
	response.OK(c, gin.H{"message": "decreased"})
}

func (h *CartHandler) Remove(c *gin.Context) {
	pid, err := getPID(c.Param("pid"))
	if err != nil {
		response.Err(c, http.StatusBadRequest, response.CodeBadRequest, "invalid product id", err.Error())
		return
	}
	uid := getUserID(c)
	if err := h.UC.Remove(c, uid, pid); err != nil {
		response.Err(c, http.StatusInternalServerError, response.CodeBadRequest, "Invalid payload", err.Error())
		return
	}
	response.OK(c, gin.H{"message": "removed"})
}
