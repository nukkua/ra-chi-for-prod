package models

import (

)

type User struct {
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"password"`
}

type UserWithoutHash struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//its a gorm hook BeforeSave and BeforeCreate

// func (u *User) BeforeSave(tx *gorm.DB) (err error) {
//     if len(u.Password) > 0 {
//         // Convertimos la contraseña en texto plano a un hash
//         bytePassword := []byte(u.Password)
//         passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
//         if err != nil {
//             return err
//         }
//         u.PasswordHash = passwordHash
//         // No almacenamos la contraseña en texto plano
//         u.Password = ""
//     }
//     return nil
// }
