# NO.1

# A

INSERT INTO operators (id, nama) VALUES (1, 'udin');
INSERT INTO operators (id, nama) VALUES (2, 'ucup');
INSERT INTO operators (id, nama) VALUES (3, 'ucok');
INSERT INTO operators (id, nama) VALUES (4, 'adil');
INSERT INTO operators (id, nama) VALUES (5, 'dani');

### B

INSERT INTO product_types (id, name) VALUES (1,'sports');
INSERT INTO product_types (id, name) VALUES (2,'otomotif');
INSERT INTO product_types (id, name) VALUES (3,'kids');

### C

INSERT INTO products (id, product_type_id, operator_id, code, name, status) VALUES (1,1,3, 'GRTS1', 'Sepatu Bola Adidas' , 0);
INSERT INTO products (id, product_type_id, operator_id, code, name, status) VALUES (2,1,3, 'GRTS2', 'T-Shirts Manchester United 2019' , 0);

### D

INSERT INTO products (id, product_type_id, operator_id, code, name, status) VALUES (3,2,1, 'GRTS3', 'Motor NMAX 2019' , 0);
INSERT INTO products (id, product_type_id, operator_id, code, name, status) VALUES (4,2,1, 'GRTS4', 'Oli AHM Matic' , 0);
INSERT INTO products (id, product_type_id, operator_id, code, name, status) VALUES (5,2,1, 'GRTS5', 'Velg Original Nissan' , 0);

### E

INSERT INTO products (id, product_type_id, operator_id, code, name, status) VALUES (6,3,4, 'GRTS3', 'Rc Mobil F1 ' , 0);
INSERT INTO products (id, product_type_id, operator_id, code, name, status) VALUES (7,3,4, 'GRTS4', 'Hot Wheels BMW seri 4' , 0);
INSERT INTO products (id, product_type_id, operator_id, code, name, status) VALUES (8,3,4, 'GRTS5', 'Layangan Peteng' , 0);

### F

INSERT INTO product_descriptions (id, description) VALUES (1, 'Ini Sepatu Adidas BAGUS');
INSERT INTO product_descriptions (id, description) VALUES (2, 'Ini tshirts Manchester United ');
INSERT INTO product_descriptions (id, description) VALUES (3, 'Ini Motor NMAX 2019');
INSERT INTO product_descriptions (id, description) VALUES (4, 'Ini Oli AHM Bagus');
INSERT INTO product_descriptions (id, description) VALUES (5, 'Ini Velg Nissan Bagus');
INSERT INTO product_descriptions (id, description) VALUES (6, 'Ini Rc Mobil F1');
INSERT INTO product_descriptions (id, description) VALUES (7, 'Ini Hot Wheels BMW seri 4');
INSERT INTO product_descriptions (id, description) VALUES (8, 'Ini Layangan Peteng');

### G

INSERT INTO payment_methods (id, name, status) VALUES (1, 'COD', 0);
INSERT INTO payment_methods (id, name, status) VALUES (2, 'Gerai', 0);
INSERT INTO payment_methods (id, name, status) VALUES (3, 'Transfer', 0);

### H

INSERT INTO users (id, status, dob, gender) VALUES (1, 0, '27-08-2003', 'L');
INSERT INTO users (id, status, dob, gender) VALUES (2, 0, '26-08-2003', 'P');
INSERT INTO users (id, status, dob, gender) VALUES (3, 0, '25-08-2003', 'L');
INSERT INTO users (id, status, dob, gender) VALUES (4, 0, '24-08-2003', 'P');
INSERT INTO users (id, status, dob, gender) VALUES (5, 0, '23-08-2003', 'L');

### I

#### 1

INSERT INTO transactions (id, user_id, payment_method_id, status, total_qty, total_price) VALUES (1, 1, 1, 0, 5, 50000);
INSERT INTO transactions (id, user_id, payment_method_id, status, total_qty, total_price) VALUES (2, 1, 1, 0, 3, 50000);
INSERT INTO transactions (id, user_id, payment_method_id, status, total_qty, total_price) VALUES (3, 1, 1, 0, 7, 50000);

#### 2

INSERT INTO transactions (id, user_id, payment_method_id, status, total_qty, total_price) VALUES (4, 2, 2, 0, 5, 50000);
INSERT INTO transactions (id, user_id, payment_method_id, status, total_qty, total_price) VALUES (5, 2, 2, 0, 3, 50000);
INSERT INTO transactions (id, user_id, payment_method_id, status, total_qty, total_price) VALUES (6, 2, 2, 0, 7, 50000);

#### 3

INSERT INTO transactions (id, user_id, payment_method_id, status, total_qty, total_price) VALUES (7, 3, 3, 0, 5, 50000);
INSERT INTO transactions (id, user_id, payment_method_id, status, total_qty, total_price) VALUES (8, 3, 3, 0, 3, 50000);
INSERT INTO transactions (id, user_id, payment_method_id, status, total_qty, total_price) VALUES (9, 3, 3, 0, 7, 50000);

#### 4

INSERT INTO transactions (id, user_id, payment_method_id, status, total_qty, total_price) VALUES (10, 4,3,0,5,50000);
INSERT INTO transactions (id, user_id, payment_method_id, status, total_qty, total_price) VALUES (11, 4,3, 0, 5, 50000);
INSERT INTO transactions (id, user_id, payment_method_id, status, total_qty, total_price) VALUES (12, 4,3,0,5,60000);

#### 5

INSERT INTO transactions (id, user_id, payment_method_id, status, total_qty, total_price) VALUES (13 5,5, 50000);
INSERT INTO transactions (id, user_id, payment_method_id, status, total_qty, total_price) VALUES (14 5,0 3, 50000);
INSERT INTO transactions (id, user_id, payment_method_id, status, total_qty, total_price) VALUES (15 5, 7, 50000);

# NO.2

### A

SELECT nama FROM `users` WHERE gender='P';

### B

SELECT \* FROM products WHERE id=3;

### C

SELECT \* FROM users WHERE created_at > CURRENT_DATE - INTERVAL 7 DAY AND nama LIKE '%a%';

### D

SELECT COUNT(\*) AS total_users FROM users WHERE gender='P';

### E

SELECT \* FROM users ORDER BY nama ASC ;

### F

SELECT \* FROM products LIMIT 5 ;

# NO.3

### A

UPDATE products SET name='product dummy' WHERE id=1;

### B

UPDATE transactions_details SET qty=3 WHERE product_id=1;

# NO.4

### A

DELETE FROM products WHERE id=1;

### B

DELETE FROM products WHERE product_type_id=1;

## Join, Union, Sub query, Function

### 1

SELECT transactions.id, transactions.created_at, transactions.user_id FROM transactions JOIN users ON transactions.user_id = users.id WHERE transactions.id IN (1,2);

### 2

SELECT price\*qty FROM transactions_details WHERE transaction_id=1;

### 3

SELECT price\*qty FROM transactions_details WHERE product_type=2

### 4

SELECT \*, product_types.name FROM products INNER JOIN product_types ON products.id = product_types.id;

### 5

SELECT \*, products.name FROM transactions_details INNER JOIN products ON products.id = transactions_details.product_id
