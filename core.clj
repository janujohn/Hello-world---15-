(ns luminus.core
(:require [ring.adapter.jetty :as jetty]))
 
 (defn handler [request]
 {:status 200
 :headers {"Content-Type" "text/html"}
 :body "<h2>Hello World!</h2>"})
  
  (defn -main []
  (jetty/run-jetty handler {:port 9005}))

