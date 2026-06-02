package models

type Aluno struct {
	Name string `json:"name"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

var Alunos []Aluno
