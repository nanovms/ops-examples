(defn handler
  "Simple handler for connections."
  [stream]
  (defer (:close stream)
    (def id (gensym))
    (def b @"")
    (print "Connection " id "!")
    (while (:read stream 1024 b)
      (printf " %v -> %v" id b)
      (:write stream b)
      (buffer/clear b))
    (printf "Done %v!" id)
    (ev/sleep 0.5)))

(net/server "0.0.0.0" "8000" handler)
