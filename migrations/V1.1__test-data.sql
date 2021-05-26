-- ATTENTION: for test purposes (don't run this in production)

INSERT INTO job_execution (object_id, sleep, status, created_at, updated_at)
VALUES('1', 15, 'SUCCESS', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO job_execution (object_id, sleep, status, created_at, updated_at)
VALUES('2', null, 'PROCESSING', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO job_execution (object_id, sleep, status, created_at, updated_at)
VALUES('3', null, 'CANCELLED', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO job_execution (object_id, sleep, status, created_at, updated_at)
VALUES('4', null, 'FAILED', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO job_execution (object_id, sleep, status, created_at, updated_at)
VALUES('5', 35, 'SUCCESS', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO job_execution (object_id, sleep, status, created_at, updated_at)
VALUES('6', 30, null, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP - interval '5 minutes' );


