package facade

type InvalidData struct {
	Frequency int
	Type      string
}

type InvalidKey struct {
	Endpoint   string
	Expression string
}

type InvalidExpression map[InvalidKey]InvalidData
