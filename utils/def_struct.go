package utils

// config.yml 结构定义

type Config struct {
	Base *Base    `yml:"base"`
	Eq   []string `yml:"eq"`
	Le   []string `yml:"le"`
}

type Base struct {
	WarDir string   `yml:"wardir"`
	Exts   []string `yml:"exts"`
}

// 存储war、ear、jar 路径及查看包内jar

type FileJar struct {
	FilePath string
	Jars     []string
}

// 提取jar关联信息

type JarInfo struct {
	ParPath string
	JarPath string
	Name    string
	Ver     string
}
