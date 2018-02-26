import "github.com/jinzhu/gorm"

// Has one way
// User has one CreditCard, CreditCardID is the foreign key

type User struct {
	gorm.Model
	CreditCard   CreditCard
	CreditCardID uint
}

type CreditCard struct {
	gorm.Model
	Number string
}

// Working with has one
var card CreditCard
db.Model(&user).Related(&card, "CreditCard")
db.Model(&user).Related(&card)

