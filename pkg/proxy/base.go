package proxy

type Base struct {
	Name   string `yaml:"name" json:"name"`
	Server string `yaml:"server" json:"server"`
	Port   int    `yaml:"port" json:"port"`
	Type   string `yaml:"type" json:"type"`
	UDP    bool   `yaml:"udp,omitempty" json:"udp,omitempty"`
}

func (b *Base) TypeName() string {
	if b.Type == "" {
		return "unknown"
	}
	return b.Type
}

func (b *Base) SetName(name string) {
	b.Name = name
}

func (b *Base) SetIP(ip string) {
	b.Server = ip
}

func (b *Base) BaseInfo() *Base {
	return b
}

func (b *Base) Clone() Base {
	c := *b
	return c
}

type Proxy interface {
	String() string
	ToClash() string
	ToSurge() string
	Identifier() string
	SetName(name string)
	SetIP(ip string)
	TypeName() string
	BaseInfo() *Base
	Clone() Proxy
}
