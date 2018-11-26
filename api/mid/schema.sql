-- auto generated definition
-- source ____
-- generate time 2018-11-26 10:36:09

-- Ereignisse welche von der Galerie erkannt wurden 
-- Die Funktion ist nur gefaket 
-- Das ist ein test von mutilines
CREATE TABLE "fotos"
(
  `id`          BLOB NOT NULL PRIMARY KEY, -- für IDs werden wir intern ulid verwenden
  `data`        TEXT, -- Base64 encodetes Bild des Ereigniss
  `description` TEXT, -- Beschreibung des Tasks
  PRIMARY KEY (`id`)
)
