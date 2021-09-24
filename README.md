# waktu
CLI to parse natural language time and print as timestamp.

Only a thin wrapper for `github.com/tj/go-naturaldate` library.


Example usages:

```
❯ waktu today 11am
2021-09-24T11:00:00+07:00
```

```
❯ waktu yesterday 1:30 pm
2021-09-23T13:30:00+07:00
```
