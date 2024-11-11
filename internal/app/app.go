package app

import (
	"time"
	"github.com/gintokos/vpagrSite/internal/config"
	"github.com/gintokos/vpagrSite/internal/data/db"
	"github.com/gintokos/vpagrSite/internal/data/db/grpcdb"
	"github.com/gintokos/vpagrSite/internal/services"
	hhttp "github.com/gintokos/vpagrSite/internal/transport/http"
	"github.com/gintokos/vpagrSite/pkg/telegramauth"

	"github.com/gintokos/serverdb/pkg/logger"
)

func MustStartApp() {
	cfg := config.MustInitConfig("config.json")

	logger.MustSetupLogger()
	lg := logger.GetLogger()
	lg.Info("Config and logger was initilized")

	StartTelegramBot(&cfg.Tconfig)
	lg.Info("TelegramBot started work")

	InitDBconnection(&cfg.GrpcdbConfig)
	lg.Info("Connection with db was created")

	services.InitServices(lg)
	lg.Info("Services was inited")

	srv := hhttp.NewServer(lg, &cfg.Sconfig)
	srv.MustStartServer()
}

func StartTelegramBot(cfg *config.TelegramBotConfig) {
	botOpt := telegramauth.BotOptions{
		TokenBot:      cfg.Token,
		Link:          cfg.Link,
		UserTokenSize: cfg.Usertokensize,
		TTLusertoken:  time.Second * time.Duration(cfg.Ttlusertoken),
	}
	bot, err := telegramauth.NewAuthBot(botOpt)
	if err != nil {
		panic(err)
	}

	telegramauth.SetTbot(bot)

	go bot.Start()
}

func InitDBconnection(cfg *config.GrpcdbConfig) {
	database := grpcdb.NewGrpcDb(cfg)
	err := database.MakeConnection()
	if err != nil {
		panic(err)
	}

	db.SetDataBase(&database)
}

func Stop(isgrpcdb bool) {
	if isgrpcdb {
		db.Database.(*grpcdb.Grpcdb).Stop()
	}
}
