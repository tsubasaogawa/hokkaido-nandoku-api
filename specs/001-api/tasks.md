# Tasks: 北海道の難読地名を取得するAPI

**Input**: Design documents from `/specs/001-api/`
**Prerequisites**: plan.md, research.md, data-model.md, contracts/

## Phase 3.1: プロジェクトセットアップ
- [x] T001: `plan.md` に基づき、`cmd/api`, `internal/handler`, `internal/model`, `internal/repository`, `pkg/csvloader`, `tests/unit`, `tests/integration`, `data`, `terraform` のディレクトリを作成する。
- [x] T002: プロジェクトルートで `go mod init github.com/your-username/hokkaido-nandoku-api` を実行し、Goモジュールを初期化する。
- [x] T003: `data/nandoku_chimei.csv` にサンプルデータを3件以上追加する。

## Phase 3.2: Infrastructure as Code (Terraform)
- [x] T004: `terraform/variables.tf` を作成し、`aws_region`, `project_name` などの変数を定義する。
- [x] T005: `terraform/outputs.tf` を作成し、API GatewayのURLを出力するように定義する。
- [x] T006: `terraform/main.tf` を作成し、`terraform-aws-modules/lambda/aws` モジュールを使用して、ビルド済みのZIPファイルをデプロイするLambda関数とAPI Gatewayの定義を記述する。

## Phase 3.3: 実装 (TDD)

### Core Logic & Data Layer
- [x] T007 [P]: `internal/model/placename.go` に `PlaceName` struct を定義する。
- [x] T008 [P]: `pkg/csvloader/loader_test.go` に、CSVファイルを正しく読み込めるかテストするユニットテストを記述する。
- [x] T009: T008のテストが失敗することを確認後、`pkg/csvloader/loader.go` にCSVファイルを読み込み `[]model.PlaceName` スライスを返す `LoadPlaceNames` 関数を実装する。
- [x] T010 [P]: `internal/repository/placename_repository_test.go` に、`FindRandom` メソッドがスライスからランダムな要素を1つ返すことを確認するユニットテストを記述する。
- [x] T011: T010のテストが失敗することを確認後、`internal/repository/placename_repository.go` に `PlaceNameRepository` インターフェースと、CSVローダーを利用する実装を記述する。

### Handler & API Layer
- [x] T012 [P]: `internal/handler/handler_test.go` に、リポジトリから取得した地名を正しくJSONレスポンスとして返すことを確認するユニットテストを記述する (`httptest` を使用)。
- [x] T013: T012のテストが失敗することを確認後、`internal/handler/handler.go` に `RandomPlaceNameHandler` を実装する。

### Application Entrypoint & Build Script
- [x] T014: `cmd/api/main.go` に、依存関係（リポジトリ、ハンドラ）を初期化し、Lambdaの実行を開始する `main` 関数を実装する。
- [x] T015: `source` ディレクトリに、GoバイナリとデータファイルをZIP化する `build.sh` スクリプトを作成する。

## Phase 3.4: 統合とデプロイ
- [x] T016: `tests/integration/api_test.go` に、デプロイされたエンドポイントに対して実際にHTTPリクエストを送り、動作確認を行うインテグレーションテストを記述する。
- [x] T017a: `source/build.sh` を実行し、`hokkaido-nandoku-api.zip` を作成する。
- [x] T017b: `terraform apply` を実行し、AWSへのデプロイを完了させる。
- [x] T018: T016のインテグレーションテストを実行し、パスすることを確認する。

## Phase 3.5: ドキュメント
- [ ] T019: プロジェクトルートの `README.md` を作成または更新し、`quickstart.md` の内容を参考に、プロジェクトの概要、ローカルでの実行方法、Terraformを使ったデプロイ手順を記述する。
