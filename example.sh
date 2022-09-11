#!/usr/bin/env bash
last_timestamp=$(cat last_timestamp.txt)
echo "Filter after=${last_timestamp}"
go run . -sourceType rss_podcast_itunes -sourceUrl http://127.0.0.1/example.xml -sourceFilter "after=${last_timestamp}" -targets echo_stdout=pub_date -targetProxy socks5://10.32.15.1:1082 -offset 0 -limit 5 -desc true | tail --lines 1 > last_timestamp.txt
if [ ! -s last_timestamp.txt ]; then
  echo "${last_timestamp}" > last_timestamp.txt
fi
