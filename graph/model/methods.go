package model

import (
	"fmt"

	"github.com/kauewa/graphql_crossxhub/graph/db"
)

func (moto *Motos) GetByID(idmoto int64) {
	query := fmt.Sprintf("SELECT * FROM crossxhub.motos WHERE id = %v", idmoto)
	fmt.Println(query)
	rows, err := db.QueryDB(query)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&moto.ID, &moto.Nome, &moto.Foto)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

	}
}

func (tipo *Tipos) GetByID(idtipo int64) {
	query := fmt.Sprintf("SELECT * FROM crossxhub.tipos WHERE id = %v", idtipo)
	fmt.Println(query)
	rows, err := db.QueryDB(query)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&tipo.ID, &tipo.Nome, &tipo.Multiplo)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}
}

func (campeonato *Campeonatos) GetByID(idcampeonato int64) {
	query := fmt.Sprintf("SELECT id, nome, idpais FROM crossxhub.campeonatos WHERE id = %v", idcampeonato)
	fmt.Println(query)
	rows, err := db.QueryDB(query)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var idpais interface{}
		err := rows.Scan(&campeonato.ID, &campeonato.Nome, &idpais)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		if i, isint := idpais.(int64); isint {
			campeonato.Pais = &Pais{}
			campeonato.Pais.GetByID(i)
		}
	}
}

func (equipe *Equipes) GetByID(idequipe int64) {
	query := fmt.Sprintf("SELECT * FROM crossxhub.equipes WHERE id = %v", idequipe)
	fmt.Println(query)
	rows, err := db.QueryDB(query)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var idmoto, idpais, idcampeonato interface{}

		err := rows.Scan(&equipe.ID, &equipe.Nome, &idmoto, &idpais, &idcampeonato)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		if i, isint := idmoto.(int64); isint {
			equipe.Moto = &Motos{}
			equipe.Moto.GetByID(i)
		}
		if i, isint := idpais.(int64); isint {
			equipe.Pais = &Pais{}
			equipe.Pais.GetByID(i)
		}
		if i, isint := idcampeonato.(int64); isint {
			equipe.Campeonato = &Campeonatos{}
			equipe.Campeonato.GetByID(i)
		}
	}
}

func (pais *Pais) GetByID(idpais int64) {
	query := fmt.Sprintf("SELECT * FROM crossxhub.paises WHERE id = %v", idpais)
	fmt.Println(query)
	rows, err := db.QueryDB(query)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&pais.ID, &pais.Nome, &pais.Bandeira)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}
}
func Iterar_Moto_Pais_equipes(motosids []*interface{}, paisesids []*interface{}, equipes []*Equipes) {
	if len(motosids) == 0 {
		return
	}
	query := "SELECT * FROM crossxhub.paises"
	fmt.Println(query)
	rows, err := db.QueryDB(query)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer rows.Close()
	var paises []*Pais
	for rows.Next() {
		var pais Pais
		err := rows.Scan(&pais.ID, &pais.Nome, &pais.Bandeira)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		paises = append(paises, &pais)
	}

	for p_id, idpais := range paisesids {
		if v, isint := (*idpais).(int64); isint {
			for _, pais := range paises {
				if fmt.Sprintf("%d", v) == pais.ID {
					equipes[p_id].Pais = pais
				}
			}
		}

	}
	query = "SELECT * FROM crossxhub.motos"
	fmt.Println(query)
	rows, err = db.QueryDB(query)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer rows.Close()
	var motos []*Motos
	for rows.Next() {
		var moto Motos
		err := rows.Scan(&moto.ID, &moto.Nome, &moto.Foto)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		motos = append(motos, &moto)
	}
	for p_id, idmoto := range motosids {
		if v, isint := (*idmoto).(int64); isint {
			for _, moto := range motos {
				if fmt.Sprintf("%d", v) == moto.ID {
					equipes[p_id].Moto = moto
				}
			}
		}
	}
}

func Iterar_Equipes_Paises_pilotos(equipesids []*interface{}, paisesid []*interface{}, pilotos []*Pilotos) {
	if len(equipesids) == 0 {
		return
	}
	query := "SELECT id, nome, idmoto, idpais FROM crossxhub.equipes"
	fmt.Println(query)
	rows, err := db.QueryDB(query)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer rows.Close()
	var equipes []*Equipes
	var motosids []*interface{}
	var paisesids []*interface{}
	for rows.Next() {
		var idmoto, idpais interface{}
		var equipe Equipes
		err := rows.Scan(&equipe.ID, &equipe.Nome, &idmoto, &idpais)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		motosids = append(motosids, &idmoto)
		paisesids = append(paisesids, &idpais)
		equipes = append(equipes, &equipe)
	}
	Iterar_Moto_Pais_equipes(motosids, paisesids, equipes)
	for p_id, idequipe := range equipesids {
		if v, isint := (*idequipe).(int64); isint {
			for _, equipe := range equipes {
				if fmt.Sprintf("%d", v) == equipe.ID {
					pilotos[p_id].Equipe = equipe
				}
			}
		}
	}

	query = "SELECT * FROM crossxhub.paises"
	fmt.Println(query)
	rows, err = db.QueryDB(query)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer rows.Close()
	var paises []*Pais
	for rows.Next() {
		var pais Pais
		err := rows.Scan(&pais.ID, &pais.Nome, &pais.Bandeira)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		paises = append(paises, &pais)
	}
	for p_id, idpais := range paisesid {
		if v, isint := (*idpais).(int64); isint {
			for _, pais := range paises {
				if fmt.Sprintf("%d", v) == pais.ID {
					pilotos[p_id].Pais = pais
				}
			}
		}
	}
}

