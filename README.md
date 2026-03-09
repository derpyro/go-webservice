# First GO Project

This is a simple go project to experiment with webservices (REST-APIs).
When started you can access the REST-API on [localhost:8080/animal](http://localhost:8080/animal).

> [!IMPORTANT]  
> Critical content demanding immediate user attention due to potential risks.

## How to run

If you want to run this code you can simply run it on your machine:

```bash
$ go run animal-service.go
```

> [!NOTE]  
> To run the application like this you must have [go](https://go.dev/doc/install) installed.

If you do not want to install [go](https://go.dev/doc/install) on your machine, you can also get this application as docker container [here](https://hub.docker.com/r/fehuworks/animal-service) and start it like this:

```bash
$ docker run -p 8080:8080 --name animal-service -d fehuworks/animal-service
3bde87e6dd620051b0a904c3e5278e28c6514a029fb7c5fc31faac12f447d67e
```

## Endpoints

> [!TIP]
> All endpoints produce and consume JSON.

### Entity

```json
{
  "id": "8a44da9e-83d4-4878-954c-7cf0aa479ad9",
  "type": "cat",
  "gender": "F",
  "name": "Mira",
  "weight": 3.3
}
```

### Get all animals

<code>GET</code> on <code>/animal</code> will return an array of all known animals in the system (or an empty array when no animals are known)

```bash
$ curl -X GET http://localhost:8080/animal
[{"id":"c7cd5ef2-46f0-4174-b1cd-037bb3f8e2bd","type":"cat","gender":"F","name":"Mira","weight":3.3}, {"id":"b64fa3af-b3c1-4a2f-8982-4768f7d99b57","type":"cat","gender":"M","name":"Pommes","weight":6.66}]
```

### Create an animal

<code>POST</code> on <code>/animal</code> will create (and save) the received entity if no validation errors occur. It is not necessary to provide an <code>id</code> in the entity, it will be ignored.
This action will return the created entity with the generated <code>id</code>.

```bash
$ curl -X POST -H "Content-Type: application/json" -d '{"type":"cat","gender":"M","name":"Pommes","weight":6.66}' http://localhost:8080/animal
{"id":"b64fa3af-b3c1-4a2f-8982-4768f7d99b57","type":"cat","gender":"M","name":"Pommes","weight":6.66}
```

### Edit an animal

<code>PUT</code> on <code>/animal</code> will override an entity by <code>id</code>. If the sent entity does not contain an <code>id</code> or the provided <code>id</code> is not known an error ([404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status/404)) will be returned. If the <code>id</code> was known the old values of the entity will be replaced with the new provided ones.

```bash
$ curl -X PUT -H "Content-Type: application/json" -d '{"id": "c7cd5ef2-46f0-4174-b1cd-037bb3f8e2bd","type": "cat","gender": "F","name": "Mira","weight": 4.20}' http://localhost
:8080/animal
{"id":"c7cd5ef2-46f0-4174-b1cd-037bb3f8e2bd","type":"cat","gender":"F","name":"Mira","weight":4.2}
```

### Get details of an animal

<code>GET</code> on <code>/animal/{id}</code> will return all details of the corresponding entity if the given <code>id</code> is known. If the given <code>id</code> is not known an error ([404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status/404)) will be returned.

```bash
$ curl -X GET http://localhost:8080/animal/c7cd5ef2-46f0-4174-b1cd-037bb3f8e2bd
{"id":"c7cd5ef2-46f0-4174-b1cd-037bb3f8e2bd","type":"cat","gender":"F","name":"Mira","weight":4.2}
```

### Delete an animal

<code>DELETE</code> on <code>/animal/{id}</code> will delete the corresponding entity if the given <code>id</code> is known and return a copy of the deleted entity. If the given <code>id</code> is not known an error ([404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status/404)) will be returned.

```bash
$ curl -X DELETE -H "Content-Type: application/json" http://localhost:8080/animal/47d346e1-54f9-47eb-8c36-98cbd9e48390
{"id":"47d346e1-54f9-47eb-8c36-98cbd9e48390","type":"dog","gender":"M","name":"Andre","weight":25.1}
```