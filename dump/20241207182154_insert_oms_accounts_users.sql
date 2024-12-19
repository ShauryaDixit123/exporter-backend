-- +goose Up
-- +goose StatementBegin

INSERT INTO roles (
    id,
    name,
    display_value
)
VALUES
    -- (1, 'super_admin', 'Super Admin'),
    (2, 'admin', 'Admin'),
    -- (3, 'user', 'User'),
    -- (4, 'guest', 'Guest'),
    -- (5, 'account_admin', 'Account Admin'),
    (6, 'buyer', 'Buyer'),
    (7, 'supplier', 'Supplier');
)

INSERT INTO locations (
    id,
    line1,
    line2,
    area,
    city,
    state,
    pincode,
    country_id,
    is_active
)
VALUES
    ('243be22a-c33d-4717-a1e4-fbf0efc8fe7c', '123 Main St', 'Apt 101', 'Green Area', 'New Delhi', 'Delhi', 110001, 'IN', TRUE),
    ('d1c5e9d8-d599-4a87-b80e-3e1f5e6e5c1e', '456 Elm St', 'Bldg 2', 'Blue Area', 'Mumbai', 'Maharashtra', 400001, 'IN', TRUE),
    ('93b7e213-bd67-4328-b7cd-1b27189b8c8a', '789 Pine St', 'Suite 301', 'Red Area', 'Bangalore', 'Karnataka', 560001, 'IN', TRUE),
    ('46f1d226-3d44-4f2a-937e-3c604c1cc3a1', '321 Oak St', 'Floor 4', 'Yellow Area', 'Chennai', 'Tamil Nadu', 600001, 'IN', TRUE),
    ('fc8de1d9-2cc9-4779-bb5d-b7203836b5b0', '654 Maple St', 'Flat 202', 'Purple Area', 'Kolkata', 'West Bengal', 700001, 'IN', TRUE),
    ('ac2276a1-47d2-41c1-bf19-32d1e9e5e370', '987 Cedar St', 'Apartment 5', 'Orange Area', 'Hyderabad', 'Telangana', 500001, 'IN', TRUE),
    ('e2b0b949-df5e-49cf-b9b5-4ac2d64f8e7b', '321 Birch St', 'Flat 8', 'Pink Area', 'Pune', 'Maharashtra', 411001, 'IN', TRUE),
    ('7f05b22a-04b5-429d-8bb4-b1e8d64d2094', '123 Maple Ave', 'Room 303', 'Silver Area', 'Surat', 'Gujarat', 395001, 'IN', TRUE),
    ('34d08b1f-c313-4239-a309-b830d237b32d', '567 Birch Blvd', 'Suite 707', 'Gold Area', 'Ahmedabad', 'Gujarat', 380001, 'IN', TRUE),
    ('51c745f0-d073-4387-9fd0-e27a73a998c2', '890 Pine Ave', 'Room 303', 'Bronze Area', 'Lucknow', 'Uttar Pradesh', 226001, 'IN', TRUE);

