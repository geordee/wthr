# WTHR

Weather App using Go.

## Installation

```bash
git checkout github.com/geordee/wthr
go get
go run .
```

## Develop

Format the code

```bash
npx prettier --write .
```

Develop with hot reloading of assets

```bash
npm run dev
go run .
```

Build

```bash
npm run build
go build .
```

## Test

```bash
curl -H "Accept: application/json" "localhost:8080/weather?latitude=51.5072&longitude=0.1276"
curl -H "Accept: text/html" "localhost:8080/weather?latitude=51.5072&longitude=0.1276"
curl "localhost:8080/weather?latitude=51.5072&longitude=0.1276"
```
