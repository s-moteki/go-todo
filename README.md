# Go言語 TODO管理API  

Goの学習目的で作成しているTODOタスク管理アプリのAPI

## 使用技術、ライブラリ  

|  項目  |  用途  |
| ---- | ---- |
|  Go  |  言語  |
|  MySQL  |  DB  |
|  docker  |  コンテナ  |
|  docker-compose  | な |  
|  gin  |  フレームワーク  |
|  godotenv  | 環境変数 |

### あとでやる  

ルーティング以外でもginを使う  
ちゃんとinputなど分ける  
バリデーション、エラーハンドリング  

### やるやら  

gorm  
ホットリロード  
Go Modules  
デプロイ  

## 環境構築  

dockerやらをインストール後

1.クローン  
> git clone https://github.com/s-moteki/go-todo.git

2.移動  
> cd go-todo

3.コンテナ起動  
> docker-compose up --build

### るーてぃんぐ  

一覧取得
> curl localhost:8080/todos  

登録
> curl -i -H "Content-type: application/json" -X POST -d '{"Title":"test","Content":"insert test","Deadline":"2020-08-10T17:25:16Z","State":1}' localhost:8080/todos  

更新  
> curl -i -H "Content-type: application/json" -X PATCH -d '{"Title":"test","Content":"insert test","Deadline":"2020-08-10T17:25:16Z","State":1}' localhost:8080/todos/1  

取得  
> curl localhost:8080/todos/1  

タイトル検索  
> curl localhost:8080/search?title=titlename
