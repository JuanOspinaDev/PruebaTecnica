package models

// Estructura para un cliente
type Cliente struct {
    ID        string `json:"id"`        // Identificador
    Nombre    string `json:"nombre"`    // Nombre completo del cliente
    Email     string `json:"email"`     // Correo del cliente
    Telefono  string `json:"telefono"`  // Teléfono del cliente

}