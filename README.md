# bookstore-items-api
Book-store OAuth API
## Domain driven design
### Using DDD
- Start from domain and working backwords
    - Domain
    - Use cases
    - Controller
    - Devices/ external interfaces/Web/UI etc
- Dependencies works from outer to inner layer ()
    - For example `Domain` has not information about `Use-cases`, `Use-cases` won't know anything about `Controller` and so on
- Data flows from Outer layer to inner layer and then backwards
- Model entities
    - 
####
- Get access token
- Create access token
- Refresh access token

## Deploy Cassandra 
- docker network create cass-cluster-network
- docker pull cassandra
- docker run -d --name nodeA --network cass-cluster-network cassandra
    - docker run -d --name nodeA -p 9042:9042 -v /Users/amitabhprasad/my-app-data/bookstore-app/cassandra/datadir:/var/lib/cassandra -d cassandra:latest
- docker logs -f nodeA
    - look for `Startup complete`
### Test cassandra deployment
- docker pull datastaxdevs/petclinic-backend
- 
```
docker run -d \
     --name backend \
     --network cass-cluster-network \
     -p 9966:9966 \
     -e CASSANDRA_USE_ASTRA=false \
     -e CASSANDRA_USER=cassandra \
     -e CASSANDRA_PASSWORD=cassandra \
     -e CASSANDRA_LOCAL_DC=datacenter1 \
     -e CASSANDRA_CONTACT_POINTS=nodeA:9042 \
     -e CASSANDRA_KEYSPACE_CQL="CREATE KEYSPACE spring_petclinic WITH REPLICATION = {'class':'SimpleStrategy','replication_factor':1};" \
     datastaxdevs/petclinic-backend
```
- Test
```
curl -X GET "http://localhost:9966/petclinic/api/pettypes" -H "accept: application/json" | jq
```
- Add some data
```
curl -X POST \
    "http://localhost:9966/petclinic/api/pettypes" \
    -H "accept: application/json" \
    -H "Content-Type: application/json" \
    -d "{ \"id\": \"unicorn\", \"name\": \"unicorn\"}" | jq
```
- Log into cassandra
```
docker exec -it nodeA cqlsh
USE spring_petclinic;
SELECT * FROM petclinic_reference_lists WHERE list_name='pet_type';
QUIT;
```
```
 describe keyspaces;
 CREATE KEYSPACE oauth WITH REPLICATION = {'class':'SimpleStrategy','replication_factor':1};
 USE oauth;
 describe tables;
 CREATE TABLE access_tokens( access_token varchar PRIMARY KEY, user_id bigint, client_id bigint, expires bigint);
```
### Cleanup
- docker ps --format '{{.ID}}\t{{.Names}}\t{{.Image}}'
- docker stop $(docker ps -aq)
- docker rm $(docker ps -aq)
## Test
- https://stackoverflow.com/questions/16353016/how-to-go-test-all-tests-in-my-project
- go test -v -coverprofile cover.out ./...


## Commands
- git tag 1.0.0
- git push origin 1.0.0
- go mod init `unique path`
- go test ./... will take care of pulling and updating all the dependency in the module file
## Refrences
- https://www.datastax.com/learn/apache-cassandra-operations-in-kubernetes/running-a-cassandra-application-in-docker
- https://go.dev/ref/mod#private-module-privacy
