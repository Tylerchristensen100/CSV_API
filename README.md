# CSV API

An API to effortlessly convert data from any CSV file into a JSON format, with powerful options for sorting and filtering.

---

## Table of Contents

* [Overview](#overview)
* [Getting Started](#getting-started)
    * [Base URLs](#base-urls)
* [API Endpoints](#api-endpoints)
    * [Get CSV Data](#get-csv-data)
* [Responses](#responses)
* [Contact](#contact)

---

## Overview

The **CSV API** simplifies working with CSV data by providing a convenient way to transform it into JSON. You can easily fetch CSV files from a given URL and then sort or filter the data based on your specific needs, all through simple API calls.

---

## Getting Started

### Base URLs

You can access the API at the following base URLs:

* `http://localhost:4000` (for local development)
* `https://csv.freethegnomes.org/`

---

## API Endpoints

### Get CSV Data

`GET /`

Converts any CSV into a JSON format, allowing for sorting and filtering of the data.

#### Parameters

| Name       | In    | Required | Description                                                                   | Type   |
| :--------- | :---- | :------- | :---------------------------------------------------------------------------- | :----- |
| `url`      | query | optional | URL of the CSV file to fetch.                                                 | string |
| `sortBy`   | query | optional | Sort by a specific header in the CSV file. If not provided, the first header in the CSV will be used as the default. | string |
| `filterBy` | query | optional | Filter the CSV data by a specific header and value. Format: `header==value`.  | string |

#### Example Request

To get data from a CSV and sort it by a column named "Name":

`GET /?url=https://example.com/data.csv&sortBy=Name`


To filter data where the "City" column is "New York":
`GET /?url=https://example.com/data.csv&filterBy=City==New York`


---

## Responses

### Success (200 OK)

This response returns the **CSV data converted into JSON**. The `key` field indicates the header used for sorting, and the `value` field contains an array of objects, where each object represents a row from the CSV with header names as keys.

```json
{
  "key": "Name",
  "value": [
    {
      "header1": "Alice",
      "header2": "Doe",
      "...": "..."
    },
    {
      "header1": "Bob",
      "header2": "Smith",
      "...": "..."
    }
  ]
}
```
### Bad Request (400)

This error is returned when there's an issue with your request, such as an **invalid CSV URL** or an **incorrect `filterBy` format**.

```json
{
  "statusCode": 400,
  "message": "Invalid URL provided"
}
```

### Internal Server Error (500)

Indicates an unexpected error on the server side.

```json
{
  "statusCode": 500,
  "message": "An internal server error occurred"
}
```

---


### Contact
- Name: Tyler Christensen

- Website: [https://freethegnomes.org](https://freethegnomes.org)

- Email: [tylerchristensen100@gmail.com](mailto:tylerchristensen100@gmail.com)