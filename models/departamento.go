package models

type Departamento struct {
	ID					int 	`json:"id" gorm:"column:dep_id;primary_key;AUTO_INCREMENT"`
	Nombre				string	`json:"nombre" gorm:"column:dep_nombre;size:85;not null"`
	Abreviado			string	`json:"abreviado" gorm:"column:dep_abreviado;size:15;not null"`
	Email				string	`json:"email" gorm:"column:dep_email;size:128;not null"`
	Activo				bool	`json:"activo" gorm:"type:tinyint(1);column:dep_activo"`
	ReparticionCliente 	bool	`json:"reparticion" gorm:"type:tinyint(1);column:dep_reparticion_cliente"`
}