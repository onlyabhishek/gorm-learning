package main 

import (
	"fmt"
    //"github.com/google/uuid"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm-learning/Crud"
)

// type Products2 struct {
//     gorm.Model
//     ID      uint   `gorm:"primaryKey"`
//     Code    string `gorm:"unique"`
//     Price   uint
// }


// type Products3 struct {
//     gorm.Model
//     ID      uint   `gorm:"primaryKey"`
//     Code    string `gorm:"unique"`
//     Price   uint
//}



type Products4 struct {
    gorm.Model
    ID      uint   `gorm:"primaryKey"`
    Code    string `gorm:"unique"`
    Price   uint
}


// type Products5 struct
// {
	
// 	ID      uint   `gorm:"primaryKey"`
//     Code    string `gorm:"unique"`
//     Price   uint
	
//}


// func (p *Products4) BeforeCreate(tx *gorm.DB) (err error){
//     if p.Price == 0 {
//         fmt.Println("Code is empty")
//     }
//     return
// }












func main() {
    dsn := "host=localhost user=postgres password=Welcome@1234 dbname=postgres port=5433 sslmode=disable TimeZone=Asia/Kolkata"

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{QueryFields: true})
    if err != nil {
        fmt.Println("Failed to connect to database:", err)
    }

    // db.AutoMigrate(&Products2{})
    // db.AutoMigrate(&Products3{})
    db.AutoMigrate(&Products4{})

    
    //product1 := Products4{Code: "D42", Price: 100}
    // db.Create(&product1)
    // fmt.Printf("Inserted Product: %+v\n", product1.Code)

    // product2 := Products2{Code: "D43", Price: 200}
    // db.Create(&product2)
    // fmt.Printf("Inserted Product: %+v\n", product2.Code)
    // var product Products2
    // result := db.First(&product, "code = ?", "D42")

	// if result.Error != nil {
	// 	fmt.Println("Product not found:", result.Error)
	// } else {
	// 	//fmt.Printf("Found Product: %+v\n", product)

	// 	// Updating Product Price
    //     fmt.Println(product.Price)
	// 	// db.Model(&product).Update("Price", 250)
	// 	// db.Model(&product).Updates(map[string]interface{}{"Price": 300})
	// }
    // var product2 []Products2
    // result1 := db.Find(&product2)
    // fmt.Println(result1.RowsAffected)
    
    // var x Products4
    // //db.First(&x, "code = ?", "D42")  // Fetch the record with Code "D42"
    // x=Products4{Code: "D423"}
    // db.Create(&x)
    
    // db.Where("code=?","D423").First(&x)
    // db.Model(&x).Update("Price", 570) // Update the Price field
    // fmt.Println("Updated the Price of code D42 to 560")

    // var y Products4
    // y=Products4{Code:"ABCf"}
    // db.Create(&y)
    // fmt.Println("Create a code ABC")

    //Crud.ReadUsers(db)
    //Crud.CreateUser(db,"abhishek",0)
    //Crud.ReadUsers(db)
    Crud.UpdateUser(db,64,"Aravind",200)
    Crud.ReadUsers(db)
    
    
    
    

    


    
    

    
    //db.Delete(&x, 1)
    // fmt.Println("Deleted Product with ID=1")

    
    // //product3 := Products2{Code: "ABC", Price: 200}
    // //db.Create(&product3)
    // //fmt.Printf("Inserted Product: %+v\n", product3.Code)
    // product4:=Products2{Price: 300}
    // db.Create(&product4)
}
