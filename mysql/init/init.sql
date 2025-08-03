CREATE TABLE IF NOT EXISTS snippet(
  id BIGINT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  body TEXT NOT NULL,
  type_id BIGINT NOT NULL,
  created_by BIGINT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now()
);

INSERT INTO snippet(id, name, body, type_id, created_at, updated_at, created_by)
VALUES
(1, 'nxkzm', 'var body: some View { codeView(code) }', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 1),
(2, 'todaryooo', 'console.log("hoge");', 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 2);



CREATE TABLE profiles (
  id BIGINT PRIMARY KEY,
  user_id BIGINT NOT NULL,
  display_name VARCHAR(50) NULL,
  avatar_url TEXT NULL,
  bio TEXT NULL,
  locale VARCHAR(10) NULL,
  timezone VARCHAR(50) NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO profiles (id, user_id, display_name, avatar_url, bio, locale, timezone, created_at, updated_at)
VALUES
(1, 1, 'kzmkzmkzm', NULL, 'test bio', NULL, NULL, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 2, 'todaryooo', NULL, 'test bio', NULL, NULL, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 3, 'test display_name hoge', NULL, 'test bio hoge', NULL, NULL, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, 4, 'test display_name fuga', NULL, 'test bio fuga', NULL, NULL, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, 5, 'test display_name piyo', NULL, 'test bio piyo', NULL, NULL, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);



CREATE TABLE comment (
  id BIGINT PRIMARY KEY,
  snippet_id INT NOT NULL,
  parent_id INT NOT NULL,
  body VARCHAR(255) NOT NULL,
  created_by BIGINT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now()
);

INSERT INTO comment (id, snippet_id, parent_id, body, created_by, created_at)
VALUES
(1, 1, 1, 'Test Comment 1', 1, CURRENT_TIMESTAMP),
(2, 2, 2, 'Test Comment 2', 2, CURRENT_TIMESTAMP),
(3, 3, 3, 'Test Comment 3', 3, CURRENT_TIMESTAMP),
(4, 4, 4, 'Test Comment 4', 4, CURRENT_TIMESTAMP);
