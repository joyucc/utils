package godate


//EqualTo 比较两个时间是否一样
func (d *goDate) EqualTo(carbon *goDate) bool {
	if d.Timestamp() == carbon.Timestamp() {
		return true
	}
	return false
}

//NotEqualTo 比较两个时间是否不一样
func (d *goDate) NotEqualTo(carbon *goDate) bool {

	if d.Timestamp() != carbon.Timestamp() {
		return true
	}
	return false
}

//GreaterThan 比较时间是否比目标大
func (d *goDate) GreaterThan(carbon *goDate) bool {
	if d.Timestamp() > carbon.Timestamp() {
		return true
	}
	return false
}

//GreaterThanOrEqualTo 比较时间是否比目标大于或者等于
func (d *goDate) GreaterThanOrEqualTo(carbon *goDate) bool {
	if d.Timestamp() >= carbon.Timestamp() {
		return true
	}
	return false
}

//LessThan 比较时间是否比目标小
func (d *goDate) LessThan(carbon *goDate) bool {
	if d.Timestamp() < carbon.Timestamp() {
		return true
	}
	return false
}

//LessThanOrEqualTo 比较时间是否比目标小或者等于
func (d *goDate) LessThanOrEqualTo(carbon *goDate) bool {
	if d.Timestamp() <= carbon.Timestamp() {
		return true
	}
	return false
}

//Between 比较当前值是否是在 first 和second 之间
func (d *goDate) Between(first, second *goDate) bool {
	if d.Timestamp() < first.Timestamp() || d.Timestamp() > second.Timestamp() {
		return false
	}

	return true
}

//After 如果c代表的时间点在u之后，返回真；否则返回假。
func (d *goDate) After(u *goDate) bool {
	return d.Time.After(u.Time)
}

//Before 如果c代表的时间点在u之前，返回真；否则返回假。
func (d *goDate) Before(u *goDate) bool {
	return d.Time.Before(u.Time)
}
