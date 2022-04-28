// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"time"
)

const TableNameSubjectField = "chii_subject_fields"

// SubjectField mapped from table <chii_subject_fields>
type SubjectField struct {
	Sid      uint32    `gorm:"column:field_sid;type:mediumint(8) unsigned;primaryKey;autoIncrement:true;index:query_date,priority:1" json:"field_sid"`
	Tid      uint16    `gorm:"column:field_tid;type:smallint(6) unsigned;not null;index:sort_id,priority:1;default:0" json:"field_tid"`
	Tags     []byte    `gorm:"column:field_tags;type:mediumtext;not null" json:"field_tags"`
	Rate1    uint32    `gorm:"column:field_rate_1;type:mediumint(8) unsigned;not null;default:0" json:"field_rate_1"`
	Rate2    uint32    `gorm:"column:field_rate_2;type:mediumint(8) unsigned;not null;default:0" json:"field_rate_2"`
	Rate3    uint32    `gorm:"column:field_rate_3;type:mediumint(8) unsigned;not null;default:0" json:"field_rate_3"`
	Rate4    uint32    `gorm:"column:field_rate_4;type:mediumint(8) unsigned;not null;default:0" json:"field_rate_4"`
	Rate5    uint32    `gorm:"column:field_rate_5;type:mediumint(8) unsigned;not null;default:0" json:"field_rate_5"`
	Rate6    uint32    `gorm:"column:field_rate_6;type:mediumint(8) unsigned;not null;default:0" json:"field_rate_6"`
	Rate7    uint32    `gorm:"column:field_rate_7;type:mediumint(8) unsigned;not null;default:0" json:"field_rate_7"`
	Rate8    uint32    `gorm:"column:field_rate_8;type:mediumint(8) unsigned;not null;default:0" json:"field_rate_8"`
	Rate9    uint32    `gorm:"column:field_rate_9;type:mediumint(8) unsigned;not null;default:0" json:"field_rate_9"`
	Rate10   uint32    `gorm:"column:field_rate_10;type:mediumint(8) unsigned;not null;default:0" json:"field_rate_10"`
	Airtime  int8      `gorm:"column:field_airtime;type:tinyint(1) unsigned;not null;index:subject_airtime,priority:1" json:"field_airtime"`
	Rank     uint32    `gorm:"column:field_rank;type:int(10) unsigned;not null;index:field_rank,priority:1;default:0" json:"field_rank"`
	Year     int32     `gorm:"column:field_year;type:year(4);not null;index:field_year_mon,priority:1;index:field_year,priority:1" json:"field_year"` // 放送年份
	Mon      int8      `gorm:"column:field_mon;type:tinyint(2);not null;index:field_year_mon,priority:2" json:"field_mon"`                            // 放送月份
	WeekDay  int8      `gorm:"column:field_week_day;type:tinyint(1);not null" json:"field_week_day"`                                                  // 放送日(星期X)
	Date     time.Time `gorm:"column:field_date;type:date;not null;index:field_date,priority:1;index:query_date,priority:2" json:"field_date"`        // 放送日期
	Redirect uint32    `gorm:"column:field_redirect;type:mediumint(8) unsigned;not null;default:0" json:"field_redirect"`
}

// TableName SubjectField's table name
func (*SubjectField) TableName() string {
	return TableNameSubjectField
}
