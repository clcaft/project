SELECT 
    ib.department_id,
    d.name as department_name,
    d.department_type,
    p.name as product_name,
    p.sku,
    ib.quantity
FROM inventory_balances ib
JOIN departments d ON ib.department_id = d.id
JOIN products p ON ib.product_id = p.id
WHERE ib.product_id = 1  -- заменить на нужный ID
ORDER BY d.department_type;

SELECT 
    p.name as product_name,
    p.sku,
    pu.short_name as unit,
    ib.quantity,
    ib.updated_at
FROM inventory_balances ib
JOIN products p ON ib.product_id = p.id
JOIN product_units pu ON p.unit_id = pu.id
WHERE ib.department_id = 1  -- заменить на нужный ID
ORDER BY p.name;

SELECT 
    'Приход' as operation_type,
    di.invoice_date as date,
    s.name as counterparty,
    dii.quantity,
    dii.price,
    dii.quantity * dii.price as amount,
    d.name as department_name
FROM delivery_invoices di
JOIN delivery_invoice_items dii ON di.id = dii.delivery_invoice_id
JOIN suppliers s ON di.supplier_id = s.id
JOIN departments d ON di.recipient_department_id = d.id
WHERE dii.product_id = 1
  AND di.status = 'confirmed'
  AND di.invoice_date BETWEEN '2026-06-01' AND '2026-06-30'

UNION ALL

SELECT 
    'Отгрузка' as operation_type,
    iti.transfer_date as date,
    d_from.name as counterparty,
    itii.quantity,
    0 as price,
    0 as amount,
    d_to.name as department_name
FROM internal_transfer_invoices iti
JOIN internal_transfer_invoice_items itii ON iti.id = itii.internal_transfer_invoice_id
JOIN departments d_from ON iti.from_department_id = d_from.id
JOIN departments d_to ON iti.to_department_id = d_to.id
WHERE itii.product_id = 1
  AND iti.status = 'shipped'
  AND iti.transfer_date BETWEEN '2026-06-01' AND '2026-06-30'

ORDER BY date;

SELECT 
    p.name as product_name,
    p.sku,
    SUM(dii.quantity) as total_quantity,
    SUM(dii.quantity * dii.price) as total_amount,
    COUNT(DISTINCT di.id) as invoices_count
FROM delivery_invoices di
JOIN delivery_invoice_items dii ON di.id = dii.delivery_invoice_id
JOIN products p ON dii.product_id = p.id
WHERE di.supplier_id = 1  -- заменить на нужный ID
  AND di.status = 'confirmed'
  AND di.invoice_date BETWEEN '2026-06-01' AND '2026-06-30'
GROUP BY p.id, p.name, p.sku
ORDER BY total_quantity DESC;

SELECT 
    p.name as product_name,
    p.sku,
    SUM(itii.quantity) as total_quantity,
    COUNT(DISTINCT iti.id) as transfers_count
FROM internal_transfer_invoices iti
JOIN internal_transfer_invoice_items itii ON iti.id = itii.internal_transfer_invoice_id
JOIN products p ON itii.product_id = p.id
WHERE iti.to_department_id = 2  -- заменить на нужный ID магазина
  AND iti.status = 'shipped'
  AND iti.transfer_date BETWEEN '2026-06-01' AND '2026-06-30'
GROUP BY p.id, p.name, p.sku
ORDER BY total_quantity DESC;

SELECT 
    p.id,
    p.name,
    p.sku,
    pu.short_name as unit,
    pc.name as category,
    s.name as supplier_name,
    s.phone as supplier_phone
FROM products p
JOIN product_units pu ON p.unit_id = pu.id
LEFT JOIN product_categories pc ON p.category_id = pc.id
JOIN supplier_products sp ON p.id = sp.product_id
JOIN suppliers s ON sp.supplier_id = s.id
WHERE p.is_active = TRUE
  AND s.is_active = TRUE
ORDER BY p.name, s.name;

SELECT 
    pr.id,
    pr.request_date,
    pr.status,
    s.name as supplier_name,
    d.name as recipient_department,
    pri.product_id,
    p.name as product_name,
    pri.quantity as requested_qty,
    COALESCE(SUM(dii.quantity), 0) as delivered_qty,
    pri.quantity - COALESCE(SUM(dii.quantity), 0) as remaining_qty
FROM purchase_requests pr
JOIN purchase_request_items pri ON pr.id = pri.purchase_request_id
JOIN products p ON pri.product_id = p.id
LEFT JOIN suppliers s ON pr.supplier_id = s.id
JOIN departments d ON pr.recipient_department_id = d.id
LEFT JOIN delivery_invoices di ON di.purchase_request_id = pr.id AND di.status = 'confirmed'
LEFT JOIN delivery_invoice_items dii ON di.id = dii.delivery_invoice_id AND dii.product_id = pri.product_id
WHERE pr.status IN ('approved', 'submitted')
GROUP BY pr.id, pri.id, pr.request_date, pr.status, s.name, d.name, p.name, pri.quantity
HAVING remaining_qty > 0
ORDER BY pr.request_date;