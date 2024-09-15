package structs

import (
    "fmt"
    "time"
    "encoding/json" 
)

type Propiedades struct {
	Data Data `json:"data"`
}


func (p Propiedades) ToString() string {
	return fmt.Sprintf("Propiedades: %+v", p)
}

func UnmarshalPropiedades(data []byte) (Propiedades, error) {
	var r Propiedades
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Propiedades) Marshal() ([]byte, error) {
	return json.Marshal(r)
}


type Data struct {
	HomeSearch HomeSearch `json:"home_search"`
}

type HomeSearch struct {
	Typename string   `json:"__typename"`
	Count    int64    `json:"count"`
	Total    int64    `json:"total"`
	Results  []Result `json:"results"`
}

type Result struct {
	Typename           ResultTypename      `json:"__typename"`
	PropertyID         string              `json:"property_id"`
	ListingID          string              `json:"listing_id"`
	PlanID             interface{}         `json:"plan_id"`
	Status             Status              `json:"status"`
	PhotoCount         int64               `json:"photo_count"`
	Branding           []BrandingElement   `json:"branding"`
	Location           Location            `json:"location"`
	OpenHouses         []OpenHouse         `json:"open_houses"`
	Description        Description         `json:"description"`
	VirtualTours       []PrimaryPhoto      `json:"virtual_tours"`
	Matterport         bool                `json:"matterport"`
	Advertisers        []Advertiser        `json:"advertisers"`
	Flags              Flags               `json:"flags"`
	Source             Source              `json:"source"`
	PetPolicy          *PetPolicy          `json:"pet_policy"`
	Community          interface{}         `json:"community"`
	PrimaryPhoto       PrimaryPhoto        `json:"primary_photo"`
	Href               string              `json:"href"`
	ListPrice          int64               `json:"list_price"`
	ListPriceMin       interface{}         `json:"list_price_min"`
	ListPriceMax       interface{}         `json:"list_price_max"`
	PriceReducedAmount *int64              `json:"price_reduced_amount"`
	Estimate           *Estimate           `json:"estimate"`
	LeadAttributes     LeadAttributesClass `json:"lead_attributes"`
	LastSoldDate       *string             `json:"last_sold_date"`
	ListDate           time.Time           `json:"list_date"`
	Products           *Products           `json:"products"`
	LastSoldPrice      *int64              `json:"last_sold_price"`
}

type Advertiser struct {
	Typename      AdvertiserTypename `json:"__typename"`
	FulfillmentID string             `json:"fulfillment_id"`
	Name          string             `json:"name"`
	Email         *string            `json:"email"`
	Href          *string            `json:"href"`
	Slogan        *string            `json:"slogan"`
	Type          AdvertiserType     `json:"type"`
}

type BrandingElement struct {
	Typename BrandingTypename `json:"__typename"`
	Photo    *string          `json:"photo"`
	Name     string           `json:"name"`
	Phone    interface{}      `json:"phone"`
	Link     interface{}      `json:"link"`
}

type Description struct {
	Typename         DescriptionTypename `json:"__typename"`
	SubType          *SubType            `json:"sub_type"`
	Type             DescriptionType     `json:"type"`
	Beds             *int64              `json:"beds"`
	Baths            *int64              `json:"baths"`
	LotSqft          *int64              `json:"lot_sqft"`
	Sqft             *int64              `json:"sqft"`
	BedsMax          interface{}         `json:"beds_max"`
	BedsMin          interface{}         `json:"beds_min"`
	SqftMax          interface{}         `json:"sqft_max"`
	SqftMin          interface{}         `json:"sqft_min"`
	BathsFull        *int64              `json:"baths_full"`
	BathsHalf        *int64              `json:"baths_half"`
	BathsMin         interface{}         `json:"baths_min"`
	BathsMax         interface{}         `json:"baths_max"`
	BathsFullCalc    *int64              `json:"baths_full_calc"`
	BathsPartialCalc *int64              `json:"baths_partial_calc"`
}