INSERT INTO users (
    id,
    name,
    email,
    password,
    is_parent,
    primary_location_id,
    access_token,
    role_id,
    is_active
)
VALUES
    ('5d74f8e4-8b2f-400f-93f2-1d2e973f2a01', 'Alice', 'alice@example.com', 'hashed_password1', FALSE, '243be22a-c33d-4717-a1e4-fbf0efc8fe7c', 'e89c69c9-74bf-4c59-bdf0-1ebdd3fd7d73', 1, TRUE),
    ('00fc8cb7-2f25-4f20-9aa4-c1b8ab5627b5', 'Bob', 'bob@example.com', 'hashed_password2', FALSE, 'e2b0b949-df5e-49cf-b9b5-4ac2d64f8e7b', '9877cb4b-1abf-4429-bc2f-e0f75da9266a', 2, TRUE),
    ('1d5c7a6a-91f4-48f6-94c1-6f94a3b1be29', 'Charlie', 'charlie@example.com', 'hashed_password3', FALSE, 'e2b0b949-df5e-49cf-b9b5-4ac2d64f8e7b', '30d876a7-4d2a-45d7-a70a-71b07ec2730b', 3, TRUE),
    ('ab3ed97e-dc3b-47f3-87fa-62b5ed37740f', 'Diana', 'diana@example.com', 'hashed_password4', FALSE, 'd1c5e9d8-d599-4a87-b80e-3e1f5e6e5c1e', 'b5d36972-7d95-4fa1-84cf-b9f5bd01c8d1', 1, TRUE),
    ('89f5ac69-7cd8-4de5-a3df-8f209ba17d91', 'Eve', 'eve@example.com', 'hashed_password5', FALSE, 'd1c5e9d8-d599-4a87-b80e-3e1f5e6e5c1e', '65f16e8e-2da5-4850-a71b-265c408c4958', 2, TRUE),
    ('e89c69c9-74bf-4c59-bdf0-1ebdd3fd7d73', 'Frank', 'frank@example.com', 'hashed_password6', FALSE, 'd1c5e9d8-d599-4a87-b80e-3e1f5e6e5c1e', '5d74f8e4-8b2f-400f-93f2-1d2e973f2a01', 3, TRUE),
    ('9877cb4b-1abf-4429-bc2f-e0f75da9266a', 'Grace', 'grace@example.com', 'hashed_password7', FALSE, '234d08b1f-c313-4239-a309-b830d237b32d', '00fc8cb7-2f25-4f20-9aa4-c1b8ab5627b5', 1, TRUE),
    ('30d876a7-4d2a-45d7-a70a-71b07ec2730b', 'Hank', 'hank@example.com', 'hashed_password8', FALSE, '234d08b1f-c313-4239-a309-b830d237b32d', '1d5c7a6a-91f4-48f6-94c1-6f94a3b1be29', 2, TRUE),
    ('b5d36972-7d95-4fa1-84cf-b9f5bd01c8d1', 'Ivy', 'ivy@example.com', 'hashed_password9', FALSE, '51c745f0-d073-4387-9fd0-e27a73a998c2', 'ab3ed97e-dc3b-47f3-87fa-62b5ed37740f', 3, TRUE),
    ('65f16e8e-2da5-4850-a71b-265c408c4958', 'Jack', 'jack@example.com', 'hashed_password10', FALSE, '243be22a-c33d-4717-a1e4-fbf0efc8fe7c', '89f5ac69-7cd8-4de5-a3df-8f209ba17d91', 1, TRUE);

