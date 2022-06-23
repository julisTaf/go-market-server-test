package route

import (
	"Go-market-test/pkg/comment"
	"Go-market-test/pkg/controllers"
	"Go-market-test/pkg/deal"
	"Go-market-test/pkg/imagechild"
	"Go-market-test/pkg/middlewares"
	"Go-market-test/pkg/user"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	//r.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"*"},
	//	AllowMethods:     []string{"PUT, POST, GET, DELETE, OPTIONS"},
	//	AllowHeaders:     []string{"Origin"},
	//	ExposeHeaders:    []string{"Content-Length,Content-Type"},
	//	AllowCredentials: true,
	//	AllowOriginFunc: func(origin string) bool {
	//		return origin == "*"
	//	},
	//	MaxAge: 12 * time.Hour,
	//}))

	r.Use(CORS())

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	api := r.Group("/api")
	{
		public := api.Group("/public")
		{
			public.POST("/login", controllers.Login)
			public.POST("/signup", controllers.Signup)
			public.POST("/upload/:id", imagechild.Upload)
			public.GET("/image/:xid", imagechild.UserFileDownloadCommonService)
		}

		// here
		protected := api.Group("/protected").Use(middlewares.Authz())
		{
			protected.GET("/profile/:id", controllers.Profile)
			protected.GET("/user/:id", user.GetUser)
		}

		dealTree := api.Group("/deal").Use(middlewares.Authz())
		{
			dealTree.GET("/:id", deal.GetDeal)
			dealTree.GET("/author/:id", deal.GetUserDeals)
			dealTree.GET("/all", deal.GetAllDeals)
			dealTree.GET("/delete/:id", deal.DeleteDeal)

			dealTree.POST("/new", deal.NewDeal)
			dealTree.POST("/update", deal.UpdateDeal)
		}

		CommentTree := api.Group("/comment").Use(middlewares.Authz())
		{
			CommentTree.GET("/:id", comment.GetComment)
			CommentTree.GET("/delete/:id", comment.DeleteComment)
			CommentTree.GET("/response/:id", comment.GetRestComment)
			CommentTree.POST("/new", comment.NewComment)
		}
	}

	return r
}

func CORS() gin.HandlerFunc {
	// TO allow CORS
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		/*
		   c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		   c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		   c.Writer.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
		   c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")
		*/

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
