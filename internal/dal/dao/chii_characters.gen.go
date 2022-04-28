// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNameCharacter = "chii_characters"

// Character mapped from table <chii_characters>
type Character struct {
	ID       uint32      `gorm:"column:crt_id;type:mediumint(8) unsigned;primaryKey;autoIncrement:true" json:"crt_id"`
	Name     string      `gorm:"column:crt_name;type:varchar(255);not null" json:"crt_name"`
	Role     uint8       `gorm:"column:crt_role;type:tinyint(4) unsigned;not null;index:crt_role,priority:1" json:"crt_role"` // 角色，机体，组织。。
	Infobox  string      `gorm:"column:crt_infobox;type:mediumtext;not null" json:"crt_infobox"`
	Summary  string      `gorm:"column:crt_summary;type:mediumtext;not null" json:"crt_summary"`
	Img      string      `gorm:"column:crt_img;type:varchar(255);not null" json:"crt_img"`
	Comment  uint32      `gorm:"column:crt_comment;type:mediumint(9) unsigned;not null;default:0" json:"crt_comment"`
	Collects uint32      `gorm:"column:crt_collects;type:mediumint(8) unsigned;not null" json:"crt_collects"`
	Dateline uint32      `gorm:"column:crt_dateline;type:int(10) unsigned;not null" json:"crt_dateline"`
	Lastpost uint32      `gorm:"column:crt_lastpost;type:int(11) unsigned;not null" json:"crt_lastpost"`
	Lock     int8        `gorm:"column:crt_lock;type:tinyint(4);not null;index:crt_lock,priority:1;default:0" json:"crt_lock"`
	ImgAnidb string      `gorm:"column:crt_img_anidb;type:varchar(255);not null" json:"crt_img_anidb"`
	AnidbID  uint32      `gorm:"column:crt_anidb_id;type:mediumint(8) unsigned;not null" json:"crt_anidb_id"`
	Ban      uint8       `gorm:"column:crt_ban;type:tinyint(3) unsigned;not null;index:crt_ban,priority:1;default:0" json:"crt_ban"`
	Redirect uint32      `gorm:"column:crt_redirect;type:int(10) unsigned;not null;default:0" json:"crt_redirect"`
	Nsfw     bool        `gorm:"column:crt_nsfw;type:tinyint(1) unsigned;not null" json:"crt_nsfw"`
	Fields   PersonField `gorm:"foreignKey:crt_id;polymorphic:Owner;polymorphicValue:crt" json:"fields"`
}

// TableName Character's table name
func (*Character) TableName() string {
	return TableNameCharacter
}
