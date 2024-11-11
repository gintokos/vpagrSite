package grpcdb

import (
	"context"
	"fmt"
	"time"
	"github.com/gintokos/vpagrSite/internal/config"
	"github.com/gintokos/vpagrSite/internal/domain/models"

	generated "github.com/gintokos/serverdb/protos/gen/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Grpcdb struct {
	con    *grpc.ClientConn
	cfg    *config.GrpcdbConfig
	client generated.DBClient
}

func NewGrpcDb(cfg *config.GrpcdbConfig) Grpcdb {
	return Grpcdb{
		cfg: cfg,
	}
}

func (g *Grpcdb) Stop() {
	if g.con != nil {
		g.con.Close()
	}
}

func (g *Grpcdb) MakeConnection() error {
	target := g.cfg.Domen + g.cfg.Port
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	client := generated.NewDBClient(conn)
	g.client = client

	return nil
}

func (g *Grpcdb) CreateUser(ctx context.Context, id int64) error {
	req := &generated.CreateUserRequest{
		UserId: id,
	}

	res, err := g.client.CreateUser(ctx, req)
	if err != nil {
		return err
	}
	return g.HandleErrInResult(res.Error)
}

func (g *Grpcdb) GetUser(ctx context.Context, id int64) (user models.User, err error) {
	req := &generated.GetUserRequest{
		UserId: id,
	}

	res, err := g.client.GetUser(ctx, req)
	if err != nil {
		return models.User{}, err
	}

	format := "02.01.2006 15:04:05"
	createdat, err := time.Parse(format, res.CreatedAt)
	if err != nil {
		return models.User{}, err
	}
	user = models.User{
		CreatedAt: createdat,
		Tid:       id,
	}

	return user, g.HandleErrInResult(res.Error)
}

func (g *Grpcdb) HandleErrInResult(errstring string) error {
	if errstring == "" {
		return nil
	}
	return fmt.Errorf("found error in result of responce: %s", errstring)
}
