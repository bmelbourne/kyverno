package internal

import (
	"github.com/go-logr/logr"
	"github.com/kyverno/kyverno/pkg/imageverifycache"
)

func setupImageVerifyCache(logger logr.Logger) imageverifycache.Client {
	logger = logger.WithName("image-verify-cache").WithValues("enabled", imageVerifyCacheEnabled, "maxsize", imageVerifyCacheMaxSize, "ttl", imageVerifyCacheTTLDuration)
	logger.V(2).Info("setup image verify cache...")
	opts := []imageverifycache.Option{
		imageverifycache.WithLogger(logger),
		imageverifycache.WithCacheEnableFlag(imageVerifyCacheEnabled),
		imageverifycache.WithMaxSize(imageVerifyCacheMaxSize),
		imageverifycache.WithTTLDuration(imageVerifyCacheTTLDuration),
	}
	imageVerifyCache, err := imageverifycache.New(opts...)
	checkError(logger, err, "failed to create image verify cache client")
	return imageVerifyCache
}
