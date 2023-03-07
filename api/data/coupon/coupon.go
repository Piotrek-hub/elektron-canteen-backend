package coupon

type Coupon struct {
	Value float32 `json:"value" bson:"value"`
	Code  string  `json:"code" bson:"code"`
}
