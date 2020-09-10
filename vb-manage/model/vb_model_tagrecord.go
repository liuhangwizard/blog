package model

import (
	"time"
	. "vb-server/database/mysql"
	"vb-server/response"
)

type VBModelTagRecord struct {
	Id           uint32    `json:"id"`
	TagId        uint32    `json:"tagid"`
	ReaderIp     string    `json:"readerip"`
	ReaderCity   string    `json:"readercity"`
	ReaderCityId string    `json:"readercityid"`
	ReaderClient string    `json:"readerclient"`
	ReaderDate   time.Time `json:"readerdate"`
}

func NewModelTagRecord() *VBModelTagRecord {
	return &VBModelTagRecord{}
}

func (m *VBModelTagRecord) DataChartTagList()([]response.VBResultTagRecordChartList,error){

	//data
	data := Main().Table("vb_tag").Select([]string{
		"vb_tag.id",
		"vb_tag.name",
		"COUNT(vb_tag_record.id) AS record",
	})

	//join
	data = data.Joins("LEFT JOIN vb_tag_record ON vb_tag.id=vb_tag_record.tag_id AND vb_tag_record.reader_date")

	//group order
	data=data.Group("vb_tag.id,vb_tag.name").Order("record DESC")

	//result
	result:=[]response.VBResultTagRecordChartList{}
	err:=data.Limit(12).Find(&result).Error

	return result,err

}