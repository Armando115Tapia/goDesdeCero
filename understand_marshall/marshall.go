package understand_marshall

// Processing struct of hierarchy processing configuration.
type Processing struct {
	BusinessRules string `json:"businessRules,omitempty"`
	Deferred      string `json:"deferred,omitempty"`
	Processors    string `json:"processors,omitempty"`
	SecurityRules string `json:"securityRules,omitempty"`
}

// HierarchyConfig struct of hierarchy configuration.
type HierarchyConfig struct {
	Processing Processing `json:"processing,omitempty"`
	RootID     string     `json:"rootId,omitempty"`
}

type OtpData struct {
	ChargeData ChargeData `json:"chargeData"`
}

type ChargeData struct {
	OtpID                string `json:"secureServiceId"`
	IsSubscriptionCharge bool   `json:"isSubscriptionCharge"`
}

type Customer struct {
	DocumentType   string `json:"documentType"`
	DocumentNumber string `json:"documentNumber"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phoneNumber"`
	Name           string `json:"firstName"`
	LastName       string `json:"lastName"`
	SecondLastName string `json:"secondLastName"`
}

type TransferCustomer struct {
	CustomerName     string `json:"customerName"`
	CustomerLastName string `json:"customerLastName"`
	DocumentType     string `json:"documentType"`
	DocumentNumber   string `json:"documentNumber"`
}

type BiometricData struct {
	SecureID string `json:"secureId"`
}

type BillingAndShippingDetails struct {
	Address          string `json:"address,omitempty"`
	City             string `json:"city,omitempty"`
	Country          string `json:"country,omitempty"`
	Name             string `json:"name,omitempty"`
	Phone            string `json:"phone,omitempty"`
	Region           string `json:"region,omitempty"`
	SecondaryAddress string `json:"secondaryAddress,omitempty"`
	UserEmail        string `json:"userEmail,omitempty"`
	ZipCode          string `json:"zipCode,omitempty"`
}

type OrderDetails struct {
	SiteDomain      string                     `json:"siteDomain"`
	ShippingDetails *BillingAndShippingDetails `json:"shippingDetails,omitempty"`
	BillingDetails  *BillingAndShippingDetails `json:"billingDetails,omitempty"`
}

type Product struct {
	Category string      `json:"category"`
	ID       interface{} `json:"id"`
	Price    interface{} `json:"price"`
	Quantity interface{} `json:"quantity"`
	SKU      interface{} `json:"sku,omitempty"`
	Tags     []string    `json:"tags"`
	Title    string      `json:"title"`
}

type ProductDetails struct {
	Product []Product `json:"product"`
}

type ThreeDSecureData struct {
	CardType                              string  `json:"cardType"`
	Currency                              string  `json:"currency"`
	LastName                              string  `json:"lastName"`
	Name                                  string  `json:"name"`
	Url                                   string  `json:"url"`
	Jwt                                   string  `json:"jwt"`
	ReferenceId                           string  `json:"referenceId"`
	Email                                 string  `json:"email"`
	MerchantName                          string  `json:"merchantName"`
	Country                               string  `json:"country"`
	Brand                                 string  `json:"brand"`
	LoginID                               string  `json:"loginId"`
	RequestorId                           string  `json:"requestorId"`
	RequestorName                         string  `json:"requestorname"`
	BillToPhoneNumber                     *string `json:"billTo_phoneNumber"`
	PayerAuthEnrollServiceMessageCategory *string `json:"payerAuthEnrollService_messageCategory"`
	InvoiceHeaderPurchaserCode            *string `json:"invoiceHeader_purchaserCode"`
	PayerAuthEnrollServiceDeviceChannel   *string `json:"payerAuthEnrollService_deviceChannel"`
	PayerAuthEnrollServiceTransactionMode *string `json:"payerAuthEnrollService_transactionMode"`
}

type PayloadProcessorRequest struct {
	TransactionType          string                 `json:"transactionType"`
	Bank                     string                 `json:"bank"`
	Processor                string                 `json:"processor"`
	Brand                    string                 `json:"brand"`
	MaskedCardNumber         string                 `json:"maskedCardNumber"`
	TransactionCardID        string                 `json:"transactionCardId,omitempty"`
	IsDeferred               string                 `json:"isDeferred"`
	UserIP                   string                 `json:"ip"`
	Country                  string                 `json:"country"`
	CityCode                 string                 `json:"cityCode"`
	DepartmentCode           string                 `json:"departmentCode"`
	IsCreditCard             string                 `json:"isCreditCard"`
	TotalAmount              string                 `json:"totalAmount"`
	Bin                      string                 `json:"bin"`
	LastFourDigits           string                 `json:"lastFourDigits"`
	Currency                 string                 `json:"currency"`
	OtpData                  OtpData                `json:"otpData"`
	Customer                 Customer               `json:"customer"`
	TransferCustomer         TransferCustomer       `json:"transferCustomer"`
	BiometricData            BiometricData          `json:"biometricData"`
	IgnoreWarnings           bool                   `json:"ignoreWarnings"`
	CredentialID             string                 `json:"credentialId"`
	IgnoreSqsSubscription    bool                   `json:"ignoreSqsSubscription"`
	VaultToken               string                 `json:"vaultToken"`
	CardHolderName           string                 `json:"cardHolderName"`
	ChargeType               string                 `json:"chargeType"`
	UserID                   string                 `json:"userId"`
	SessionID                string                 `json:"sessionId"`
	OrderDetails             OrderDetails           `json:"orderDetails"`
	TransactionReference     string                 `json:"transactionReference"`
	ProductDetails           ProductDetails         `json:"productDetails"`
	SellerUserID             string                 `json:"sellerUserId"`
	ThreeDSecureData         ThreeDSecureData       `json:"3Ds"`
	MerchantCountry          string                 `json:"merchantCountry"`
	TransactionCreated       int64                  `json:"transactionCreated"`
	Deferred                 rep.Deferred           `json:"deferred,omitempty"`
	Metadata                 map[string]interface{} `json:"metadata,omitempty"`
	Email                    string                 `json:"email,omitempty"`
	PhoneNumber              string                 `json:"phoneNumber,omitempty"`
	MerchantName             string                 `json:"merchantName"`
	ContactDetails           *ContactDetailsObject  `json:"contactDetails"`
	CybersourceApiValidation bool                   `json:"cybersourceApiValidation,omitempty"`
}

type ProcessorRequestInput struct {
	MerchantCountry string                  `json:"merchantCountry"`
	MerchantID      string                  `json:"merchantId"`
	HierarchyConfig HierarchyConfig         `json:"hierarchyConfig,omitempty"`
	IsTokenCharge   bool                    `json:"isTokenCharge,omitempty"`
	Payload         PayloadProcessorRequest `json:"detail"`
	SandboxEnable   bool                    `json:"sandboxEnable,omitempty"`
	TransactionKind string                  `json:"transactionKind"`
	KushkiInfo      rep.KushkiInfo          `json:"kushkiInfo,omitempty"`
	AuthValidation  string                  `json:"authValidation,omitempty"`
	CallbackUrl     string                  `json:"callbackUrl,omitempty"`
}

func understandMarshall() {
	var bodyRequest ProcessorRequestInput
}
