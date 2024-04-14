CREATE TABLE banners (
                         id SERIAL PRIMARY KEY,
                         feature_id INT NOT NULL,
                         content JSONB NOT NULL,
                         is_active BOOLEAN NOT NULL DEFAULT TRUE,
                         created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tags (
                      id SERIAL PRIMARY KEY,
                      name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE features (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE banner_tags (
                             banner_id INT NOT NULL REFERENCES banners(id) ON DELETE CASCADE,
                             tag_id INT NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
                             PRIMARY KEY (banner_id, tag_id)
);