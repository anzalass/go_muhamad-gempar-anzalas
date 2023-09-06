# Buat database alta_online_shop

CREATE DATABASE alta_online_shop;

# Masuk ke database alta_online_shop

USE alta_online_shop;

# Buat table user

CREATE TABLE user (
user_id INT NOT NULL AUTO_INCREMENT,
name VARCHAR(255) NOT NULL,
address VARCHAR(255) NOT NULL,
date_of_birth DATE NOT NULL,
status_user VARCHAR(255) NOT NULL,
gender VARCHAR(255) NOT NULL,
created_at DATETIME NOT NULL,
updated_at DATETIME NOT NULL,
PRIMARY KEY (user_id)
);

# Buat table product

CREATE TABLE product (
product_id INT NOT NULL AUTO_INCREMENT,
product_name VARCHAR(255) NOT NULL,
product_type VARCHAR(255) NOT NULL,
product_description TEXT NOT NULL,
operator VARCHAR(255) NOT NULL,
payment_methods VARCHAR(255) NOT NULL,
PRIMARY KEY (product_id)
);

# Buat table transaction

CREATE TABLE transaction (
transaction_id INT NOT NULL AUTO_INCREMENT,
customer_id INT NOT NULL,
product_id INT NOT NULL,
transaction_date DATETIME NOT NULL,
transaction_total INT NOT NULL,
transaction_status VARCHAR(255) NOT NULL,
PRIMARY KEY (transaction_id),
FOREIGN KEY (customer_id) REFERENCES user (user_id) ON DELETE CASCADE,
FOREIGN KEY (product_id) REFERENCES product (product_id) ON DELETE CASCADE
);

# Buat table transaction_detail

CREATE TABLE transaction_detail (
transaction_detail_id INT NOT NULL AUTO_INCREMENT,
transaction_id INT NOT NULL,
product_id INT NOT NULL,
quantity INT NOT NULL,
price INT NOT NULL,
PRIMARY KEY (transaction_detail_id),
FOREIGN KEY (transaction_id) REFERENCES transaction (transaction_id) ON DELETE CASCADE,
FOREIGN KEY (product_id) REFERENCES product (product_id) ON DELETE CASCADE
);

# Buat table kurir

CREATE TABLE kurir (
id INT NOT NULL AUTO_INCREMENT,
name VARCHAR(255) NOT NULL,
created_at DATETIME NOT NULL,
updated_at DATETIME NOT NULL,
PRIMARY KEY (id)
);

# Tambahkan kolom ongkos_dasar pada tabel kurir

ALTER TABLE kurir ADD ongkos_dasar INT NOT NULL;

# Ganti nama tabel kurir menjadi shipping

RENAME TABLE kurir TO shipping;

# Hapus tabel shipping

DROP TABLE shipping;

# Buat table payment_method_description

CREATE TABLE payment_method_description (
payment_method_id INT NOT NULL,
description VARCHAR(255) NOT NULL,
PRIMARY KEY (payment_method_id),
FOREIGN KEY (payment_method_id) REFERENCES product (product_id) ON DELETE CASCADE
);

# Buat table alamat

CREATE TABLE alamat (
alamat_id INT NOT NULL AUTO_INCREMENT,
user_id INT NOT NULL,
provinsi VARCHAR(255) NOT NULL,
kota VARCHAR(255) NOT NULL,
kecamatan VARCHAR(255) NOT NULL,
kelurahan VARCHAR(255) NOT NULL,
kode_pos VARCHAR(255) NOT NULL,
alamat_lengkap VARCHAR(255) NOT NULL,
PRIMARY KEY (alamat_id),
FOREIGN KEY (user_id) REFERENCES user (user_id) ON DELETE CASCADE
);

# Buat table user_payment_method_detail

CREATE TABLE user_payment_method_detail (
user_id INT NOT NULL,
payment_method_id INT NOT NULL,
PRIMARY KEY (user_id, payment_method_id),
FOREIGN KEY (user_id) REFERENCES user (user_id) ON DELETE CASCADE,
FOREIGN KEY (payment_method_id) REFERENCES payment_method (payment_method_id) ON DELETE CASCADE
);
