package request

type CreateShipmentRequest struct {
	Profile   string     `json:"profile"`
	Shipments []shipment `json:"shipments"`
}

type shipment struct {
	Product          string    `json:"product"`
	BillingNumber    string    `json:"billingNumber"`
	RefNumber        string    `json:"refNo"`
	CostCenter       string    `json:"costCenter"`
	CreationSoftware string    `json:"creationSoftware"`
	ShipDate         string    `json:"shipDate"`
	Shipper          shipper   `json:"shipper"`
	Consignee        consignee `json:"consignee"`
	Details          details   `json:"details"`
	Services         services  `json:"services"`
	Customs          customs   `json:"customs"`
}

type shipper struct {
	Name1         string `json:"name1"`
	Name2         string `json:"name2"`
	Name3         string `json:"name3"`
	AddressStreet string `json:"addressStreet"`
	AddressHouse  string `json:"addressHouse"`
	PostalCode    string `json:"postalCode"`
	City          string `json:"city"`
	Country       string `json:"country"`
	ContactName   string `json:"contactName"`
	Email         string `json:"email"`
}

type consignee struct {
	Name1                         string `json:"name1"`
	Name2                         string `json:"name2"`
	Name3                         string `json:"name3"`
	DispatchingInformation        string `json:"dispatchingInformation"`
	AddressStreet                 string `json:"addressStreet"`
	AddressHouse                  string `json:"addressHouse"`
	AdditionalAddressInformation1 string `json:"additionalAddressInformation1"`
	AdditionalAddressInformation2 string `json:"additionalAddressInformation2"`
	PostalCode                    string `json:"postalCode"`
	City                          string `json:"city"`
	State                         string `json:"state"`
	Country                       string `json:"country"`
	ContactName                   string `json:"contactName"`
	PhoneNumber                   string `json:"phone"`
	Email                         string `json:"email"`
}

type details struct {
	Dimension dimension `json:"dim"`
	Weight    weight    `json:"weight"`
}

type dimension struct {
	OUM    string `json:"uom"`
	Height int    `json:"height"`
	Length int    `json:"length"`
	Width  int    `json:"width"`
}

type weight struct {
	OUM   string `json:"uom"`
	Value int    `json:"value"`
}

type services struct {
	PrefNeighbour               string              `json:"preferredNeighbour"`
	PrefLocation                string              `json:"preferredLocation"`
	VisualAgeCheck              string              `json:"visualCheckOfAge"`
	NamedPersonOnly             bool                `json:"namedPersonOnly"`
	IdentCheck                  identCheck          `json:"identCheck"`
	SignedForByRecipient        bool                `json:"signedForByRecipient"`
	Endorsement                 string              `json:"endorsement"`
	PreferredDay                string              `json:"preferredDay"`
	NoNeighbourDelivery         bool                `json:"noNeighbourDelivery"`
	AdditionalInsurance         additionalInsurance `json:"additionalInsurance"`
	BulkyGoods                  bool                `json:"bulkyGoods"`
	CashOnDelivery              cashOnDelivery      `json:"cashOnDelivery"`
	IndividualSenderRequirement string              `json:"individualSenderRequirement"`
	Premium                     bool                `json:"premium"`
	ClosestDropPoint            bool                `json:"closestDropPoint"`
	ParcelOutletRouting         string              `json:"parcelOutletRouting"`
	GoGreenPlus                 bool                `json:"goGreenPlus"`
	DhlRetoure                  dhlRetoure          `json:"dhlRetoure"`
	PostalDeliveryDutyPaid      bool                `json:"postalDeliveryDutyPaid"`
}

type identCheck struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
	MinimumAge  string `json:"minimumAge"`
}

type additionalInsurance struct {
	Currency string `json:"currency"`
	Value    int    `json:"value"`
}

type cashOnDelivery struct {
	Amount           money       `json:"amount"`
	BankAccount      bankAccount `json:"bankAccount"`
	AccountReference string      `json:"accountReference"`
	TransferNote1    string      `json:"transferNote1"`
	TransferNote2    string      `json:"transferNote2"`
}

type bankAccount struct {
	AccountHolder string `json:"accountHolder"`
	BankName      string `json:"bankName"`
	IBAN          string `json:"iban"`
	BIC           string `json:"bic"`
}

type dhlRetoure struct {
	BillingNumber string    `json:"billingNumber"`
	RefNumber     string    `json:"refNo"`
	ReturnAddress consignee `json:"returnAddress"`
	GoGreenPlus   bool      `json:"goGreenPlus"`
}

type customs struct {
	InvoiceNo                       string        `json:"invoiceNo"`
	ExportType                      string        `json:"exportType"`
	ExportDescription               string        `json:"exportDescription"`
	ShippingConditions              string        `json:"shippingConditions"`
	PermitNo                        string        `json:"permitNo"`
	AttestationNo                   string        `json:"attestationNo"`
	HasElectronicExportNotification bool          `json:"hasElectronicExportNotification"`
	MRN                             string        `json:"MRN"`
	PostalCharges                   money         `json:"postalCharges"`
	OfficeOfOrigin                  string        `json:"officeOfOrigin"`
	ShipperCustomsRef               string        `json:"shipperCustomsRef"`
	ConsigneeCustomsRef             string        `json:"consigneeCustomsRef"`
	Items                           []customsItem `json:"items"`
}

type customsItem struct {
	ItemDescription  string `json:"itemDescription"`
	HsCode           string `json:"hscode"`
	CountryOfOrigin  string `json:"countryOfOrigin"`
	PackagedQuantity int    `json:"packagedQuantity"`
}

type money struct {
	Currency string `json:"currency"`
	Value    int    `json:"value"`
}
