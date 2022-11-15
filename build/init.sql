CREATE TABLE `records` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `url` varchar(255) NOT NULL,
  `display_name` varchar(255) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_time` datetime(6) NOT NULL,
  `updated_time` datetime(6) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO crud_demo.records (display_name, url, description, created_time, updated_time) VALUES
('Cupcake', '/url/cupcake', 'Description of Cupcake', '2022-10-29 18:31:22.180788000', '2022-10-29 18:31:22.180788000'),
('Donut', '/url/donut', 'Description of Donut', '2022-10-29 18:31:22.337593000', '2022-10-29 18:31:22.337593000'),
('Eclair', '/url/eclair', 'Description of Eclair', '2022-10-29 18:31:22.359053000', '2022-10-29 18:31:22.359053000'),
('Cheesecake', '/url/cheesecake', 'Description of Cheesecake', '2022-10-29 18:31:22.378373000', '2022-10-29 18:31:22.378373000'),
('Gingerbread', '/url/gingerbread', 'Description of Gingerbread', '2022-10-29 18:31:22.402472000', '2022-10-29 18:31:22.402472000'),
('Honeycomb', '/url/honeycomb', 'Description of Honeycomb', '2022-10-29 18:31:22.429566000', '2022-10-29 18:31:22.429566000'),
('IceCreamSandwich', '/url/icecreamsandwich', 'Description of IceCreamSandwich', '2022-10-29 18:31:22.449591000', '2022-10-29 18:31:22.449591000'),
('JellyBean', '/url/jellybean', 'Description of JellyBean', '2022-10-29 18:31:22.474969000', '2022-10-29 18:31:22.474969000'),
('KitKat', '/url/kitkat', 'Description of KitKat', '2022-10-29 18:31:22.498620000', '2022-10-29 18:31:22.498620000'),
('Lollipop', '/url/lollipop', 'Description of Lollipop', '2022-10-29 18:31:22.525241000', '2022-10-29 18:31:22.525241000'),
('Marshmallow', '/url/marshmallow', 'Description of Marshmallow', '2022-10-29 18:31:22.544983000', '2022-10-29 18:31:22.544983000'),
('Oreo', '/url/oreo', 'Description of Oreo', '2022-10-29 18:31:22.568580000', '2022-10-29 18:31:22.568580000'),
('Pie', '/url/pie', 'Description of Pie', '2022-10-29 18:31:22.590239000', '2022-10-29 18:31:22.590239000'),
('Tiramisu', '/url/tiramisu', 'Description of Tiramisu', '2022-10-29 18:31:22.612282000', '2022-10-29 18:31:22.612282000'),
('RedVelvetCake', '/url/redvelvetcake', 'Description of RedVelvetCake', '2022-10-29 18:31:22.638185000', '2022-10-29 18:31:22.638185000');