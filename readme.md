This is a simple cache exercise. 

For purposes of this exercise: 

1. The cache is a fixed size.
2. Connection.getItem always returns a new Rankable object.


A lot of the fun stuff happens in the tests files. 

# Testing the app

To run the tests, run `make test`

To run the tests with the race detector, run `make test_race`. The race detector will check for data races when accessing the cache concurrently.




