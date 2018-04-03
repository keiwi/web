package main

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/apex/log"
	"github.com/keiwi/utils"
	"github.com/keiwi/web/api"
	"github.com/keiwi/web/models"
	"github.com/keiwi/web/routes"
	"github.com/nats-io/go-nats"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
)

func main() {
	// TODO: Maybe add support for re-reading config because viper supports it?
	configType := os.Getenv("KeiwiServerConfigType")
	if configType == "" {
		configType = "json"
	}
	viper.SetConfigType(configType)

	viper.SetConfigFile("config." + configType)
	viper.AddConfigPath(".")

	viper.SetDefault("log_dir", "./logs")
	viper.SetDefault("log_syntax", "%date%_web.log")
	viper.SetDefault("log_level", "info")
	viper.SetDefault("nats_delay", 10)

	if err := viper.ReadInConfig(); err != nil {
		utils.Log.Debug("Config file not found, saving default")
		if err = viper.WriteConfigAs("config." + configType); err != nil {
			utils.Log.WithField("error", err.Error()).Fatal("Can't save default config")
		}
	}

	level := strings.ToLower(viper.GetString("log_level"))
	utils.Log = utils.NewLogger(utils.NameToLevel[level], &utils.LoggerConfig{
		Dirname: viper.GetString("log_dir"),
		Logname: viper.GetString("log_syntax"),
	})

	db, err := models.NewHandler(nats.DefaultURL)
	if err != nil {
		utils.Log.WithError(err).Fatal("error when initializing handler")
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
		Logger:    utils.Log,
		Level:     log.DebugLevel,
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

	utils.Log.Handler.HandleLog(entry)
}
