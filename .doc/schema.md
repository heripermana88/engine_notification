# Engine Notification

## API 
### Agents (GET)
#### response
    [
        {
            "id":"67764c4d1b5a837459257201",
            "name":"sms"
        },
        {
            "id":"a83745925720167764c4d1b5",
            "name":"whatsapp"
        },
        {
            "id":"a83745925720167764c4d1b5",
            "name":"email"
        },
    ]
### Tempalates (GET)
#### response
    [
        {
            "id":"67764c4d1b5a837459257201",
            "name":"notification expiration Subscription"
        },
        {
            "id":"a83745925720167764c4d1b5",
            "name":"notification Lebaran"
        },
    ]
### Request Notification (POST)
#### payload
    {
        "recipient": [
            {
                "key": "userId",
                "value": "12345"
            }
        ],
        //or
        "criteria": [
            {
                "field":"ssoId",
                "operation":"in",
                "value":[1312,123123,123123]
            }
        ],
        "template_id": "67764c4d1b5a837459257201",
        "quota": 200,
        "agent": "email",
        "meta_data": {
            "priority": "high",
            "retries":0,
            "execute_at": 1698637555585,
            "max_execute_at":1698637555585
        }
    }

## Schema

<b>message_notifications</b> is an  <i>Abstraction</i> of <b>request_notifications</b> per <u>user</u>.

request_notifications contains:<br/>
<b>recipient</b> array or <b>criteria</b> array that is used to get users to whom notifications will be sent<br>
<b>template_id</b> that is used to get message content</br>
<b>agent_id</b> that is used to get agent notification</br>
<b>metadata</b> that is used to details execute notification</br>
<b>status</b> as progress status request notification</br>


### Schema Design
```mermaid
erDiagram
    agents {
        A string PK "_id"
        A string "name"
        ANA string "providers[].name"
        ANAA string "providers[].config.host"
        ANAA string "providers[].config.port"
        ANAA string "providers[].config.username"
        ANAA string "providers[].config.password"
        A string "default_provider"
    }
    templates {
        A string PK "_id"
        A string "client_id"
        A string "title"
        A string "subject"
        A string "content"
    }
    request_notifications {
        A string PK "_id"
        ANA string "receipients[].key"
        ANA string "receipients[].value"
        ANA string "criteria[].key"
        ANA string "criteria[].operation"
        ANA string "criteria[].value"
        A string "template_id"
        A number "quota"
        AA string "agent._id"
        AA string "agent.name"
        AA string "metadata.priority"
        AA number "metadata.retries"
        AA number "metadata.execute_at"
        AA number "metadata.max_execute_at"
        A string "status"
        A number "created_at"
        A number "updated_at"
        A number "deleted_at"
    }
    message_notifications {
        A string PK "_id"
        A number "receipient.ssoId"
        A number "receipient.name"
        A number "receipient.email"
        A string "receipient.msisdn"
        A string "receipient.deviceId"
        A string "agent"
        AA string "message.title"
        AA string "message.subject"
        AA string "message.content"
        AA string "message.deepLink"
        AA string "metadata.priority"
        AA number "metadata.retries"
        AA number "metadata.execute_at"
        AA number "metadata.max_execute_at"
        A string "status"
        A number "created_at"
        A number "updated_at"
        A number "deleted_at"
    }
    logs {
        A string PK "_id"
        A string PK "notification_id"
        A number "date"
        A string "event"
        AA string "detail.provider"
        AA string "detail.status"
        AA string "detail.error_message"
    }
    agents }o--|| "request_notifications": "has agents"
    templates }o--|| "request_notifications": "has templates"
    agents }o--|| "message_notifications": "has agents"
    message_notifications }o--|| "logs": "has logs"
```

## QUEUE
### Exchange
    engine.notification
#### RoutingKey
#### 
    create
##### payload 
    {
        "recipient": [
            {
                "key": "userId",
                "value": "12345"
            }
        ],
        "criteria": [
            {

            }
        ],
        "template_id": {
            "_id": "6654d406a0361ef5331c6b42",
            "name": "template notif A",
        },
        "quota": 200,
        "agent": {
            "_id": "6654d406a0361ef5331c6b42",
            "name": "email",
        },
        "meta_data": {
            "priority": "high",
            "retries":0,
            "execute_at": 1698637555585,
            "max_execute_at":1698637555585
        }
    }
####
    attempt
##### payload
    {
        "_id": "6654d4aea0361ef5331c7b23",
        "recipient": {
            "ssoId": 123455,
            "email": "jhon.pantau@gmail.com",
            "name": "Jhon Pantau",
        },
        "message": {
            "title": "lalala",
            "subject": "yeyeyeye",
            "content": "content message",
        },
        "agent": "email",
        "meta_data": {
            "priority": "high",
            "retries":0,
            "execute_at": 1698637555585,
            "max_execute_at":1698637555585
        },
        "status": "waiting"
    }
####
    update_attempt
##### payload success
    {
        "_id": "6654d4aea0361ef5331c7b23",
        "result": {
            "success": true,
            "status": "success",
            "code":"",
            "error_message":""
        } 
    }
##### payload failed
    {
        "_id": "6654d4aea0361ef5331c7b23",
        "result": {
            "success": false,
            "status": "failed",
            "code":"",
            "error_message":""
        } 
    }