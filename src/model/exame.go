package model

type ExameClinico struct {
    InspecaoColo string `json:"inspecao_colo"`
    SinaisDST    string `json:"sinais_dst"`
    Observacoes  string `json:"observacoes"`
}