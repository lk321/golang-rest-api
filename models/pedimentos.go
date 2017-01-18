package models

import "time"

type Pedimentos struct {
    PedimentoID         int         `json:"pedimentoid" gorm:"column:pedimentoid;primary_key"`
    RemesaID            int         `json:"remesaid" gorm:"column:remesaid;"`
    ClienteID           int         `json:"clienteid" gorm:"column:clienteid;not null"`
    Cliente             string      `json:"cliente" gorm:"column:cliente;size:64"`
    CruceID             int         `json:"cruceid" gorm:"column:cruceid"`
    Aduana              int         `json:"aduana" gorm:"column:aduana"`
    Patente             int         `json:"patente" gorm:"column:patente"`
    Pedimento           string      `json:"pedimento" gorm:"column:pedimento;size:64"`
    NumPedimento        string      `json:"num_pedimento" gorm:"column:num_pedimento;size:64"`
    Remesa              int         `json:"remesa" gorm:"column:remesa"`
    Tipo                string      `json:"tipo" gorm:"column:tipo;size:255"`
    Partidas            int         `json:"partidas" gorm:"column:partidas"`
    Placas              string      `json:"placas" gorm:"column:placas;size:64"`
    Factura             string      `json:"factura" gorm:"column:factura;size:64"`
    Clase               string      `json:"clase" gorm:"column:clase;size:64"`
    Clave               string      `json:"clave" gorm:"column:clave;size:64"`
    Status              string      `json:"status" gorm:"column:status;size:64"`
    Costos              bool        `json:"costos" gorm:"column:costos;type:tinyint(1);"`
    Checklist           bool        `json:"checklist" gorm:"column:checklist;type:tinyint(1);"`
    Cobros              bool        `json:"cobros" gorm:"column:cobros;type:tinyint(1);"`
    PedimentoCol        string      `json:"pedimentocol" gorm:"column:pedimentocol;size:64"`
    Ubicacion           string      `json:"ubicacion" gorm:"column:ubicacion;size:124"`
    Observaciones       string      `json:"observaciones" gorm:"column:observaciones;size:500"`
    Autorizado          bool        `json:"autorizado" gorm:"column:autorizado;type:tinyint(1);"`
    CodigoAutorizacion  string      `json:"codigo_autorizacion" gorm:"column:codigo_autorizacion;size:128"`
    Archivado           bool        `json:"archivado" gorm:"column:archivado;type:tinyint(1);"`
    SoiaRemesa          string      `json:"soiaremesa" gorm:"column:soiaremesa;size:45"`
    FechaApertura       time.Time   `json:"fecha_apertura" gorm:"column:fecha_apertura;type:date"`
    FechaPedimento      time.Time   `json:"fecha_pedimento" gorm:"column:fecha_pedimento;type:date"`
    FechaCierre         time.Time   `json:"fecha_cierre" gorm:"column:fecha_cierre;type:date"`
    FechaDocumentacion  time.Time   `json:"fecha_documentacion" gorm:"column:fecha_documentacion;type:timestamp;not null"`
    fecha_checklist     time.Time   `json:"fecha_checklist" gorm:"column:fecha_checklist;type:datetime"`
    fecha_autorizacion  time.Time   `json:"fecha_autorizacion" gorm:"column:fecha_autorizacion;type:datetime"`
}