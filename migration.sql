CREATE TABLE IF NOT EXISTS products (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(255) NOT NULL,
  `description` TEXT NOT NULL,
  `seller` VARCHAR(255) NOT NULL,
  `rating` DECIMAL(2, 1) NOT NULL CHECK (rating >= 1.0 AND rating <= 5.0),
  `createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

INSERT INTO products (title, description, seller, rating) VALUES
('Wireless Headphones', 'High-quality wireless headphones with noise cancellation.', 'AudioTech', 5),
('Smartphone X', 'Latest smartphone with advanced features and sleek design.', 'TechWorld', 4),
('Gaming Mouse', 'Ergonomic gaming mouse with customizable buttons and RGB lighting.', 'GamerZone', 5),
('Bluetooth Speaker', 'Portable Bluetooth speaker with excellent sound quality.', 'SoundWave', 4),
('Laptop Stand', 'Adjustable laptop stand for better ergonomics and airflow.', 'OfficeEssentials', 3),
('Smartwatch', 'Feature-rich smartwatch with fitness tracking capabilities.', 'WearableTech', 4),
('4K Monitor', 'Ultra HD monitor with vibrant colors and high refresh rate.', 'DisplayMasters', 5),
('USB-C Hub', 'Multi-port USB-C hub for connecting various devices.', 'TechAccessories', 3),
('Fitness Tracker', 'Accurate fitness tracker with heart rate monitoring.', 'HealthGear', 4),
('Portable Charger', 'High-capacity portable charger for on-the-go charging.', 'PowerUp', 5);
