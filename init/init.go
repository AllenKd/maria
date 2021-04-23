package sysInit

import (
	"maria/internal/pkg/configs"
)

func Init() {
	configs.New()
	//db.New()
}
