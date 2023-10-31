# ScyllaCloudGoAPI
Project is to demo an example of scyllacloud with go API

head to scylla cloud and create a free cluster and download the connection certificate. Place it in db folder. conn.yaml here in this project

in cqlsh terminal

CREATE KEYSPACE catalog WITH REPLICATION = { 'class' : 'NetworkTopologyStrategy','DC1' : 3};

use catalog;

CREATE TABLE mutant_data (
   first_name text,
   last_name text,
   address text,
   picture_location text,
   PRIMARY KEY((first_name, last_name)));


insert into mutant_data ("first_name","last_name","address","picture_location") VALUES ('Bob','Loblaw','1313 Mockingbird Lane', 'http://www.facebook.com/bobloblaw');
insert into mutant_data ("first_name","last_name","address","picture_location") VALUES ('Bob','Zemuda','1202 Coffman Lane', 'http://www.facebook.com/bzemuda');
insert into mutant_data ("first_name","last_name","address","picture_location") VALUES ('Jim','Jeffries','1211 Hollywood Lane', 'http://www.facebook.com/jeffries');


Run the project now: 

go get .

go run main.go

localhost:8080


note: don't forget to change the certificate path in db/scylladb.go

const (
	connectionBundlePath = "/Users/preethamu/Desktop/code/golang/gitgolangscylla/ScyllaCloudGoAPI/db/conn.yaml"
)

change this to your path


