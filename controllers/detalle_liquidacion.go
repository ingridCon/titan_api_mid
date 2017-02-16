package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"titan_api_mid/models"

	"github.com/astaxie/beego"
)

// operations for Liquidar
type DetalleLiquidacionController struct {
	beego.Controller
}

func (c *DetalleLiquidacionController) URLMapping() {
	c.Mapping("Detalle_liquidacion", c.InsertarDetallePreliquidacion)
}

func (c *DetalleLiquidacionController) InsertarDetallePreliquidacion() {
	fmt.Println("detalle")
	var v []int
	var tam int
	var IdPreliquidacion int
	var idPersonaString string
	var idPreliquidacionString string
	var d []models.DetallePreliquidacion

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		tam = len(v)
		IdPreliquidacion = v[tam-1]
		for i := 0; i < len(v)-1; i++ {
			idPersonaString = strconv.Itoa(v[i])
			idPreliquidacionString = strconv.Itoa(IdPreliquidacion)
			if err := getJson("http://"+beego.AppConfig.String("Urlcrud")+":"+beego.AppConfig.String("Portcrud")+"/"+beego.AppConfig.String("Nscrud")+"/detalle_preliquidacion?limit=0&query=Preliquidacion:"+idPreliquidacionString+",Persona:"+idPersonaString+"", &d); err == nil {

			} else {

			}
		}

		//http://localhost:8082/v1/detalle_preliquidacion?limit=0&query=Preliquidacion:7,Persona:184
	}
}
