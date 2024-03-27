## Weather-Service

### Running Application
- to run the application, go to root directory and run
 
```go run main.go```

### Future Considerations

- Better logging implemented for the application
- Wrap the muxer with an error handling middleware for logging
- Add granular validation for the incoming lat/lon params
- Potentially, wrap the main() functionality into a command for better usabliity 
- The configuration package needs to have a single environment reader
- Add tests for the WeatherAPI interface and the server handler