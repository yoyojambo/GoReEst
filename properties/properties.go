package properties

import (
	"evaluator/users"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

var once sync.Once
var rates [4]decimal.Decimal 

// Returns the last 4 years of inflation, from 2020 to 2023
func Rates() [4]decimal.Decimal {
	once.Do(func() {
		rates[0], _ = decimal.NewFromString("1.2")
		rates[1], _ = decimal.NewFromString("4.7")
		rates[2], _ = decimal.NewFromString("8.0")
		rates[3], _ = decimal.NewFromString("4.1")
	})
	
	return rates
}

type Property struct  {
	Latitude float64
	Longitude float64
	State string
	City string
	Address string
	Sqft int32
	Type string
	Value decimal.Decimal
	Updated int64 // time.DateTime is the constant that defines the format
	Status string
}

type Evaluation struct {
	Property Property
	Val_total decimal.Decimal
	Val_sqft decimal.Decimal
	Confidence float32
	Adjs_evaluated int 
}

func mockPropertiesEvaluated() []Evaluation {
	props := []Property{}
	props = append(props, Property{34.081376, -118.327534, "California", "Los Angeles",
		"585 N Rossmore Ave Apt 506", 17955, "condo", decimal.NewFromInt(649000),
		time.Date(2024,time.September,12, 0, 0, 0, 0, &time.Location{}).Unix(), "for sale"})
	props = append(props, Property{34.080074, -118.318977, "California", "Los Angeles",
		"610 N Gramercy Pl", 5090, "single_family", decimal.NewFromInt(120500),
		time.Date(2024,time.September,12, 0, 0, 0, 0, &time.Location{}).Unix(), "for sale"})
	props = append(props, Property{34.08206, -118.312273, "California", "Los Angeles",
		"537 N Irving Blvd", 2334, "single_family", decimal.NewFromInt(3100000),
		time.Date(2024,time.September,12, 0, 0, 0, 0, &time.Location{}).Unix(), "for sale"})


	evaluated := make([]Evaluation, len(props))
	evaluated[0].Property = props[0]
	evaluated[0].Confidence = 0.9
	evaluated[0].Val_sqft = props[0].Value.Div(decimal.NewFromInt32(props[0].Sqft))
	evaluated[0].Val_total = props[0].Value
	evaluated[0].Adjs_evaluated = 123
	
	evaluated[1].Property = props[1]
	evaluated[1].Confidence = 0.5
	evaluated[1].Val_sqft = props[1].Value.Div(decimal.NewFromInt32(props[1].Sqft))
	evaluated[1].Val_total = props[1].Value
	evaluated[1].Adjs_evaluated = 123
	
	evaluated[2].Property = props[2]
	evaluated[2].Confidence = 0.7
	evaluated[2].Val_sqft = props[2].Value.Div(decimal.NewFromInt32(props[2].Sqft))
	evaluated[2].Val_total = props[2].Value
	evaluated[2].Adjs_evaluated = 123
	

	return evaluated
}

func yourPropertiesFunc(ctx *gin.Context) {
	user := ctx.MustGet("AuthUser")
	fmt.Println("Properties for user:", user)

	props := mockPropertiesEvaluated()
	props_iter := make([]gin.H, len(props))
	for i, p := range props {
		lat := strconv.FormatFloat(p.Property.Latitude, 'g', -1, 64)
		lon := strconv.FormatFloat(p.Property.Longitude, 'g', -1, 64)
		loc := lat + ", " + lon
		props_iter[i] = gin.H{"Location": loc, "Value": p.Val_total}
	}

	ctx.HTML(http.StatusOK, "properties.html", props_iter)
}

func postNewProperty(ctx *gin.Context) {
	user := ctx.MustGet("AuthUser")

	lat := ctx.PostForm("Latitude")
	lon := ctx.PostForm("Longitude")
	addr := ctx.PostForm("Adress")
	sqft := ctx.PostForm("Area")
}

// EvaluateProperty takes a `Property` object and its adjacents, and
// creates an `Evaluation` that takes into account historic data.
func EvaluateProperty(prop Property, adjacents []Property) Evaluation {
	var res Evaluation
	res.Property = prop
	res.Adjs_evaluated = len(adjacents)
	var sum_value decimal.Decimal
	var sum_sqft int64
	var sum_confidence int // 100% confidence would be len(adjacents) * 5

	for _, p := range adjacents {
		cur_value := p.Value
		sum_sqft += int64(p.Sqft)
		
		date_sold := time.Unix(p.Updated, 0)

		// Apply inflation to get current value
		year_diff := 2024 - date_sold.Year()
		for i := 4-year_diff; i < 4; i++ {
			cur_value = cur_value.Mul(rates[i])
		}

		sum_value = sum_value.Add(cur_value)
		sum_confidence += 5 - year_diff
	}

	res.Val_sqft = sum_value.Div(decimal.NewFromInt(sum_sqft))

	// Currently total value is the estimated value of the land per sqft * the sqft of this property.
	res.Val_total = res.Val_sqft.Mul(decimal.NewFromInt32(res.Property.Sqft))
	res.Confidence = float32(sum_confidence) / float32(res.Adjs_evaluated * 5) // Value from 0 to 1 to reflect confidence

	return res
}


func LoadPropertiesHandlers(r *gin.Engine) {
	authed := r.Group("/", users.AuthMiddleware())
	authed.GET("/my_properties", yourPropertiesFunc)
}
