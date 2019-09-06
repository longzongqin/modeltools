package generate

type Field struct {
	Field string `gorm:"column:Field"`
	Type string `gorm:"column:Type"`
	Null string `gorm:"column:Null"`
	Key string `gorm:"column:Key"`
	Default string `gorm:"column:Default"`
	Extra string `gorm:"column:Extra"`
	Privileges string `gorm:"column:Privileges"`
	Comment string `gorm:"column:Comment"`
}
