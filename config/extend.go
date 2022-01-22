package config

var ExtConfig Extend

// Extend 扩展配置
//  extend:
//    demo:
//      name: demo-name
// 使用方法： config.ExtConfig......即可！！
type Extend struct {
	AMap AMap // 这里配置对应配置文件的结构即可
	Gen  Gen  // 这里配置对应配置文件的结构即可
}

type AMap struct {
	Key string
}

type GenModelFlag struct {
	Models  bool
	Apis    bool
	Router  bool
	Dto     bool
	Service bool
	Api_js  bool
	Vue     bool
}

type Gen struct {
	Serverpath string
	ModelFlag  GenModelFlag
}
