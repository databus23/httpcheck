httpcheck
=========

A dead simple tool to check multiple urls in parallel and report on the result.


```
$> httpcheck https://google.com http://10.10.10.10 https://www.microsoft.com
  https://google.com             : 302 Found (235.59528ms)
  https://www.microsoft.com      : 302 Found (252.012142ms)
  http://10.10.10.10             : Get http://10.10.10.10: net/http: request canceled (Client.Timeout exceeded while awaiting headers)
```
