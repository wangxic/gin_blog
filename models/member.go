package models

import (
	"time"
)

type Member struct {
	Mid      uint
	Groupid  int
	Mobile   string
	Account  string
	Password string
	Salt     string
	Ctime    time.Time
	Edittime time.Time
	Nick     string
}

//RecordNotFound 我查询到数据是不报错！

func (member *Member) First(account string) *Member {
	orm.Where("account=?", account).First(member).RecordNotFound()
	return member
}
func (member *Member) GetFirst(account string, password string) *Member {

	orm.Where("account=?", account).Where("password=?", password).First(member).RecordNotFound()
	return member
}
