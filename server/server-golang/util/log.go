package util

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorLogData struct {
	Url string
	Err string
}

type Manager struct {
	ErrorLog chan *ErrorLogData
}

var ManagerInstance = Manager{
	ErrorLog: make(chan *ErrorLogData, 128),
}

func (manager *Manager) Start() {
	log.Println("Manage Started !")
	for {
		select {
		case errorLog := <-manager.ErrorLog:
			go LogError(errorLog)
		}
	}
}

func LogError(errorLogData *ErrorLogData) {
	log.Println("URL:", errorLogData.Url, "  Info: ", errorLogData.Err)
}

func (manager *Manager) SendError(url, err string) {
	data := &ErrorLogData{
		Url: url,
		Err: err,
	}
	manager.ErrorLog <- data
}

func (manager *Manager) FailWithoutData(ctx *gin.Context, msg string) {
	Response(ctx, http.StatusOK, 400, gin.H{}, msg)
	manager.SendError(ctx.Request.URL.String(), msg)
}