type Estimate struct {
	Typename EstimateTypename `json:"__typename"`
	Estimate int64            `json:"estimate"`
}

type Flags struct {
	Typename          FlagsTypename `json:"__typename"`
	IsPriceReduced    *bool         `json:"is_price_reduced"`
	IsNewConstruction *bool         `json:"is_new_construction"`
	IsForeclosure     interface{}   `json:"is_foreclosure"`
	IsPlan            interface{}   `json:"is_plan"`
	IsNewListing      bool          `json:"is_new_listing"`
	IsComingSoon      *bool         `json:"is_coming_soon"`
	IsContingent      *bool         `json:"is_contingent"`
	IsPending         *bool         `json:"is_pending"`
}

type LeadAttributesClass struct {
	Typename             LeadAttributesTypename `json:"__typename"`
	LeadType             LeadType               `json:"lead_type"`
	ShowContactAnAgent   bool                   `json:"show_contact_an_agent"`
	OpcityLeadAttributes OpcityLeadAttributes   `json:"opcity_lead_attributes"`
}

type OpcityLeadAttributes struct {
	Typename             OpcityLeadAttributesTypename `json:"__typename"`
	FlipTheMarketEnabled bool                         `json:"flip_the_market_enabled"`
}

type Location struct {
	Typename      LocationTypename `json:"__typename"`
	Address       Address          `json:"address"`
	StreetViewURL string           `json:"street_view_url"`
	County        County           `json:"county"`
}

type Address struct {
	Typename     AddressTypename `json:"__typename"`
	City         City            `json:"city"`
	Line         *string         `json:"line"`
	StreetName   string          `json:"street_name"`
	StreetNumber string          `json:"street_number"`
	StreetSuffix StreetSuffix    `json:"street_suffix"`
	Country      Country         `json:"country"`
	PostalCode   string          `json:"postal_code"`
	StateCode    StateCode       `json:"state_code"`
	State        State           `json:"state"`
	Coordinate   CoordinateClass `json:"coordinate"`
}

type CoordinateClass struct {
	Typename CoordinateTypename `json:"__typename"`
	Lat      float64            `json:"lat"`
	Lon      float64            `json:"lon"`
	Accuracy interface{}        `json:"accuracy"`
}

type County struct {
	Typename CountyTypename `json:"__typename"`
	FIPSCode string         `json:"fips_code"`
}

type OpenHouse struct {
	Typename    OpenHouseTypename `json:"__typename"`
	StartDate   time.Time         `json:"start_date"`
	EndDate     time.Time         `json:"end_date"`
	Description *string           `json:"description"`
	TimeZone    TimeZone          `json:"time_zone"`
}

type PetPolicy struct {
	Typename string `json:"__typename"`
	Cats     bool   `json:"cats"`
	Dogs     bool   `json:"dogs"`
}

type PrimaryPhoto struct {
	Typename PrimaryPhotoTypename `json:"__typename"`
	Href     string               `json:"href"`
}

type Products struct {
	Typename  ProductsTypename `json:"__typename"`
	BrandName BrandName        `json:"brand_name"`
	Products  []LeadType       `json:"products"`
}

type Source struct {
	Typename    SourceTypename `json:"__typename"`
	Agents      []Agent        `json:"agents"`
	ID          ID             `json:"id"`
	Type        SourceType     `json:"type"`
	SpecID      interface{}    `json:"spec_id"`
	PlanID      interface{}    `json:"plan_id"`
	ListingHref interface{}    `json:"listing_href"`
	ListingID   string         `json:"listing_id"`
}

type Agent struct {
	Typename   AgentTypename `json:"__typename"`
	ID         ID            `json:"id"`
	AgentID    string        `json:"agent_id"`
	AgentName  string        `json:"agent_name"`
	OfficeID   string        `json:"office_id"`
	OfficeName *string       `json:"office_name"`
}

