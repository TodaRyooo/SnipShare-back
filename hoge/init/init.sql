CREATETABLEIFNOTEXISTSsnippet(
  id BIGINT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  body TEXT NOT NULL,
  type_id BIGINT NOT NULL,
  created_by BIGINT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now()
);
INSERT INTO snippet(
id,
name,
body,
type_id,
created_at,
updated_at,
created_by
)VALUES(
1,
'nxkzm',
'var body: some View { codeView(code) }',
1,
now(

),
now(

),
1
);
