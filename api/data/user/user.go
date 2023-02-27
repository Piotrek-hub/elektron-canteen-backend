package user 

type User struct {
  Email string 
  Name string
  Surname string
  Password string 
}

type NewUser struct {
  Email string `json:"email"`
  Name string`json:"name"`
  Surname string`json:"surname"`
  Password string `json:"password"`
}

type UpdateUser struct {}
