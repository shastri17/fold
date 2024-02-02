CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150),
    createdat TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS projects (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150),
    slug VARCHAR(150),
    description text,
    createdat TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );


CREATE TABLE IF NOT EXISTS user_projects (
    user_id int,
    project_id int,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_project FOREIGN KEY(project_id) REFERENCES projects(id)
);

CREATE TABLE IF NOT EXISTS hashtags (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150),
    createdat TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS project_hashtags(
    hashtag_id int,
    project_id int,
    CONSTRAINT fk_user FOREIGN KEY(hashtag_id) REFERENCES hashtags(id),
    CONSTRAINT fk_project FOREIGN KEY(project_id) REFERENCES projects(id)
);

insert into users(name)values ('Test user 1');
insert into projects(name,slug,description)values('Test project','test-project','description of test prokect');
insert into user_projects(user_id,project_id)values (1,1);

