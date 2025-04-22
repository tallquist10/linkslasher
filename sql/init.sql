-- Create the table 'links'
CREATE TABLE links (
    hash VARCHAR(10) NOT NULL PRIMARY KEY, -- Hash column with a maximum length of 10 characters
    original TEXT NOT NULL, -- Original URL
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp for creation time
    last_modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Timestamp for last modified time
);