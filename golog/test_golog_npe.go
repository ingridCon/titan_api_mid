package golog

import (
	"fmt"
	"strconv"
	models "titan_api_mid/models"
	. "github.com/mndrix/golog"
	"strings"
)

func CargarReglasPE(reglas string, pensionado models.InformacionPensionado, /*beneficiarios string*/) (rest []models.Respuesta) {

	var resultado []models.Respuesta
	temp := models.Respuesta{}
	var lista_descuentos []models.ConceptosResumen

	var cedulaProveedor = strconv.Itoa(pensionado.InformacionProveedor)
	var valorpension = strconv.Itoa(pensionado.ValorPensionAsignada)
	var pen string
	var lugarResidencia = pensionado.PensionadoEnExterior
	var lugar string
	var tpensionado = strconv.Itoa(pensionado.TipoPensionado)
	//var pensionadoFallecido = pensionado.PersonaFallecido

	if lugarResidencia == "S"{
		lugar = "1"
		}else{
			lugar = "2"
		}

		reglas = reglas + "valor_mesada_pensionado(" + cedulaProveedor + "," + valorpension + ")." + "\n"
		reglas = reglas + "tipo_pensionado(" + cedulaProveedor +","+ tpensionado +")." + "\n"
		reglas = reglas + "categoria_beneficiario(" + cedulaProveedor +","+ "0" +")." + "\n"
		reglas = reglas + "numHijos("+ "0" +")." + "\n"
		reglas = reglas + "residencia(" + cedulaProveedor + "," + lugar + ")." + "\n"

		/*if beneficiarios == " "{
		reglas = reglas + "numero_beneficiarios(" + cedulaProveedor +","+ "0"+")." + "\n"
		}else{
		reglas = reglas + "numero_beneficiarios(" + cedulaProveedor +","+ beneficiarios +")." + "\n"
		}*/

		reglas = reglas + "tipo_valor(" + "1" + ")." + "\n"
		//reglas = reglas + "valor_incremento_vigencia(" + incremento + ","+ año +")." + "\n"

		m := NewMachine().Consult(reglas)
		//if pensionadoFallecido == "N"{
		//fmt.Println("No fallecido")
		pension := m.ProveAll("pension_asignada(" + cedulaProveedor +",P).")
		for _, solution := range pension {
			Valor, _ := strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("P")), 64)
			pen = strconv.FormatFloat(Valor, 'f', 6, 64)
			temp_conceptos := models.ConceptosResumen{Nombre: "pensionadoPension",
				Valor: fmt.Sprintf("%.0f", Valor),
			}
			codigo := m.ProveAll("codigo_concepto(" + temp_conceptos.Nombre + ",C).")
			for _, cod := range codigo {
				temp_conceptos.Id, _ = strconv.Atoi(fmt.Sprintf("%s", cod.ByName_("C")))
			}
			lista_descuentos = append(lista_descuentos, temp_conceptos)
			temp.Conceptos = &lista_descuentos
			resultado = append(resultado, temp)
		}

		valor := m.ProveAll("aporte_fondoSoli(" + cedulaProveedor +",W).")
		for _, solution := range valor {
			Valor, _ := strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("W")), 64)
			var v int = int(Valor)
			c := strconv.Itoa(v)
			numero:= strings.Split(c,"")
			a, _ := strconv.Atoi(numero[len(numero)-1])
			b, _ := strconv.Atoi(numero[len(numero)-2])
			d, _ := strconv.Atoi(numero[len(numero)-3])
			fmt.Println("a")
			fmt.Println(a)
			var val string
			if  a < 5 || a > 5 || a == 5{				
				numero[len(numero)-1] = "0"
			}
			if  b < 5 {
				numero[len(numero)-2] = "0"
			}else{
				if  b > 5{
					numero[len(numero)-2] = "0"
					numero[len(numero)-3] = strconv.Itoa(d+1)
				}
			}
			for i := 0; i < len(numero); i++ {
				val = val + numero[i]
			}
			Valor, _ = strconv.ParseFloat(val,64)
			temp_conceptos := models.ConceptosResumen{Nombre: "fondoSolidaridad",
				Valor: fmt.Sprintf("%.0f", Valor),
			}

			codigo := m.ProveAll("codigo_concepto(" + temp_conceptos.Nombre + ",C).")
			for _, cod := range codigo {
				temp_conceptos.Id, _ = strconv.Atoi(fmt.Sprintf("%s", cod.ByName_("C")))
			}
			lista_descuentos = append(lista_descuentos, temp_conceptos)
			temp.Conceptos = &lista_descuentos
			resultado = append(resultado, temp)
		}

		aporte_salud := m.ProveAll("aporte_salud(" + cedulaProveedor +",S).")
		for _, solution := range aporte_salud {
			Valor, _ := strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("S")), 64)
			temp_conceptos := models.ConceptosResumen{Nombre: "salud",
				Valor: fmt.Sprintf("%.0f", Valor),
			}
			codigo := m.ProveAll("codigo_concepto(" + temp_conceptos.Nombre + ",C).")
			for _, cod := range codigo {
				temp_conceptos.Id, _ = strconv.Atoi(fmt.Sprintf("%s", cod.ByName_("C")))
			}
			lista_descuentos = append(lista_descuentos, temp_conceptos)
			temp.Conceptos = &lista_descuentos
			resultado = append(resultado, temp)
		}
		/*
		pago_subfamiliar:= m.ProveAll("pago_subfamiliar(" + cedulaProveedor +",F).")
		for _, solution := range pago_subfamiliar {
		Valor, _ := strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("F")), 64)
		temp_conceptos := models.ConceptosResumen{Nombre: "subsidioFamiliar",
		Valor: fmt.Sprintf("%.0f", Valor),
	}
	fmt.Println("pago_subfamiliar")
	fmt.Println(Valor)
	codigo := m.ProveAll("codigo_concepto(" + temp_conceptos.Nombre + ",C).")
	for _, cod := range codigo {
	temp_conceptos.Id, _ = strconv.Atoi(fmt.Sprintf("%s", cod.ByName_("C")))
}
lista_descuentos = append(lista_descuentos, temp_conceptos)
temp.Conceptos = &lista_descuentos
resultado = append(resultado, temp)
}

pago_subsidio_libros := m.ProveAll("pago_subsidio_libros(" + cedulaProveedor +",F).")
for _, solution := range pago_subsidio_libros {
Valor, _ := strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("F")), 64)
temp_conceptos := models.ConceptosResumen{Nombre: "subsidioLibros",
Valor: fmt.Sprintf("%.0f", Valor),
}
fmt.Println("pago_subsidio_libros")
fmt.Println(Valor)

codigo := m.ProveAll("codigo_concepto(" + temp_conceptos.Nombre + ",C).")
for _, cod := range codigo {
temp_conceptos.Id, _ = strconv.Atoi(fmt.Sprintf("%s", cod.ByName_("C")))
}
lista_descuentos = append(lista_descuentos, temp_conceptos)
temp.Conceptos = &lista_descuentos
resultado = append(resultado, temp)
}*/

