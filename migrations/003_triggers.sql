<<<<<<< HEAD
USE procurement_db;

SET GLOBAL log_bin_trust_function_creators = 1;

DELIMITER //

DROP TRIGGER IF EXISTS trg_delivery_invoice_confirmed//
DROP TRIGGER IF EXISTS trg_transfer_invoice_shipped//
DROP TRIGGER IF EXISTS trg_transfer_invoice_shipped_after//

=======

DELIMITER //
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
CREATE TRIGGER trg_delivery_invoice_confirmed
AFTER UPDATE ON delivery_invoices
FOR EACH ROW
BEGIN
    IF OLD.status != 'confirmed' AND NEW.status = 'confirmed' THEN
<<<<<<< HEAD
        INSERT INTO inventory_balances (department_id, product_id, quantity)
        SELECT 
            NEW.recipient_department_id,
            dii.product_id,
            dii.quantity
        FROM delivery_invoice_items dii
        WHERE dii.delivery_invoice_id = NEW.id
        ON DUPLICATE KEY UPDATE
=======
        INSERT INTO inventory_balances (department_id, product_id, quantity, avg_price)
        SELECT 
            NEW.recipient_department_id,
            dii.product_id,
            dii.quantity,
            dii.price
        FROM delivery_invoice_items dii
        WHERE dii.delivery_invoice_id = NEW.id
        ON DUPLICATE KEY UPDATE
            avg_price = ((inventory_balances.quantity * inventory_balances.avg_price) + (dii.quantity * dii.price)) 
                        / (inventory_balances.quantity + dii.quantity),
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
            quantity = inventory_balances.quantity + dii.quantity;
    END IF;
END//

<<<<<<< HEAD
=======
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
CREATE TRIGGER trg_transfer_invoice_shipped
BEFORE UPDATE ON internal_transfer_invoices
FOR EACH ROW
BEGIN
    DECLARE v_insufficient INT DEFAULT 0;
    
    IF OLD.status != 'shipped' AND NEW.status = 'shipped' THEN
        SELECT COUNT(*) INTO v_insufficient
        FROM internal_transfer_invoice_items itii
        LEFT JOIN inventory_balances ib ON ib.product_id = itii.product_id 
            AND ib.department_id = NEW.from_department_id
        WHERE itii.internal_transfer_invoice_id = NEW.id
          AND (ib.quantity IS NULL OR ib.quantity < itii.quantity);
        
        IF v_insufficient > 0 THEN
            SIGNAL SQLSTATE '45000'
            SET MESSAGE_TEXT = 'Insufficient inventory balance for transfer';
        END IF;
    END IF;
END//

<<<<<<< HEAD
-- Триггер: при отгрузке — списать с отправителя, приходовать получателю
=======
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
CREATE TRIGGER trg_transfer_invoice_shipped_after
AFTER UPDATE ON internal_transfer_invoices
FOR EACH ROW
BEGIN
    IF OLD.status != 'shipped' AND NEW.status = 'shipped' THEN
        -- Списываем с отправителя
        UPDATE inventory_balances ib
        JOIN internal_transfer_invoice_items itii ON ib.product_id = itii.product_id
        SET ib.quantity = ib.quantity - itii.quantity
        WHERE itii.internal_transfer_invoice_id = NEW.id
          AND ib.department_id = NEW.from_department_id;
        
<<<<<<< HEAD
        -- Приходуем получателю
        INSERT INTO inventory_balances (department_id, product_id, quantity)
        SELECT 
            NEW.to_department_id,
            itii.product_id,
            itii.quantity
        FROM internal_transfer_invoice_items itii
        WHERE itii.internal_transfer_invoice_id = NEW.id
        ON DUPLICATE KEY UPDATE
=======
        INSERT INTO inventory_balances (department_id, product_id, quantity, avg_price)
        SELECT 
            NEW.to_department_id,
            itii.product_id,
            itii.quantity,
            COALESCE(ib.avg_price, 0)
        FROM internal_transfer_invoice_items itii
        LEFT JOIN inventory_balances ib ON ib.product_id = itii.product_id 
            AND ib.department_id = NEW.from_department_id
        WHERE itii.internal_transfer_invoice_id = NEW.id
        ON DUPLICATE KEY UPDATE
            avg_price = ((inventory_balances.quantity * inventory_balances.avg_price) + (itii.quantity * COALESCE(ib.avg_price, 0))) 
                        / (inventory_balances.quantity + itii.quantity),
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
            quantity = inventory_balances.quantity + itii.quantity;
    END IF;
END//

DELIMITER ;
<<<<<<< HEAD

SELECT 'Triggers created successfully!' AS Status;
SHOW TRIGGERS;
=======
>>>>>>> fc07f468f8ab1a3e8bbde8aad30dcf077a584766
