package main

// Simple tool to write test configs to redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cisco-open/synthetic-heart/common/proto"
	"github.com/cisco-open/synthetic-heart/common/storage"
	"github.com/hashicorp/go-hclog"
	"os"
	"path/filepath"
	"time"
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:            "redis-poster",
		Level:           hclog.Trace,
		IncludeLocation: true,
		Color:           hclog.ForceColor,
	})
	redis := storage.NewRedisSynHeartStore(storage.SynHeartStoreConfig{
		Type:       "redis",
		BufferSize: 1000,
		Address:    "localhost:6379",
	}, logger)

	ctx := context.Background()

	// push the test configs to redis
	yamlFiles, _ := filepath.Glob("../configs/syntest-configs/*.yaml")
	for _, file := range yamlFiles {
		rawData, err := os.ReadFile(filepath.Clean(file))
		if err != nil {
			fmt.Printf("Error reading config file: %s, %v\n", file, err)
			continue
		}
		// Unmarshal the YAML into a Config struct
		config := proto.SynTestConfig{}
		err = json.Unmarshal(rawData, &config)
		if err != nil {
			fmt.Printf("Error unmarshaling YAML file to SynTestConfig struct: %s, %v\n", file, err)
			continue
		}

		logger.Info("pushing test to redis", "name", config.Name, "version", config.Version, "file", file)
		config.Version = time.Now().String()
		err = redis.WriteTestConfig(ctx, config, string(rawData))

		if err != nil {
			logger.Error("error sending command to redis", "err", err)
		}
	}
}
