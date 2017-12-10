## go言語でHTTPクライアントをつくる
Alexaに喋らせる用の
JSONを返すAPIのローカルでのテストツールとして使用したい。

## 作る予定の機能
- [x] GETリクエストするだけ
- [ ] POSTできる
- [ ] ファイルに指定した内容をPOSTできる
- [ ] Alexa仕様のリクエストボディをPOSTできる
- [ ] 引数でリクエストタイプを指定できる
- [ ] リクエストボディに実行時日時でタイムスタンプを設定する


* Alexaのリクエスト仕様に沿ったリクエストの送信
  * リクエストの種類を選べる（Launch,Intent,SessionEnded）
  * タイムスタンプを自動で現在時刻を設定
  * レスポンスのSessionAttributeの内容を 次のリクエストに自動で設定

* APIのテスト用機能
  * レスポンスがAlexaのレスポンス仕様を満たすか検証する
  * 一連の会話シナリオを設定してテスト可能

## ローカルで動かす
### ◆ JSONサーバーの用意
nodeのjson-serverを導入
```
$ npm install -g json-server
```
応答用jsonファイルの作成
```db.jsono
{
  "id": 1,
  "name": "The Godfather",
  "director": "Francis Ford Coppola",
  "rating": 9.1
},
{
  "id": 2,
  "name": "Casablanca",
  "director": "Michael Curtiz",
  "rating": 8.8
}
```
立ち上げ
```
$ json-server --watch db.json

  \{^_^}/ hi!

  Loading db.json
  Done
```

### ◆ httpCliantの実行
```
go run httpCliant.go
```
