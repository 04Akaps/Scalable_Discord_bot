# Scalable_Discord_bot
확장성있게 구현하는 DisCode Bot



<h1> MySQL </h1>

MySQL을 통해서 Bot이 핸들링 해야 하는 부분을 관리하고 있다.
데이터 추가에 따른 이벤트를 스트리밍 하는 형태는 아니기 떄문에,

데이터가 추가되어서 추가적인 Bot 구동을 위해서는 단순히 재기동만 하는 형태로 구성을 해보았다.

```
CREATE database discord_bot_db;
use discord_bot_db;

create table bot_info (
    `t_id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `channel_name` VARCHAR(100) NOT NULL,
    `bot_name` VARCHAR(100) NOT NULL,
    `bot_token` VARCHAR(500) NOT NULL
);

create table  bot_handler (
    `t_id` BIGINT PRIMARY KEY  AUTO_INCREMENT,
    `bot_name` VARCHAR(100) NOT NULL,
    `content_match` VARCHAR(100) NOT NULL,
    `type` INT NOT NULL DEFAULT 0,
    `message` VARCHAR(500) NOT NULL DEFAULT ''
)
```

- 추후 발전 가능성 있는 Schema 형태이다.
- 예를들면, Image 정보등을 Type을 통해서 핸들링 하는 형태를 생각하고 작성하였다.