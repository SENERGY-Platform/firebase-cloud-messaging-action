# Firebase Cloud Messaging Action

Send a message via Firebase Cloud Messaging (FCM) in your GitHub workflows.

## Inputs

## `CREDENTIALS`

**Required** Your FCM credentials in JSON format. This includes the project_id, private_key, etc. Use a secret for this!

## `MESSAGE`

**Required** The message to be sent in JSON format. Example: {"topic": "my-topic", "data": {"foo": "bar"}}

## Example usage
    - name: FCM Action
      uses: senergy-platform/firebase-cloud-messaging-action@latest
      with:
        CREDENTIALS: ${{ secrets.FCM_CREDENTIALS }}
        MESSAGE: '{"topic": "android", "data": {"foo": "bar"}}'