INSERT INTO accounts (
    id,
    gst_no,
    primary_user_id,
    default_workflow_post_order,
    default_workflow_pre_order,
    is_active
)
VALUES
    (11, 123456, '5d74f8e4-8b2f-400f-93f2-1d2e973f2a01', '102a889c-cd78-4a53-913b-4f64aeb8f738', '102a889c-cd78-4a53-913b-4f64aeb8f738', TRUE),
    (22, 123457, '00fc8cb7-2f25-4f20-9aa4-c1b8ab5627b5', '102a889c-cd78-4a53-913b-4f64aeb8f738', '102a889c-cd78-4a53-913b-4f64aeb8f738', TRUE),
    (43, 123458, '1d5c7a6a-91f4-48f6-94c1-6f94a3b1be29', '102a889c-cd78-4a53-913b-4f64aeb8f738', '102a889c-cd78-4a53-913b-4f64aeb8f738', TRUE),
    (45, 123459, 'ab3ed97e-dc3b-47f3-87fa-62b5ed37740f', '102a889c-cd78-4a53-913b-4f64aeb8f738', '102a889c-cd78-4a53-913b-4f64aeb8f738', TRUE),
    (59, 123460, '89f5ac69-7cd8-4de5-a3df-8f209ba17d91', '102a889c-cd78-4a53-913b-4f64aeb8f738', '102a889c-cd78-4a53-913b-4f64aeb8f738', TRUE),
    (60, 123461, 'e89c69c9-74bf-4c59-bdf0-1ebdd3fd7d73', '102a889c-cd78-4a53-913b-4f64aeb8f738', '102a889c-cd78-4a53-913b-4f64aeb8f738', TRUE),
    (73, 123462, '9877cb4b-1abf-4429-bc2f-e0f75da9266a', '102a889c-cd78-4a53-913b-4f64aeb8f738', '102a889c-cd78-4a53-913b-4f64aeb8f738', TRUE),
    (83, 123463, '30d876a7-4d2a-45d7-a70a-71b07ec2730b', '102a889c-cd78-4a53-913b-4f64aeb8f738', '102a889c-cd78-4a53-913b-4f64aeb8f738', TRUE),
    (91, 123464, 'b5d36972-7d95-4fa1-84cf-b9f5bd01c8d1', '102a889c-cd78-4a53-913b-4f64aeb8f738', '102a889c-cd78-4a53-913b-4f64aeb8f738', TRUE),
    (102, 123465, '65f16e8e-2da5-4850-a71b-265c408c4958', '102a889c-cd78-4a53-913b-4f64aeb8f738', '102a889c-cd78-4a53-913b-4f64aeb8f738', TRUE);
   
   INSERT INTO accounts_users_map (
    user_id,
    account_id
)
VALUES
    ('5d74f8e4-8b2f-400f-93f2-1d2e973f2a01', 11),
    ('00fc8cb7-2f25-4f20-9aa4-c1b8ab5627b5', 22),
    ('1d5c7a6a-91f4-48f6-94c1-6f94a3b1be29', 43),
    ('ab3ed97e-dc3b-47f3-87fa-62b5ed37740f', 45),
    ('89f5ac69-7cd8-4de5-a3df-8f209ba17d91', 59),
    ('e89c69c9-74bf-4c59-bdf0-1ebdd3fd7d73', 60),
    ('9877cb4b-1abf-4429-bc2f-e0f75da9266a', 73),
    ('30d876a7-4d2a-45d7-a70a-71b07ec2730b', 83),
    ('b5d36972-7d95-4fa1-84cf-b9f5bd01c8d1', 91),
    ('65f16e8e-2da5-4850-a71b-265c408c4958', 102);

-- User 1: e89c69c9-74bf-4c59-bdf0-1ebdd3fd7d73
INSERT INTO users_locations_map (user_id, location_id, is_active)
VALUES 
    ('e89c69c9-74bf-4c59-bdf0-1ebdd3fd7d73', '244be22a-c33d-4717-a1e4-fbf0efc8fe7c', TRUE),
    ('e89c69c9-74bf-4c59-bdf0-1ebdd3fd7d73', '93b7e213-bd67-4328-b7cd-1b27189b8c8a', TRUE),
    ('e89c69c9-74bf-4c59-bdf0-1ebdd3fd7d73', 'ac2276a1-47d2-41c1-bf19-32d1e9e5e370', TRUE),
    ('e89c69c9-74bf-4c59-bdf0-1ebdd3fd7d73', '7f05b22a-04b5-429d-8bb4-b1e8d64d2094', TRUE),
    ('e89c69c9-74bf-4c59-bdf0-1ebdd3fd7d73', '34d08b1f-c313-4239-a309-b830d237b32d', TRUE);

-- User 2: 5d74f8e4-8b2f-400f-93f2-1d2e973f2a01
INSERT INTO users_locations_map (user_id, location_id, is_active)
VALUES 
    ('5d74f8e4-8b2f-400f-93f2-1d2e973f2a01', '46f1d226-3d44-4f2a-937e-3c604c1cc3a1', TRUE),
    ('5d74f8e4-8b2f-400f-93f2-1d2e973f2a01', 'fc8de1d9-2cc9-4779-bb5d-b7203836b5b0', TRUE),
    ('5d74f8e4-8b2f-400f-93f2-1d2e973f2a01', 'e2b0b949-df5e-49cf-b9b5-4ac2d64f8e7b', TRUE),
    ('5d74f8e4-8b2f-400f-93f2-1d2e973f2a01', '51c745f0-d073-4387-9fd0-e27a73a998c2', TRUE),
    ('5d74f8e4-8b2f-400f-93f2-1d2e973f2a01', '244be22a-c33d-4717-a1e4-fbf0efc8fe7c', TRUE);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
