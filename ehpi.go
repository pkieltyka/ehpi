package ehpi

import (
	"github.com/pkieltyka/ehpi/config"
	"github.com/pkieltyka/ehpi/data"
)

type Ehpi struct {
	Config *config.Config
	DB     *data.DB
}
