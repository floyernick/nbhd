CREATE TABLE sessions (
  id      UUID PRIMARY KEY,
  user_id UUID NOT NULL
);

CREATE TABLE users (
  id       UUID PRIMARY KEY,
  name     VARCHAR(150) NOT NULL,
  photo    VARCHAR(250) NOT NULL,
  phone    VARCHAR(10)  NOT NULL UNIQUE,
  location VARCHAR(50)  NOT NULL,
  password VARCHAR(60)  NOT NULL
);

CREATE TABLE tasks (
  id               UUID PRIMARY KEY,
  title            VARCHAR(75)       NOT NULL,
  category         SMALLINT          NOT NULL,
  location         DECIMAL ARRAY [2] NOT NULL,
  description      VARCHAR(280)      NOT NULL,
  time             INT               NOT NULL,
  creator          UUID              NOT NULL,
  performer        UUID              NOT NULL,
  encouragement    SMALLINT          NOT NULL,
  pay              DECIMAL           NOT NULL,
  creator_rating   SMALLINT          NOT NULL,
  performer_rating SMALLINT          NOT NULL,
  created          TIMESTAMP         NOT NULL,
  archived         BOOLEAN           NOT NULL
);


CREATE TABLE task_performers (
  task_id UUID NOT NULL,
  user_id UUID NOT NULL,
  PRIMARY KEY (task_id, user_id)
);