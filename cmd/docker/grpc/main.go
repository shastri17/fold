package main

import (
	"database/sql"
	"fmt"
	"fold/protobuf/golang/grpc/hashtag"
	"fold/protobuf/golang/grpc/project"
	user "fold/protobuf/golang/grpc/user"
	"fold/service"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/jinzhu/gorm"
	_ "go.elastic.co/apm/module/apmgorm/dialects/postgres"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

var server *grpc.Server

const maxMsgSize = 30 << 20

func main() {
	initialize()
	run()
}

func run() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", 8001))
	if err != nil {
		log.Println("error in initilizing grpc: ", err)
		return
	}
	err = server.Serve(l)
	if err != nil {
		log.Println("error in serving grpc: ", err)
		return
	}
}
func initialize() {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%v",
		"postgres", "folduser", "folduser_password", "fold_db", 5432, "disable")
	baseDb, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	db, err := gorm.Open("postgres", baseDb)
	if err != nil {
		fmt.Println(err)
		return
	}
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://elasticsearch:9200"},
	})
	if err != nil {
		os.Exit(1)
	}
	userServ := service.NewUserProject(db, esClient)
	projServ := service.NewProject(db)
	hashServ := service.NewHashTag(db)
	server = grpc.NewServer(grpc.MaxSendMsgSize(maxMsgSize),
		grpc.MaxRecvMsgSize(maxMsgSize))
	user.RegisterUserServer(server, userServ)
	project.RegisterProjectServer(server, projServ)
	hashtag.RegisterHashtagServer(server, hashServ)

}
