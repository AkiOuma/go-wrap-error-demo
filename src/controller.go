package src

type controller struct {
	dao *dao
}

func NewController(dao *dao) *controller {
	return &controller{
		dao: dao,
	}
}

// find user name using id of user
func (c *controller) FindUserNameByID(id int) (string, error) {
	var name string
	row, err := c.dao.FindUserByID(id)
	if err != nil {
		return name, err
	}
	name = row["name"].(string)
	return name, err
}
