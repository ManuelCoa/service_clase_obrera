package vehiculo_obrero

type VehiculoObrero struct {
	VehiculoID      int    `json:"id_vehiculo"`
	MarcaVehiculoID int    `json:"id_marca_vehiculo"`
	ObreroCedula    int    `json:"cedula_obrero"`
	Placa           string `json:"placa_vehiculo"`
	Modelo          string `json:"modelo_vehiculo"`
	Color           string `json:"color_vehiculo"`
	NivelGasolina   string `json:"nivel_gasolina"`
}