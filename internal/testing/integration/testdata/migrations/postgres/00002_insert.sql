-- +goose Up
-- +goose StatementBegin
INSERT INTO owners(owner_name, owner_type) 
    VALUES ('lucas', 'user'), ('space', 'organization');
-- +goose StatementEnd

INSERT INTO owners(owner_name, owner_type) 
    VALUES ('james', 'user'), ('ido177', 'organization');

INSERT INTO repos(repo_full_name, repo_owner_id) 
    VALUES ('james/rover', 3), ('ido177/goose', 4);

-- +goose Down
-- +goose StatementBegin
DELETE FROM owners;
-- +goose StatementEnd
