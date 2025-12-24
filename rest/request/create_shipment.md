# How to Build a Request JSON


```json
{
  "profile": "string",
  "shipments": [
    {
      "product": "V01PAK",
      "billingNumber": "33333333330101 or 333333333362aa",
      "refNo": "stringst",
      "costCenter": "string",
      "creationSoftware": "string",
      "shipDate": "2019-08-24",
      "shipper": {
        "name1": "Blumen Krause",
        "name2": "To the attention of Erna.",
        "name3": "Backdrawer all the way back.",
        "addressStreet": "Hauptstrasse",
        "addressHouse": "1a",
        "postalCode": "53113",
        "city": "Berlin",
        "country": "ABW",
        "contactName": "Konrad Kontaktmann",
        "email": "mustermann@example.com"
      },
      "consignee": {
        "name1": "Blumen Krause",
        "name2": "To the attention of Erna.",
        "name3": "Backdrawer all the way back.",
        "dispatchingInformation": "PO Box, bpack 24/7",
        "addressStreet": "Hauptstrasse",
        "addressHouse": "1a",
        "additionalAddressInformation1": "3. Etage",
        "additionalAddressInformation2": "Apartment 12",
        "postalCode": "53113",
        "city": "Berlin",
        "state": "NRW",
        "country": "ABW",
        "contactName": "Konrad Kontaktmann",
        "phone": "+49 170 1234567",
        "email": "mustermann@example.com"
      },
      "details": {
        "dim": {
          "uom": "cm",
          "height": 10,
          "length": 20,
          "width": 15
        },
        "weight": {
          "uom": "g",
          "value": 500
        }
      },
      "services": {
        "preferredNeighbour": "Please ring at Meier next door",
        "preferredLocation": "Please leave in carport",
        "visualCheckOfAge": "A18",
        "namedPersonOnly": true,
        "identCheck": {
          "firstName": "Max",
          "lastName": "Mustermann",
          "dateOfBirth": "2019-08-24",
          "minimumAge": "A18"
        },
        "signedForByRecipient": true,
        "endorsement": "RETURN",
        "preferredDay": "2019-08-24",
        "noNeighbourDelivery": true,
        "additionalInsurance": {
          "currency": "AED",
          "value": 0
        },
        "bulkyGoods": true,
        "cashOnDelivery": {
          "amount": {
            "currency": "AED",
            "value": 0
          },
          "bankAccount": {
            "accountHolder": "John D. Rockefeller",
            "bankName": "The Iron Bank, Braavos",
            "iban": "DE02100100100006820101",
            "bic": "DEUTDEFFXXX"
          },
          "accountReference": "string",
          "transferNote1": "string",
          "transferNote2": "string"
        },
        "individualSenderRequirement": "ZZ",
        "premium": true,
        "closestDropPoint": true,
        "parcelOutletRouting": "max.mustermann@example.com",
        "goGreenPlus": true,
        "dhlRetoure": {
          "billingNumber": "string",
          "refNo": "string",
          "returnAddress": {
            "name1": "Blumen Krause",
            "name2": "To the attention of Erna.",
            "name3": "Backdrawer all the way back.",
            "dispatchingInformation": "PO Box, bpack 24/7",
            "addressStreet": "Hauptstrasse",
            "addressHouse": "1a",
            "additionalAddressInformation1": "3. Etage",
            "additionalAddressInformation2": "Apartment 12",
            "postalCode": "53113",
            "city": "Berlin",
            "state": "NRW",
            "country": "ABW",
            "contactName": "Konrad Kontaktmann",
            "phone": "+49 170 1234567",
            "email": "mustermann@example.com"
          },
          "goGreenPlus": true
        },
        "postalDeliveryDutyPaid": true
      },
      "customs": {
        "invoiceNo": "string",
        "exportType": "OTHER",
        "exportDescription": "Detailed description for OTHER goods.",
        "shippingConditions": "DAP",
        "permitNo": "string",
        "attestationNo": "string",
        "hasElectronicExportNotification": true,
        "MRN": "abcd1234567890",
        "postalCharges": {
          "currency": "AED",
          "value": 0
        },
        "officeOfOrigin": "string",
        "shipperCustomsRef": "DE73282932000074",
        "consigneeCustomsRef": "GB73282932000074",
        "items": [
          {
            "itemDescription": "T-Shirt",
            "hscode": "61099090",
            "countryOfOrigin": "DE",
            "packagedQuantity": 3
          },
          {
            "itemDescription": "Book",
            "hscode": "49019900",
            "packagedQuantity": 1
          }
        ]
      }
    }
  ]
}