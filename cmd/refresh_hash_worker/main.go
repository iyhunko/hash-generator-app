package main

import (
	"encoding/json"
	"fmt"
	"github.com/iyhunko/hash-generation-app/internal/config"
	"github.com/iyhunko/hash-generation-app/internal/entity"
	"github.com/iyhunko/hash-generation-app/internal/store"
	"github.com/iyhunko/hash-generation-app/pkg/logger"
	"log"
	"time"
)

const (
	workerStartMsg = "Starting hash_refresher worker"
	updatedMsg     = "Updated: %s"
)

func main() {
	lgr, err := logger.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	lgr.Info(workerStartMsg)
	conf := config.NewConfig(lgr)
	storage := store.NewStore(lgr)

	for range time.Tick(conf.HashGenerationInterval) {
		hash := entity.NewHash()
		marshaledHash, err := json.Marshal(hash)
		if err != nil {
			lgr.ErrorWithExit(err.Error())
		}
		err = storage.Set(conf.HashFilePath, marshaledHash)
		if err != nil {
			lgr.ErrorWithExit(err.Error())
		}
		lgr.Info(fmt.Sprintf(updatedMsg, hash))
	}
}
