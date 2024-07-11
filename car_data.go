package irdata

import (
	"encoding/json"
	"time"
)

type Car struct {
	AiEnabled               bool                    `json:"ai_enabled"`
	AllowNumberColors       bool                    `json:"allow_number_colors"`
	AllowNumberFont         bool                    `json:"allow_number_font"`
	AllowSponsor1           bool                    `json:"allow_sponsor1"`
	AllowSponsor2           bool                    `json:"allow_sponsor2"`
	AllowWheelColor         bool                    `json:"allow_wheel_color"`
	AwardExempt             bool                    `json:"award_exempt"`
	CarDirPath              string                  `json:"car_dirpath"`
	CarID                   int                     `json:"car_id"`
	CarName                 string                  `json:"car_name"`
	CarNameAbbreviated      string                  `json:"car_name_abbreviated"`
	CarTypes                []CarType               `json:"car_types"`
	CarWeight               int                     `json:"car_weight"`
	Categories              []string                `json:"categories"`
	Created                 time.Time               `json:"created"`
	FirstSale               time.Time               `json:"first_sale"`
	ForumURL                string                  `json:"forum_url,omitempty"`
	FreeWithSubscription    bool                    `json:"free_with_subscription"`
	HasHeadlights           bool                    `json:"has_headlights"`
	HasMultipleDryTireTypes bool                    `json:"has_multiple_dry_tire_types"`
	HasRainCapableTireTypes bool                    `json:"has_rain_capable_tire_types"`
	Hp                      int                     `json:"hp"`
	IsPsPurchasable         bool                    `json:"is_ps_purchasable"`
	MaxPowerAdjustPct       int                     `json:"max_power_adjust_pct"`
	MaxWeightPenaltyKg      int                     `json:"max_weight_penalty_kg"`
	MinPowerAdjustPct       int                     `json:"min_power_adjust_pct"`
	PackageID               int                     `json:"package_id"`
	Patterns                int                     `json:"patterns"`
	Price                   float64                 `json:"price"`
	PriceDisplay            string                  `json:"price_display,omitempty"`
	RainEnabled             bool                    `json:"rain_enabled"`
	Retired                 bool                    `json:"retired"`
	SearchFilters           string                  `json:"search_filters"`
	Sku                     int                     `json:"sku"`
	CarMake                 string                  `json:"car_make,omitempty"`
	CarModel                string                  `json:"car_model,omitempty"`
	PaintRules              map[string]CarPaintRule `json:"paint_rules,omitempty"`
	SiteURL                 string                  `json:"site_url,omitempty"`
}

type CarType struct {
	CarType string `json:"car_type"`
}

type CarPaintRule struct {
	RestrictCustomPaint bool   `json:"RestrictCustomPaint,omitempty"`
	PaintCarAvailable   bool   `json:"PaintCarAvailable"`
	Color1              string `json:"Color1"`
	Color2              string `json:"Color2"`
	Color3              string `json:"Color3"`
	Sponsor1Available   bool   `json:"Sponsor1Available"`
	Sponsor2Available   bool   `json:"Sponsor2Available"`
	Sponsor1            string `json:"Sponsor1"`
	Sponsor2            string `json:"Sponsor2"`
	RulesExplanation    string `json:"RulesExplanation"`
}

func (r *CarPaintRule) MarshalJSON() ([]byte, error) {
	if r.RestrictCustomPaint {
		return json.Marshal(true)
	}

	type alias CarPaintRule
	return json.Marshal((*alias)(r))
}

func (r *CarPaintRule) UnmarshalJSON(data []byte) error {
	if string(data) == "true" {
		r.RestrictCustomPaint = true
		return nil
	} else if string(data) == "false" {
		r.RestrictCustomPaint = false
		return nil
	}

	type alias CarPaintRule
	return json.Unmarshal(data, (*alias)(r))
}
