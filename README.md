# ReadersLounge
●読書好きが繋がるというコンセプトのSNSアプリです。<br>
●読んだ本の感想を投稿したり、他の人が書いた投稿を読んでコメントするなど、面白いと思った本の感想を他者と共有することができます。<br>
●また、趣味が合うユーザーをフォローしたり、チャット機能を用いて会話をしたりなど、ユーザー同士での交流が可能です。<br>
<br>
![login](https://github.com/kato0209/ReadersLounge/assets/89386373/44d7079c-beb8-4c8a-970a-3e704c5a9463)
![home](https://github.com/kato0209/ReadersLounge/assets/89386373/fc5b708c-94b2-4f55-862b-a1bc36c29d6f)


# URL
https://readerslounge-server.com/  <br>

# 使用技術
- Golang(echo)
- Next.js(TypeScript)
- Docker
- OpenAPI
- AWS
- Terraform
- GitHubActions

# AWS構成図
![ReadersLounge_AWS drawio (4)](https://github.com/kato0209/ReadersLounge/assets/89386373/3e4f8fe7-b403-4de6-ab89-84029a453a06)

# 機能一覧
- ユーザー登録、ログイン機能(+GoogleOAuth)
- 感想投稿機能
  - 5段階評価機能
  - 感想を投稿したい本をジャンルや、キーワード検索で探すことができる
- コメント機能
- チャット機能
- ユーザー検索機能
- いいね機能
- フォロー機能
- ユーザープロフィール画像変更機能

## 機能一覧画像例
### 感想投稿の流れ↓
- 感想を投稿したい本をキーワード or ジャンルで検索
![bookSearch](https://github.com/kato0209/ReadersLounge/assets/89386373/4311a73d-f63d-4afe-8abc-2f93a9686d4a)<br>

- 対象の本を選択
![book](https://github.com/kato0209/ReadersLounge/assets/89386373/00e9b18c-faa7-426a-a08d-45f1bb64a6bf)<br>

- 投稿
![post](https://github.com/kato0209/ReadersLounge/assets/89386373/78e74c31-5b7f-4edd-8b59-3f39466c305c)<br>


### チャット機能
![chat](https://github.com/kato0209/ReadersLounge/assets/89386373/a501fdf7-8d0e-435e-95da-39b700d3b857)

