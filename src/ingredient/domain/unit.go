package domain

type Unit int

const (
	Unit_Undefined Unit = iota
	Unit_Gram
	Unit_Unit
	Unit_Cas
	Unit_Cac
	Unit_Milliliter
)

func (u Unit) String() string {
	switch u {
	case Unit_Undefined:
		return "?"
	case Unit_Gram:
		return "g"
	case Unit_Unit:
		return ""
	case Unit_Cas:
		return "cas"
	case Unit_Cac:
		return "cac"
	case Unit_Milliliter:
		return "ml"
	}

	return "undefined"
}

func (u *Unit) MarshalJSON() ([]byte, error) {
	return []byte("\"" + u.String() + "\""), nil
}
