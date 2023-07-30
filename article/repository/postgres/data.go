package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"nossobr/domain"
)

type PGArticle struct {
	DB *sql.DB
}

func (pg *PGArticle) Get(ctx context.Context, uf, slug *string) (res *domain.Article, err error) {
	res = new(domain.Article)

	var info []byte
	if err = pg.DB.QueryRowContext(ctx, `
		SELECT TA.id, TA.conteudo, TA.cidade_id, TA.info::JSON, TA.criacao, TA.atualizacao 
		FROM t_artigos TA
		JOIN t_cidades TC ON TC.id = TA.cidade_id 
		JOIN t_estados TE ON TE.id = TC.estado_id
		WHERE TC.slug = $1 AND TE.sigla = $2`, slug, uf,
	).Scan(&res.ID, &res.Conteudo, &res.CidadeID, &info, &res.Criacao, &res.Atualizacao); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(info, &res.Info); err != nil {
		return nil, err
	}

	return
}
