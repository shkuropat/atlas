package common

func Init(defaultConfigFile string) {
	InitLog()
	InitConfig(defaultConfigFile)
}
