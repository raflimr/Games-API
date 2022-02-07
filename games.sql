CREATE TABLE `publisher` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `logopub` varchar(255) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `deskripsi` varchar(255) NOT NULL,
  `website` varchar(255) NOT NULL
);

CREATE TABLE `developer` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `logodev` varchar(255) NOT NULL,
  `kantor_pusat` varchar(255) NOT NULL,
  `pendiri` varchar(255) NOT NULL,
  `tahun_pendirian` varchar(255) NOT NULL
);

CREATE TABLE `kategori` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `nama` varchar(255) NOT NULL,
  `deskripsi` varchar(255) NOT NULL
);

CREATE TABLE `game` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `judul` varchar(255) NOT NULL,
  `deskripsi` varchar(255) NOT NULL,
  `penerbit` varchar(255) NOT NULL,
  `platform` varchar(255) NOT NULL,
  `website` varchar(255) NOT NULL,
  `status_game` varchar(255) NOT NULL,
  `tanggal_rilis` varchar(255) NOT NULL,
  `developer_id` integer NOT NULL,
  `publisher_id` integer NOT NULL
);

CREATE TABLE `gambar` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `url` varchar(255) NOT NULL,
  `game_id` integer NOT NULL
);

CREATE TABLE `kategori_game` (
  `kategori_id` integer NOT NULL,
  `game_id` integer NOT NULL,
  PRIMARY KEY (`kategori_id`, `game_id`)
);

CREATE TABLE `user` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(255) NOT NULL UNIQUE,
  `password` varchar(255) NOT NULL,
  `role` varchar(255) NOT NULL
);

CREATE TABLE `review` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `konten_review` varchar(255) NOT NULL,
  `tanggal_review` timestamp NOT NULL,
  `tipe` varchar(255) NOT NULL,
  `game_id` integer NOT NULL,
  `user_id` integer NOT NULL
);

CREATE TABLE `reaksi_review` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `tipe_reaksi` varchar(255) NOT NULL,
  `review_id` integer NOT NULL,
  `user_id` integer NOT NULL
);

CREATE TABLE `rating` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `rating` integer NOT NULL,
  `game_id` integer NOT NULL,
  `user_id` integer NOT NULL
);

ALTER TABLE `game` ADD FOREIGN KEY (`developer_id`) REFERENCES `developer` (`id`);

ALTER TABLE `game` ADD FOREIGN KEY (`publisher_id`) REFERENCES `publisher` (`id`);

ALTER TABLE `gambar` ADD FOREIGN KEY (`game_id`) REFERENCES `game` (`id`);

ALTER TABLE `kategori_game` ADD FOREIGN KEY (`kategori_id`) REFERENCES `kategori` (`id`);

ALTER TABLE `kategori_game` ADD FOREIGN KEY (`game_id`) REFERENCES `game` (`id`);

ALTER TABLE `review` ADD FOREIGN KEY (`game_id`) REFERENCES `game` (`id`);

ALTER TABLE `review` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `reaksi_review` ADD FOREIGN KEY (`review_id`) REFERENCES `review` (`id`);

ALTER TABLE `reaksi_review` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `rating` ADD FOREIGN KEY (`game_id`) REFERENCES `game` (`id`);

ALTER TABLE `rating` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

