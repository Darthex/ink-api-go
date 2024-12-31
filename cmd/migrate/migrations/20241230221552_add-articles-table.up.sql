CREATE TABLE articles (
    id SERIAL PRIMARY KEY, -- Auto-incrementing primary key
    owner_id VARCHAR NOT NULL, -- Owner ID, indexed
    owner_name VARCHAR NOT NULL, -- Owner name, indexed
    title VARCHAR NOT NULL, -- Article title
    content TEXT NOT NULL, -- Article content
    description TEXT, -- Optional description
    cover VARCHAR, -- Optional cover image URL
    tags TEXT[] DEFAULT ARRAY[]::TEXT[], -- Array of tags, defaults to empty array
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for owner_id and owner_name
CREATE INDEX idx_owner_id ON articles (owner_id);
CREATE INDEX idx_owner_name ON articles (owner_name);