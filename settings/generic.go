package settings

// GenericSettings exposes project's generic settings
type GenericSettings struct {
	Debug  bool   `mapstructure:"debug"`
	Env    string `mapstructure:"env" default:"development"`
	Expiry int64  `mapstructure:"expiry" default:"10080"` // minutes for 7 days
	Port   int64  `mapstructure:"port" default:"9000"`
}

func init() {
	readSettings("generic", &GetSettings().Generic, false)
}
