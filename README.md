# front-go

API wrapper for the [Front Public API](https://dev.frontapp.com).

## USAGE

```
client, err := front.NewClientWithResponses(
  "https://api2.frontapp.com/",
  front.WithAuthorizationToken("<YOUR_AUTH_TOKEN>"),
)
if err != nil {
  // ...
}

response, err := client.ListAccountsWithResponse(context.Background())
if err == nil && response.JSON200 != nil {
  // ...
}
```

## LICENSE

MIT
