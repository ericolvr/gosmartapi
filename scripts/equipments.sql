CREATE TABLE equipments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    identifier VARCHAR(100) NOT NULL,
    uniorg VARCHAR(10) NOT NULL,
    code VARCHAR(10) NOT NULL,
    used_code BOOLEAN DEFAULT FALSE,
    environment VARCHAR(100) NOT NULL,
    system_name VARCHAR(100) NOT NULL, 
    description TEXT NULL,
    completed BOOLEAN DEFAULT FALSE,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);