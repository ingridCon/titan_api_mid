package controllers

import (
	"encoding/json"
	"fmt"
	"time"
	"titan_api_mid/models"

	"github.com/astaxie/beego"
)

// operations for Liquidar
type LiquidarController struct {
	beego.Controller
}

func (c *LiquidarController) URLMapping() {
	c.Mapping("Liquidar", c.Liquidar)
}

func (c *LiquidarController) Liquidar() {
	var idLiquidacion interface{}
	var v models.Preliquidacion

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		liquidacion := models.Liquidacion{Id: v.Id, NombreLiquidacion: v.Nombre, Nomina: &models.Nomina{Id: v.Nomina.Id}, EstadoLiquidacion: v.Estado, FechaLiquidacion: time.Now(), FechaInicio: v.FechaInicio, FechaFin: v.FechaFin}

		fmt.Println(liquidacion)

		if err := sendJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/liquidacion", "POST", &idLiquidacion, &liquidacion); err == nil {

		}
	}
}
