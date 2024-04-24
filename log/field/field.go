package field

// 日志字段的类型
type FieldType int

// 日志字段类型常量
const (
	UnknownType FieldType = iota
	BoolType
	Int64Type
	Float64Type
	StringType
	ObjectType
)

type Field interface {
	Name() string
	Type() FieldType
	Value() any
}
