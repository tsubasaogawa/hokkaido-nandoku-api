# Tasks: API Gatewayに/listエンドポイントを追加

**Input**: Design documents from `/specs/003-api-gateway-list/`
**Prerequisites**: plan.md (required), research.md, data-model.md, contracts/

## Phase 3.1: Setup
- [X] T001: [Terraform] API Gatewayに`/list`リソースと`GET`メソッドを追加する。
    - **File**: `terraform/main.tf`
    - **Details**: `aws_api_gateway_resource` を使用して `/list` パスを作成し、`aws_api_gateway_method` を使用して `GET` メソッドを追加します。このメソッドを既存のLambda関数に統合します。

## Phase 3.2: Core Implementation
- [X] T002: [Go] `placename.go` のJSONタグを修正する。 [P]
    - **File**: `source/internal/model/placename.go`
    - **Details**: `NandokuChimei` 構造体のフィールドタグを `json:"name"` と `json:"yomi"` に更新します。
- [X] T003: [Go] `handler.go` のロジックを、すべての地名リストを返すように修正する。
    - **File**: `source/internal/handler/handler.go`
    - **Details**: `placename_repository` を呼び出してすべての地名データを取得し、それをJSON形式で返すように `HandleRequest` 関数を修正します。
- [X] T004: [Go] `handler_test.go` を修正し、`/list`エンドポイントのレスポンスを検証する。 [P]
    - **File**: `source/internal/handler/handler_test.go`
    - **Details**: ハンドラが `name` と `yomi` キーを持つJSONオブジェクトの配列を正しく返すことを確認するようにテストケースを更新します。

## Phase 3.3: Build & Deploy
- [X] T005: [Build] アプリケーションをビルドし、ローカルでテストする。
    - **File**: `source/build.sh`
    - **Details**: `build.sh` を実行してGoアプリケーションをビルドし、Dockerコンテナを起動して `curl` コマンドでローカルエンドポイントの動作を確認します。
- [ ] T006: [Deploy] Terraformの変更を適用する。
    - **Directory**: `terraform/`
    - **Details**: `terraform apply` を実行して、API Gatewayの変更をAWSにデプロイします。

## Phase 3.4: Validation
- [ ] T007: デプロイされた`/list`エンドポイントをテストする。
    - **Details**: `terraform output` からAPI GatewayのURLを取得し、`curl <URL>/list` を実行して、地名リストが正しく返されることを確認します。

## Dependencies
- T002 と T004 は並行して実行可能です。
- T003 は T002 の完了後に実行する必要があります。
- T005 (ビルド) は、T003 の完了後に行います。
- T006 (デプロイ) は、T001 と T005 の完了後に行います。
- T007 (検証) は、T006 の完了後に行います。

## Parallel Example
```
# T002とT004は同時に開始できます
Task: "[Go] placename.go のJSONタグを修正する。"
Task: "[Go] handler_test.go を修正し、/listエンドポイントのレスポンスを検証する。"
```
