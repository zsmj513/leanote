package api

type SyncState struct {
	Usn  int `json:"LastSyncUsn"`
	Time int `json:"LastSyncTime"`
}

type ReUpdate struct {
	Ok  bool   `json:"Ok"`
	Msg string `json:"Msg"`
	Usn int    `json:"Usn"`
}
