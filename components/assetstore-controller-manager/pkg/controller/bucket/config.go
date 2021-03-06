package bucket

import (
	"github.com/kyma-project/kyma/components/assetstore-controller-manager/pkg/store"
	"github.com/vrischmann/envconfig"
	"time"
)

type Config struct {
	Store                 store.Config
	BucketRequeueInterval time.Duration `envconfig:"default=5m"`
}

func loadConfig(prefix string) (Config, error) {
	cfg := Config{}
	err := envconfig.InitWithPrefix(&cfg, prefix)
	return cfg, err
}
