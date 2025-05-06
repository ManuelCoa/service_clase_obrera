package interfaz

import paredes_vivienda "claseobrera/domains/object/tipo_paredes_vivienda"

type TipoParedInterface func() ([]paredes_vivienda.TipoParedesVivienda, error)

func TipoParedGetService() TipoParedInterface {
    return func() ([]paredes_vivienda.TipoParedesVivienda, error) {
        return paredes_vivienda.GetTiposParedService()
    }
}

func TipoParedGetServiceID(id int) TipoParedInterface {
    return func() ([]paredes_vivienda.TipoParedesVivienda, error) {
        return paredes_vivienda.GetTipoParedServiceID(id)
    }
}

func TipoParedPostService(data interface{}) error {
    return paredes_vivienda.PostTipoParedService(data)
}

func TipoParedPutService(data interface{}) error {
    return paredes_vivienda.PutTipoParedService(data)
}

func TipoParedDeleteService(id int) error {
    return paredes_vivienda.DeleteTipoParedService(id)
}