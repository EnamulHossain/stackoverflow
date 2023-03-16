package main

import (
	"embed"
	"fmt"
	"log"
	"net"
	"strings"

	categorypb "stackoverflow/gunk/v1/category"
	userpb "stackoverflow/gunk/v1/user"
	questionpb "stackoverflow/gunk/v1/question"
	answerepb "stackoverflow/gunk/v1/answere"
	cc "stackoverflow/usermgm/core/category"
	cq "stackoverflow/usermgm/core/question"
	cu "stackoverflow/usermgm/core/user"
	ca "stackoverflow/usermgm/core/answere"
	"stackoverflow/usermgm/service/category"
	"stackoverflow/usermgm/service/question"
	"stackoverflow/usermgm/service/user"
	"stackoverflow/usermgm/service/answere"
	"stackoverflow/usermgm/storage/postgres"

	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//go:embed migrations
var migrationFiles embed.FS

func main() {
	config := viper.NewWithOptions(
		viper.EnvKeyReplacer(
			strings.NewReplacer(".", "_"),
		),
	)
	config.SetConfigFile("config")
	config.SetConfigType("ini")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("error loading configuration: %v", err)
	}

	port := config.GetString("server.port")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("unable to listen port: %v", err)
	}

	postgreStorage, err := postgres.NewPostgresStorage(config)
	if err != nil {
		log.Fatalln(err)
	}

	goose.SetBaseFS(migrationFiles)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalln(err)
	}

	if err := goose.Up(postgreStorage.DB.DB, "migrations"); err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	userCore := cu.NewCoreUser(postgreStorage)
	userSvc := user.NewUserSvc(userCore)
	userpb.RegisterUserServiceServer(grpcServer, userSvc)

	categoryCore:= cc.NewCoreCategory(postgreStorage)
	categorySvc :=category.NewCatrgorySvc(categoryCore)
	categorypb.RegisterCategoryServiceServer(grpcServer,categorySvc)


	questionCore:=cq.NewCoreQuestion(postgreStorage)
	questionSvc := question.NewQuestionSvc(questionCore)
	questionpb.RegisterQuestionServiceServer(grpcServer,questionSvc)


	answereCore := ca.NewCoreAnswere(postgreStorage)
	answereSvc := answere.NewAnswereSvc(answereCore)
	answerepb.RegisterAnswereServiceServer(grpcServer,answereSvc)

	// start reflection server
	reflection.Register(grpcServer)

	fmt.Println("usermgm server running on: ", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("unable to serve: %v", err)
	}
}
