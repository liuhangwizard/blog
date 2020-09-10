package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "vb-server/constant"
	"vb-server/model"
	"vb-server/response"
	VBLog "vb-server/utils/log"
)

type VBControllerTagRecord struct {
	VBControllerBase
	modelTagRecord *model.VBModelTagRecord
}



func (c *VBControllerTagRecord) RecordTagList(ctx *gin.Context){


	result,err:=c.modelTagRecord.DataChartTagList()
	if err !=nil{
		VBLog.Print.Info("[TagRecord][查询标签记录失败]" + err.Error())
		ctx.JSON(http.StatusInternalServerError, response.Error(COMMON_ERROR, map[string]interface{}{}))
		return;
	}

	//success
	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS,gin.H{
		"info":result,
	}))
}


func NewVBControllerTagRecord() VBControllerTagRecord {

	controllerTag := VBControllerTagRecord{
		modelTagRecord:model.NewModelTagRecord(),
	}
	return controllerTag
}

