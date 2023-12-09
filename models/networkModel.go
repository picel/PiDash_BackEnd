package models

type NetInfo struct {
	Interface string `json:"interface"`
	Mac       string `json:"mac"`
}

type NetStats struct {
	Interface string `json:"interface"`
	TxSpeed   uint64 `json:"txSpeed"`
	RxSpeed   uint64 `json:"rxSpeed"`
}
