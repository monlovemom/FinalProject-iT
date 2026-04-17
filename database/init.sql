-- 1. ต้องสร้าง Sequence (ตัวนับ) ขึ้นมาก่อนเป็นอันดับแรก
CREATE SEQUENCE pid_seq START 1;

-- 2. สร้างตาราง users โดยเรียกใช้ pid_seq ที่เพิ่งสร้าง
CREATE TABLE users (
    id            VARCHAR(10) PRIMARY KEY DEFAULT 'PID' || LPAD(nextval('pid_seq')::TEXT, 2, '0'),
    username      VARCHAR(100) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    full_name     VARCHAR(200),
    role          VARCHAR(20) NOT NULL DEFAULT 'employee',
    created_at    TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE permissions (
    id              SERIAL PRIMARY KEY,
    permission_code VARCHAR(100) NOT NULL UNIQUE,
    description     TEXT
);

-- 3. ตารางนี้ต้องแก้ประเภทข้อมูลของ user_id จาก UUID ให้เป็น VARCHAR(10) เพื่อให้ตรงกับ users(id)
CREATE TABLE user_permissions (
    user_id       VARCHAR(10) REFERENCES users(id) ON DELETE CASCADE,
    permission_id INT REFERENCES permissions(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, permission_id)
);

---------------------------------------------------------
-- ส่วนของการเพิ่มข้อมูลเบื้องต้น (Seed Data)
---------------------------------------------------------

INSERT INTO permissions (permission_code, description) VALUES
('SALE',      'ระบบขาย'),
('PURCHASE',  'ระบบซื้อ'),
('WAREHOUSE', 'ระบบคลัง');

INSERT INTO users (username, password_hash, full_name, role) VALUES
('owner1',  'hash...', 'สมพร เจ้าของร้าน', 'owner'),
('admin1',  'hash...', 'วิชัย ผู้ดูแลระบบ', 'admin'),
('emp1',    'hash...', 'สมชาย พนักงาน A',  'employee'),
('emp2',    'hash...', 'สมศรี พนักงาน B',  'employee');

-- emp1 → ได้แค่ซื้อ
INSERT INTO user_permissions (user_id, permission_id)
SELECT u.id, p.id FROM users u, permissions p
WHERE u.username = 'emp1' AND p.permission_code = 'PURCHASE';

-- emp2 → ได้ซื้อ + ขาย + คลัง
INSERT INTO user_permissions (user_id, permission_id)
SELECT u.id, p.id FROM users u, permissions p
WHERE u.username = 'emp2' AND p.permission_code IN ('PURCHASE', 'SALE', 'WAREHOUSE');