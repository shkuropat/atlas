package common

func Init(rootPaths, homeRelativePaths []string, defaultConfigFile string) {
	InitLog()
	InitConfig(rootPaths, homeRelativePaths, defaultConfigFile)
}
