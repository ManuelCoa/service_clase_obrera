package politico_obrero

type PoliticoObrero struct {
	PoliticoObreroID int    `json:"id_politico_obrero"`
	ObreroCedula     int    `json:"cedula_obrero"`
	FuerzaPoliticaID int    `json:"id_fuerza_politica"`
	PPID             int    `json:"id_pp"`
	MSID             int    `json:"id_ms"`
	MisionesID       int    `json:"id_misiones"`
	RegistroUnoxDiez string `json:"registro_unoxdiez"`
	Observacion      string `json:"observacion"`
}