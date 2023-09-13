make docker-rabbit:
	sudo docker pull sofiwi/queue
	sudo docker run -d --name rabbitmq-container -p 5672:5672 -p 15672:15672 rabbitmq:3-management
	sudo docker run -d --name queue-container sofiwi/queue:latest

make docker-center:
	sudo docker pull sofiwi/center
	sudo docker run -d --name center-container sofiwi/center:latest


make docker-asia:
	sudo docker pull sofiwi/asia
	sudo docker run -d -p 50053:50053 --name asia-container sofiwi/asia:latest

make docker-america:
	sudo docker pull sofiwi/america
    sudo docker run -d -p 50054:50054 --name america-container sofiwi/america:latest

make docker-europa:
	sudo docker pull sofiwi/europa
    sudo docker run -d -p 50052:50052 --name europa-container sofiwi/europa:latest

make docker-oceania:
	sudo docker pull sofiwi/oceania
	sudo docker run -d -p 50051:50051 --name oceania-container sofiwi/oceania:latest
