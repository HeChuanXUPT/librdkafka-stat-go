package main

type BrokerCount struct {
	Min   int64 `json:"min"`
	Max   int64 `json:"max"`
	Avg   int64 `json:"avg"`
	Sum   int64 `json:"sum"`
	Count int64 `json:"cnt"`
}

type TopicPartition struct {
	Topic     string `json:"topic"`
	Partition int    `json:"partition"`
}

type Broker struct {
	Name           string                    `json:"name"`
	NodeId         int                       `json:"nodeid"`
	State          string                    `json:"state"`
	StateAge       int64                     `json:"stateage"`
	OutbufCnt      int64                     `json:"outbuf_cnt"`
	OutbufMsgCnt   int64                     `json:"outbuf_msg_cnt"`
	WaitRespCnt    int64                     `json:"waitresp_cnt"`
	WaitRespMsgCnt int64                     `json:"waitresp_msg_cnt"`
	Tx             int64                     `json:"tx"`
	TxBytes        int64                     `json:"txbytes"`
	TxErrs         int64                     `json:"txerrs"`
	TxrEtries      int64                     `json:"txretries"`
	ReqTimeouts    int64                     `json:"req_timeouts"`
	Rx             int64                     `json:"rx"`
	RxBytes        int64                     `json:"rxbytes"`
	RxErrs         int64                     `json:"rxerrs"`
	RxCorriderrs   int64                     `json:"rxcorriderrs"`
	Rxpartial      int64                     `json:"rxpartial"`
	ZBufGrow       int64                     `json:"zbuf_grow"`
	BufGrow        int64                     `json:"buf_grow"`
	Rtt            BrokerCount               `json:"rtt"`
	Throttle       BrokerCount               `json:"throttle"`
	Toppars        map[string]TopicPartition `json:"toppars"`
}

type PartitionInfo struct {
	Partition       int    `json:"partition"`
	Leader          int    `json:"leader"`
	Desired         bool   `json:"desired"`
	Unknown         bool   `json:"unknown"`
	MsgqCnt         int64  `json:"msgq_cnt"`
	MsgqBytes       int64  `json:"msgq_bytes"`
	XmitMsgqCnt     int64  `json:"xmit_msgq_cnt"`
	XmitMsgqBytes   int64  `json:"xmit_msgq_bytes"`
	FetchqCnt       int64  `json:"fetchq_cnt"`
	FetchqSize      int64  `json:"fetchq_size"`
	FetchState      string `json:"fetch_state"`
	QueryOffset     int    `json:"query_offset"`
	NextOffset      int    `json:"next_offset"`
	AppOffset       int    `json:"app_offset"`
	StoredOffset    int    `json:"stored_offset"`
	CommitedOffset  int    `json:"commited_offset"`
	CommittedOffset int    `json:"committed_offset"`
	EofOffset       int    `json:"eof_offset"`
	LoOffset        int    `json:"lo_offset"`
	HiOffset        int    `json:"hi_offset"`
	ConsumerLag     int    `json:"consumer_lag"`
	TxMsgs          int64  `json:"txmsgs"`
	TxBytes         int64  `json:"txbytes"`
	Msgs            int64  `json:"msgs"`
	RxVerDrops      int64  `json:"rx_ver_drops"`
}

type Topic struct {
	Topic       string                   `json:"topic"`
	MetadataAge uint                     `json:"metadata_age"`
	Partitions  map[string]PartitionInfo `json:"partitions"`
}

type Stat struct {
	Name       string            `json:"name"`
	Type       string            `json:"type"`
	TS         int64             `json:"ts"`
	Time       int64             `json:"time"`
	ReplyQ     int64             `json:"replyq"`
	MsgCnt     int64             `json:"msg_cnt"`
	MsgSize    int64             `json:"msg_size"`
	MsgMax     int64             `json:"msg_max"`
	MsgSizeMax int64             `json:"msg_size_max"`
	SimpleCnt  int64             `json:"simple_cnt"`
	Brokers    map[string]Broker `json:"brokers"`
	Topics     map[string]Topic  `json:"topics"`
}
