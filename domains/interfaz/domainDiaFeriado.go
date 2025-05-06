package interfaz

import (
	diasferiados "claseobrera/domains/object/dia_feriado"
)

type DiaFeriadoInterface func() ([]diasferiados.DiaFeriado, error)

func DiaFeriadoGetService() DiaFeriadoInterface {
	return func() ([]diasferiados.DiaFeriado, error) {
		return diasferiados.GetDiaFeriadoService()
	}
}

func DiaFeriadoGetServiceID(id int) DiaFeriadoInterface {
	return func() ([]diasferiados.DiaFeriado, error) {
		return diasferiados.GetDiaFeriadoServiceID(id)
	}
}

func DiaFeriadoPostService(data interface{}) error {
	return diasferiados.PostDiaFeriadoService(data)
}

func DiaFeriadoPutService(data interface{}) error {
	return diasferiados.PutDiaFeriadoService(data)
}

func DiaFeriadoDeleteService(id int) error {
	return diasferiados.DeleteDiaFeriadoService(id)
}