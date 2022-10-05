package api

import (
	"api-gateway/model"
	"api-gateway/proto"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (c *Controller) GetAll(gc *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := c.Services.TodosService().GetAll(ctx, &proto.GetAllTodosRequest{})

	if err != nil {
		gc.JSON(http.StatusInternalServerError, err)
		return
	}

	c.ResponseProtoJson(gc, response)
}

func (c *Controller) Create(gc *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()
	var newTodo model.Todos
	if err := gc.BindJSON(&newTodo); err != nil {
		panic(err)
	}
	response, err := c.Services.TodosService().Create(ctx, &proto.CreateTodoRequest{
		UserId:      newTodo.UserId,
		Description: newTodo.Description,
		Title:       newTodo.Title,
	})

	if err != nil {
		gc.JSON(http.StatusInternalServerError, err)
		return
	}

	c.ResponseProtoJson(gc, response)
}

func (c *Controller) Update(gc *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()
	id := gc.Param("id")
	var updTodo model.Todos
	if err := gc.BindJSON(&updTodo); err != nil {
		panic(err)
	}
	response, err := c.Services.TodosService().Update(ctx, &proto.UpdateTodoRequest{
		Id:          id,
		UserId:      updTodo.UserId,
		Description: updTodo.Description,
		Title:       updTodo.Title,
	})

	if err != nil {
		gc.JSON(http.StatusInternalServerError, err)
		return
	}

	c.ResponseProtoJson(gc, response)
}
func (c *Controller) Delete(gc *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()
	id := gc.Param("id")

	response, err := c.Services.TodosService().Delete(ctx, &proto.DeleteTodoRequset{
		Id: id,
	})

	if err != nil {
		gc.JSON(http.StatusInternalServerError, err)
		return
	}

	c.ResponseProtoJson(gc, response)
}

func (c *Controller) GetByID(gc *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()
	id := gc.Param("id")

	response, err := c.Services.TodosService().GetById(ctx, &proto.GetByIdRequest{Id: id})

	if err != nil {
		gc.JSON(http.StatusInternalServerError, err)
		return
	}

	c.ResponseProtoJson(gc, response)
}

// Create(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error)
//	GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetTodoResponse, error)
//	GetAll(ctx context.Context, in *GetAllTodosRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
//	Delete(ctx context.Context, in *DeleteTodoRequset, opts ...grpc.CallOption) (*DeleteTodoResponse, error)
//	Update(ctx context.Context, in *UpdateTodoRequest, opts ...grpc.CallOption) (*UpdateTodoResponse, error)
