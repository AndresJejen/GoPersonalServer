package main

import (
	"fmt"
	"net/http"

	"github.com/AndresJejen/GoPersonalServer/controller"
	"github.com/AndresJejen/GoPersonalServer/router"
)

var (
	httpRouter     router.Router             = router.NewMuxRouter()
	postController controller.PostController = controller.NewPostController()
)

func main() {

	const port string = ":8000"

	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Your system is Online")
	})

	httpRouter.GET("/post", postController.GetPost)
	httpRouter.POST("/post", postController.AddPost)

	httpRouter.SERVE(port)
}
