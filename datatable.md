### snippet
- `id`
	- ID
- `name`
	- スニペット名
- `body`
	- スニペットの中身
- `type_id`
    - どの言語のスニペットなのか(`javascript`, `swift`)
- `created_at`
	- 作成日
- `updated_at`
	- 更新日
- `created_by`
	- 作成者の`user_id`

```sql
CREATE TABLE snippet (
  id           BIGSERIAL PRIMARY KEY,         -- 一意のスニペットID
  name         VARCHAR(255) NOT NULL,         -- スニペット名
  body         TEXT         NOT NULL,         -- コード本体
  type_id      BIGINT       NOT NULL,         -- 言語ID
  created_by   BIGINT       NOT NULL,         -- 作成者 (user.id)
  created_at   TIMESTAMP    NOT NULL DEFAULT now(),
  updated_at   TIMESTAMP    NOT NULL DEFAULT now()
);
```

---
### profile
- `id`
	- ID
- `username`
	- ユーザーID
- `email`
	- メール
- `password`
	- パスワード
- `user_id`
	- ユーザーID
- `is_deleted`
	- 退会したかどうか
- `last_login_at`
	- 最終ログイン日
- `created_at`
	- 作成日
- `update_at`
	- 更新日

```sql
CREATE TABLE profiles (
  id            BIGSERIAL PRIMARY KEY,              -- プロフィール固有ID
  user_id       BIGINT     NOT NULL,                -- users.id を参照
  display_name  VARCHAR(50)  NULL,                  -- 表示名（ユーザー名とは別に）  
  avatar_url    TEXT         NULL,                  -- アバター画像 URL:contentReference[oaicite:4]{index=4}
  bio           TEXT         NULL,                  -- 自己紹介文
  locale        VARCHAR(10)  NULL,                  -- 言語・地域設定  
  timezone      VARCHAR(50)  NULL,                  -- タイムゾーン設定  
  created_at    TIMESTAMP    NOT NULL DEFAULT now(),
  updated_at    TIMESTAMP    NOT NULL DEFAULT now(),

  CONSTRAINT fk_profiles_user
    FOREIGN KEY(user_id) REFERENCES users(id)
    ON DELETE CASCADE
);
```

---
### thread
- `id`
	- ID
- `snippet_id`
	- 紐づくスニペットID
- `title`
	- スレッドタイトル
- `created_by`
	- スレッド作成者
- `created_at`
	- スレッド作成日
- `updated_at`
	- スレッド更新日

```sql
-- スニペットに紐づくスレッド
CREATE TABLE thread (
  id           BIGSERIAL PRIMARY KEY,         -- 一意のスレッドID
  snippet_id   BIGINT       NOT NULL,         -- 紐づくスニペットID
  title        VARCHAR(255) NOT NULL,         -- スレッドタイトル
  created_by   BIGINT       NOT NULL,         -- スレッド作成者 (user.id)
  created_at   TIMESTAMP    NOT NULL DEFAULT now(),
  updated_at   TIMESTAMP    NOT NULL DEFAULT now(),
  CONSTRAINT fk_thread_snippet
    FOREIGN KEY (snippet_id)
    REFERENCES snippet(id)
    ON DELETE CASCADE                        -- スニペット削除時にスレッドも削除
);
```

---
### comment
- `id`
	- ID
- `thread_id`
	- 紐づくスレッドID
- `parent_id`
	- ネスト時の親コメントID
- `body`
	- コメント内容
- `created_by`
	- コメント作成者
- `created_at`
	- コメント作成日
- ~~updated_at~~
	- ~~コメント更新日~~

```sql
-- スニペットに紐づくスレッド
CREATE TABLE thread (
  id           BIGSERIAL PRIMARY KEY,         -- 一意のスレッドID
  snippet_id   BIGINT       NOT NULL,         -- 紐づくスニペットID
  title        VARCHAR(255) NOT NULL,         -- スレッドタイトル
  created_by   BIGINT       NOT NULL,         -- スレッド作成者 (user.id)
  created_at   TIMESTAMP    NOT NULL DEFAULT now(),
  updated_at   TIMESTAMP    NOT NULL DEFAULT now(),
  CONSTRAINT fk_thread_snippet
    FOREIGN KEY (snippet_id)
    REFERENCES snippet(id)
    ON DELETE CASCADE                        -- スニペット削除時にスレッドも削除
);
```
 
