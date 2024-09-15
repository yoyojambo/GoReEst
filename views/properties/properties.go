package properties

import (
	"sync"
	"time"

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
	Updated string // time.DateTime is the constant that defines the format
	Status string
}

type Evaluation struct {
	Property Property
	Val_total decimal.Decimal
	Val_sqft decimal.Decimal
	Confidence float32
	Adjs_evaluated int 
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
		
		date_sold, err := time.Parse(time.DateTime, p.Updated)
		if err != nil {
			panic("Could not evaluate property! Invalid date on adjacent: " + err.Error())
		}

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
