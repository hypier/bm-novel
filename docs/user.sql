CREATE TABLE `user` (
  `user_id` varchar(32) NOT NULL,
  `create_at` datetime DEFAULT NULL,
  `is_lock` bit(1) DEFAULT NULL,
  `need_reset_password` bit(1) DEFAULT NULL,
  `password` varchar(64) NOT NULL,
  `real_name` varchar(12) NOT NULL,
  `role_code` varchar(12) NOT NULL,
  `update_at` datetime DEFAULT NULL,
  `user_name` varchar(32) NOT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `UK_lqjrcobrh9jc8wpcar64q1bfh` (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
