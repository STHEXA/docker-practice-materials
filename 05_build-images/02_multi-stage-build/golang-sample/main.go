package main

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type App struct {
	Logger *logrus.Logger
	Router *mux.Router
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := &App{
		Logger: initLogger(),
		Router: mux.NewRouter(),
	}

	// ルーティング設定
	app.setRoutes()

	// ログ設定
	loggedRouter := app.logRequestMiddleware(app.Router)

	// サーバーを起動
	app.Logger.Infof("Server started ( http://localhost:%s )", port)
	log.Fatal(http.ListenAndServe(":"+port, loggedRouter))
}

// Loggerを初期化する関数
func initLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
	return logger
}

// ルーティングの設定
func (a *App) setRoutes() {
	// 各ページ用のルーティング設定
	a.Router.HandleFunc("/", a.handleWithView("index", map[string]interface{}{
		"Message": "Hello! Docker",
	}))

	// 静的ファイルの提供設定
	a.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
}

// 汎用ハンドラ
func (a *App) handleWithView(tmplName string, data map[string]interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a.renderTemplate(w, r, tmplName, data)
	}
}

// テンプレートからレンダー
func (a *App) renderTemplate(w http.ResponseWriter, r *http.Request, tmplName string, data interface{}) {
	templatePath := filepath.Join("templates", tmplName+".html")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		a.Logger.Errorf("Template not found: %v", err)
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, data); err != nil {
		a.Logger.Errorf("Error executing template: %v", err)
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
	}
}

// ロギングの設定
func (a *App) logRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		a.Logger.WithFields(logrus.Fields{
			"method": r.Method,
			"url":    r.URL.String(),
			"remote": r.RemoteAddr,
		}).Info("Request received")

		next.ServeHTTP(w, r)

		a.Logger.WithFields(logrus.Fields{
			"method":      r.Method,
			"url":         r.URL.String(),
			"remote":      r.RemoteAddr,
			"duration_ms": time.Since(start).Milliseconds(),
		}).Info("Request processed")
	})
}
