package model

type Metadata struct {
	Id             int64  `gorm:"column:id;primaryKey" json:"id"`
	MimeType       string `gorm:"column:mime_type" json:"mime_type"`
	Path           string `gorm:"column:path" json:"path"`
	SubDomain      string `gorm:"column:sub_domain" json:"sub_domain"`
	TableReference string `gorm:"column:table_reference" json:"table_reference"`
	MarkAsDelete   bool   `gorm:"column:mark_as_delete" json:"mark_as_delete"`
	CreateDate     int64  `gorm:"column:create_date" json:"create_date"`
}

func (Metadata) TableName() string {
	return "metadata"
}
