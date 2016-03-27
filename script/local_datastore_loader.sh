#!/usr/bin/env bash

source=~/.githubcompare/store/part-test

while IFS=, read repository day stars
do
    curl 'http://localhost:8000/datastore/edit' -H 'Origin: http://localhost:8000' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: en-US,en;q=0.8,fr;q=0.6,sv;q=0.4' -H 'Upgrade-Insecure-Requests: 1' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.116 Safari/537.36' -H 'Content-Type: application/x-www-form-urlencoded' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8' -H 'Cache-Control: max-age=0' -H 'Referer: http://localhost:8000/datastore/edit?kind=RepositoryStarEvent&next=http%3A%2F%2Flocalhost%3A8000%2Fdatastore%3Fkind%3DRepositoryStarEvent' -H 'Connection: keep-alive' --data "xsrf_token=ZoKkZreCbf&next=http%3A%2F%2Flocalhost%3A8000%2Fdatastore%3Fkind%3DRepositoryStarEvent&kind=RepositoryStarEvent&datetime%7CDate=${day}+00%3A00%3A00&string%7CRepositoryName=${repository}&int%7CStars=${stars}" --compressed
done < $source