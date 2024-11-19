package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xprasetio/go_simple_forum.git/internal/configs"
	"github.com/xprasetio/go_simple_forum.git/internal/handlers/memberships"
	"github.com/xprasetio/go_simple_forum.git/internal/handlers/posts"
	membershipRepo "github.com/xprasetio/go_simple_forum.git/internal/repository/memberships"
	postRepo "github.com/xprasetio/go_simple_forum.git/internal/repository/posts"
	membershipSvc "github.com/xprasetio/go_simple_forum.git/internal/service/memberships"
	postSvc "github.com/xprasetio/go_simple_forum.git/internal/service/posts"
	"github.com/xprasetio/go_simple_forum.git/pkg/internalsql"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	if err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml")); err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}
	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	defer db.Close()

	memberRepo := membershipRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)

	memberService := membershipSvc.NewService(cfg, memberRepo)
	postService := postSvc.NewService(cfg, postRepo)

	membershipHandler := memberships.NewHandler(r, memberService)
	membershipHandler.RegisterRoutes()

	postHandler := posts.NewHandler(r, postService)
	postHandler.RegisterRoutes()

	r.Run(cfg.Service.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
