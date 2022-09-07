package internal

import (
	"net/http"
	"time"

	"github.com/Oybek-uzb/mana-taxi-bot-service/internal/config"
	"github.com/Oybek-uzb/mana-taxi-bot-service/pkg/logging"

	tele "gopkg.in/telebot.v3"
)

type app struct {
	cfg        *config.Config
	logger     *logging.Logger
	httpServer *http.Server
}

type App interface {
	Run()
}

func NewApp(logger *logging.Logger, cfg *config.Config) (App, error) {
	logger.Println("NewApp initializing")
	return &app{
		cfg:    cfg,
		logger: logger,
	}, nil
}

func (a *app) Run() {
	a.startBot()
}

func (a *app) startBot() {
	pref := tele.Settings{
		Token:  a.cfg.Telegram.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		a.logger.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Start()
}
