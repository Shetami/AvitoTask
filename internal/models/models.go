package models

type User struct {
	Id      int `json:"-" db:"id"`
	Balance int `json:"balance"`
}

type Reserve struct {
	Id          int    `json:"id"`
	Cash        int    `json:"cash"`
	Description string `json:"description"`
}

type Buffer struct {
	Balance     int    `json:"balance" db:"balance"`
	UserId      int    `json:"user_id" db:"user_id"`
	Description string `json:"description" db:"description"`
}

type ConfirmationStruct struct {
	Id    int  `json:"id"`
	Value bool `json:"value"`
}

type Avito struct {
	Id      int `json:"id"`
	Balance int `json:"balance"`
}

type Transaction struct {
	Id   int `json:"id"`
	Idtr int `json:"idtr"`
	Cash int `json:"cash"`
}

type AddMoneyStruct struct {
	Id    int `json:"id"`
	Money int `json:"money"`
}
type Report struct {
	UserId      int    `csv:"userId"`
	Description string `csv:"description"`
	Price       int    `csv:"price"`
}
