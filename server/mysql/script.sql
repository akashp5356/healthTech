
CREATE DATABASE IF NOT EXISTS healthtech;
USE healthtech;

-- roleMaster
CREATE TABLE IF NOT EXISTS roleMaster (
    id INT AUTO_INCREMENT PRIMARY KEY,
    role_name VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO roleMaster (role_name) VALUES ('admin'), ('user');

-- registerDetails
CREATE TABLE IF NOT EXISTS registerDetails (
    id INT AUTO_INCREMENT PRIMARY KEY,
    full_name VARCHAR(100) NOT NULL,
    email VARCHAR(120) UNIQUE NOT NULL,
    phone VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- loginMaster
CREATE TABLE IF NOT EXISTS loginMaster (
    id INT AUTO_INCREMENT PRIMARY KEY,
    register_id INT NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (register_id) REFERENCES registerDetails(id),
    FOREIGN KEY (role_id) REFERENCES roleMaster(id)
);

-- documentMaster
CREATE TABLE IF NOT EXISTS documentMaster (
    id INT AUTO_INCREMENT PRIMARY KEY,
    document_type VARCHAR(100) NOT NULL, -- prescription/report
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO documentMaster (document_type) VALUES 
('prescription'),
('diagnostic_report');

-- documentDetails
CREATE TABLE IF NOT EXISTS documentDetails (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    document_id INT NOT NULL,
    file_path VARCHAR(255) NOT NULL,
    original_filename VARCHAR(255) NOT NULL,
    file_size BIGINT NOT NULL,
    description TEXT,
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES registerDetails(id),
    FOREIGN KEY (document_id) REFERENCES documentMaster(id)
);

-- test user
INSERT INTO registerDetails (full_name, email, phone) VALUES ('Test Patient', 'patient@example.com', '1234567890');