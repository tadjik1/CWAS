# [K]WAS - Curl With AWS Signing

Extremely useful CLI utility that helps you send HTTP requests to AWS. 
 
## About
A very simple tool that sign your requests to AWS API. It was written to [sign](http://docs.aws.amazon.com/general/latest/gr/signing_aws_api_requests.html)
HTTP requests to AWS API.

## Install
1. Download executable file for your platform from latest release.
2. Put path to file into $PATH environment variable (optional).

## Usage
```bash
./kwas -X PUT https://es-service.eu-west-1.es.amazonaws.com/_snapshot/backup/snapshot -d'{
    "indices": "my_index"
}'
```

### Options
```bash
usage: kwas --region=REGION --service=SERVICE [<flags>] <url>

[K]WAS - Curl With AWS Signing.

Flags:
      --help               Show context-sensitive help (also try --help-long and --help-man).
      --region=REGION      AWS region
      --service=SERVICE    AWS Service short name
  -X, --method="GET"       Request method.
  -d, --body=""            Request body.
  -H, --header=HEADER ...  Request headers (could be an array).
      --version            Show application version.

Args:
  <url>  Request url
```

### Credentials
This library is trying to find credentials in your session or in ENV variables `AWS_AWS_ACCESS_KEY_ID`, `AWS_ACCESS_KEY`.
