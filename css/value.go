package css

type UnitType string

const (
	Pixels   UnitType = "px"
	Points            = "pt"
	Em                = "em"
	Rem               = "rem"
	Percent           = "%"
	Viewport          = "v_" // ??
)

type Value interface {
	Info() string
}

type StringValue struct {
	value string
}

type NumberValue struct {
	value string
	unit  UnitType
}
