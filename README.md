# Backend

- Go (gin)
- MongoDB

## API

- `/create-room` ? `owner`
- `/join-room` ? `user` & `room`
- `/exit-room` ? `user` & `room`
- `/send-message` ? `user` & `room` & `text`
- `/get-room` ? `room`

## Room structure

```javascript
{
  _id: ObjectId("64df18873f4d4638c8bb2c24"),
  sessionId: '870d0ae3-0632-4639-ab8f-f22ace84bf1a',
  owner: 'Anon',
  participants: [
    { userId: 'cjfhh1ovbiv3iu26hqig', name: 'Anon' },
    { userId: 'cjfhie8vbiv3iu26hqj0', name: 'splash' },
    { userId: 'cjfhin0vbiv3iu26hqjg', name: 'K' }
  ],
  messages: [
    {
      userId: 'cjfhh1ovbiv3iu26hqig',
      text: '"Hey, have you heard about this awesome new chat application?"',
      sent_at: ISODate("2023-08-18T07:13:33.955Z")
    },
    {
      userId: 'cjfhie8vbiv3iu26hqj0',
      text: 'Oh really? Tell me more about it!',
      sent_at: ISODate("2023-08-18T07:14:04.871Z")
    },
    {
      userId: 'cjfhh1ovbiv3iu26hqig',
      text: 'It's a chat application!',
      sent_at: ISODate("2023-08-18T07:13:33.955Z")
    },
    {
      userId: 'cjfhin0vbiv3iu26hqjg',
      text: 'Wait, that sounds interesting.',
      sent_at: ISODate("2023-08-18T07:15:02.550Z")
    }
  ],
  created_at: ISODate("2023-08-18T07:06:47.078Z"),
  updated_at: ISODate("2023-08-18T07:06:47.078Z")
}
```
<sup>lame</sup>

### Features

- If zero participants, destroy room after five minutes
- Perform cleanup of rooms if inactive for more than 30 minutes
