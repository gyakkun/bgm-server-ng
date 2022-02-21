// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNamePersonSubjects = "chii_person_cs_index"

// PersonSubjects mapped from table <chii_person_cs_index>
type PersonSubjects struct {
	PrsnType      string `gorm:"column:prsn_type;type:enum('prsn','crt');primaryKey;autoIncrement:false" json:"prsn_type"`
	PersonID      uint32 `gorm:"column:prsn_id;type:mediumint(9) unsigned;primaryKey;autoIncrement:false;index:prsn_id,priority:1" json:"prsn_id"`
	PrsnPosition  uint16 `gorm:"column:prsn_position;type:smallint(5) unsigned;primaryKey;autoIncrement:false;index:prsn_position,priority:1" json:"prsn_position"` // 监督，原案，脚本,..
	SubjectID     uint32 `gorm:"column:subject_id;type:mediumint(9) unsigned;primaryKey;autoIncrement:false;index:subject_id,priority:1" json:"subject_id"`
	SubjectTypeID uint8  `gorm:"column:subject_type_id;type:tinyint(4) unsigned;not null;index:subject_type_id,priority:1" json:"subject_type_id"`
	Summary       string `gorm:"column:summary;type:mediumtext;not null" json:"summary"`
	PrsnAppearEps string `gorm:"column:prsn_appear_eps;type:mediumtext;not null" json:"prsn_appear_eps"` // 可选，人物参与的章节
}

// TableName PersonSubjects's table name
func (*PersonSubjects) TableName() string {
	return TableNamePersonSubjects
}
