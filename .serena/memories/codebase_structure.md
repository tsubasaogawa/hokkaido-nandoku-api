# Codebase Structure

- `source/`: Contains the Go source code for the API.
  - `cmd/api/`: The main application entry point.
  - `data/`: Contains the CSV data file for place names.
  - `internal/`: Contains the core application logic, separated into:
    - `handler/`: HTTP handlers.
    - `model/`: Data models.
    - `repository/`: Data access logic.
  - `pkg/csvloader/`: A utility package for loading data from CSV files.
- `terraform/`: Contains the Terraform code for managing AWS infrastructure.
- `specs/`: Contains feature specifications and design documents.
- `tools/`: Contains Python scripts for development and data generation purposes.
- `.gemini/`: Contains configuration for the Gemini CLI agent.
