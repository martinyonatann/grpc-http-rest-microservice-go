package v1

import (
"database/sql"
	"context"
	"fmt"
	"time"

"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	//apiVersion is version of API is providedby server
	apiVersion = "v1"
)

//toDoServiceServer is implementation of v1. ToDoServiceServer proto interface
type toDoServiceServer struct {
	db *sql.DB
}

//NewToDoServiceServer creates ToDo service
func NewToDoServiceServer(db *sql.DB) v1.ToDoServiceServer{
	return &toDoServiceServer{db: db}
}

// checkAPI check if the API version requested by client is supported by server
func (s *toDoServiceServer) checkAPI (api string) error{
	//API version is "" means use current version of the service
	if len (api) > 0 {
		if apiVersion != api{
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version
			'%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}


//connecct returns SQL database connection from the pool
func (s *toDoServiceServer) connect (ctx context.Context)(*sql.Conn, err){
	c, err := s.db.Conn(ctx)
	if err != nil{
		return nil, status.Error(codes.Unknown, "failed to connect to database->"+err.Error())
	}
	return c, nil
}

//Create new todo task
func (s *toDoServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	//Check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil{
		return niil, err
	}

	//get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil{
		return nil, err
	}
	defer c.Close()

	reminder, err := 
}