USE procurement_db;

SET GLOBAL log_bin_trust_function_creators = 1;

DELIMITER //
-- удаление старых триггеров и процедур
DROP TRIGGER IF EXISTS trg_delivery_invoice_confirmed//
DROP TRIGGER IF EXISTS trg_delivery_invoice_confirmed_insert//
DROP TRIGGER IF EXISTS trg_delivery_invoice_cancelled//
DROP TRIGGER IF EXISTS trg_transfer_invoice_shipped//
DROP TRIGGER IF EXISTS trg_transfer_invoice_shipped_after//
DROP TRIGGER IF EXISTS trg_transfer_invoice_cancelled//

DROP PROCEDURE IF EXISTS post_delivery_invoice//
DROP PROCEDURE IF EXISTS confirm_delivery_invoice//
-- приходная накладная - проведение (AFTER UPDATE)
-- при смене статуса draft на confirmed - приходуем товар
CREATE TRIGGER trg_delivery_invoice_confirmed
AFTER UPDATE ON delivery_invoices
FOR EACH ROW
BEGIN
    IF OLD.status != 'confirmed' AND NEW.status = 'confirmed' THEN
        INSERT INTO inventory_balances (department_id, product_id, quantity)
        SELECT 
            NEW.recipient_department_id,
            dii.product_id,
            dii.quantity
        FROM delivery_invoice_items dii
        WHERE dii.delivery_invoice_id = NEW.id
        ON DUPLICATE KEY UPDATE
            quantity = inventory_balances.quantity + VALUES(quantity);
    END IF;
END//
-- приходная накладная - проведение (AFTER INSERT)
-- если накладная создаётся сразу со статусом 'confirmed'
CREATE TRIGGER trg_delivery_invoice_confirmed_insert
AFTER INSERT ON delivery_invoices
FOR EACH ROW
BEGIN
    IF NEW.status = 'confirmed' THEN
        INSERT INTO inventory_balances (department_id, product_id, quantity)
        SELECT 
            NEW.recipient_department_id,
            dii.product_id,
            dii.quantity
        FROM delivery_invoice_items dii
        WHERE dii.delivery_invoice_id = NEW.id
        ON DUPLICATE KEY UPDATE
            quantity = inventory_balances.quantity + VALUES(quantity);
    END IF;
END//
-- приходная накладная - отмена (AFTER UPDATE)
-- При смене статуса confirmed на cancelled - сторнируем приход
CREATE TRIGGER trg_delivery_invoice_cancelled
AFTER UPDATE ON delivery_invoices
FOR EACH ROW
BEGIN
    IF OLD.status = 'confirmed' AND NEW.status = 'cancelled' THEN
        UPDATE inventory_balances ib
        JOIN delivery_invoice_items dii ON ib.product_id = dii.product_id
        SET ib.quantity = ib.quantity - dii.quantity
        WHERE dii.delivery_invoice_id = NEW.id
          AND ib.department_id = NEW.recipient_department_id;
    END IF;
END//
-- внутреннее перемещение - проверка остатков (BEFORE UPDATE)
-- проверка хватает ли товара на складе-отправителе
CREATE TRIGGER trg_transfer_invoice_shipped
BEFORE UPDATE ON internal_transfer_invoices
FOR EACH ROW
BEGIN
    DECLARE v_insufficient INT DEFAULT 0;
    
    IF OLD.status != 'shipped' AND NEW.status = 'shipped' THEN
        SELECT COUNT(*) INTO v_insufficient
        FROM internal_transfer_invoice_items itii
        LEFT JOIN inventory_balances ib 
            ON ib.product_id = itii.product_id 
            AND ib.department_id = NEW.from_department_id
        WHERE itii.internal_transfer_invoice_id = NEW.id
          AND (ib.quantity IS NULL OR ib.quantity < itii.quantity);
        
        IF v_insufficient > 0 THEN
            SIGNAL SQLSTATE '45000'
            SET MESSAGE_TEXT = 'Insufficient inventory balance for transfer';
        END IF;
    END IF;
END//
-- внутреннее перемещение - проведение (AFTER UPDATE)
-- списываем с отправителя, приходуем получателю
CREATE TRIGGER trg_transfer_invoice_shipped_after
AFTER UPDATE ON internal_transfer_invoices
FOR EACH ROW
BEGIN
    IF OLD.status != 'shipped' AND NEW.status = 'shipped' THEN
        -- Списываем с отправителя
        INSERT INTO inventory_balances (department_id, product_id, quantity)
        SELECT 
            NEW.from_department_id,
            itii.product_id,
            -itii.quantity
        FROM internal_transfer_invoice_items itii
        WHERE itii.internal_transfer_invoice_id = NEW.id
        ON DUPLICATE KEY UPDATE
            quantity = inventory_balances.quantity + VALUES(quantity);
        
        -- Приходуем получателю
        INSERT INTO inventory_balances (department_id, product_id, quantity)
        SELECT 
            NEW.to_department_id,
            itii.product_id,
            itii.quantity
        FROM internal_transfer_invoice_items itii
        WHERE itii.internal_transfer_invoice_id = NEW.id
        ON DUPLICATE KEY UPDATE
            quantity = inventory_balances.quantity + VALUES(quantity);
    END IF;
END//
-- внутреннее перемещение - отмена (AFTER UPDATE)
-- при смене shipped на cancelled - сторнируем движение
CREATE TRIGGER trg_transfer_invoice_cancelled
AFTER UPDATE ON internal_transfer_invoices
FOR EACH ROW
BEGIN
    IF OLD.status = 'shipped' AND NEW.status = 'cancelled' THEN
        INSERT INTO inventory_balances (department_id, product_id, quantity)
        SELECT 
            NEW.from_department_id,
            itii.product_id,
            itii.quantity
        FROM internal_transfer_invoice_items itii
        WHERE itii.internal_transfer_invoice_id = NEW.id
        ON DUPLICATE KEY UPDATE
            quantity = inventory_balances.quantity + VALUES(quantity);
        
        -- Списываем со склада-получателя
        INSERT INTO inventory_balances (department_id, product_id, quantity)
        SELECT 
            NEW.to_department_id,
            itii.product_id,
            -itii.quantity
        FROM internal_transfer_invoice_items itii
        WHERE itii.internal_transfer_invoice_id = NEW.id
        ON DUPLICATE KEY UPDATE
            quantity = inventory_balances.quantity + VALUES(quantity);
    END IF;
END//

DELIMITER ;

-- проверка
SELECT 'Triggers created successfully!' AS Status;
SELECT TRIGGER_NAME, EVENT_OBJECT_TABLE, ACTION_TIMING, EVENT_MANIPULATION 
FROM information_schema.TRIGGERS 
WHERE TRIGGER_SCHEMA = 'procurement_db';