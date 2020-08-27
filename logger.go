package gologger

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber"
)

//LogStruct - logger structure
type LogStruct struct {
	IP        string
	URL       string
	StartTime string
	EndTime   string
	Duration  int64
	Agent     string
	Status    int
	Method    string
}

//Log - logger will print JSON formatted logs onto STDOUT
func Log(ctx *fiber.Ctx) {
	t := time.Now()
	logger := LogStruct{
		IP:        ctx.IP(),
		URL:       ctx.OriginalURL(),
		StartTime: t.String(),
		Method:    string(ctx.Fasthttp.Method()),
		Agent:     string(ctx.Fasthttp.UserAgent()),
	}
	ctx.Next()
	logger.Status = ctx.Fasthttp.Response.StatusCode()
	logger.EndTime = time.Now().String()
	logger.Duration = time.Since(t).Milliseconds()
	logStr, err := json.Marshal(logger)
	if err == nil {
		log.Printf("%s", string(logStr))
	}
}
