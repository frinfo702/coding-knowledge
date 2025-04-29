package constants

import "fmt"

//go:generate stringer -type=Role
type Role int

const (
    _ Role = iota
    RoleAdmin
    RoleEditor
    RoleViewer
)

func (r Role) IsValid() bool { return r >= RoleAdmin && r <= RoleViewer }

func (r Role) MarshalJSON() ([]byte, error) {
    if !r.IsValid() {
        return nil, fmt.Errorf("invalid role: %d", r)
    }
    return []byte(fmt.Sprintf("%q", r.String())), nil
}
