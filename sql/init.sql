-- Create the table 'links'
CREATE TABLE links (
    hash VARCHAR(10) NOT NULL PRIMARY KEY, -- Hash column with a maximum length of 10 characters
    -- id INTEGER PRIMARY KEY AUTOINCREMENT, -- Autoincrementing integer id
    original TEXT NOT NULL, -- Original URL
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp for creation time
    last_modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Timestamp for last modified time
);

-- Create a trigger to update 'last_modified' timestamp on any row update
CREATE TRIGGER update_last_modified
AFTER UPDATE ON links
FOR EACH ROW
BEGIN
    UPDATE links
    SET last_modified = CURRENT_TIMESTAMP
    WHERE id = OLD.id;
END;

-- Create a composite index on 'id' and 'hash' columns
-- CREATE INDEX idx_links_id_hash ON links(id, hash); 