package models

type Classifys struct {
	Id   int
	Name string
	Pid  int
	Sort int
}

func (_ *Classifys) List(like string, page int) ([]Classifys, int) {
	pageSize := 20
	total := 0

	var classifyList []Classifys
	DBCount := orm.Model(&classifyList)
	if like != "" {
		DBCount = DBCount.Where("name like ?", "%"+like+"%")
	}

	DBCount.Count(&total)

	DB := orm.Model(&classifyList).Limit(pageSize).Offset((page - 1) * pageSize)
	DB = DB.Select("id,name,pid,sort").Order("sort DESC")
	if like != "" {
		DB = DB.Where("name like ?", "%"+like+"%")
	}
	DB.Order("id DESC").Find(&classifyList)

	return classifyList, total
}

//获取单条
func (classify *Classifys) GetRow(id int) *Classifys {
	orm.Where(&Classifys{Id: id}).First(classify)
	return classify
}

//返回影响的行数
func (classify *Classifys) Del(id int) int64 {
	ret := classify.GetRow(id)
	if ret.Id == 0 {
		return 0
	}
	rowsAffected := orm.Delete(ret).RowsAffected
	return rowsAffected
}

func (classify *Classifys) Update(id, pid, sort int, name string) int64 {
	ret := classify.GetRow(id)
	if ret.Id == 0 {
		return 0
	}
	RowsAffecte := orm.Model(ret).Updates(map[string]interface{}{"pid": pid, "sort": sort, "name": name}).RowsAffected
	return RowsAffecte
}

func (_ *Classifys) Insert(pid, sort int, name string) int {
	classifyd := &Classifys{Pid: pid, Sort: sort, Name: name}

	orm.Create(&classifyd)
	return classifyd.Id
}
