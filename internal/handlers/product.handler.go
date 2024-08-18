package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"backend_coffeeShop.go/internal/models"
	"backend_coffeeShop.go/internal/repository"
	"backend_coffeeShop.go/pkg"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type ProductHandler struct {
	repository.ProductRepositoryInterface
	pkg.Cloudinary
}

func NewProductHandler(r repository.ProductRepositoryInterface, cld pkg.Cloudinary) *ProductHandler {
	return &ProductHandler{r,cld}
}

func (h *ProductHandler) Post(ctx *gin.Context){
	response := pkg.NewResponse(ctx)
	product := models.Product{}

	if err := ctx.ShouldBind(&product); err!= nil {
		response.BadRequest("create data failed,", err.Error())
		return
	}

	file , header , err := ctx.Request.FormFile("productImg")
	if err != nil {
		response.BadRequest("create data failed, upload file failed", err.Error())
		return
	}
	
	mimeType := header.Header.Get("Content-Type")
	if mimeType != "image/jpg" && mimeType != "image/png" {
		response.BadRequest("create data failed, upload file failed, wrong file type", err)
		return
	}

	const MaxFileSize = 5242880
	fileSize := header.Size
    if fileSize > MaxFileSize {
        response.BadRequest("create data failed, file too large", nil)
        return
    }

	randomNumber := rand.Int()
    fileName := fmt.Sprintf("go-product-%d", randomNumber)
	uploadResult, err := h.UploadFile(ctx, file, fileName)
	if err != nil {
		response.BadRequest("create data failed, upload file failed", err.Error())
		return
	}
	product.Image_url = uploadResult.SecureURL

	result, err := h.CreatedProduct(&product)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			response.BadRequest("create data failed, product name already exists", err.Error())
		} else {
			response.BadRequest("create data failed,", err.Error())
		}
		return
	}

	response.Success("create data success", result)
}

func (h *ProductHandler) FetchAll(ctx *gin.Context){
	response := pkg.NewResponse(ctx)
	category := ctx.Query("category")
	favorite := ctx.Query("favoriteNpromo")
	searchText := ctx.Query("searchText")
	promo := ctx.Query("promo")
	limit := ctx.Query("limit")
	page := ctx.Query("page")
	sortBy := ctx.Query("sortBy")

	promoBool := promo == "true"
	

	limits , _ := strconv.Atoi(limit)
	pages , _ := strconv.Atoi(page)

	params := &models.Filter{       
		Category: category,
        Favorite: favorite,
        SearchText: searchText,
		Promo: promoBool,
		Limit: limits,
		Page: pages,
		SortBy: sortBy,
    }

	result, err := h.GetAllProduct(params)
    if err != nil {
        response.BadRequest("Get data failed", err.Error())
        return
    }

    response.Success("Get data success", result)

}

func (h *ProductHandler) FetchDetail(ctx *gin.Context){
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")

	result , err := h.GetDetailProduct(id)
	if err != nil {
		response.BadRequest("get data failed,", err.Error())
		return
	}
	
	response.Success("Get data success", result)
}

func (h *ProductHandler) Update(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	product := models.EditProduct{}
	id := ctx.Param("id")

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/* file */
	file, header, err := ctx.Request.FormFile("productImg")
	if err != nil {
		if err.Error() == "http: no such file" {
			fmt.Println("No file uploaded, skipping image update.")
			product.Image_url = "" 
		} else {
			response.BadRequest("update data failed, upload file failed", err.Error())
			return
		}
	} else {
		fmt.Println(header.Size)
		mimeType := header.Header.Get("Content-Type")
		if mimeType != "image/jpg" && mimeType != "image/png" {
			response.BadRequest("update data failed, upload file failed, wrong file type", err)
			return
		}

		const MaxFileSize = 5242880
		fileSize := header.Size
    	if fileSize > MaxFileSize {
    	    response.BadRequest("update data failed, file too large", nil)
    	    return
    	}

		randomNumber := rand.Int()
		fileName := fmt.Sprintf("go-product-%d", randomNumber)
		uploadResult, err := h.UploadFile(ctx, file, fileName)
		if err != nil {
			response.BadRequest("update data failed, upload file failed", err.Error())
			return
		}
		product.Image_url = uploadResult.SecureURL
	}

	result, err := h.EditProduct(&product, id)
	if err != nil {
		response.BadRequest("update data failed", err)
		return
	}

	response.Success("update data success", result)
}


func (h *ProductHandler) Delete(ctx *gin.Context){
	id := ctx.Param("id")
	data , err := h.DeleteProduct(id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200 , data)
}