type AdvertiserType string

const (
	Seller AdvertiserType = "seller"
)

type AdvertiserTypename string

const (
	HomeAdvertiser AdvertiserTypename = "HomeAdvertiser"
)

type BrandingTypename string

const (
	Branding BrandingTypename = "Branding"
)

type SubType string

const (
	Condo     SubType = "condo"
	Townhouse SubType = "townhouse"
)

type DescriptionType string

const (
	Condos        DescriptionType = "condos"
	DuplexTriplex DescriptionType = "duplex_triplex"
	Land          DescriptionType = "land"
	MultiFamily   DescriptionType = "multi_family"
	SingleFamily  DescriptionType = "single_family"
	Townhomes     DescriptionType = "townhomes"
)

type DescriptionTypename string

const (
	SearchHomeDescription DescriptionTypename = "SearchHomeDescription"
)

type EstimateTypename string

const (
	LatestEstimate EstimateTypename = "LatestEstimate"
)

type FlagsTypename string

const (
	HomeFlags FlagsTypename = "HomeFlags"
)

type LeadType string

const (
	CoBroke                 LeadType = "co_broke"
	CoreAgent               LeadType = "core.agent"
	CoreBroker              LeadType = "core.broker"
	ListingAgentProduct     LeadType = "listing_agent_product"
	ListingOwnerBrandBroker LeadType = "listing_owner_brand.broker"
)

type OpcityLeadAttributesTypename string

const (
	OpCityLeadAttributes OpcityLeadAttributesTypename = "OpCityLeadAttributes"
)

type LeadAttributesTypename string

const (
	LeadAttributes LeadAttributesTypename = "LeadAttributes"
)

type City string

const (
	LosAngeles City = "Los Angeles"
)

type CoordinateTypename string

const (
	Coordinate CoordinateTypename = "Coordinate"
)

type Country string

const (
	Usa Country = "USA"
)

type State string

const (
	California State = "California"
)

type StateCode string

const (
	CA StateCode = "CA"
)

type StreetSuffix string

const (
	Ave  StreetSuffix = "Ave"
	Blvd StreetSuffix = "Blvd"
	DR   StreetSuffix = "Dr"
	Pl   StreetSuffix = "Pl"
	RD   StreetSuffix = "Rd"
	St   StreetSuffix = "St"
)

type AddressTypename string

const (
	SearchHomeAddress AddressTypename = "SearchHomeAddress"
)

type CountyTypename string

const (
	HomeCounty CountyTypename = "HomeCounty"
)

type LocationTypename string

const (
	SearchHomeLocation LocationTypename = "SearchHomeLocation"
)

type TimeZone string

const (
	Pst TimeZone = "PST"
)

type OpenHouseTypename string

const (
	HomeOpenHouse OpenHouseTypename = "HomeOpenHouse"
)

type PrimaryPhotoTypename string

const (
	HomePhoto   PrimaryPhotoTypename = "HomePhoto"
	VirtualTour PrimaryPhotoTypename = "VirtualTour"
)

type BrandName string

const (
	AdvantageBrand BrandName = "advantage_brand"
	BasicOptIn     BrandName = "basic_opt_in"
	Essentials     BrandName = "essentials"
)

type ProductsTypename string

const (
	ProductSummary ProductsTypename = "ProductSummary"
)

type ID string

const (
	Mrca ID = "MRCA"
	Nsny ID = "NSNY"
	Sdca ID = "SDCA"
	Toca ID = "TOCA"
	Weca ID = "WECA"
)

type AgentTypename string

const (
	MlsAgent AgentTypename = "MlsAgent"
)

type SourceType string

const (
	Mls SourceType = "mls"
)

type SourceTypename string

const (
	MlsSource SourceTypename = "MlsSource"
)

type Status string

const (
	ForSale Status = "for_sale"
)

type ResultTypename string

const (
	SearchHome ResultTypename = "SearchHome"
)
