# 大学生の履修登録を簡易化するアプリ「ClassPlanner」(Backend)

## 概要
大学生の履修登録を簡易化するアプリ「ClassPlanner」のバックエンドリポジトリです。

## リプレイスのきっかけ
以前**Next.js**を使用してフルスタックなアプリケーションを作成しました。ですが、その際にフロントエンドとバックエンドのコードが混在してしまい、管理が煩雑になってしまいました。そのため、今回はフロントエンドとバックエンドを分離し、バックエンドのリポジトリを新たに作成しました。
またTypescript以外の言語を学ぶことで将来的な技術選定の際に幅広い選択肢を持つことができると考え、今回は**Go**を使って何かを作成してみようと思いリプレイスを行いました。

## 使用技術
- *Language*
  - **Go**(version 1.22.2)
- *Framework*
  - **Echo**(version 4.12.0)
- *DB*
  - **PostgreSQL**(image: postgres:15.1-alpine)
- *Other*
  - **GORM**(version 1.25.11)
  - **ozzo-validation**(version 4.3.0)
  - **Docker**

## ディレクトリ構成
```
.
├── auth        # 認証関連の処理
├── controller  # リクエストを受け取り、レスポンスを返す層
├── db          # データベースの初期化などの処理
├── middleware  # ミドルウェア
├── migrate     # マイグレーション処理
├── model       # DBのテーブル定義やレスポンスとして返すデータの構造体
├── repository  # DB操作
├── router      # ルーティング
├── usecase     # ビジネスロジック
└── validator   # バリデーション
```
## アーキテクチャと採用理由
- **Clean Architecture**
  - ビジネスロジックとデータベースアクセスを分離することで、ビジネスロジックの変更がデータベースアクセスに影響を与えないようにするため。

## 処理の流れ
![backend architecture](https://github.com/user-attachments/assets/c0e3a336-8140-46e6-a9b1-43cdfe9d1fd2)



## その他
- フロントエンドリポジトリ(https://github.com/kameiryohei/Ie-ClassPro)