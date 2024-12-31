package entity

type Recipient struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MetaData struct {
	Priority   string `json:"priority"`
	Retries    int    `json:"retries"`
	ExecuteAt  int64  `json:"execute_at" bson:"execute_at"`
	MaxExecute int64  `json:"max_execute_at" bson:"max_execute_at"`
}

type RequestNotification struct {
	ID         string      `json:"id,omitempty" bson:"_id,omitempty"`
	Recipient  []Recipient `json:"recipient"`
	TemplateId string      `json:"template_id" bson:"template_id"`
	Quota      int64       `json:"quota"`
	Agent      string      `json:"agent"`
	MetaData   MetaData    `json:"meta_data" bson:"meta_data"`
	Status     string      `bson:"status"`
}
