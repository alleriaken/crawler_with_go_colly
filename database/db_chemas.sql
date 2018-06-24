CREATE TABLE `words` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `word` VARCHAR(255) NOT NULL DEFAULT "",
  `w_type` VARCHAR(64) NOT NULL DEFAULT "",
  `w_def` TEXT NULL DEFAULT NULL,
  `w_pron` TEXT NULL DEFAULT NULL,
  `examword_id` VARCHAR(255) NOT NULL UNIQUE,
  PRIMARY KEY (`id`)
);

CREATE TABLE `examples` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `word_id` INT(11) NOT NULL,
  `example` TEXT NOT NULL,
  `created_by` INT(11) NOT NULL DEFAULT 0,
  `approved_by` INT(11) NOT NULL DEFAULT 0,
  `approved_at` INT(11) NOT NULL DEFAULT 0,
  `created_at` INT(11) NOT NULL DEFAULT 0,
  `vote_up` INT(11) NOT NULL DEFAULT 0,
  `vote_down` INT(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
);

CREATE TABLE `synonyms` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `word_id` INT(11) NOT NULL,
  `definition` TEXT NOT NULL,
  `synonyms` TEXT NOT NULL,
  PRIMARY KEY (`id`)
);