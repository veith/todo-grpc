-- auto generated definition
-- source ____
-- generate time 2018-11-27 02:56:46


-- Ereignisse welche von der Galerie erkannt wurden 
-- Die Funktion ist nur gefaket 
-- Das ist ein test von mutilines
CREATE TABLE "fotos"
(
  `id`       BLOB,        -- für IDs werden wir intern ulid verwenden
  `data`       TEXT,        -- Base64 encodetes Bild des Ereigniss
  `description`       TEXT,        -- Beschreibung des Tasks
  PRIMARY KEY (`id`)
)

