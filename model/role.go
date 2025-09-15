package model

type Role int

const (
	MEMBER Role = iota
	ADMIN
	ADMIN1
	ADMIN2
	MANAGER
)

func (r Role) String() string {
	return []string{
		"MEMBER",
		"ADMIN",
		"ADMIN1",
		"ADMIN2",
	}[r]
}
