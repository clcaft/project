SET NAMES utf8mb4;
SET CHARACTER SET utf8mb4;

INSERT INTO departments (name, department_type, is_active) VALUES
('Центральный склад', 'warehouse', 1),
('Магазин №1 (Ленина)', 'store', 1),
('Магазин №2 (Пушкина)', 'store', 1),
('Офис закупок', 'office', 1),
('Магазин №3 (Онлайн)', 'store', 1);

INSERT INTO suppliers (name, phone, email, is_active, address) VALUES
('ООО ТехноПоставка', '+74951234567', 'info@techno.ru', 1, 'Москва, ул. Техническая 10'),
('ИП Иванов А.В.', '+79169876543', 'ivanov@mail.ru', 1, 'Москва, ул. Ивановская 5'),
('ООО ПродуктОпт', '+74959876543', 'opt@products.ru', 1, 'Москва, ул. Складская 3'),
('ООО БюроСнаб', '+74951112233', 'snab@office.ru', 0, 'Москва, ул. Офисная 1');

INSERT INTO product_units (name, short_name) VALUES
('Штука', 'шт'),
('Килограмм', 'кг'),
('Литр', 'л'),
('Метр', 'м'),
('Упаковка', 'упак');

INSERT INTO product_categories (name) VALUES
('Электроника'),
('Продукты питания'),
('Канцелярия'),
('Бытовая химия'),
('Офисная техника');

INSERT INTO products (name, sku, unit_id, category_id, is_active) VALUES
('Ноутбук Dell XPS 15', 'NB-DELL-001', 1, 1, 1),
('Мышь Logitech MX Master', 'MOUSE-LOG-001', 1, 1, 1),
('Клавиатура механическая', 'KEYB-MECH-001', 1, 1, 1),
('Бумага A4 500л', 'PAP-A4-500', 5, 3, 1),
('Ручка шариковая синяя', 'PEN-BLUE-001', 1, 3, 1),
('Молоко 3.2% 1л', 'MILK-3.2-1L', 3, 2, 1),
('Хлеб белый', 'BREAD-WHT-001', 1, 2, 1),
('Моющее средство 5л', 'CLEAN-5L-001', 3, 4, 1),
('Принтер лазерный HP', 'PRINT-HP-001', 1, 5, 1),
('Монитор 27" Samsung', 'MON-SAM-001', 1, 1, 0);  -- неактивный

INSERT INTO warehouses (address, warehouse_type, department_id) VALUES
('Москва, Складская 1', 'main', 1),
('Москва, Складская 2', 'retail', 1);

INSERT INTO stores (address, store_type, department_id) VALUES
('Москва, Ленина 1', 'retail', 2),
('Москва, Пушкина 10', 'retail', 3),
('Интернет-магазин', 'online', 5);

INSERT INTO supplier_products (supplier_id, product_id) VALUES
(1, 1), (1, 2), (1, 3), (1, 9), 
(2, 4), (2, 5),                  
(3, 6), (3, 7),                   
(4, 4), (4, 5);                   

INSERT INTO inventory_balances (department_id, product_id, quantity) VALUES
(1, 1, 50),   
(1, 2, 100),  
(1, 3, 30),   
(1, 4, 200),  
(1, 5, 500),  
(1, 6, 100),  
(1, 7, 50),   
(1, 8, 20),   
(1, 9, 10),   

(2, 1, 10),
(2, 2, 20),
(2, 4, 50),
(2, 6, 30),

(3, 3, 5),
(3, 5, 100),
(3, 7, 20),
(3, 8, 10),

(5, 1, 5),
(5, 2, 15),
(5, 9, 3);


INSERT INTO purchase_requests (request_date, planned_delivery_date, supplier_id, status, recipient_department_id) VALUES
('2026-06-01', '2026-06-05', 1, 'approved', 1),   
('2026-06-03', '2026-06-08', 2, 'submitted', 1),   
('2026-06-05', '2026-06-10', 3, 'approved', 1),   
('2026-06-10', '2026-06-15', 1, 'draft', 1),      
('2026-06-12', '2026-06-18', 2, 'approved', 2);   

INSERT INTO purchase_request_items (purchase_request_id, product_id, quantity) VALUES
(1, 1, 20),
(1, 2, 50),
(1, 9, 5),

(2, 4, 100),
(2, 5, 200),

(3, 6, 50),
(3, 7, 30),

(4, 3, 10),

(5, 4, 30),
(5, 5, 50);

INSERT INTO delivery_invoices (invoice_date, supplier_id, purchase_request_id, status, recipient_department_id) VALUES
('2026-06-05', 1, 1, 'confirmed', 1),
('2026-06-08', 2, 2, 'confirmed', 1),
('2026-06-10', 3, 3, 'draft', 1),
('2026-06-15', 1, NULL, 'confirmed', 1),
('2026-06-20', 1, NULL, 'draft', 2);

INSERT INTO delivery_invoice_items (delivery_invoice_id, product_id, quantity, price) VALUES

(1, 1, 20, 85000.00),
(1, 2, 50, 3500.00),
(1, 9, 5, 25000.00),

(2, 4, 80, 250.00),

(3, 6, 50, 65.00),
(3, 7, 30, 35.00),

(4, 3, 15, 5000.00),
(4, 8, 10, 450.00),

(5, 4, 30, 250.00),
(5, 5, 50, 15.00);

INSERT INTO internal_transfer_invoices (transfer_date, from_department_id, to_department_id, status) VALUES
('2026-06-07', 1, 2, 'shipped'),
('2026-06-09', 1, 3, 'shipped'),
('2026-06-11', 1, 5, 'shipped'),
('2026-06-14', 1, 2, 'draft'),
('2026-06-16', 2, 3, 'draft');

INSERT INTO internal_transfer_invoice_items (internal_transfer_invoice_id, product_id, quantity) VALUES
(1, 1, 5),
(1, 2, 10),
(1, 4, 20),

(2, 3, 3),
(2, 5, 50),
(2, 7, 15),
(2, 8, 5),

(3, 1, 3),
(3, 9, 2),

(4, 6, 20),
(4, 7, 10),

(5, 2, 5),
(5, 4, 10);

SELECT 'Seed data inserted successfully!' AS Status;