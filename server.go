package main

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/keiwi/utils/log"
	"github.com/keiwi/utils/log/handlers/cli"
	"github.com/keiwi/utils/log/handlers/file"
	"github.com/keiwi/web/api"
	"github.com/keiwi/web/models"
	"github.com/keiwi/web/routes"
	"github.com/nats-io/go-nats"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
)

func main() {
	log.Log = log.NewLogger(log.DEBUG, []log.Reporter{
		cli.NewCli(),
		file.NewFile("./logs", "%date%_web.log"),
	})

	// TODO: Maybe add support for re-reading config because viper supports it?
	configType := os.Getenv("KeiwiServerConfigType")
	if configType == "" {
		configType = "json"
	}
	viper.SetConfigType(configType)

	viper.SetConfigFile("config." + configType)
	viper.AddConfigPath(".")

	viper.SetDefault("log.dir", "./logs")
	viper.SetDefault("log.syntax", "%date%_web.log")
	viper.SetDefault("log.level", "info")

	viper.SetDefault("nats.delay", 10)
	viper.SetDefault("nats.url", nats.DefaultURL)

	if err := viper.ReadInConfig(); err != nil {
		log.Debug("Config file not found, saving default")
		if err = viper.WriteConfigAs("config." + configType); err != nil {
			log.WithField("error", err.Error()).Fatal("Can't save default config")
		}
	}

	level := strings.ToLower(viper.GetString("log.level"))
	log.Log = log.NewLogger(log.GetLevelFromString(level), []log.Reporter{
		cli.NewCli(),
		file.NewFile(viper.GetString("log.dir"), viper.GetString("log.syntax")),
	})

	db, err := models.NewHandler(viper.GetString("nats.url"))
	if err != nil {
		log.WithError(err).Fatal("error when initializing handler")
		return
	}

	apiHandler := api.NewAPI(db)
	routesHandler := routes.NewRoutes(apiHandler)
	n := negroni.New(negroni.NewRecovery(), NewLogger())
	n.UseHandler(routesHandler)
	n.Run(":3000")
}

// NewLogger returns a new Logger instance
func NewLogger() *Logger {
	logger := &Logger{}
	return logger
}

// Logger is a middleware handler that logs the request as it goes in and the response as it goes out.
// Modified version of the default negroni logger.
type Logger struct {
}

func (l *Logger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()

	next(rw, r)

	res := rw.(negroni.ResponseWriter)
	entry := &log.Entry{
		Logger:    log.Log,
		Level:     log.DEBUG,
		Timestamp: start,
		Message:   "[negroni] HTTP",
		Fields: log.Fields{
			"Status":   res.Status(),
			"Duration": time.Since(start),
			"Hostname": r.Host,
			"Method":   r.Method,
			"Path":     r.URL.Path,
		},
	}

	log.Log.Write(log.DEBUG, entry, "[negroni] HTTP", 1)
}
