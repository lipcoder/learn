package magazine

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

type Subscriber struct {
    Name        string
    Rate        float64
    Active      bool
    HomeAddress Address
}

type Employee struct {
    Name        string
    Address	//匿名字段（=嵌入字段）：struct 里只写类型，可以少写一层HomeAddress
}