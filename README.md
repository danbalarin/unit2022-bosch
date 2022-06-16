# Links

* [Trello](https://trello.com/b/FWsgLMCR/unit2022-bosch)
* [UML Chart](https://lucid.app/lucidchart/3be4c301-b037-4a87-8d42-6ab0436b62bf/edit?invitationId=inv_413067ca-ff31-488c-bf39-bb7f6eb15592)
* [FE deployment](https://unit2022-bosch.vercel.app/)
* [BE deployment](https://unit2022.herokuapp.com/api)


## Theory

### TCP, UDP

TCP is a connection-oriented protocol. First it establishes connection with the server, then it sends request and receives response. After each response there is also sent confirmation message from client.

UDP is a connection less protocol. It sends request and receives response without establishing connection.

That means if you have slow internet connection, TCP will be way slower, because of the needed handshake, but you'll be guaranteed to receive all data.

### TLS

TLS is a protocol that is used to encrypt data.
It is used to protect data from eavesdropping and to ensure that data is not tampered with.
It also requires handshake to be done before data is encrypted and sent.

## HTTP

Built on top of TCP and TLS, so for each _host_ there is a need for at least two handshakes.

- HTTP 1.0: each request needs separate connection -> handshakes
- HTTP 1.1: multiple files from the same host can be sent over the same connection, and can be requested in parallel
- HTTP 2.0: parallel request can be completed in no particular order of execution
- HTTP 3.0: replaces TCP and TLS with UDP and QUIC, which is faster and more reliable and requires less handshakes

### Methods

- _CONNECT: establishes connection_
- DELETE: deletes resource
- GET: gets resource
  - GET is the only one from the common methods that is cached by browser by default*
- _HEAD: gets the same headers as GET request_
- _OPTIONS: get list of permitted methods for url/host_
- PATCH: updates resource (partial update)
- POST: general purpose method for manipulating with data (used as create most of the time)
- PUT: updates resource (full override)
- _TRACE: ping with message_

*You can cache other requests as well via the Cache API, if the browser supports it.

### Status codes

- [1xx: Informational messages](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#information_responses)
- [2xx: Successful responses](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#successful_responses)
- [3xx: Redirection messages](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#redirection_messages)
- [4xx: Client error response](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#client_error_responses)
- [5xx: Server error response](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#server_error_responses)


## API

We will talk about API in two different contexts.
First is API as some sort of server that provides you needed information about entities.
And the second API is set of functions/methods that you'll be provided by some library.

There are several ways of defining API (server), but most used ones are REST and GraphQL, both are using JSON format for data exchange.

### REST

It's non standardized way of communication between server and client.
Most of the time, server defines endpoints for each entity as well as some other endpoints, like login, search, etc.
Let's take a Eventio API as an example.
The server is hosted on https://testproject-api-v2.strv.com/, this is the **root URL**.
The API provides us with endpoints for two entities, **Events** and **Users**.
Here is the list of endpoints for Events:

- GET /events - returns list of events
- POST /events - creates new event
- GET /events/{id} - returns event with given id
- PATCH /events/{id} - updates event with given id
- DELETE /events/{id} - deletes event with given id

As you can see, the API is defined by url and method.
It's nicely readable and easy to use.

_The Eventio API gets you attendees in the list detail, but for sake of the example imagine following_
Imagine that you need to get list of attendees that would be interested in the event. That would most likely be exposed on `GET /events/{id}/attendees`.
So if you want to get all the information for the detail page, you'd need to make two requests, detail of event and list of attendees.
This problem is called **under fetching**, we are getting less data from one request than we really need to.

Now let's go the other way. Imagine that list of events also gets you list of attendees for each event.
But you don't need this information, because you are not showing it anywhere, but you basically get huge list for every event.
The bigger the response is, the slower it is as well obviously.
This problem is called **over fetching**, we are getting more data from one request than we need to.

And last problem of REST is that it's not standardized and is detached from the data type.
So basically anyone can chose their own style of defining REST API and you can't even be sure of the type you are getting.

### GraphQL

GraphQL solves all of the problems mentioned above.
It's a query language, so it's standardized, strongly typed.
The drawback is that it's a bit more complicated than REST.
But with this steeper learning curve comes a lot of benefits.
Unlike REST, GraphQL is exposed on a single endpoint, and the requests are made with POST method.
The body of the request contains query that is then executed on the server and you get data in a structure according to the request query.
To differentiate between queries that mutate data there are two types of queries, query and mutation.

If we return to the example with Eventio, let's make query to get detail of event with attendees, just like we did with REST.
```graphql
query GetEvent($id: ID!){
  event(id: $id){
    name
    description
    ...
    attendees{
      name
      ...
    }
  }
}
```

As you can see, we defined query that if executed with id of event, will return event with attendees.
The result of this query is a structure that looks like this:
```json
{
    data: {
        event: {
            name: "Event name",
            description: "Event description",
            attendees: [
                {
                    name: "Attendee name",
                },
                {
                    name: "Attendee name",
                },
            ]
        }
    }
}
```

So you can change the query so it suits your needs.
This solves the under and over fetching problems of the REST.

But GraphQL comes with some problems on its own as well.
First of them comes by design.

<details>
    <summary>Caching</summary>
    <br/>
    Because every request is made via POST method and to the same endpoint, it's not cached by the browser.
    You can cache it by using the Cache API, but that's some additional work.
    Another workaround is to use **persistent queries**.
    What that means is that some queries are exposed on another endpoint, so with this example we could have endpoint like this:

    - GET /events/{id}

    that would internally execute the query above.
</details>


And the second problem is backend only, so if you don't care, just skip next section.
<details>
  <summary>N+1 problem</summary>
  <br/>
  If the executor is not optimized, it will generate
</details>

<details>
  <summary>Current Image Formats</summary>
  <br/>

  *  JPEG
     * Ideal for photographic content
     * Significantly smaller than PNGs
  *  PNG
     * Ideal for graphics that have sharp contrast
     * Supports transparency
  *  GIF
     * Predecessor to PNG
     * Known for animation
  *  WEBP
     * ~30% smaller than JPEG
     * 26% smaller than PNG
     * Supports transparency
     * [Partially supported](https://caniuse.com/?search=webp) on all major browsers
     * Preferred option to use in 2022 with JPEG/PNG fallback
  *  SVG
     * Vector based format
     * Best for icons
</details>

