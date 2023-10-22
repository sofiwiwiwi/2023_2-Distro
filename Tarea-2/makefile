docker-ONU:
	sudo docker pull sofiwi/onu
	sudo docker run -d --name onu-container sofiwi/onu:latest

docker-OMS:
	sudo docker pull sofiwi/oms
	sudo docker run -d -p 50051:50051 --name oms-container sofiwi/oms:latest

docker-datanode1:
	sudo docker pull sofiwi/datanode1
	sudo docker run -d -p 50052:50052 --name dn1-container sofiwi/datanode1:latest

docker-datanode2:
	sudo docker pull sofiwi/datanode2
	sudo docker run -d -p 50053:50053 --name dn2-container sofiwi/datanode2:latest

docker-america:
	sudo docker pull sofiwi/america
	sudo docker run -d --name america-container sofiwi/america:latest

docker-asia:
	sudo docker pull sofiwi/asia
	sudo docker run -d --name asia-container sofiwi/asia:latest

docker-europa:
	sudo docker pull sofiwi/europa
	sudo docker run -d --name europa-container sofiwi/europa:latest

docker-oceania:
|	sudo docker pull sofiwi/oceania
	sudo docker run -d --name oceania-container sofiwi/oceania:latest