//idProveedorString := strconv.Itoa(idProveedor)
novedades := m.ProveAll("info_concepto(" + cedulaProveedor + ","+ pen +",T,2017,N,R).")
for _, solution := range novedades {
	Valor, _ := strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("R")), 64)
	temp_conceptos := models.ConceptosResumen{Nombre: fmt.Sprintf("%s", solution.ByName_("N")),
		Valor: fmt.Sprintf("%.0f", Valor),
	}
	codigo := m.ProveAll("codigo_concepto(" + temp_conceptos.Nombre + ",C).")
	for _, cod := range codigo {
		temp_conceptos.Id, _ = strconv.Atoi(fmt.Sprintf("%s", cod.ByName_("C")))
	}

	lista_descuentos = append(lista_descuentos, temp_conceptos)
	temp.Conceptos = &lista_descuentos
	resultado = append(resultado, temp)
}
/*}else{
if persona = "S"{
pension := m.ProveAll("pension_asignada(" + cedulaProveedor +",P).")
for _, solution := range pension {
Valor, _ := strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("P")), 64)
pen = strconv.FormatFloat(Valor, 'f', 6, 64)
temp_conceptos := models.ConceptosResumen{Nombre: "pensionadoPension",
Valor: fmt.Sprintf("%.0f", Valor),
}
fmt.Println("pensionadoPension")
fmt.Println(Valor)
codigo := m.ProveAll("codigo_concepto(" + temp_conceptos.Nombre + ",C).")
for _, cod := range codigo {
temp_conceptos.Id, _ = strconv.Atoi(fmt.Sprintf("%s", cod.ByName_("C")))
}
lista_descuentos = append(lista_descuentos, temp_conceptos)
temp.Conceptos = &lista_descuentos
resultado = append(resultado, temp)
}

valor := m.ProveAll("aporte_fondoSoli(" + cedulaProveedor +",W).")
for _, solution := range valor {
Valor, _ := strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("W")), 64)
temp_conceptos := models.ConceptosResumen{Nombre: "fondoSolidaridad",
Valor: fmt.Sprintf("%.0f", Valor),
}
fmt.Println("fondoSolidaridad")
fmt.Println(Valor)
codigo := m.ProveAll("codigo_concepto(" + temp_conceptos.Nombre + ",C).")
for _, cod := range codigo {
temp_conceptos.Id, _ = strconv.Atoi(fmt.Sprintf("%s", cod.ByName_("C")))
}
lista_descuentos = append(lista_descuentos, temp_conceptos)
temp.Conceptos = &lista_descuentos
resultado = append(resultado, temp)
}

aporte_salud := m.ProveAll("aporte_salud(" + cedulaProveedor +",S).")
for _, solution := range aporte_salud {
Valor, _ := strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("S")), 64)
fmt.Println("salud")
fmt.Println(Valor)
temp_conceptos := models.ConceptosResumen{Nombre: "salud",
Valor: fmt.Sprintf("%.0f", Valor),
}
codigo := m.ProveAll("codigo_concepto(" + temp_conceptos.Nombre + ",C).")
for _, cod := range codigo {
temp_conceptos.Id, _ = strconv.Atoi(fmt.Sprintf("%s", cod.ByName_("C")))
}
lista_descuentos = append(lista_descuentos, temp_conceptos)
temp.Conceptos = &lista_descuentos
resultado = append(resultado, temp)
}
}
}*/
/*
}else{
if pensionadoFallecido == "S"{
fmt.Println("Persona Fallecida SI")
pension := m.ProveAll("pension_asignada(" + cedulaProveedor +",P).")
for _, solution := range pension {
Valor, _ := strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("P")), 64)
pen = strconv.FormatFloat(Valor, 'f', 6, 64)
temp_conceptos := models.ConceptosResumen{Nombre: "pensionadoPension",
Valor: fmt.Sprintf("%.0f", Valor),
}
fmt.Println("pensionadoPension")
fmt.Println(Valor)
codigo := m.ProveAll("codigo_concepto(" + temp_conceptos.Nombre + ",C).")
for _, cod := range codigo {
temp_conceptos.Id, _ = strconv.Atoi(fmt.Sprintf("%s", cod.ByName_("C")))
}
lista_descuentos = append(lista_descuentos, temp_conceptos)
temp.Conceptos = &lista_descuentos
resultado = append(resultado, temp)
}

valor := m.ProveAll("aporte_fondoSoli(" + cedulaProveedor +",W).")
for _, solution := range valor {
Valor, _ := strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("W")), 64)
temp_conceptos := models.ConceptosResumen{Nombre: "fondoSolidaridad",
Valor: fmt.Sprintf("%.0f", Valor),
}
fmt.Println("fondoSolidaridad")
fmt.Println(Valor)
codigo := m.ProveAll("codigo_concepto(" + temp_conceptos.Nombre + ",C).")
for _, cod := range codigo {
temp_conceptos.Id, _ = strconv.Atoi(fmt.Sprintf("%s", cod.ByName_("C")))
}
lista_descuentos = append(lista_descuentos, temp_conceptos)
temp.Conceptos = &lista_descuentos
resultado = append(resultado, temp)
}

aporte_salud := m.ProveAll("aporte_salud(" + cedulaProveedor +",S).")
for _, solution := range aporte_salud {
Valor, _ := strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("S")), 64)
fmt.Println("salud")
fmt.Println(Valor)
temp_conceptos := models.ConceptosResumen{Nombre: "salud",
Valor: fmt.Sprintf("%.0f", Valor),
}
codigo := m.ProveAll("codigo_concepto(" + temp_conceptos.Nombre + ",C).")
for _, cod := range codigo {
temp_conceptos.Id, _ = strconv.Atoi(fmt.Sprintf("%s", cod.ByName_("C")))
}
lista_descuentos = append(lista_descuentos, temp_conceptos)
temp.Conceptos = &lista_descuentos
resultado = append(resultado, temp)
}
}
}*/

