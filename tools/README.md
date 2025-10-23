# Tools

このディレクトリには、プロジェクトで利用する補助的なスクリプトを格納します。

## `generate_hokkaido_nandoku_list.py`

Wikipedia の「[北海道の難読地名一覧](https://ja.wikipedia.org/wiki/%E5%8C%97%E6%B5%B7%E9%81%93%E3%81%AE%E9%9B%A3%E8%AA%AD%E5%9C%B0%E5%90%8D%E4%B8%80%E8%A6%A7)」から地名と読みがなの一覧をCSV形式で取得し、標準出力へ出力するスクリプトです。

### 使い方

1. 依存パッケージをインストールします。
   ```shell
   uv pip install requests
   ```

2. スクリプトを実行します。
   ```shell
   uv run python generate_hokkaido_nandoku_list.py > nandoku.csv
   ```
   実行すると、`nandoku.csv` が生成されます。
