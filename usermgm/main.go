package main

import (
	"embed"
	"fmt"
	"log"
	"net"
	"strings"

	userpb "stackoverflow/gunk/v1/user"
	cu "stackoverflow/usermgm/core/user"
	"stackoverflow/usermgm/service/user"
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

	// start reflection server
	reflection.Register(grpcServer)

	fmt.Println("usermgm server running on: ", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("unable to serve: %v", err)
	}
}
