package dish 

type Dish struct {
  Name string`json:"name"`
  Additions []string `json:"additions"`
  Salads []string `json:"salads"`
}


type Addition struct {
  Name string`json:"name"`
}

