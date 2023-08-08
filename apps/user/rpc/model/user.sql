CREATE TABLE `user` (
                        `user_id` bigint(20) unsigned NOT NULL COMMENT '唯一用户ID，使用雪花算法生成',
                        `username` varchar(32) NOT NULL COMMENT '用户注册用户名，最长32个字符，唯一',
                        `password` varchar(32) NOT NULL COMMENT '密码，最长32个字符',
                        `avatar` varchar(256) NOT NULL COMMENT '用户头像',
                        `background_image` varchar(256) NOT NULL COMMENT '用户个人页顶部大图',
                        `signature` varchar(32) NOT NULL COMMENT '个人简介',
                        PRIMARY KEY (`user_id`),
                        UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='TikTok用户表';