return resultado
}

func CargarReglasSustitutosPE(reglas string, sustituto models.Sustituto, cedulaPensionado string, pension string) (rest []models.Respuesta) {

	var resultado []models.Respuesta
	temp := models.Respuesta{}
	var lista_descuentos []models.ConceptosResumen


	var cedulaProveedor = strconv.Itoa(sustituto.Beneficiario)
	var porcent = strconv.Itoa(sustituto.Porcentaje)
	fmt.Println("porcentajeeeee")
	fmt.Println(porcent)

	//var lugarResidencia = pensionado.PensionadoEnExterior
	//var lugar string

	/*
	if lugarResidencia == "S"{
	lugar = "1"
	}else{
	lugar = "2"
}
*/
reglas = reglas + "valor_mesada_pensionado(" + cedulaPensionado + "," + pension + ")." + "\n"
reglas = reglas + "porcentaje(" + cedulaProveedor + "," + porcent + ")." + "\n"

//reglas = reglas + "residencia(" + cedulaProveedor + "," + lugar + ")." + "\n"

m := NewMachine().Consult(reglas)
pensionSust := m.ProveAll("pension_asignada_sust(" + cedulaProveedor +",P).")
for _, solution := range pensionSust {
	Valor, _ := strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("P")), 64)
	//pen = strconv.FormatFloat(Valor, 'f', 6, 64)
	temp_conceptos := models.ConceptosResumen{Nombre: "pensionadoPension",
		Valor: fmt.Sprintf("%.0f", Valor),
	}
	fmt.Println("pension sust")
	fmt.Println(Valor)
	codigo := m.ProveAll("codigo_concepto(" + temp_conceptos.Nombre + ",C).")
	for _, cod := range codigo {
		temp_conceptos.Id, _ = strconv.Atoi(fmt.Sprintf("%s", cod.ByName_("C")))
	}
	lista_descuentos = append(lista_descuentos, temp_conceptos)
	temp.Conceptos = &lista_descuentos
	resultado = append(resultado, temp)
}

