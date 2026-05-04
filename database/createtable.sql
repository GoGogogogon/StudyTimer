

CREATE TABLE IF NOT EXISTS study_logs (
        ID          INT AUTO_INCREMENT PRIMARY KEY, 
        subject     VARCHAR(255) NOT NULL,
        time  INTEGER NOT NULL,
        created_at  DATETIME DEFAULT CURRENT_TIMESTAMP
    );