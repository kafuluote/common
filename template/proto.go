package template

var (
	ImproveProto = `protoc --proto_path=. --micro_out=. --go_out=. ./*.proto
`
)
