import "github.com/jinzhu/gorm"

// User has and belongs to many languages, use `user_languages` as join table

type User struct {
	gorm.Model
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}

// It will create a many2many relationship for those two structs,
// and their relations will be saved into join table PersonAccount
// with foreign keys customize_person_id_person AND customize_account_id_account

type CustomizePerson struct {
	IdPerson string             `gorm:"primary_key:true"`
	Accounts []CustomizeAccount `gorm:"many2many:PersonAccount;association_foreignkey:idAccount;foreignkey:idPerson"`
}

type CustomizeAccount struct {
	IdAccount string `gorm:"primary_key:true"`
	Name      string
}
