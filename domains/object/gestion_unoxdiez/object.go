package gestion_unoxdiez

type GestionUnoxDiez struct {
	UnoxDiezID       int    `json:"id_unoxdiez"`
	ObreroCedula     int    `json:"cedula_obrero"`
	CentroVotacionID int    `json:"id_centro_votacion"`
	NombresApellidos string `json:"nombres_apellidos"`
	Cedula           string `json:"cedula"`
	Telefono         string `json:"telefono"`
	Observacion      string `json:"observacion"`
}