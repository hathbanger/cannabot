package controllers

import (
//	"encoding/json"

	"github.com/astaxie/beego"
	models "cannabot/cannabot_lib/models"
)

type MachineOperator struct {
	beego.Controller
	Instructions    []interface{}   `json:"instructions"`
	CurrentReport 	*models.Report 	`json:"report"`
}

func (mo *MachineOperator) SendReport(/*r *Report*/) {
	report := models.InitTestReport()
	mo.Data["json"] = report
//	mo.ServeJSON()
}