fondo := m.ProveAll("aporte_fondoSoli_sust(" + cedulaProveedor +",W).")
for _, solution := range fondo {
	Valor, _ := strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("W")), 64)
	var v int = int(Valor)
	c := strconv.Itoa(v)
	numero:= strings.Split(c,"")
	a, _ := strconv.Atoi(numero[len(numero)-1])
	b, _ := strconv.Atoi(numero[len(numero)-2])
	d, _ := strconv.Atoi(numero[len(numero)-3])
	var val string
	if  a < 5 || a > 5{
		numero[len(numero)-1] = "0"
	}
	if  b < 5 {
		numero[len(numero)-2] = "0"
	}else{
		if  b > 5{
			numero[len(numero)-2] = "0"
			numero[len(numero)-3] = strconv.Itoa(d+1)
		}
	}
	for i := 0; i < len(numero); i++ {
		val = val + numero[i]
	}
	Valor, _ = strconv.ParseFloat(val,64)
	fmt.Println(Valor)
	temp_conceptos := models.ConceptosResumen{Nombre: "fondoSolidaridad",
		Valor: fmt.Sprintf("%.0f", Valor),
	}
	fmt.Println("fondosust")
	fmt.Println(Valor)
	codigo := m.ProveAll("codigo_concepto(" + temp_conceptos.Nombre + ",C).")
	for _, cod := range codigo {
		temp_conceptos.Id, _ = strconv.Atoi(fmt.Sprintf("%s", cod.ByName_("C")))
	}
	lista_descuentos = append(lista_descuentos, temp_conceptos)
	temp.Conceptos = &lista_descuentos
	resultado = append(resultado, temp)
}

aporte_salud := m.ProveAll("aporte_salud_sust(" + cedulaProveedor +",S).")
for _, solution := range aporte_salud {
	Valor, _ := strconv.ParseFloat(fmt.Sprintf("%s", solution.ByName_("S")), 64)
	temp_conceptos := models.ConceptosResumen{Nombre: "salud",
		Valor: fmt.Sprintf("%.0f", Valor),
	}
	codigo := m.ProveAll("codigo_concepto(" + temp_conceptos.Nombre + ",C).")
	for _, cod := range codigo {
		temp_conceptos.Id, _ = strconv.Atoi(fmt.Sprintf("%s", cod.ByName_("C")))
	}
	lista_descuentos = append(lista_descuentos, temp_conceptos)
	temp.Conceptos = &lista_descuentos
	resultado = append(resultado, temp)
}
return resultado
}