func Iterar_Pilotos_resultados(pilotosid []*interface{}, resultados []*ResultadoEtapas) {
	if len(pilotosid) == 0 {
		return
	}
	query := "SELECT * FROM crossxhub.pilotos"
	fmt.Println(query)
	rows, err := db.QueryDB(query)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer rows.Close()
	var pilotos []*Pilotos
	var equipesids, paisesids []*interface{}
	for rows.Next() {
		var idequipe interface{}
		var idpais interface{}
		var piloto Pilotos
		err := rows.Scan(&piloto.ID, &piloto.Nome, &idpais, &piloto.Numero, &piloto.Foto, &piloto.Mxon, &piloto.Datanascimento, &piloto.Altura, &idequipe, &piloto.Titulosconquistados, &piloto.Video, &piloto.Fotorecente, &piloto.Galeriafotoss, &piloto.Status)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		equipesids = append(equipesids, &idequipe)
		paisesids = append(paisesids, &idpais)
		pilotos = append(pilotos, &piloto)
	}
	Iterar_Equipes_Paises_pilotos(equipesids, paisesids, pilotos)
	for p_id, idpiloto := range pilotosid {
		if v, isint := (*idpiloto).(int64); isint {
			for _, piloto := range pilotos {
				if fmt.Sprintf("%d", v) == piloto.ID {
					resultados[p_id].Piloto = piloto
				}
			}
		}
	}

}

func Iterar_Paises_campeonatos(paisesid []*interface{}, campeonatos []*Campeonatos) {
	if len(paisesid) == 0 {
		return
	}
	query := "SELECT * FROM crossxhub.paises"
	fmt.Println(query)

	rows, err := db.QueryDB(query)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer rows.Close()
	var paises []*Pais
	for rows.Next() {
		var pais Pais
		err := rows.Scan(&pais.ID, &pais.Nome, &pais.Bandeira)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		paises = append(paises, &pais)
	}
	for p_id, idpais := range paisesid {
		if v, isint := (*idpais).(int64); isint {
			for _, pais := range paises {
				if fmt.Sprintf("%d", v) == pais.ID {
					campeonatos[p_id].Pais = pais
				}
			}
		}

	}
}

func Iterar_tipos_etapas(tiposid []*interface{}, etapas []*Etapas) {
	if len(tiposid) == 0 {
		return
	}
	query := "SELECT * FROM crossxhub.tipos"
	fmt.Println(query)
	rows, err := db.QueryDB(query)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer rows.Close()
	var tipos []*Tipos
	for rows.Next() {
		var tipo Tipos
		err := rows.Scan(&tipo.ID, &tipo.Nome, &tipo.Multiplo)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		tipos = append(tipos, &tipo)
		for p_id, idtipo := range tiposid {
			if v, isint := (*idtipo).(int64); isint {
				for _, tipo := range tipos {
					if fmt.Sprintf("%d", v) == tipo.ID {
						etapas[p_id].Tipos = tipo
					}
				}
			}
		}
	}
}

func (piloto *Pilotos) GetByID(idpiloto int64) {
	query := fmt.Sprintf("SELECT * FROM crossxhub.pilotos WHERE id = %v", idpiloto)
	fmt.Println(query)
	rows, err := db.QueryDB(query)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var idequipe interface{}
		var idpais interface{}

		err := rows.Scan(&piloto.ID, &piloto.Nome, &idpais, &piloto.Numero, &piloto.Foto, &piloto.Mxon, &piloto.Datanascimento, &piloto.Altura, &idequipe, &piloto.Titulosconquistados, &piloto.Video, &piloto.Fotorecente, &piloto.Galeriafotoss, &piloto.Status)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		if i, isint := idequipe.(int64); isint {
			piloto.Equipe = &Equipes{}
			piloto.Equipe.GetByID(i)

		}
		if i, isint := idpais.(int64); isint {
			piloto.Pais = &Pais{}
			piloto.Pais.GetByID(i)
		}
	}
}

func (etapa *Etapas) GetResults() {
	query := fmt.Sprintf("SELECT id, idpiloto, posicao FROM crossxhub.resultado_etapas WHERE idetapa = %s", etapa.ID)
	fmt.Println(query)
	rows, err := db.QueryDB(query)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer rows.Close()

	var resultados []*ResultadoEtapas
	var pilotosid []*interface{}
	for rows.Next() {
		var idpiloto interface{}

		var resultado ResultadoEtapas
		err := rows.Scan(&resultado.ID, &idpiloto, &resultado.Posicao)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		pilotosid = append(pilotosid, &idpiloto)

		resultados = append(resultados, &resultado)
	}
	Iterar_Pilotos_resultados(pilotosid, resultados)

	pointer := &etapa.Resultados
	*pointer = resultados
}
