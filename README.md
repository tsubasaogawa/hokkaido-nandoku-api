# 北海道の難読地名API

北海道の難読地名をランダムに返すAPIです。

## Quickstart

このガイドは、APIをローカル環境でビルドし、実行するための手順を説明します。

### 前提条件
- Go (1.22 or later) がインストールされていること

### 1. データファイルの準備
`source/data/nandoku_chimei.csv` に地名のCSVデータが配置されていることを確認します。

### 2. ビルド
`source` ディレクトリに移動し、アプリケーションをビルドします。

```bash
cd source
go build -o bootstrap ./cmd/api
```
これにより、`source` ディレクトリ直下に `bootstrap` という名前の実行可能ファイルが生成されます。

### 3. ローカルでの実行
`source` ディレクトリ内で、以下のコマンドを実行してAPIサーバーをローカルで起動します。

```bash
./bootstrap
```
サーバーはデフォルトでポート `8080` をリッスンします。

### 4. 動作確認
別のターミナルから `curl` コマンドを使ってAPIにリクエストを送信し、動作を確認します。

```bash
curl http://localhost:8080/random
```

成功すると、以下のようなJSONレスポンスが返却されます。
```json
{"name":"国縫","yomi":"くんぬい"}
```

## API Routes

### GET /random

ランダムに選ばれた北海道の難読地名とその読みを返します。

**レスポンスサンプル:**
```json
{
  "name": "国縫",
  "yomi": "くんぬい"
}
```

### GET /list

北海道の難読地名とその読みの一覧を返します。

**レスポンスサンプル:**
```json
[
  {
    "name": "足寄",
    "yomi": "あしょろ"
  },
  {
    "name": "神恵内",
    "yomi": "かもえない"
  },
  {
    "name": "国縫",
    "yomi": "くんぬい"
  }
 :
]
```

## AWSへのデプロイ (Terraform)
Terraformを使用して、GoアプリケーションのビルドからAWSへのデプロイまでを実行します。

### 前提条件
- Terraform (1.0 or later) がインストールされていること
- AWS CLIがインストールされ、認証情報が設定済みであること
- `zip` コマンドがインストールされていること

### 手順
1. **デプロイパッケージを作成します。**
   `source` ディレクトリに移動し、ビルドスクリプトを実行します。
   ```bash
   cd source
   ./build.sh
   ```
   これにより、`source` ディレクトリに `hokkaido-nandoku-api.zip` が作成されます。

2. **Terraformを実行してインフラを構築・デプロイします。**
   `terraform/` ディレクトリに移動し、以下のコマンドを実行します。
   ```bash
   cd ../terraform

   # 初期化
   terraform init

   # 実行計画の確認
   terraform plan

   # 適用
   terraform apply
   ```

   `apply` が完了すると、出力としてAPI GatewayのエンドポイントURLが表示されます。

3. **デプロイされたAPIの動作確認**
   出力されたエンドポイントURLに対して `curl` でリクエストを送信します。
   ```bash
   curl https://<api-gateway-id>.execute-api.ap-northeast-1.amazonaws.com/random
   ```
   成功すれば、ローカルでの実行時と同様のJSONレスポンスが返ってきます。
