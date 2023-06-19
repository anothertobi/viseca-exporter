# viseca-exporter

Little helper to get transactions from Viseca One and print them in CSV format.

## Usage

1. Log in to [one.viseca.ch](https://one.viseca.ch)
1. Go to "Transaktionen" on [one.viseca.ch](https://one.viseca.ch)
1. Save the card ID from the path (between `/v1/card/` and `/transactions`)

1.  ```
    export VISECA_CLI_USERNAME=<email>
    export VISECA_CLI_PASSWORD=<password>
    go run cmd/viseca-cli/main.go transactions <cardID>
    ```

### CLI Output

```csv
"TransactionID","Date","Merchant","Amount","PFMCategoryID","PFMCategoryName"
"AUTH8c919db2-1c23-43f1-8862-61c31336d9b6","2021-10-20T17:05:44","ALDI","50.550000","cv_groceries","Groceries"
```

## API

### Known issues

The API returns `500 Internal Server Error` without any additional information when a request doesn't meet the API requirements.

Large page sizes (e.g. 1000) lead to an error.

### API Output

```json
{
    "totalCount": 1,
    "list": [
        {
            "transactionId": "AUTH8c919db2-1c23-43f1-8862-61c31336d9b6",
            "cardId": "0000000AAAAA0000",
            "maskedCardNumber": "XXXXXXXXXXXX0000",
            "cardName": "Mastercard",
            "date": "2021-10-20T17:05:44",
            "showTimestamp": true,
            "amount": 50.55,
            "currency": "CHF",
            "originalAmount": 50.55,
            "originalCurrency": "CHF",
            "merchantName": "Aldi Suisse 00",
            "prettyName": "ALDI",
            "merchantPlace": "",
            "isOnline": false,
            "pfmCategory": {
                "id": "cv_groceries",
                "name": "Lebensmittel",
                "lightColor": "#E2FDD3",
                "mediumColor": "#A5D58B",
                "color": "#51A127",
                "imageUrl": "https://api.one.viseca.ch/v1/media/categories/icon_with_background/ic_cat_tile_groceries_v2.png",
                "transparentImageUrl": "https://api.one.viseca.ch/v1/media/categories/icon_without_background/ic_cat_tile_groceries_v2.png"
            },
            "stateType": "authorized",
            "details": "Aldi Suisse 00",
            "type": "merchant",
            "isBilled": false,
            "links": {
                "transactiondetails": "/v1/card/0000000AAAAA0000/transaction/AUTH8c919db2-1c23-43f1-8862-61c31336d9b6"
            }
        }
    ]
}
```
