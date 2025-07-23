package server

import (
	"fmt"
	"github.com/maxviazov/audittrail/config"
	"log"
)

func main() {
	fmt.Println("==> Starting server <==")

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

}
