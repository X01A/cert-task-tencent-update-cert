package main

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

func init() {
	flag.Parse()

	if len(os.Getenv("TENCENT_VERBOSE")) > 0 {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func envValue(keys []string, fallback string) string {
	if len(fallback) > 0 {
		return fallback
	}

	for _, key := range keys {
		v := os.Getenv(key)
		if len(v) > 0 {
			return v
		}
	}

	return ""
}

func getCertDomain() string {
	return envValue([]string{"TENCENT_OVERRIDE_DOMAIN", "Le_Domain"}, "")
}

func getTencentClient() *common.Credential {
	credential := common.NewCredential(
		envValue([]string{"Tencent_SecretId", "TENCENT_SECRET_ID"}, ""),
		envValue([]string{"Tencent_SecretKey", "TENCENT_SECRET_KEY"}, ""))

	return credential
}

func readCertKey() string {
	return os.Getenv("TASK_KEY")
}

func readCertFullchain() string {
	return os.Getenv("TASK_CERT")
}
