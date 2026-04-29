# viseca-exporter

Little helper to get transactions from Viseca One and print them in CSV format.

## Usage

While the "auth w/o browser" method is a very comfortable way to get the transactions, it has a higher risk to error due to changes to the auth flow of Viseca One. If it fails, try the "auth w/ browser" method.

### Auth w/o browser

This method processes the auth flow in the CLI and will trigger a 2FA request like the login in a browser would.

1. Log in to [one.viseca.ch](https://one.viseca.ch)
1. Go to "Transaktionen" on [one.viseca.ch](https://one.viseca.ch)
1. Save the card ID from the path (between `/v1/card/` and `/transactions`)

1.  ```
    export VISECA_CLI_USERNAME=<email>
    export VISECA_CLI_PASSWORD=<password>
    go run cmd/viseca-cli/main.go transactions <cardID>
    ```

### Auth w/ browser

This method requires a valid session cookie obtained from an authenticated browser session.

1. Log in to [one.viseca.ch](https://one.viseca.ch)
1. Open the developer tools of your browser and navigate to the network tab
1. Go to "Transaktionen" on [one.viseca.ch](https://one.viseca.ch)
1. Filter the URLs in the network tab of the developer tools for `transactions`
1. Save the card ID from the path (between `/v1/card/` and `/transactions`) to an env file (see examples)
1. Save the session cookie (`AL_SESS-S=AAAAAA...`) to an env file (see examples)

1.  ```
    source .env
    go run main.go "$VISECA_CARD" "$VISECA_SESS" > data/export.csv
    ```

## Example .env 

```
VISECA_SESS=AL_SESS-S=xxxxxxxxxxxxxxxyyyyyy
VISECA_CARD=443592xxxxxxxxxx
```

## Example Output

```csv
"TransactionID","Date","Merchant","Amount","PFMCategoryID","PFMCategoryName"
"AUTH8c919db2-1c23-43f1-8862-61c31336d9b6","2021-10-20T17:05:44","ALDI","50.550000","cv_groceries","Groceries"
```

## Foregin currencies

To include foreign currencies, true can be provided as the third argument to `main.go`:
```shell
go run main.go "$VISECA_CARD" "$VISECA_SESS" "true" > data/export.csv
```

or the foreignCurrencies flag set for the CLI:
```shell
go run cmd/viseca-cli/main.go transactions --foreign-currency=true
```

This will result in an output like the following, additionally including the currency, original amount and original currency:
```csv
"TransactionID","Date","Merchant","Amount","Currency","OriginalAmount","OriginalCurrency","PFMCategoryID","PFMCategoryName"
"AUTH8c919db2-1c23-43f1-8862-61c31336d9b6","2021-10-20T17:05:44","ALDI","50.55","CHF","20.00","USD","cv_groceries","Groceries"
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
