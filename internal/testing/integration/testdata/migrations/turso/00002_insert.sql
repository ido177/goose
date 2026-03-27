-- +goose Up
-- +goose StatementBegin
INSERT INTO owners(owner_name, owner_type) 
VALUES 
    ('lucas', 'user'),
    ('space', 'organization'),
    ('james', 'user'),
    ('ido177', 'organization');
-- +goose StatementEnd

INSERT INTO repos(repo_full_name, repo_owner_id) 
VALUES 
    ('james/rover', (SELECT owner_id FROM owners WHERE owner_name = 'james')),
    ('ido177/goose', (SELECT owner_id FROM owners WHERE owner_name = 'ido177'));

-- +goose Down
-- +goose StatementBegin
DELETE FROM owners;
-- +goose StatementEnd
