package order

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type ListRequest struct {
	Page int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type ListResponseItem struct {
	Id string `json:"id"`
	Name string `json:"name"`

}

type ListResponse struct {
	Items []*ListResponseItem
}

func List(c *gin.Context) {
	req := ListRequest{}
	c.ShouldBindJSON(&req)
	

	resp := ListResponse{}
	for i := 0; i < 5; i++ {
		resp.Items = append(resp.Items, &ListResponseItem{
			Id: fmt.Sprintf("%d", i),
			Name: "name",
		})	
	}




	c.JSON(http.StatusOK, &resp)
}