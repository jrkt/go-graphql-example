package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/jrkt/go-graphql-example/basic-plus/graphql"
	"log"
	"net/http"
)

func main() {

	schema, err := graphql.CompileSchema()
	if err != nil {
		log.Fatalln(err)
	}

	// create router
	router := mux.NewRouter()
	router.Handle("/graphiql", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(GraphIQL)
	}))
	router.Handle("/query", &relay.Handler{Schema: schema})

	fmt.Println("serving http on :8001")
	err = http.ListenAndServe(":8001", router)
	if err != nil {
		log.Fatalln(err)
	}
}

var GraphIQL = []byte(`
<!DOCTYPE html>
  <html>
       <head>
               <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.11.10/graphiql.css" />
               <script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.0/fetch.min.js"></script>
               <script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
               <script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
               <script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.11.10/graphiql.js"></script>
               <script src="//unpkg.com/subscriptions-transport-ws@0.8.3/browser/client.js"></script>
               <script src="//unpkg.com/graphiql-subscriptions-fetcher@0.0.2/browser/client.js"></script>
       </head>
       <body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
               <div id="graphiql" style="height: 100vh;">Loading...</div>
               <script>
                       function graphQLFetcher(graphQLParams) {
                               return fetch("/query", {
                                       method: "post",
                                       body: JSON.stringify(graphQLParams),
                                       credentials: "include",
                               }).then(function (response) {
                                       return response.text();
                               }).then(function (responseBody) {
                                       try {
                                               return JSON.parse(responseBody);
                                       } catch (error) {
                                               return responseBody;
                                       }
                               });
                       }
                       var subscriptionsClient = new window.SubscriptionsTransportWs.SubscriptionClient('ws://localhost:8001/query', { reconnect: true });
                       var subscriptionsFetcher = window.GraphiQLSubscriptionsFetcher.graphQLFetcher(subscriptionsClient, graphQLFetcher);
                       ReactDOM.render(
                               React.createElement(GraphiQL, {fetcher: subscriptionsFetcher}),
                               document.getElementById("graphiql")
                       );
               </script>
       </body>
  </html>
`)
