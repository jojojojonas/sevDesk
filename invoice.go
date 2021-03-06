//**********************************************************
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede UG (haftungsbeschränkt) <info@jj-ideenschmiede.de>
//
// This file is part of sevdesk.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor
//
//**********************************************************

package sevdesk

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

// The data that the function uses
type Invoice struct {
	ContactID     string
	InvoiceDate   string
	Status        string
	InvoiceType   string
	ContactPerson string
	Token         string
}

// To return the data from invoices
type InvoicesReturn struct {
	Objects []InvoiceObjects `json:"objects"`
}

// For returning the data
type NewInvoiceReturn struct {
	Objects InvoiceObjects `json:"objects"`
}

type InvoiceObjects struct {
	ID                                  string     `json:"id"`
	ObjectName                          string     `json:"objectName"`
	AdditionalInformation               string     `json:"additionalInformation"`
	Contact                             ObjectName `json:"contact"`
	Create                              string     `json:"create"`
	Update                              string     `json:"update"`
	SevClient                           ObjectName `json:"sevClient"`
	InvoiceDate                         string     `json:"invoiceDate"`
	Header                              string     `json:"header"`
	HeadText                            string     `json:"headText"`
	FootText                            string     `json:"footText"`
	TimeToPlay                          string     `json:"timeToPlay"`
	DiscountTime                        string     `json:"discountTime"`
	Discount                            string     `json:"discount"`
	AddressName                         string     `json:"addressName"`
	AddressStreet                       string     `json:"addressStreet"`
	AddressZip                          string     `json:"addressZip"`
	AddressCity                         string     `json:"addressCity"`
	PayDate                             string     `json:"payDate"`
	CreateUser                          ObjectName `json:"createUser"`
	DeliveryDate                        string     `json:"deliveryDate"`
	Status                              string     `json:"status"`
	SmallSettlement                     string     `json:"smallSettlement"`
	ContactPerson                       ObjectName `json:"contactPerson"`
	TaxRate                             string     `json:"taxRate"`
	TaxText                             string     `json:"taxText"`
	DunningLevel                        string     `json:"dunningLevel"`
	AddressParentName                   string     `json:"addressParentName"`
	TaxType                             string     `json:"taxType"`
	SendDate                            string     `json:"sendDate"`
	OriginLastInvoice                   string     `json:"originLastInvoice"`
	InvoiceType                         string     `json:"invoiceType"`
	AccountIntervall                    string     `json:"accountIntervall"`
	AccountLastInvoice                  string     `json:"accountLastInvoice"`
	AccountNextInvoice                  string     `json:"accountNextInvoice"`
	ReminderTotal                       string     `json:"reminderTotal"`
	ReminderDebit                       string     `json:"reminderDebit"`
	ReminderDeadline                    string     `json:"reminderDeadline"`
	ReminderCharge                      string     `json:"reminderCharge"`
	AddressParentName2                  string     `json:"addressParentName2"`
	AddressName2                        string     `json:"addressName2"`
	AddressGender                       string     `json:"addressGender"`
	AccountStartDate                    string     `json:"accountStartDate"`
	AccountEndDate                      string     `json:"accountEndDate"`
	Address                             string     `json:"address"`
	Currency                            string     `json:"currency"`
	SumNet                              string     `json:"sumNet"`
	SumTax                              string     `json:"sumTax"`
	SumGross                            string     `json:"sumGross"`
	SumDiscounts                        string     `json:"sumDiscounts"`
	SumNetForeignCurrency               string     `json:"sumNetForeignCurrency"`
	SumTaxForeignCurrency               string     `json:"sumTaxForeignCurrency"`
	SumGrossForeignCurrency             string     `json:"sumGrossForeignCurrency"`
	SumDiscountsForeignCurrency         string     `json:"sumDiscountsForeignCurrency"`
	SumNetAccounting                    string     `json:"sumNetAccounting"`
	SumTaxAccounting                    string     `json:"sumTaxAccounting"`
	SumGrossAccounting                  string     `json:"sumGrossAccounting"`
	PaidAmount                          float64    `json:"paidAmount"`
	CustomerInternalNote                string     `json:"customerInternalNote"`
	ShowNet                             string     `json:"showNet"`
	Enshrined                           string     `json:"enshrined"`
	SendType                            string     `json:"sendType"`
	DeliveryDateUntil                   string     `json:"deliveryDateUntil"`
	SendPaymentReceivedNotificationDate string     `json:"sendPaymentReceivedNotificationDate"`
	SumDiscountNet                      string     `json:"sumDiscountNet"`
	SumDiscountGross                    string     `json:"sumDiscountGross"`
	SumDiscountNetForeignCurrency       string     `json:"sumDiscountNetForeignCurrency"`
	SumDiscountGrossForeignCurrency     string     `json:"sumDiscountGrossForeignCurrency"`
}

// Uses for invoices and positions
type ObjectName struct {
	Id         string `json:"id"`
	ObjectName string `json:"objectName"`
}

// Check invoices
func Invoices(token string) (InvoicesReturn, error) {

	// Define client
	client := &http.Client{}

	// New http request
	request, err := http.NewRequest("GET", "https://my.sevdesk.de/api/v1/Invoice", nil)
	if err != nil {
		return InvoicesReturn{}, err
	}

	// Set headers
	request.Header.Set("Authorization", token)

	// Response to sevDesk
	response, err := client.Do(request)
	if err != nil {
		return InvoicesReturn{}, err
	}

	// Close response
	defer response.Body.Close()

	// Decode data
	var decode InvoicesReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return InvoicesReturn{}, err
	}

	// Return data
	return decode, nil

}

// Create a new invoice
func NewInvoice(config Invoice) (NewInvoiceReturn, error) {

	// Define client
	client := &http.Client{}

	// Define body data
	body := url.Values{}
	body.Set("invoiceNumber", "")
	body.Set("contact[id]", config.ContactID)
	body.Set("contact[objectName]", "Contact")
	body.Set("invoiceDate", config.InvoiceDate)
	body.Set("header", "")
	body.Set("status", config.Status)
	body.Set("invoiceType", config.InvoiceType)
	body.Set("currency", "EUR")
	body.Set("mapAll", "true")
	body.Set("objectName", "Invoice")
	body.Set("discount", "false")
	body.Set("contactPerson[id]", config.ContactPerson)
	body.Set("contactPerson[objectName]", "SevUser")
	body.Set("taxType", "default")
	body.Set("taxRate", "")
	body.Set("taxText", "0")
	body.Set("showNet", "false")

	// New http request
	request, err := http.NewRequest("POST", "https://my.sevdesk.de/api/v1/Invoice", strings.NewReader(body.Encode()))
	if err != nil {
		return NewInvoiceReturn{}, err
	}

	// Set header
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", config.Token)

	// Response to sevDesk
	response, err := client.Do(request)
	if err != nil {
		return NewInvoiceReturn{}, err
	}

	// Close response
	defer response.Body.Close()

	// Decode data
	var decode NewInvoiceReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return NewInvoiceReturn{}, err
	}

	// Return data
	return decode, nil

}
