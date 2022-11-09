package titaniumcrud

import "strconv"

// EnumOP op enum
type EnumOP string

const (
	// EQ op string
	EQ EnumOP = "eq"
	// NE op string
	NE EnumOP = "ne"
	// GT op string
	GT EnumOP = "gt"
	// GE op string
	GE EnumOP = "ge"
	// LT op string
	LT EnumOP = "lt"
	// LE op string
	LE EnumOP = "le"
	// Limit op string
	Limit EnumOP = "limit"
)

// Condition for curd
type Condition struct {
	conditions map[string]map[EnumOP]string
}

// NewCondition create new condition
func NewCondition() *Condition {
	return &Condition{conditions: make(map[string]map[EnumOP]string)}
}

// EQ operation
func (c *Condition) EQ(key string, value string) {
	newMap := make(map[EnumOP]string)
	newMap[EQ] = value
	c.conditions[key] = newMap
}

// NE operation
func (c *Condition) NE(key string, value string) {
	newMap := make(map[EnumOP]string)
	newMap[NE] = value
	c.conditions[key] = newMap
}

// GT operation
func (c *Condition) GT(key string, value string) {
	newMap := make(map[EnumOP]string)
	newMap[GT] = value
	c.conditions[key] = newMap
}

// GE operation
func (c *Condition) GE(key string, value string) {
	newMap := make(map[EnumOP]string)
	newMap[GE] = value
	c.conditions[key] = newMap
}

// LT operation
func (c *Condition) LT(key string, value string) {
	newMap := make(map[EnumOP]string)
	newMap[LT] = value
	c.conditions[key] = newMap
}

// LE operation
func (c *Condition) LE(key string, value string) {
	newMap := make(map[EnumOP]string)
	newMap[LE] = value
	c.conditions[key] = newMap
}

// Limit operation
func (c *Condition) Limit(count int) {
	c.limit(0, count)
}

// LimitOffset operation
func (c *Condition) LimitOffset(offset int, count int) {
	c.limit(offset, count)
}

func (c *Condition) limit(offset int, count int) {
	newMap := make(map[EnumOP]string)
	if offset < 0 {
		offset = 0
	}
	if count < 0 {
		count = 0
	}
	newMap[Limit] = strconv.Itoa(offset) + "," + strconv.Itoa(count)
	c.conditions["limit"] = newMap
}

// GetConditions return conditions
func (c *Condition) GetConditions() map[string]map[EnumOP]string {
	return c.conditions
}

// SetConditions set conditions
func (c *Condition) SetConditions(newMap map[string]map[EnumOP]string) {
	c.conditions = newMap
}
