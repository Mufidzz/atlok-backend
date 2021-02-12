package jwt

type Header struct {
	Alg string
	Typ string
}

type Body struct {
	Uid uint
	Iat int64
	Acs uint
}
