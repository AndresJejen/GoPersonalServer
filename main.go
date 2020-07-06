package main

import (
	"fmt"
	"net/http"

	"github.com/AndresJejen/GoPersonalServer/controller"
	"github.com/AndresJejen/GoPersonalServer/repository"
	"github.com/AndresJejen/GoPersonalServer/router"
	"github.com/AndresJejen/GoPersonalServer/service"
)

var (
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewMuxRouter()
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
