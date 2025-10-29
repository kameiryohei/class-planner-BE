# 大学生の履修登録を簡易化するアプリ「ClassPlanner」(Backend)

## 概要

大学生の履修登録を簡易化するアプリ「ClassPlanner」のバックエンドリポジトリです。

## リプレイスのきっかけ

以前**Next.js**を使用してフルスタックなアプリケーションを作成しました。ですが、その際にフロントエンドとバックエンドのコードが混在してしまい、管理が煩雑になってしまいました。そのため、今回はフロントエンドとバックエンドを分離し、バックエンドのリポジトリを新たに作成することで管理を簡単にしました。
また Typescript 以外の言語を学ぶことで将来的な技術選定の際に幅広い選択肢を持つことができると考え、今回は**Go**を使って何かを作成してみようと思いリプレイスを行いました。

## 使用技術

- _Language_
  - **Go**(version 1.22.2)
- _Framework_
  - **Echo**(version 4.12.0)
- _DB_
  - **PostgreSQL**(image: postgres:15.1)
- _Other_
  - **GORM**(version 1.25.11)
  - **ozzo-validation**(version 4.3.0)
  - **Docker**(開発環境用)
  - **Github Actions**(CI/CD)
  - **Render**(本番環境ホスティング)

## ディレクトリ構成

```
.
├── auth        # 認証関連の処理
├── controller  # httpレスポンス作成層
├── db          # データベースの初期化
├── middleware  # jwt認証などのミドルウェア
├── migrate     # マイグレーション処理
├── model       # DBのテーブル定義やレスポンスとして返すデータの構造体を定義
├── repository  # DB操作
├── router      # ルーティング
├── seed        # 初期データ投入
├── usecase     # ビジネスロジック
└── validator   # バリデーション
```

## アーキテクチャと採用理由

- **クリーンアーキテクチャ**
  - ある企業のインターンに行った際に、レイヤードアーキテクチャを採用したプロジェクトを経験しました。その際に、各レイヤーが明確に分離されていることで、コードの可読性や保守性が向上していることを実感しました。なので、今回のプロジェクトを通してクリーンアーキテクチャの理解を深め、特徴の違いを学び今後のアーキテクチャ選定の際に役立てたいと考え採用しました。

## セットアップ手順

環境変数を設定。内容を適切に編集する。

```bash
cp .env.example .env
```

クローン後、以下の Make コマンドを使用してセットアップを行う。

````bash
make setup
```bash
make setup/first
````

その後サーバー起動

```bash
make run
```

以下の画像のように表示されれば成功。
![run-sever](https://github.com/user-attachments/assets/c650405e-d6b0-432c-bb71-eeafbb647364)

## その他

- フロントエンドリポジトリ(https://github.com/kameiryohei/Ie-ClassPro)
