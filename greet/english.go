package greet

type course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []int
	Clases  map[uint]string
}

func (c *course) GetName() string {
	return c.Name
}

func (c *course) GetPrice() float64 {
	return c.Price
}
func (c *course) GetIsFree() bool {
	return c.IsFree
}

func (c *course) GetUserIDs() []int {
	return c.UserIDs
}

func (c *course) GetClases() map[uint]string {
	return c.Clases
}
func (c *course) SetName(name string) {
	c.Name = name
}

func (c *course) SetPrice(price float64) {
	c.Price = price
}

func (c *course) SetIsFree(isFree bool) {
	c.IsFree = isFree
}

func (c *course) SetUserIDs(userIDs []int) {
	c.UserIDs = userIDs
}

func (c *course) SetClases(clases map[uint]string) {
	c.Clases = clases
}

func NewCourse(name string, price float64, isFree bool, userIDs []int, clases map[uint]string) *course {
	return &course{
		Name:    name,
		Price:   price,
		IsFree:  isFree,
		UserIDs: userIDs,
		Clases:  clases,
	}
}

func English() string {
	return "Hello"
}
