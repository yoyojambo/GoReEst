package structs

type Sort struct {
    Direction string `json:"direction"`
    Field     string `json:"field"`
}

type Boundary struct {
    Coordinates [][][]float64 `json:"coordinates"`
}

type Payload struct {
    Limit    int      `json:"limit"`
    Offset   int      `json:"offset"`
    Sort     Sort     `json:"sort"`
    Boundary Boundary `json:"boundary"`
}


//Estructura para insert a base de datos
type PropiedadAsociada struct {
	IdPropiedadAsociada int     `json:"id_propiedad_asociada"`
	Lat                 float64 `json:"lat"`
	Long                float64 `json:"long"`
	Ciudad              string  `json:"ciudad"`
	Estado              string  `json:"estado"`
	Address             string  `json:"address"`
	Precio              float64 `json:"precio"`
	SqFt                int     `json:"sq_ft"`
	TipoTerreno         string  `json:"tipo_terreno"`
	Status              string  `json:"status"`
}
