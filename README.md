# front-go

API wrapper for the [Front Public API](https://dev.frontapp.com).

## USAGE

```
client := front.NewClientWithResponses(
  "https://api2.frontapp.com/",
  front.WithAuthorizationToken("<YOUR_AUTH_TOKEN>"),
)
response, err := client.ListAccountsWithResponse(context.Background())
if err == nil && response.JSON200 != nil {
  ...
}
```

## LICENSE

MIT
