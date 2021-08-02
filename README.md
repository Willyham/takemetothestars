# takemetothestars
A tool to help city people find places where they can see the stars


## Development

### Backend

The backend is a simple Go app. It currently serves a single 'candidates' endpoint as JSON over HTTP.
It comes with a stubbed implementation for working on the frontend independently.

Run the backend with:

```
cd backend
go run .
```

You can the query with curl:

```
➜  ~ curl "localhost:8080/candidates?lat=51.509865&long=-0.118092&limit=2" | jq .
[
  {
    "Point": [
      51.566110948888195,
      -0.11809823922080634
    ],
    "DistanceMeters": 6254,
    "TravelTimeMinutes": 12
  },
  {
    "Point": [
      51.42498610628127,
      -0.11796948432380247
    ],
    "DistanceMeters": 9438,
    "TravelTimeMinutes": 18
  }
]
```

### Web

The web is a simple react app. You can start with:

```
cd web
yarn
```

then `yarn start`.