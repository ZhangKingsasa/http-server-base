package cmd

import (
	"git.garena.com/shopee/MLP/aip/platform/aip-user-service/cmd/app"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("Server app failed to run: %v", err)
	}
}