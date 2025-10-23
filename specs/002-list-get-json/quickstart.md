# Quickstart for `/list` Endpoint

This document provides a quick guide on how to use the new `GET /list` endpoint.

## Endpoint

`GET /list`

## Description

Retrieves a complete list of all difficult-to-read place names in Hokkaido and their corresponding readings.

## Successful Response (200 OK)

### Example Request
```bash
curl -X GET http://localhost:8080/list
```

### Example Response Body
```json
{
  "placenames": [
    {
      "name": "支笏湖",
      "yomi": "しこつこ"
    },
    {
      "name": "積丹",
      "yomi": "しゃこたん"
    }
  ]
}
```

## Error Responses

- This endpoint does not have any specific error responses under normal conditions. If the server is down or a catastrophic error occurs, a standard `500 Internal Server Error` might be returned.
- If the data source is empty, the endpoint will return a `200 OK` with an empty `placenames` array:
  ```json
  {
    "placenames": []
  }
  ```
