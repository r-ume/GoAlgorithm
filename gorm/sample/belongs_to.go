// belongs to association

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string
}

// Profile belongs to User.
// Profile struct must have UserID, which is the foreign key to users table and User struct
type Profile struct {
	gorm.Model
	UserID int
	User   User
	Name   string
}

// Working with Belong to 
db.Model(&user).Related(&profile)
// SELECT * FROM profiles WHERE user_id = 111; // 111 is user's ID

