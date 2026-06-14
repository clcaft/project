SET NAMES utf8mb4;
SET CHARACTER SET utf8mb4;
SET SESSION collation_connection = utf8mb4_unicode_ci;

SET NAMES utf8mb4;
SET CHARACTER SET utf8mb4;
CREATE TABLE IF NOT EXISTS departments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    department_type ENUM('warehouse', 'store', 'office', 'production') NOT NULL,
    is_active TINYINT(1) DEFAULT 1
);

CREATE TABLE IF NOT EXISTS suppliers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(50),
    email VARCHAR(255),
    is_active TINYINT(1) DEFAULT 1,
    address VARCHAR(500)
);

CREATE TABLE IF NOT EXISTS product_categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS product_units (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    short_name VARCHAR(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    sku VARCHAR(100) NOT NULL UNIQUE,
    unit_id INT NOT NULL,
    category_id INT,
    is_active TINYINT(1) DEFAULT 1,
    FOREIGN KEY (unit_id) REFERENCES product_units(id),
    FOREIGN KEY (category_id) REFERENCES product_categories(id)
);

CREATE TABLE IF NOT EXISTS warehouses (
    id INT AUTO_INCREMENT PRIMARY KEY,
    address VARCHAR(500),
    warehouse_type ENUM('main', 'retail', 'distribution', 'cold_storage') NOT NULL,
    department_id INT NOT NULL,
    FOREIGN KEY (department_id) REFERENCES departments(id)
);

CREATE TABLE IF NOT EXISTS stores (
    id INT AUTO_INCREMENT PRIMARY KEY,
    address VARCHAR(500),
    store_type ENUM('retail', 'online', 'franchise', 'outlet') NOT NULL,
    department_id INT NOT NULL,
    FOREIGN KEY (department_id) REFERENCES departments(id)
);

CREATE TABLE IF NOT EXISTS supplier_products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    supplier_id INT NOT NULL,
    product_id INT NOT NULL,
    FOREIGN KEY (supplier_id) REFERENCES suppliers(id),
    FOREIGN KEY (product_id) REFERENCES products(id),
    UNIQUE KEY uk_supplier_product (supplier_id, product_id)
);

CREATE TABLE IF NOT EXISTS purchase_requests (
    id INT AUTO_INCREMENT PRIMARY KEY,
    request_date DATE NOT NULL,
    planned_delivery_date DATE,
    supplier_id INT,
    status ENUM('draft', 'submitted', 'approved', 'rejected', 'cancelled') DEFAULT 'draft',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    cancelled_at DATETIME,
    recipient_department_id INT NOT NULL,
    FOREIGN KEY (supplier_id) REFERENCES suppliers(id),
    FOREIGN KEY (recipient_department_id) REFERENCES departments(id)
);

CREATE TABLE IF NOT EXISTS purchase_request_items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    purchase_request_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity DECIMAL(15,3) NOT NULL,
    FOREIGN KEY (purchase_request_id) REFERENCES purchase_requests(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE IF NOT EXISTS delivery_invoices (
    id INT AUTO_INCREMENT PRIMARY KEY,
    invoice_date DATE NOT NULL,
    supplier_id INT NOT NULL,
    purchase_request_id INT,
    status ENUM('draft', 'confirmed', 'received', 'cancelled') DEFAULT 'draft',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    cancelled_at DATETIME,
    recipient_department_id INT NOT NULL,
    FOREIGN KEY (supplier_id) REFERENCES suppliers(id),
    FOREIGN KEY (purchase_request_id) REFERENCES purchase_requests(id),
    FOREIGN KEY (recipient_department_id) REFERENCES departments(id)
);

CREATE TABLE IF NOT EXISTS delivery_invoice_items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    delivery_invoice_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity DECIMAL(15,3) NOT NULL,
    price DECIMAL(15,2) NOT NULL,
    FOREIGN KEY (delivery_invoice_id) REFERENCES delivery_invoices(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE IF NOT EXISTS internal_transfer_invoices (
    id INT AUTO_INCREMENT PRIMARY KEY,
    transfer_date DATE NOT NULL,
    from_department_id INT NOT NULL,
    to_department_id INT NOT NULL,
    status ENUM('draft', 'confirmed', 'shipped', 'received', 'cancelled') DEFAULT 'draft',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    cancelled_at DATETIME,
    FOREIGN KEY (from_department_id) REFERENCES departments(id),
    FOREIGN KEY (to_department_id) REFERENCES departments(id),
    CHECK (from_department_id != to_department_id)
);

CREATE TABLE IF NOT EXISTS internal_transfer_invoice_items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    internal_transfer_invoice_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity DECIMAL(15,3) NOT NULL,
    FOREIGN KEY (internal_transfer_invoice_id) REFERENCES internal_transfer_invoices(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE IF NOT EXISTS inventory_balances (
    id INT AUTO_INCREMENT PRIMARY KEY,
    department_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity DECIMAL(15,3) NOT NULL DEFAULT 0,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY uk_inventory (department_id, product_id),
    FOREIGN KEY (department_id) REFERENCES departments(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);