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

<h1> DisCord Complex </h1>

아무래도 Discord의 복잡한 메시지 프로토콜을 따라주며 작업이 들어가야 한다.
그러다보니 각각 메시지마다 원하는 타입을 선언하고, 언마살하면 사용을 하는 방향으로 작업이 되었다.

```
{
    "content": "This is a message with components",
    "components": [
        {
            "type": 1,
            "components": [
                {
                    "type": 2,
                    "label": "Click me!",
                    "style": 1,
                    "custom_id": "click_one"
                }
            ]

        }
    ]
}
```
- 이러한 형태에 대해서, `type/bot/helloBotChannel`과 같이 선언해서 사용하게 된다.