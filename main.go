package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	rpio "github.com/stianeikeland/go-rpio/v4"
)

const (
	kSleepTime = time.Second
)

func triggerGarage(context *gin.Context) {
	pin := rpio.Pin(4)

	defer pin.Low()
	pin.High()
	time.Sleep(kSleepTime)
}

func main() {
	listenAddr := flag.String("address", "", "Address and port to listen on.")
	ginDebug := flag.Bool("gin_debug", false, "Run Gin in Debug Mode.")

	flag.Parse()

	if len(*listenAddr) == 0 {
		fmt.Println("Flag Validation Failure: Empty Address")
		return
	}

	// Setup GPIO pins
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		return
	}
	pin := rpio.Pin(4)
	pin.Output() // Ensure GPIO pin is in output mode
	pin.Low()    // Ensure relay is off

	ginMode := gin.ReleaseMode
	if *ginDebug {
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)

	router := gin.Default()
	router.GET("/garage", triggerGarage)
	router.Run(*listenAddr)
}
