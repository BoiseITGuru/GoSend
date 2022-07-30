# GoSend

GoSend is a lightweight email system utilizing SendGrid to send and receive emails for multiple domains/mailboxes. It is designed to be run either embedded in another application or can be run as a separate application to keep web services separated.

## Config

Copy the config-excample.json file to config.json in your root directory and configure the needed fields.

```json
{
    "version": "0.0.1",
    "database": {
        "gorm_engine": "mysql",
        "gorm_connection": "<username>:<password>@tcp(<db url/ip>:<db port>)/<db name>?parseTime=true"
    },
    "server": {
        "port": 8080
    },
    "aws": {
        "region": "us-west-2",
        "kms": {
            "encryption_key": "arn:aws:kms:us-west-2:",
            "jwt_key": "arn:aws:kms:us-west-2:"
        }
    }
}
```
