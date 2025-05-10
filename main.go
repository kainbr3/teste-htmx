package main

import (
	cfg "crypto-braza-tokens-admin/configs"
	w "crypto-braza-tokens-admin/web"
)

func main() {
	cfg.InitRequirements()
	w.Start()
}
