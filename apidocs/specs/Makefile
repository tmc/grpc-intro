a/ll:
	cat *.swagger.json | jq --slurp 'reduce .[] as $$item ({}; . * $$item)' > apis.json
