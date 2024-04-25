package params

func LoadParams() []Param {
	return []Param{{Id: 1}}
}

func GetParams(p Param) map[string]string {
	// Implementation to convert Param struct to map
	return map[string]string{}
}

type Param struct {
	Id int
}
