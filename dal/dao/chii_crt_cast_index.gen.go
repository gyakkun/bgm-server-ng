// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNameCast = "chii_crt_cast_index"

// Cast mapped from table <chii_crt_cast_index>
type Cast struct {
	CharacterID   uint32    `gorm:"column:crt_id;type:mediumint(9) unsigned;primaryKey"`
	PersonID      uint32    `gorm:"column:prsn_id;type:mediumint(9) unsigned;primaryKey"`
	SubjectID     uint32    `gorm:"column:subject_id;type:mediumint(9) unsigned;primaryKey"`
	SubjectTypeID uint8     `gorm:"column:subject_type_id;type:tinyint(3) unsigned;not null"` // 根据人物归类查询角色，动画，书籍，游戏
	Summary       string    `gorm:"column:summary;type:varchar(255);not null"`                // 幼年，男乱马，女乱马，变身形态，少女形态。。
	Character     Character `gorm:"foreignKey:crt_id;references:crt_id" json:"character"`
	Subject       Subject   `gorm:"foreignKey:subject_id;references:subject_id" json:"subject"`
	Person        Person    `gorm:"foreignKey:prsn_id;references:prsn_id" json:"person"`
}

// TableName Cast's table name
func (*Cast) TableName() string {
	return TableNameCast
}
