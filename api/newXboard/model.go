package newXboard

import "encoding/json"

// NodeInfoResponse is the response of node
type NodeInfoResponse struct {
	SpeedLimit float64 `json:"speedlimit"`
	Method     string  `json:"method"`
	Port       uint32  `json:"outPort"`
	Routes     []route `json:"routes"`
}

// UserResponse is the response of user
type UserResponse struct {
	ID          int     `json:"id"`
	Passwd      string  `json:"passwd"`
	SpeedLimit  float64 `json:"nodeSpeedlimit"`
	DeviceLimit int     `json:"nodeConnector"`
}

// Response is the common response
type Response struct {
	Code uint            `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

// PostData is the data structure of post data
type PostData struct {
	Type    string      `json:"type"`
	NodeId  int         `json:"nodeId"`
	Users   interface{} `json:"users"`
	Onlines interface{} `json:"onlines"`
}

// SystemLoad is the data structure of systemload
type SystemLoad struct {
	Uptime string `json:"uptime"`
	Load   string `json:"load"`
}

// OnlineUser is the data structure of online user
type OnlineUser struct {
	UID int    `json:"user_id"`
	IP  string `json:"ip"`
}

// UserTraffic is the data structure of traffic
type UserTraffic struct {
	UID      int    `json:"id"`
	Upload   int64  `json:"up"`
	Download int64  `json:"down"`
	Ip       string `json:"ip"`
}

type route struct {
	Id          int      `json:"id"`
	Match       []string `json:"matchs"`
	Action      string   `json:"actionFun"`
	ActionValue string   `json:"actionValue"`
}

type IllegalItem struct {
	ID  int `json:"list_id"`
	UID int `json:"user_id"`
}
