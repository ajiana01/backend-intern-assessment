Question 3.1

1.
type User struct {
    ID           int `json:"id"`
    FullName     string `json:"full_name"`
    Email        string `json:"email"`
    PasswordHash string `json:"-"`
}

2.
json:"-" Tag on GO encoding/json package will tell to ignore that field during JSON serialization. So, the PasswordHash will not include in JSON output when json.Marshal() is called.

===

Question 3.2

For User Wallets & Financial Balances, i would choose MySQL because this module require consistency, accurate data, and reliable transaction handling. MySQL also support ACID transaction, it mean if in the middle process transaction is fail, it can be roll back for prevent partial or inconsistent data.

For Product Catalog & Dynamic Spesifications, i preffer using MongoDB because product attributes can be different each product and types. For example, laptops may have CPU and RAM fields, while T-Shirts may have size, color, fabric fields. Because MongoDB is document base schema, it will more easier to store and manage dynamic product specs without require same structure every product.