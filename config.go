package main

// Config keeps the global configuration
var Config = struct {
	NodeNetwork struct {
		Port      int    `default:"3000" env:"NODE_PORT"`
		Interface string `default:"127.0.0.1"`
	}
	Web struct {
		Port      int    `default:"4000"`
		Interface string `default:"127.0.0.1"`
	}
	Logger struct {
		PrintEmoji bool `default:"true"`
	}
}{}
