package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func ExecuteSQLMigrations(db *sql.DB) {
	// busca todos os arquivos sql na pasta migration, se der algum erro para a aplicacao e manda um erro
	files, err := filepath.Glob(filepath.Join("src/configuration/database/migrations", "*.sql"))
	if err != nil {
		log.Fatalf("Erro ao listar arquivos .sql: %v", err)
	}

	// organiza os arquivos em ordem
	sort.Strings(files)

	// faz um loop em cada arquivo
	for _, file := range files {
		// so um log pra aparecer no terminal qual arquivo ta rodando a migration pra caso der erro saber em qual foi
		log.Printf("Executando migração: %s", file)

		// le o conteudo do arquivo, se der algum erro para a aplicacao e manda um erro
		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("Erro ao ler o arquivo %s: %v", file, err)
		}

		// tira os espacamentos dos arquivos, e depois testa se ele esta vazio, se estiver vazio passa pro proximo arquivo
		queries := strings.TrimSpace(string(content))
		if queries == "" {
			continue
		}
		fmt.Printf("Executando SQL de %s:\n%s\n", file, queries)

		// aqui executa o sql rresente no arquivo, caso de erro apresenta o erro e para a aplicacao
		_, err = db.Exec(queries)
		if err != nil {
			log.Fatalf("Erro ao executar o SQL de %s: %v", file, err)
		}
	}
}
