# Quickstart

## ローカルでの実行

1. `source`ディレクトリに移動します。
   ```bash
   cd source
   ```
2. ビルドします。
   ```bash
   ./build.sh
   ```
3. Dockerコンテナを起動します。
   ```bash
   docker run -p 9000:8080 hokkaido-nandoku-api:latest
   ```
4. 別のターミナルからAPIを叩きます。
   ```bash
   curl http://localhost:9000/2015-03-31/functions/function/invocations -d '{}'
   ```

## デプロイ

1. `terraform`ディレクトリに移動します。
   ```bash
   cd terraform
   ```
2. `terraform apply`を実行します。
   ```bash
   terraform apply
   ```
3. 出力されたAPI GatewayのURLにアクセスします。
   ```bash
   curl <API GatewayのURL>/list
   ```
