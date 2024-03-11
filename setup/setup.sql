-- Create or use the specified database
CREATE DATABASE IF NOT EXISTS messaging_system;
USE messaging_system;

-- Create the 'messages' table
CREATE TABLE IF NOT EXISTS messages (
    id INT AUTO_INCREMENT PRIMARY KEY,
    content VARCHAR(255) NOT NULL,
    recipient_phone VARCHAR(20) NOT NULL,
    sent_status BOOL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Insert sample data into the 'messages' table
INSERT INTO messages (content, recipient_phone, sent_status)
VALUES
    ('Insider - Project', '+905551111111', 0),
    ('Hello, this is a test message', '+905552222222', 0),
    ('Another test message', '+905553333333', 0),
    ('Sent message', '+905554444444', 1),
    ('Yet another test message', '+905555555555', 0),
    ('This is the last test message', '+905556666666', 0);
