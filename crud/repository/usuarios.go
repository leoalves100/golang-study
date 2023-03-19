package repository

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

type Erro struct {
	Erro string `json:"erro"`
}

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		fmt.Println(erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco"))
		fmt.Println(erro)
	}
	defer db.Close()

	// PREPARE STATEMENT
	statemant, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar statemant"))
		fmt.Println(erro)
	}
	defer statemant.Close()

	insercao, erro := statemant.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statemant"))
		fmt.Println(erro)
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido"))
		fmt.Println(erro)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserido)))
}

// BuscarUsuarios Busca todos os usuários
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
	}
	defer db.Close()

	linhas, err := db.Query("select * from usuarios")
	if err != nil {
		w.Write([]byte("Erro ao buscar os usuários!"))
	}
	defer linhas.Close()

	var usuarios []usuario

	// Intera sobre todos os usuários retornados
	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter usuários para json!"))
	}

}

// BuscarUsuario Busca um usuário específico
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	ID, erro := strconv.ParseUint(param["id"], 10, 16)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados!"))
		return
	}

	linha, err := db.Query("select * from usuarios where id = ?", ID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}
	}

	if usuario.ID == 0 {
		erro := &Erro{
			Erro: "Usuário não encontrado!",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(erro)
	} else {
		w.WriteHeader(http.StatusOK)
		if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
			w.Write([]byte("Erro ao coverter o usuário para JSON!"))
			return
		}
	}

}

// AtualizarUsuario Atualizar os dados de um usuário específico
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	ID, err := strconv.ParseUint(param["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	corpoRequest, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario

	if err := json.Unmarshal(corpoRequest, &usuario); err != nil {
		w.Write([]byte("Erro ao converter usuários para json!"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar ao banco"))
		return
	}

	defer db.Close()

	statement, err := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if err != nil {
		w.Write([]byte("Erro ao criar um statement"))
		return
	}

	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuário!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeletarUsuario Deleta um usuário específico
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	ID, err := strconv.ParseUint(param["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar ao banco"))
		return
	}

	defer db.Close()

	stat, err := db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		w.Write([]byte("Erro ao criar um statement"))
		return
	}

	defer stat.Close()

	if _, err := stat.Exec(ID); err != nil {
		w.Write([]byte("Erro ao deletar usuário!"))
	}

	w.WriteHeader(http.StatusNoContent)
}
