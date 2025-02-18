package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/haruotsu/grpc-test/internal/pb"
	"github.com/haruotsu/grpc-test/internal/repository"
	"github.com/haruotsu/grpc-test/internal/server"
	"github.com/haruotsu/grpc-test/internal/service"

	_ "github.com/go-sql-driver/mysql" // MySQL ドライバ
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// DSN の形式: <username>:<password>@tcp(<host>:<port>)/<dbname>?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "hoge:hoge_pass@tcp(localhost:3306)/test-db?charset=utf8mb4&parseTime=True&loc=Local"

	// MySQL への接続
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	// Repository と Service の初期化
	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo)

	// gRPC サーバの初期化とサービス登録
	grpcServer := grpc.NewServer()
	userServer := server.NewUserServer(userSvc)
	pb.RegisterUserServiceServer(grpcServer, userServer)

	// grpcurl でのテストを容易にするためのreflectionの登録
	reflection.Register(grpcServer)

	// サーバリスンの設定
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
