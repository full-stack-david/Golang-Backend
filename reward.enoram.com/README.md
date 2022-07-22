# reward-go

### Project Structure

  1. Model - Same as Entities, A model in Go is a set of data structures and functions, will store any Object’s Struct and its method. Example : ActivityLog, Timesheet etc.
  2. Services - This layer contains application specific business rules. It encapsulates and implements all of the use cases of the system.
  3. Repository - Repository will store any Database handler. If calling microservices, will handled create HTTP Request to other services, and sanitize the data.
  4. Controller - This layer is a set of adapters that convert data from the format most convenient for the services and models, to the format most convenient for some external interface such as REST API or grpc
  5.  Utils / Configs - This layers will conatin packages such as 
	
			|--> util    -- Utility functions
			|--> config  -- Application Configuration
			|--> comm    -- Commuication API Wrapper
	

  