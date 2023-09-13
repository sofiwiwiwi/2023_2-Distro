|     Nombre     |     Rol     |
|:--------------:|:-----------:|
|  Javier Pérez  | 202004533-k |
| Gabriel Ortiz  | 202073535-2 |
| Sofía Riquelme | 202073615-4 |

# Laboratorio 1

Estructura para correr en máquinas virutales:
Se necesitan 4 máquinas, donde no es necesario pero sí deseable que estén de la siguiente forma
máquina 1: cola rabbit y servidor asia
máquina 2: center y servidor america
máquina 3: servidor europa
máquina 4: servidor oceania

Correr los siguientes comandos:

- Para la máquina 1: 

    `sudo docker pull sofiwi/queue`

    `sudo docker pull sofiwi/asia`

    `sudo docker run -d --name rabbitmq-container -p 5672:5672 -p 15672:15672 rabbitmq:3-management` 

    `sudo docker run -d -p 50053:50053 --name asia-container sofiwi/asia:latest`
- Para la máquina 2:

    `sudo docker pull sofiwi/america`

    `sudo docker run -d -p 50054:50054 --name america-container sofiwi/america:latest`

- Para la máquina 3:

    `sudo docker pull sofiwi/europa

    `sudo docker run -d -p 50052:50052 --name europa-container sofiwi/europa:latest`

- Para la máquina 4: 
    `sudo docker pull sofiwi/oceania`

    `sudo docker run -d -p 50051:50051 --name oceania-container sofiwi/oceania:latest`


Luego: 
- Para la máquina 2:

    `sudo docker pull sofiwi/center`

    `sudo docker run -d --name center-container sofiwi/center:latest`


Para correrlo local hay que cambiar todos los puertos de rabbitmq y de grpc a los correspondientes.
Luego de eso abrir 6 terminales y correr `go run path/to/archivo.go`` en cada una, siguiendo el orden de cola, regionales y luego center