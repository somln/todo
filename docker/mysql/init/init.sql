CREATE TABLE IF NOT EXISTS Todo (
                                    todo_id INT AUTO_INCREMENT PRIMARY KEY,
                                    content TEXT NOT NULL,
                                    status ENUM('progress', 'completed') DEFAULT 'progress',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );