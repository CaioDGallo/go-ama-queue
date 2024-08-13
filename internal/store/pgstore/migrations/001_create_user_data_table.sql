-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS user_data (
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    ip varchar(255) NOT NULL,
    user_agent varchar(255) NOT NULL,
    timestamp timestamptz NOT NULL DEFAULT now(),
    location varchar(255) NOT NULL, 
    device varchar(255) NOT NULL,
    action varchar(255) NOT NULL,
    json_response_body varchar(255) NOT NULL,
    referrer varchar(255) NOT NULL,
    request_method varchar(255) NOT NULL,
    request_path varchar(255) NOT NULL
);

---- create above / drop below ----

DROP TABLE IF EXISTS user_data;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.

