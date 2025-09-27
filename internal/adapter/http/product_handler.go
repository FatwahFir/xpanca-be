package httpadapter

import (
	"net/http"
	"strconv"

	"github.com/FatwahFir/xpanca-be/internal/dto"
	"github.com/FatwahFir/xpanca-be/internal/usecase"
	"github.com/FatwahFir/xpanca-be/pkg/response"
	"github.com/FatwahFir/xpanca-be/utils"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct{ UC *usecase.ProductUsecase }

func NewProductHandler(r gin.IRouter, uc *usecase.ProductUsecase) {
	h := &ProductHandler{UC: uc}
	r.GET("/products", h.List)
	r.GET("/products/:id", h.Detail)
}

func (h *ProductHandler) List(c *gin.Context) {
	q := dto.ProductQuery{
		Page:     utils.Atoi(c.DefaultQuery("page", "1")),
		PageSize: utils.Atoi(c.DefaultQuery("page_size", "10")),
		Name:     c.Query("name"),
		Category: c.Query("category"),
		Search:   c.Query("search"),
	}
	items, total, err := h.UC.Find(c, q)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, response.CodeServerError, err.Error(), nil)
		return
	}

	dtoItems := dto.ToProductListResponse(items)
	totalPages := (total + int64(q.PageSize) - 1) / int64(q.PageSize)

	response.Paginated(c, dtoItems, response.Pagination{
		Page: q.Page, PageSize: q.PageSize, Total: total, TotalPages: totalPages,
	})
}

func (h *ProductHandler) Detail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	p, err := h.UC.GetByID(c, uint(id))
	if err != nil {
		response.Err(c, http.StatusNotFound, response.CodeNotFound, "product not found", nil)
		return
	}

	response.OK(c, dto.ToProductResponse(*p))
}
