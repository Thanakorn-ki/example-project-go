package account

import "time"

type account struct {
	ID                           int       `json:"id" gorm:"PRIMARY_KEY; type:int(10) unsigned"`
	AgencyID                     int       `json:"agency_id" gorm:"type:int(10) unsigned"`
	MainCategoryID               int       `json:"main_category_id" gorm:"type:int(10) unsigned"`
	SubCategoryID                int       `json:"sub_category_id" gorm:"type:int(10) unsigned"`
	AccountEmail                 string    `json:"account_email" gorm:"type:text"`
	AccountOffererName           string    `json:"account_offerer_name" gorm:"type:text"`
	AccountOffererTel            string    `json:"account_offerer_tel" gorm:"type:text"`
	AccountOffererTelCountryCode string    `json:"account_offerer_tel_country_code" gorm:"type:text"`
	DisplayName                  string    `json:"display_name" gorm:"type:varchar(255)"`
	ShopName                     string    `json:"shop_name" gorm:"type:varchar(255)"`
	SiteURL                      string    `json:"site_url" gorm:"type:text"`
	SiteURLFacebook              string    `json:"site_url_facebook" gorm:"type:text"`
	SiteURLInstagram             string    `json:"site_url_instagram" gorm:"type:text"`
	SiteURLTwitter               string    `json:"site_url_twitter" gorm:"type:text"`
	Tel                          string    `json:"tel" gorm:"type:text"`
	LineBotID                    string    `json:"line_bot_id" gorm:"type:varchar(50)"`
	LineInstanceID               string    `json:"line_instance_id" gorm:"type:varchar(50)"`
	IsAutoRenewPlan              bool      `json:"is_auto_renew_plan"`
	Status                       string    `json:"status" gorm:"type:enum('pending', 'error', 'completed', 'rejected', 'retried', 'deleted', 'waiting transfer')"`
	CreatedAt                    time.Time `json:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at"`
	AccountOffererPrefix         string    `json:"account_offerer_prefix"`
	IsForeignCompanies           string    `json:"is_foreign_companies"`
	ProfileImageURL              string    `json:"profile_image_url"`
	IsSubscribe                  bool      `json:"is_subscribe"`
	shop_address_1               string    `json:"ShopAddress1"`
	IsTransferredSsk             bool      `json:"is_transferred_ssk"`
	ShopProvince                 string    `json:"shop_province"`
	Note                         string    `json:"note"`
}
