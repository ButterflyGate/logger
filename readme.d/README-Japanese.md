# logger
Butterfly Gate標準ログライブラリ

## level
- ログレベル一覧

| name        | number | destination | description |
| ---:        | :----: | :---------: | :---------- |
| None        | 0      | stderr      | 出力なし |
| Emergency   | 1      | stderr      | Emergency レベルのみ出力 |
| Critical    | 2      | stderr      | Critical と Emergency レベルを出力 |
| Error       | 3      | stderr      | Error レベル以上のログを出力 |
| Alert       | 4      | stderr      | Alert レベル以上のログを出力 |
| Warning     | 5      | stderr      | Warning レベル以上のログを出力 |
| Notice      | 6      | stdout      | Notice レベル以上のログを出力 |
| Information | 7      | stdout      | Information レベル以上のログを出力 |
| Debug       | 8      | stdout      | Debug レベル以上のログを出力 |
| Trace       | 99     | stdout      | Trace レベル以上のログを出力 |

## Options
- 出力制御オプション(下記項目それぞれの表示、非表示を指定可能)
    - ログレベル
    - タイムスタンプ
    - 呼び出し位置
    - メッセージ (ソース変更の必要有り)
    - データ (ソース変更の必要有り)
    - 変数名 (ソース変更の必要有り)

- 出力形式の選択
  - 出力時のインデント使用有無 (json形式のみ)
  - 出力行数制御 (文字列またはエラーの出力時)
  - json形式、テキスト形式の選択 (未実装)
  - 独自フォーマットで出力 (未実装)

