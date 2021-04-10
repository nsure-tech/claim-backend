CREATE DATABASE `vote` CHARACTER SET 'utf8' COLLATE 'utf8_general_ci';

drop table if exists `v_config`;
CREATE TABLE `v_config` (
                            `id` bigint NOT NULL AUTO_INCREMENT,
                            `created_at` timestamp NULL DEFAULT NULL,
                            `updated_at` timestamp NULL DEFAULT NULL,
                            `key_word` varchar(255) NOT NULL,
                            `val`  varchar(255) NOT NULL,
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `idx_keyword` (`key_word`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

drop table if exists `v_transfer`;
CREATE TABLE `v_transfer` (
                              `id` bigint NOT NULL AUTO_INCREMENT,
                              `created_at` timestamp NULL DEFAULT NULL,
                              `updated_at` timestamp NULL DEFAULT NULL,
                              `currency` varchar(255) NOT NULL,
                              `from_address` varchar(255) NOT NULL,
                              `to_address` varchar(255) NOT NULL,
                              `amount` decimal(64,0) NOT NULL DEFAULT '0',
                              `raw` varchar(20000) NOT NULL,
                              `status` varchar(255) NOT NULL,
                              `settled` tinyint(1) NOT NULL DEFAULT '0',
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

drop table if exists `v_metamask`;
CREATE TABLE `v_metamask` (
                              `id` bigint NOT NULL AUTO_INCREMENT,
                              `created_at` timestamp NULL DEFAULT NULL,
                              `updated_at` timestamp NULL DEFAULT NULL,
                              `user_id` varchar(255) NOT NULL,
                              `sig_hex` varchar(255) NOT NULL,
                              `msg` varchar(20000) NOT NULL,
                              `settled` tinyint(1) NOT NULL DEFAULT '0',
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

drop table if exists `v_account`;
CREATE TABLE `v_account` (
                             `id` bigint NOT NULL AUTO_INCREMENT,
                             `created_at` timestamp NULL DEFAULT NULL,
                             `updated_at` timestamp NULL DEFAULT NULL,
                             `user_id` varchar(255) NOT NULL,
                             `currency` varchar(255) NOT NULL,
                             `available` decimal(64,0) NOT NULL DEFAULT '0',
                             `hold` decimal(64,0) NOT NULL DEFAULT '0',
                             PRIMARY KEY (`id`),
                             UNIQUE KEY `idx_uid_currency` (`user_id`,`currency`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

drop table if exists `v_qualification`;
CREATE TABLE `v_qualification` (
                                   `id` bigint NOT NULL AUTO_INCREMENT,
                                   `created_at` timestamp NULL DEFAULT NULL,
                                   `updated_at` timestamp NULL DEFAULT NULL,
                                   `arbiter_id` varchar(255) NOT NULL,
                                   `available` int NOT NULL DEFAULT '0',
                                   `used` int NOT NULL DEFAULT '0',
                                   `pending` int NOT NULL DEFAULT '0',
                                   `closed` int NOT NULL DEFAULT '0',
                                   PRIMARY KEY (`id`),
                                   UNIQUE KEY `idx_arbiter` (`arbiter_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

drop table if exists `v_pending`;
CREATE TABLE `v_pending` (
                             `id` bigint NOT NULL AUTO_INCREMENT,
                             `created_at` timestamp NULL DEFAULT NULL,
                             `updated_at` timestamp NULL DEFAULT NULL,
                             `arbiter_id` varchar(255) NOT NULL,
                             `submit_at` timestamp NULL DEFAULT NULL,
                             `pending` int NOT NULL DEFAULT '0',
                             `settled` tinyint(1) NOT NULL DEFAULT '0',
                             PRIMARY KEY (`id`),
                             KEY `idx_arbiter` (`arbiter_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

drop table if exists `v_bill`;
CREATE TABLE `v_bill` (
                          `id` bigint NOT NULL AUTO_INCREMENT,
                          `created_at` timestamp NULL DEFAULT NULL,
                          `updated_at` timestamp NULL DEFAULT NULL,
                          `user_id` varchar(255) NOT NULL,
                          `currency` varchar(255) NOT NULL,
                          `available` decimal(64,0) NOT NULL DEFAULT '0',
                          `hold` decimal(64,0) NOT NULL DEFAULT '0',
                          `type` varchar(255) NOT NULL,
                          `settled` tinyint(1) NOT NULL DEFAULT '0',
                          `notes` varchar(255) DEFAULT NULL,
                          PRIMARY KEY (`id`),
                          KEY `idx_gsoci` (`user_id`,`currency`,`settled`),
                          KEY `idx_s` (`settled`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

drop table if exists `v_claim`;
CREATE TABLE `v_claim` (
                           `id` bigint NOT NULL AUTO_INCREMENT,
                           `created_at` timestamp NULL DEFAULT NULL,
                           `updated_at` timestamp NULL DEFAULT NULL,
                           `user_id` varchar(255) NOT NULL,
                           `product` varchar(255) NOT NULL,
                           `cover_id` varchar(255) NOT NULL,
                           `cover_hash` varchar(255) NOT NULL,
                           `currency` varchar(255) NOT NULL,
                           `amount` decimal(64,0) NOT NULL DEFAULT '0',
                           `cost` decimal(64,0) NOT NULL DEFAULT '0',
                           `reward` decimal(64,0) NOT NULL DEFAULT '0',
                           `submit_at` timestamp NULL DEFAULT NULL,
                           `arbiter_at` timestamp DEFAULT NULL,
                           `vote_at` timestamp DEFAULT NULL,
                           `challenge_at` timestamp DEFAULT NULL,
                           `cover_begin_at` timestamp NULL DEFAULT NULL,
                           `cover_end_at` timestamp NULL DEFAULT NULL,
                           `status` varchar(255) NOT NULL,
                           `payment_status` varchar(255) DEFAULT NULL,
                           `apply_num` int NOT NULL DEFAULT '0',
                           `vote_num` int NOT NULL DEFAULT '0',
                           `challenged` tinyint(1) NOT NULL DEFAULT '0',
                           `loss` varchar(255) DEFAULT NULL,
                           `description` varchar(5000) DEFAULT NULL,
                           `credential` varchar(5000) DEFAULT NULL,
                           `settled` tinyint(1) NOT NULL DEFAULT '0',
                           `notes` varchar(255) DEFAULT NULL,
                           PRIMARY KEY (`id`),
                           UNIQUE KEY `id_hash` (`cover_hash`),
                           KEY `idx_uid_product` (`user_id`,`product`),
                           KEY `idx_product_status` (`product`, `status`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

# 20选5使用
drop table if exists `v_apply`;
CREATE TABLE `v_apply` (
                           `id` bigint(20) NOT NULL AUTO_INCREMENT,
                           `created_at` timestamp NULL DEFAULT NULL,
                           `updated_at` timestamp NULL DEFAULT NULL,
                           `claim_id` bigint(20) NOT NULL,
                           `submit_at` timestamp NULL DEFAULT NULL,
                           `cover_id` varchar(255) NOT NULL,
                           `cover_hash` varchar(255) NOT NULL,
                           `user_id` varchar(255) NOT NULL,
                           `product` varchar(255) NOT NULL,
                           `apply_at` timestamp NULL DEFAULT NULL,
                           `arbiter_at` timestamp DEFAULT NULL,
                           `apply_num` int NOT NULL DEFAULT '0',
                           `arbiter_id` varchar(255) NOT NULL,
                           `status` varchar(255) NOT NULL,
                           `settled` tinyint(1) NOT NULL DEFAULT '0',
                           PRIMARY KEY (`id`),
                           UNIQUE KEY `idx_claim_arbiter` (`claim_id`,`arbiter_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

drop table if exists `v_challenge`;
CREATE TABLE `v_challenge` (
                               `id` bigint NOT NULL AUTO_INCREMENT,
                               `created_at` timestamp NULL DEFAULT NULL,
                               `updated_at` timestamp NULL DEFAULT NULL,
                               `challenge_id` varchar(255) NOT NULL,
                               `challenge_at` timestamp NOT NULL,
                               `claim_id`  bigint(20) NOT NULL,
                               `cover_id` varchar(255) NOT NULL,
                               `cover_hash` varchar(255) NOT NULL,
                               `currency` varchar(255) NOT NULL,
                               `amount` decimal(64,0) NOT NULL DEFAULT '0',
                               `reward` decimal(64,0) NOT NULL DEFAULT '0',
                               `hold` decimal(64,0) NOT NULL DEFAULT '0',
                               `claim_status` varchar(255) DEFAULT NULL,
                               `status` varchar(255) DEFAULT NULL,
                               `settled` tinyint(1) NOT NULL DEFAULT '0',
                               `notes` varchar(255) DEFAULT NULL,
                               PRIMARY KEY (`id`),
                               UNIQUE KEY `id_claim` (`claim_id`),
                               UNIQUE KEY `id_hash` (`cover_hash`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

drop table if exists `v_challenge_fill`;
CREATE TABLE `v_challenge_fill` (
                                    `id` bigint NOT NULL AUTO_INCREMENT,
                                    `created_at` timestamp NULL DEFAULT NULL,
                                    `updated_at` timestamp NULL DEFAULT NULL,
                                    `challenge_id` varchar(255) NOT NULL,
                                    `challenge_at` timestamp NOT NULL,
                                    `claim_id`  bigint(20) NOT NULL,
                                    `cover_id` varchar(255) NOT NULL,
                                    `cover_hash` varchar(255) NOT NULL,
                                    `user_id` varchar(255) NOT NULL,
                                    `product` varchar(255) NOT NULL,
                                    `currency` varchar(255) NOT NULL,
                                    `amount` decimal(64,0) NOT NULL DEFAULT '0',
                                    `reward` decimal(64,0) NOT NULL DEFAULT '0',
                                    `hold` decimal(64,0) NOT NULL DEFAULT '0',
                                    `claim_status` varchar(255) DEFAULT NULL,
                                    `status` varchar(255) DEFAULT NULL,
                                    `settled` tinyint(1) NOT NULL DEFAULT '0',
                                    `notes` varchar(255) DEFAULT NULL,
                                    PRIMARY KEY (`id`),
                                    UNIQUE KEY `id_claim` (`claim_id`),
                                    UNIQUE KEY `id_hash` (`cover_hash`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

# 仲裁
drop table if exists `v_vote`;
CREATE TABLE `v_vote` (
                          `id` bigint(20) NOT NULL AUTO_INCREMENT,
                          `created_at` timestamp NULL DEFAULT NULL,
                          `updated_at` timestamp NULL DEFAULT NULL,
                          `claim_id` bigint(20) NOT NULL,
                          `arbiter_id` varchar(255) NOT NULL,
                          `user_id` varchar(255) NOT NULL,
                          `product` varchar(255) NOT NULL,
                          `cover_id` varchar(255) NOT NULL,
                          `cover_hash` varchar(255) NOT NULL,
                          `currency` varchar(255) NOT NULL,
                          `amount` decimal(64,0) NOT NULL DEFAULT '0',
                          `reward` decimal(64,0) NOT NULL DEFAULT '0',
                          `arbiter_at` timestamp NULL DEFAULT NULL,
                          `submit_at` timestamp NULL DEFAULT NULL,
                          `cover_begin_at` timestamp NULL DEFAULT NULL,
                          `status` varchar(255) NOT NULL,
                          `sign_hash` varchar(255) DEFAULT NULL,
                          `settled` tinyint(1) NOT NULL DEFAULT '0',
                          `notes` varchar(255) DEFAULT NULL,
                          PRIMARY KEY (`id`),
                          UNIQUE KEY `idx_claim_arbiter` (`claim_id`,`arbiter_id`),
                          KEY `idx_aid_product` (`arbiter_id`,`product`, `status`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

drop table if exists `v_vote_fill`;
CREATE TABLE `v_vote_fill` (
                               `id` bigint(20) NOT NULL AUTO_INCREMENT,
                               `created_at` timestamp NULL DEFAULT NULL,
                               `updated_at` timestamp NULL DEFAULT NULL,
                               `vote_id` bigint(20) NOT NULL,
                               `claim_id` bigint(20) NOT NULL,
                               `cover_id` varchar(255) NOT NULL,
                               `cover_hash` varchar(255) NOT NULL,
                               `currency` varchar(255) NOT NULL,
                               `amount` decimal(64,0) NOT NULL DEFAULT '0',
                               `reward` decimal(64,0) NOT NULL DEFAULT '0',
                               `arbiter_id` varchar(255) NOT NULL,
                               `vote_at` timestamp NOT NULL,
                               `claim_status` varchar(255) NOT NULL,
                               `vote_status` varchar(255) NOT NULL,
                               `payment_status` varchar(255) DEFAULT NULL,
                               `challenge_status` varchar(255) DEFAULT NULL,
                               `status` varchar(255) NOT NULL,
                               `reward_num` int NOT NULL DEFAULT '0',
                               `settled` tinyint(1) NOT NULL DEFAULT '0',
                               `notes` varchar(255) DEFAULT NULL,
                               PRIMARY KEY (`id`),
                               UNIQUE KEY `idx_vote` (`vote_id`),
                               UNIQUE KEY `idx_claim_arbiter` (`claim_id`,`arbiter_id`),
                               KEY `idx_s` (`settled`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

drop table if exists `v_reward`;
CREATE TABLE `v_reward` (
                            `id` bigint(20) NOT NULL AUTO_INCREMENT,
                            `created_at` timestamp NULL DEFAULT NULL,
                            `updated_at` timestamp NULL DEFAULT NULL,
                            `vote_id` bigint(20) NOT NULL,
                            `claim_id` bigint(20) NOT NULL,
                            `cover_id` varchar(255) NOT NULL,
                            `cover_hash` varchar(255) NOT NULL,
                            `arbiter_id` varchar(255) NOT NULL,
                            `claim_status` varchar(255) NOT NULL,
                            `vote_status` varchar(255) NOT NULL,
                            `currency` varchar(255) NOT NULL,
                            `amount` decimal(64,0) NOT NULL DEFAULT '0',
                            `settled` tinyint(1) NOT NULL DEFAULT '0',
                            `notes` varchar(255) DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `idx_vote` (`vote_id`),
                            UNIQUE KEY `idx_claim_arbiter` (`claim_id`,`arbiter_id`),
                            KEY `idx_s` (`settled`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

drop table if exists `v_reward_cha`;
CREATE TABLE `v_reward_cha` (
                                `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                `created_at` timestamp NULL DEFAULT NULL,
                                `updated_at` timestamp NULL DEFAULT NULL,
                                `claim_id` bigint(20) NOT NULL,
                                `cover_id` varchar(255) NOT NULL,
                                `cover_hash` varchar(255) NOT NULL,
                                `challenge_id` varchar(255) NOT NULL,
                                `claim_status` varchar(255) NOT NULL,
                                `challenge_status` varchar(255) NOT NULL,
                                `currency` varchar(255) NOT NULL,
                                `amount` decimal(64,0) NOT NULL DEFAULT '0',
                                `settled` tinyint(1) NOT NULL DEFAULT '0',
                                `notes` varchar(255) DEFAULT NULL,
                                PRIMARY KEY (`id`),
                                UNIQUE KEY `idx_claim_arbiter` (`claim_id`,`challenge_id`, `currency`),
                                KEY `idx_s` (`settled`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

drop table if exists `v_punish_cha`;
CREATE TABLE `v_punish_cha` (
                                `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                `created_at` timestamp NULL DEFAULT NULL,
                                `updated_at` timestamp NULL DEFAULT NULL,
                                `claim_id` bigint(20) NOT NULL,
                                `cover_id` varchar(255) NOT NULL,
                                `cover_hash` varchar(255) NOT NULL,
                                `challenge_id` varchar(255) NOT NULL,
                                `claim_status` varchar(255) NOT NULL,
                                `challenge_status` varchar(255) NOT NULL,
                                `currency` varchar(255) NOT NULL,
                                `amount` decimal(64,0) NOT NULL DEFAULT '0',
                                `settled` tinyint(1) NOT NULL DEFAULT '0',
                                `notes` varchar(255) DEFAULT NULL,
                                PRIMARY KEY (`id`),
                                UNIQUE KEY `idx_claim_cha` (`claim_id`,`challenge_id`, `currency`),
                                KEY `idx_s` (`settled`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

drop table if exists `v_payment`;
CREATE TABLE `v_payment` (
                             `id` bigint NOT NULL AUTO_INCREMENT,
                             `created_at` timestamp NULL DEFAULT NULL,
                             `updated_at` timestamp NULL DEFAULT NULL,
                             `user_id` varchar(255) NOT NULL,
                             `product` varchar(255) NOT NULL,
                             `cover_id` varchar(255) NOT NULL,
                             `cover_hash` varchar(255) NOT NULL,
                             `currency` varchar(255) NOT NULL,
                             `amount` decimal(64,0) NOT NULL DEFAULT '0',
                             `claim_id` bigint(20) NOT NULL,
                             `claim_status` varchar(255) DEFAULT NULL,
                             `admin_id` varchar(255) DEFAULT NULL,
                             `pay` decimal(64,0) NOT NULL DEFAULT '0',
                             `settled` tinyint(1) NOT NULL DEFAULT '0',
                             `notes` varchar(255) DEFAULT NULL,
                             PRIMARY KEY (`id`),
                             UNIQUE KEY `id_hash` (`cover_hash`),
                             KEY `idx_uid_product` (`user_id`,`product`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

drop table if exists `v_chain_claim`;
CREATE TABLE `v_chain_claim` (
                                 `id` bigint NOT NULL AUTO_INCREMENT,
                                 `created_at` timestamp NULL DEFAULT NULL,
                                 `updated_at` timestamp NULL DEFAULT NULL,
                                 `user_id` varchar(255) NOT NULL,
                                 `currency` varchar(255) NOT NULL,
                                 `amount` decimal(64,0) NOT NULL DEFAULT '0',
                                 `nonce` bigint NOT NULL,
                                 `raw` varchar(20000) NOT NULL,
                                 `status` varchar(255) NOT NULL,
                                 `settled` tinyint(1) NOT NULL DEFAULT '0',
                                 PRIMARY KEY (`id`),
                                 UNIQUE KEY `idx_user_nonce` (`user_id`,`nonce`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

drop table if exists `v_withdraw`;
CREATE TABLE `v_withdraw` (
                              `id` bigint NOT NULL AUTO_INCREMENT,
                              `created_at` timestamp NULL DEFAULT NULL,
                              `updated_at` timestamp NULL DEFAULT NULL,
                              `user_id` varchar(255) NOT NULL,
                              `currency` varchar(255) NOT NULL,
                              `amount` decimal(64,0) NOT NULL DEFAULT '0',
                              `nonce` bigint NOT NULL,
                              `end_at` timestamp NOT NULL,
                              `status` varchar(255) NOT NULL,
                              `settled` tinyint(1) NOT NULL DEFAULT '0',
                              PRIMARY KEY (`id`),
                              KEY `idx_user_nonce` (`user_id`,`nonce`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


insert into `v_config`(`created_at`,`updated_at`,`key_word`,`val`) values
(null,null,"block_number_deposit","100000"),
(null,null,"block_number_claim","100000");

insert into `v_config`(`created_at`,`updated_at`,`key_word`,`val`) values
(null,null,"currency_0","ETH"),
(null,null,"currency_1","Nsure");

insert into `v_config`(`created_at`,`updated_at`,`key_word`,`val`) values
(null,null,"admin_address_1","0x11"),
(null,null,"admin_address_2","0x12"),
(null,null,"admin_address_3","0x13"),
(null,null,"admin_address_4","0x14"),
(null,null,"admin_address_5","0x15"),
(null,null,"admin_address_6","0x16"),
(null,null,"admin_address_7","0x17"),
(null,null,"admin_address_8","0x18");

insert into `v_config`(`created_at`,`updated_at`,`key_word`,`val`) values
(null,null,"challenge_address_1","0x11"),
(null,null,"challenge_address_2","0x12"),
(null,null,"challenge_address_3","0x13"),
(null,null,"challenge_address_4","0x14"),
(null,null,"challenge_address_5","0x15");

ALTER TABLE `v_claim` ADD `loss` varchar(255) DEFAULT NULL;

20210410
ALTER TABLE  `v_payment` ADD `admin_id` varchar(255) DEFAULT NULL;
ALTER TABLE  `v_payment` ADD `pay` decimal(64,0) NOT NULL DEFAULT '0';