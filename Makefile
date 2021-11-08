swagger:
	wget 'https://developer.godaddy.com/swagger/swagger_domains.json'
client-api:swagger
	swagger-codegen generate -i swagger_domains.json \
                -l go -o gen/out/godaddy -DpackageName=godaddy
