-- 1단계: unique_id 컬럼 추가
ALTER TABLE
    school_notifications
ADD
    COLUMN unique_id VARCHAR(36) DEFAULT NULL;

-- 2단계: 기존 레코드에 UUID 부여
UPDATE
    school_notifications
SET
    unique_id = UUID()
WHERE
    unique_id IS NULL;

-- 3단계: NOT NULL 제약조건 추가
ALTER TABLE
    school_notifications
MODIFY
    COLUMN unique_id VARCHAR(36) NOT NULL;

-- 4단계: UNIQUE 제약조건 추가
ALTER TABLE
    school_notifications
ADD
    CONSTRAINT unique_id_constraint UNIQUE (unique_id);