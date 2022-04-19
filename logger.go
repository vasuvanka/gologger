package gologger

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber"
)

//LogStruct - logger structure
type LogStruct struct {
	IP        string 	`json:"ip"`
	URL       string	`json:"url"`
	StartTime string	`json:"start_time"`
	EndTime   string	`json:"end_time"`
	Duration  int64		`json:"duration"`
	Agent     string 	`json:"agent"`	
	Status    int 		`json:"status"`
	Method    string	`json:"method"`
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
	logStr, _ := json.Marshal(logger)
	log.Printf("%s", string(logStr))
}
