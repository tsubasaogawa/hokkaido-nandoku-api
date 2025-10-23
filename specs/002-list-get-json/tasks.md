# Implementation Tasks for `/list` Endpoint

This document breaks down the work required to implement the `GET /list` endpoint.

## Phase 1: Model and Repository Layer

-   **[X] Task 1.1: Update `Placename` Model**
    -   **File**: `source/internal/model/placename.go`
    -   **Change**: Modify the `Placename` struct to align with the data model. Ensure the JSON tags are correct (`name` and `yomi`).

-   **[X] Task 1.2: Implement `FindAll` Method in Repository**
    -   **File**: `source/internal/repository/placename_repository.go`
    -   **Change**: Create a new method `FindAll()` that returns all `Placename` objects from the in-memory store.
    -   **Test**: Add a unit test for `FindAll()` in `placename_repository_test.go`.

## Phase 2: Handler and Routing

-   **[X] Task 2.1: Create `ListPlaceNames` Handler**
    -   **File**: `source/internal/handler/handler.go`
    -   **Change**: Implement a new handler function `ListPlaceNames(w http.ResponseWriter, r *http.Request)` that calls the repository's `FindAll()` method and writes the result as a JSON response.
    -   **Test**: Add a unit test for the `ListPlaceNames` handler in `handler_test.go`.

-   **[X] Task 2.2: Add New Route**
    -   **File**: `source/cmd/api/main.go`
    -   **Change**: Register the `ListPlaceNames` handler for the `GET /list` route.

## Phase 3: Integration Testing

-   **[X] Task 3.1: Write Integration Test**
    -   **File**: `source/tests/integration/api_test.go`
    -   **Change**: Add an integration test that starts the API server, sends a `GET /list` request, and verifies that the response is a `200 OK` with the expected JSON structure and data.
