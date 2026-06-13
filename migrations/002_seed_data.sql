SET NAMES utf8mb4;
SET CHARACTER SET utf8mb4;

INSERT INTO departments (name, department_type, is_active) VALUES
('Центральный склад', 'warehouse', 1),
('Магазин №1', 'store', 1),
('Магазин №2', 'store', 1),
('Офис закупок', 'office', 1);

INSERT INTO suppliers (name, phone, email, is_active, address) VALUES
('ООО ТехноПоставка', '+74951234567', 'info@techno.ru', 1, 'Москва'),
('ИП Иванов', '+79169876543', 'ivanov@mail.ru', 1, 'Москва');

INSERT INTO product_units (name, short_name) VALUES
('Штука', 'шт'), ('Килограмм', 'кг'), ('Литр', 'л');

INSERT INTO product_categories (name) VALUES
('Электроника'), ('Продукты'), ('Канцелярия');

INSERT INTO products (name, sku, unit_id, category_id, is_active) VALUES
('Ноутбук Dell', 'NB-DELL-001', 1, 1, 1),
('Мышь Logitech', 'MOUSE-LOG-001', 1, 1, 1),
('Бумага A4', 'PAP-A4-500', 1, 3, 1);

INSERT INTO warehouses (address, warehouse_type, department_id, is_active) VALUES
('Москва, Складская 1', 'main', 1, 1);

INSERT INTO stores (address, store_type, department_id, is_active) VALUES
('Москва, Ленина 1', 'retail', 2, 1);INSERT INTO inventory_balances (department_id, product_id, quantity, avg_price) VALUES (1, 1, 50, 85000.00), (1, 2, 100, 8500.00), (1, 3, 30, 12000.00), (1, 6, 500, 350.00), (2, 1, 10, 85000.00), (2, 2, 20, 8500.00), (3, 3, 15, 12000.00);
INSERT INTO inventory_balances (department_id, product_id, quantity, avg_price) VALUES (1, 1, 50, 85000.00), (1, 2, 100, 8500.00), (1, 3, 30, 12000.00), (1, 6, 500, 350.00), (2, 1, 10, 85000.00), (2, 2, 20, 8500.00), (3, 3, 15, 12000.00);
