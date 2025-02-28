package Crud 


import(
	"fmt"
	//"gorm.io/driver/postgres"
    "gorm.io/gorm"
)


type Products4 struct {
    gorm.Model
    ID      uint   `gorm:"primaryKey"`
    Code    string `gorm:"unique"`
    Price   uint
}




func (p *Products4) AfterCreate(tx *gorm.DB) (err error){
    if (*p).Price == 0 {
       (*p).Price=100
	   tx.Save(p)
    }
    return
}





func ReadUsers(tx *gorm.DB) {
    var users []Products4
    result := tx.Find(&users)
    if result.Error != nil {
        fmt.Println("Error reading users:", result.Error)
    } else {
        fmt.Println("Users in database:")
        for _, user := range users {
            fmt.Printf("ID: %d | Code: %s | Price: %d\n", user.ID, user.Code,user.Price)
        }
    }
	
}

func CreateUser(tx* gorm.DB,Code string, Price uint)(*Products4,error){
    user := Products4{Code: Code,Price:Price}
    result := tx.Create(&user)
    if result.Error != nil {
        return &user,result.Error
    } 
    fmt.Println("User created successfully:", user)
    
	return &user,nil
}

// always return an error 
func UpdateUser(tx *gorm.DB,ID uint,Code string,Price uint) {
    var user Products4
    result := tx.First(&user,ID) // Find by ID
    if result.Error != nil {
        fmt.Println("Error finding user:", result.Error)
        return
    }
    user.Code = Code 
	user.Price = Price
    tx.Create(&user) // Save updated user
    fmt.Println("User updated successfully:", user)
}

func DeleteUser(tx *gorm.DB,ID uint) {
    result := tx.Delete(&Products4{}, ID)
    if result.Error != nil {
        fmt.Println("Error deleting user:", result.Error)
    } else {
        fmt.Println("User deleted successfully!")
    }
}
 




