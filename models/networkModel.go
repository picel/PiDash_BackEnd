package models

type NetInfo struct {
	Interface string `json:"interface"`
	Mac       string `json:"mac"`
}

type NetStats struct {
	Interface string `json:"interface"`
	TxSpeed   string `json:"txSpeed"`
	RxSpeed   string `json:"rxSpeed"`
}
