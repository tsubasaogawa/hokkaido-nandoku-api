#!/usr/bin/env python3
"""
wikipedia_hokkaido_nandoku_list.py

用途:
  MediaWiki API を利用して、Wikipedia の「北海道の難読地名一覧」ページから
  地名とその読み（ふりがな）を正確に取得する。

  HTMLスクレイピングではなく、API経由で各セクションの元データ (Wikitext) を
  直接取得するため、ページのレイアウト変更に強く、高い精度が期待できる。

使い方:
  1) 依存パッケージをインストール:
     uv pip install requests
  2) 実行:
     uv run python generate_hokkaido_nandoku_list.py

出力:
  - 標準出力に「漢字,読み」をCSV形式で一覧表示。
"""

import requests
import re

API_URL = "https://ja.wikipedia.org/w/api.php"
PAGE_TITLE = "北海道の難読地名一覧"
HEADERS = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"
}

def get_sections():
    """ページの全セクション情報を取得する。"""
    params = {
        "action": "parse",
        "page": PAGE_TITLE,
        "prop": "sections",
        "format": "json",
        "formatversion": 2
    }
    try:
        response = requests.get(API_URL, params=params, headers=HEADERS)
        response.raise_for_status()
        data = response.json()
        return data.get("parse", {}).get("sections", [])
    except requests.exceptions.RequestException as e:
        print(f"Error fetching sections: {e}")
        return []

def get_section_wikitext(section_index):
    """指定されたセクションのWikitextを取得する。"""
    params = {
        "action": "parse",
        "page": PAGE_TITLE,
        "prop": "wikitext",
        "section": section_index,
        "format": "json",
        "formatversion": 2
    }
    try:
        response = requests.get(API_URL, params=params, headers=HEADERS)
        response.raise_for_status()
        data = response.json()
        return data.get("parse", {}).get("wikitext", "")
    except requests.exceptions.RequestException as e:
        print(f"Error fetching wikitext for section {section_index}: {e}")
        return ""

def main():
    print("漢字,読み")

    sections = get_sections()
    if not sections:
        return

    excluded_sections = ["参考文献", "関連項目", "外部リンク"]

    # Wikitextから地名と読みを抽出するための正規表現
    # - `^*+` : 行頭の `*` (1回以上) にマッチ (入れ子リスト対応)
    # - `\s*` : 空白文字
    # - `(.+?)` : 地名部分 (非貪欲マッチ)
    # - `\s*（` : `（`
    # - `([^）]+)` : `）` 以外の文字 (読み)
    # - `）` : `）`
    line_pattern = re.compile(r"^\*+\s*(.+?)\s*（([^）]+)）")

    for section in sections:
        if section.get("level") == "2" and section.get("line") not in excluded_sections:
            wikitext = get_section_wikitext(section["index"])
            
            for line in wikitext.splitlines():
                match = line_pattern.match(line)
                if match:
                    # 地名部分からWikiリンクのマークアップを除去
                    kanji_raw = match.group(1).strip()
                    kanji = re.sub(r"\[\[(?:[^|\]]+\|)?([^\]]+)\]\]", r"\1", kanji_raw)

                    # 読み仮名から注釈タグを除去
                    reading_raw = match.group(2).strip()
                    reading = re.sub(r'<ref.*?>.*?</ref>', '', reading_raw)
                    
                    print(f"{kanji},{reading}")

if __name__ == "__main__":
    main()
