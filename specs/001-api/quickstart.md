# Quickstart:北海道の難読地名API

このガイドは、APIをローカル環境でビルドし、実行するための手順を説明します。

## 前提条件
- Go (1.22 or later) がインストールされていること
- Dockerがインストールされていること（コンテナビルドのため）

## 1. データファイルの準備
プロジェクトルートに `data` ディレクトリを作成し、その中に `nandoku_chimei.csv` という名前で以下の形式のCSVファイルを配置します。

成功すると、以下のようなJSONレス- **`data/nandoku_chimei.csv`**
    ```csv
    name,yomi
    長万部,おしゃまんべ
    国縫,くんぬい
    礼文,れぶん
    ```

## 2. ビルド
プロジェクトのルートディレクトリで以下のコマンドを実行し、アプリケーションをビルドします。

```bash
go build -o bootstrap ./cmd/api
```
これにより、`bootstrap` という名前の実行可能ファイルが生成されます。

## 3. ローカルでの実行
ビルドが完了したら、以下のコマンドでAPIサーバーをローカルで起動します。

```bash
./bootstrap
```
サーバーはデフォルトでポート `8080` をリッスンします。

## 4. 動作確認
別のターミナルから `curl` コマンドを使ってAPIにリクエストを送信し、動作を確認します。

```bash
curl http://localhost:8080/random
```

成功すると、以下のようなJSONレスポンスが返却されます。
```json
{"name":"国縫","yomi":"くんぬい"}
```

## 5. AWSへのデプロイ (Terraform)
Terraformを使用して、アプリケーションのコンテナイメージのビルドからAWSへのデプロイまでを一度に実行します。

### 前提条件
- Terraform (1.0 or later) がインストールされていること
- Dockerがインストールされていること
- AWS CLIがインストールされ、認証情報が設定済みであること

### 手順
1. **Terraformを実行してインフラを構築・デプロイします。**
   `terraform/` ディレクトリに移動し、以下のコマンドを実行します。
   ```bash
   cd terraform

   # 初期化
   terraform init

   # 実行計画の確認
   terraform plan

   # 適用（コンテナのビルドとデプロイが自動的に実行されます）
   terraform apply
   ```

   `apply` が完了すると、出力としてAPI GatewayのエンドポイントURLが表示されます。

2. **デプロイされたAPIの動作確認**
   出力されたエンドポイントURLに対して `curl` でリクエストを送信します。
   ```bash
   curl https://<api-gateway-id>.execute-api.ap-northeast-1.amazonaws.com/prod/random
   ```
   成功すれば、ローカルでの実行時と同様のJSONレスポンスが返ってきます。


