-- auto generated definition

-- Tasks items...
CREATE TABLE "tasks"
(
  `id`       BLOB NOT NULL,         -- für IDs werden wir intern ulid verwenden
  `title`       BLOB NOT NULL,         -- Titel des Tasks
  PRIMARY KEY (`id`)
)
