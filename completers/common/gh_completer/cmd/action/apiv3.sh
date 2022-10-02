#!/bin/sh
set -e -o pipefail

wget -nc https://raw.githubusercontent.com/github/rest-api-description/main/descriptions/api.github.com/api.github.com.json


echo 'package action'
echo
echo '// ./apiv3.sh > apiv3.go '
echo 'var v3Paths map[string][]string = map[string][]string{'

method_endpoints() {
echo '  "'"$1"'" : {'
cat api.github.com.json \
  | jq --raw-output '.paths | to_entries[] | select(.value.'"$1"' != null) | select(.key != "/") | "\"\(.key[1:])\",\"\(.value.'"$1"'.summary)\","' \
  |  sed -e 's#gitignore/templates/{name}#gitignore/templates/{gitignore_name}#' \
         -e 's#codes_of_conduct/{key}#codes_of_conduct/{coc_key}#' \
         -e 's#labels/{name}#labels/{label}#'
echo '  },'
}

method_endpoints delete
method_endpoints get
method_endpoints patch
method_endpoints post
method_endpoints put

echo '}'
