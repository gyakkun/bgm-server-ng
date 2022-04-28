// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNameIndex = "chii_index"

// Index mapped from table <chii_index>
type Index struct {
	ID           uint32 `gorm:"column:idx_id;type:mediumint(8);primaryKey;autoIncrement:true;uniqueIndex:mid,priority:1" json:"idx_id"` // 自动id
	Type         uint8  `gorm:"column:idx_type;type:tinyint(3) unsigned;not null;index:idx_type,priority:1;default:0" json:"idx_type"`
	Title        string `gorm:"column:idx_title;type:varchar(80);not null" json:"idx_title"`                                                // 标题
	Desc         string `gorm:"column:idx_desc;type:mediumtext;not null" json:"idx_desc"`                                                   // 简介
	Replies      uint32 `gorm:"column:idx_replies;type:mediumint(8) unsigned;not null;default:0" json:"idx_replies"`                        // 回复数
	SubjectTotal uint32 `gorm:"column:idx_subject_total;type:mediumint(8) unsigned;not null;default:0" json:"idx_subject_total"`            // 内含条目总数
	Collects     uint32 `gorm:"column:idx_collects;type:mediumint(8);not null;index:idx_collects,priority:1;default:0" json:"idx_collects"` // 收藏数
	Stats        string `gorm:"column:idx_stats;type:mediumtext;not null" json:"idx_stats"`
	Dateline     int32  `gorm:"column:idx_dateline;type:int(10);not null" json:"idx_dateline"` // 创建时间
	Lasttouch    uint32 `gorm:"column:idx_lasttouch;type:int(10) unsigned;not null" json:"idx_lasttouch"`
	UID          uint32 `gorm:"column:idx_uid;type:mediumint(8);not null;index:idx_uid,priority:1" json:"idx_uid"` // 创建人UID
	Ban          bool   `gorm:"column:idx_ban;type:tinyint(1) unsigned;not null;index:idx_ban,priority:1;default:0" json:"idx_ban"`
}

// TableName Index's table name
func (*Index) TableName() string {
	return TableNameIndex
}
