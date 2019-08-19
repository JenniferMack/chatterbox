bin      := chatterbox
bindir   := cmd/$(bin)
files    :=  server/*.go $(bindir)/*.go

$(bin) : $(files)
	go build ./$(bindir)/...

docker : $(bin)
	docker build -t chatterbox .

docker-run :
	docker run --name chatterbox --rm -d -p 5050:5050 $(bin):latest

docker-kill :
	docker kill --signal=INT chatterbox
