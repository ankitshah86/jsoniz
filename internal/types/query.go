package types

type RequestType string
type OperationType string
type JoinType string
type ConsistencyLevel string
type Operator string

const (
	SQL               RequestType      = "SQL"
	NoSQL             RequestType      = "NoSQL"
	SELECT            OperationType    = "SELECT"
	INSERT            OperationType    = "INSERT"
	UPDATE            OperationType    = "UPDATE"
	DELETE            OperationType    = "DELETE"
	QUERY             OperationType    = "QUERY"
	CREATE_DATABASE   OperationType    = "CREATE_DATABASE"
	CREATE_COLLECTION OperationType    = "CREATE_COLLECTION"
	INNER             JoinType         = "INNER"
	LEFT              JoinType         = "LEFT"
	RIGHT             JoinType         = "RIGHT"
	STRONG            ConsistencyLevel = "STRONG"
	EVENTUAL          ConsistencyLevel = "EVENTUAL"
)

type Request struct {
	RequestID   string                 `json:"requestId"`
	Type        RequestType            `json:"type"`
	Operation   OperationType          `json:"operation"`
	Payload     map[string]interface{} `json:"payload"`
	Transaction Transaction            `json:"transaction"`
	Settings    Settings               `json:"settings"`
}

type Condition struct {
	Field    string   `json:"field"`
	Operator Operator `json:"operator"`
	Value    string   `json:"value"`
}

type Join struct {
	Type        JoinType    `json:"type"`
	Table       string      `json:"table"`
	OnCondition OnCondition `json:"onCondition"`
}

type OnCondition struct {
	Field1   string   `json:"field1"`
	Operator Operator `json:"operator"`
	Field2   string   `json:"field2"`
}

type Transaction struct {
	Start    bool   `json:"start"`
	Commit   bool   `json:"commit"`
	Rollback bool   `json:"rollback"`
	ID       string `json:"id"`
}

type Settings struct {
	ConsistencyLevel  ConsistencyLevel `json:"consistencyLevel"`
	ReplicationFactor int              `json:"replicationFactor"`
	Timeout           int              `json:"timeout"`
}
