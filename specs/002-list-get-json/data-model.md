# Data Model for Feature: /list エンドpoint

## 1. Logical Data Model

### Entity: `Placename`
- **Description**: Represents a place name with its reading.
- **Attributes**:
  - `name` (string, mandatory): The official name of the place.
    - Example: "支笏湖"
  - `yomi` (string, mandatory): The phonetic reading of the name in hiragana.
    - Example: "しこつこ"

## 2. Data Flow

1.  **Source**: The raw data originates from the `source/data/nandoku_chimei.csv` file.
2.  **Loading**: At application startup, the CSV data is loaded into memory by the `placename_repository`.
3.  **API Request**: A client sends a `GET /list` request.
4.  **Handler**: The `ListPlaceNames` handler in the API receives the request.
5.  **Repository Access**: The handler calls the repository to fetch the full list of `Placename` entities.
6.  **Response**: The handler formats the list of entities into the specified JSON structure and sends it as the HTTP response.
