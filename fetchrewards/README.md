# Fetch Rewards Coding Exercise
Author  : Abel Gancsos \
Version : 1.0.0.0 

## Synopsis
Build a REST service to manage credits given to a specific user from given payers (companies).  

## Assumptions
* The system where this application will run either has Docker or Go and Python3 (with requests) installed.
* Credits can be added in any order.
* Credits need to be spent in order of timestamp.
* Payer credits should not go negative.
* A database engine is not required.
* The project must run on Windows, Linux, and macOS (assuming the dependencies are installed).
* A GUI is not required.
* The gin Go package is installed.

## Requirements
* Users will be able to add credits to a repository.
* Users will be able to spend credits from a repository.

## Design implementation
The project was implemented using all built-in packages in a single module.  Upon starting the executable, the REST endpoints are generated and the server listener is started.  When a request comes in, the appropriate handler picks it up and uses the data repository, in this case an in-memory array, to perform the requested operation.  

If the request is adding a new credit, the Repository first checks for any balances due to the payer and uses a portion of the incoming amount to pay off the balance.

When the request is spending available credit, the Repository first ensures that the data is in the appropriate order for audting then itereates through the available credit to update the balances.

As for a request that is to check the current balances, we simply iterate through the credits and calculate the balances for each payer.

### Endpoints
The following endpoints are available.  Assume that the base endpoint is in the format: http://<host-or-ip>:4441
|Endpoint|Description                     |Method|Data                                                                    |
|--      |--                              |--     |--                                                                     |
|/credit | Adds points to the repository. | POST  |{"payer":"<payer-name>", "points":<points>, "timestamp":"<timestamp>"} |
|/spend  | Use available points.          | POST  |{"points":<points>}                                                    |
|/version| Retrieves the version.         | GET   |                                                                       |
|/balance| Retrieve the balances.         | GET   |                                                                       |


## Instructions
1. Once the dependencies are installed, ensure that the Docker image is built by running the below command at the project root:
```bash
docker build -t rewards-server:latest
```

2. Start the container by running the below command:
```bash
docker run -d -p 4441:4441 --name rewards-server -t rewards-server:latest
```
NOTE: There's an assumption that port 4441 is open on the Docker host.  If not, please update the external mapping accordingly.  Example:

```bash
docker run -d -p 4445:4441 --name rewards-server -t rewards-server:latest
```

3. Start the REST service in the container by running the below command:
```bash
docker exec -it rewards-server /root/stuff/go/bin/server &
```

4. Run the test case by running one of the following commands:
```bash
## Local workstation
python3 <path-to-project-root>/tests/tests.py
```

```bash
## Container
docker exec -it rewards-server python3 /root/stuff/go/tests/tests.py
```

## Retrospective
* Overall, this was a really fun project since it was a more realistic example of what would be done during daily tasks.
* What would have made the project better is if the server was attached to a PostgreSQL database, mostly because it would be an opportunity to demonstrate database integration.
* Normally, I would split up the modules for cleaner and more mantainable code then import my custom packages along with the other dependencies.
* Normally, the standard is to use fmt.Printf; the main reason for println is to ensure the newline is generated.
* Semicolons might not be required as implemented, but were added as a preference.  If the rest of the code doesn't have it, it wouldn't be added.
* Instead of having hard-coded messages, I would have added a static resource.
* My only qualm with the project is the thousands separator in the expected result for the spend request.  Not sure if it was a typo, but it's not a valid JSON since the thousands separator depicts a string formatting.  As the structure definition defines the field as an int, I did not include it.

## Closing thoughts
Thank you again for your time in reviewing my resume as well as my project submission and the opportunity to apply for the role.
I look forward to hearing any feedback that might be shared.

## Feedback from Fetch Rewards
### What went well
* Gancsos provided excellent documentation, including a PDF rendering of some details. Very impressive! Gancsos also:
* Used Docker and Go (plus Python for tests) to implement their solution
* Appropriately used HTTP verbs
* Structured their JSON well
* What could have gone better:

### What didn't go so well
* The points balance can become negative (via submitting negative points to the crediting endpoint), and these negative points can be spent. This fails to meet the acceptance criteria.
* The project is poorly structured, with a single, large Golang file that acts as the server. The docs say they would structure their project differently, but the effort to do so upfront seems trivial with as much code as the provided.
* Additionally, Gancsos used only the standard library to implement their solution. Sometimes, this is desirable, but their solution misses some important benefits of third parties that would improve their solution. In essence, they've reinvented the wheel but for seemingly no benefit.
* Along with poor project structure, their Dockerfile is unnecessarily complex and bloated. I get the impression that candidate isn't well versed in Docker.

### What I would have replied with in a PR
* Thank you for raising this, although that's a good point, that's not what the sample expected results shows in the specifications.  During my testing, that -200 is required to get the expected balance.  Without it, the results of the spend call would be [-300, -200, -4500].  Regardless, I added an explicit check, but I still believe the sample was misleading.
* Thank you for your feedback, the thought at the time was simplicity, not neccessarily lazziness and I didn't believe creating packages would be prudent.  Regardless, I restructured the code to better reflect what I would have pushed.
* Thank you for your feedback. After looking into the gin package, I realized how much more of what I wanted to do could be done.  I will be using this in the future.
* Thank you for your feedback. I added some configuration in the Docker file in the event a non-root user needed to get in.  Yes, this could depict bloat, but this doesn't neccessarily elude to lack of knowledge with Docker.  Regardless, I cleaned up the bloat. 

## References
* https://go.dev/dl/
* https://www.python.org/downloads/
* https://stackoverflow.com/questions/69789292/how-to-sort-an-struct-array-by-dynamic-field-name-in-golang

