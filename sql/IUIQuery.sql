/* tạo bảng */
CREATE TABLE `ot`.`users` (
                              `id` VARCHAR(45) NOT NULL,
                              `password` VARCHAR(45) NOT NULL,
                              `max_todo` VARCHAR(45) NOT NULL DEFAULT '5',
                              PRIMARY KEY (`id`));
/* insert thông tin */
INSERT INTO `ot`.`users` (`id`, `password`, `max_todo`) VALUES ('1', 'abcd', '5');
INSERT INTO `ot`.`users` (`id`, `password`, `max_todo`) VALUES ('2', '1234', '5');
INSERT INTO `ot`.`users` (`id`, `password`, `max_todo`) VALUES ('3', 'firstUser', '5');

/* tạo bảng */
CREATE TABLE `ot`.`task` (
                             `id` VARCHAR(45) NOT NULL,
                             `content` VARCHAR(45) NOT NULL,
                             `user_id` VARCHAR(45) NOT NULL,
                             `created_date` VARCHAR(45) NOT NULL,
                             PRIMARY KEY (`id`));

/* insert data */
insert into task (id, content, user_id, created_date) values ('1', 'abc', '1', '23/8/2021');

/* ràng buộc giữa 2 bảng */
ALTER TABLE `ot`.`task`
    ADD INDEX `user_id_idx` (`user_id` ASC) VISIBLE;
;
ALTER TABLE `ot`.`task`
    ADD CONSTRAINT `user_id`
        FOREIGN KEY (`user_id`)
            REFERENCES `ot`.`users` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION;
