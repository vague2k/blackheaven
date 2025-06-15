CREATE TABLE inquiries (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    topic TEXT NOT NULL,
    email TEXT NOT NULL,
    name TEXT,
    order_number TEXT,
    subject TEXT,
    content TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

