# LogSearch

## Synopsis
Meant to be used strictly for troubleshooting.  Opposed to business analytics, the solution focusses on ingesting and filtering semi-formatted logs in an easy and sufficient way.

## Assumptions
* The solution may or may not be ran on Windows or Linux.
* The solution may or may not be used for mainframe logs.
* Log formats may be comma-separated (,).
* Log formats may be tab-separated (	).
* Log formats may be pipe-separated (|).
* Log formats may be JSON.
* Filter language may be SQL.
* Filter language may be key-pair search.
* Return values may be JSON format.
* A RESTful API will be required.

## Requirements
* The solution must be able to ingest different formats of logs.
* The solution must be able to store ingested data.
* The solution must be able to accept RESTful requests.
* The solution must be able to return RESTful responses.
* The solution must be able to filter table columns.
* The solution must be able to filter child structures.

## Implementation Details

### Endpoints
|Endpoint                       | Action | Description                                                          |
|--|--|--|--|
|/test/                         | GET    | Simple HELLO WORLD response.                                         |
|/purge/                        | POST   | Purge the database of events.                                        |
|/query/                        | GET    | Query events.                                                        |
|/add/                          | POST   | Add event to data source.                                            |


### Flags
#### API
|Flag                           | Description                                                                   |
|--|--|
|--port                         | The port to use for the RESTful API.                                          |

#### CLI
|Flag                           | Description                                                                   |
|--|--|
|--end                          | Based endpoint for the API.                                                   |
|--op                           | Operation to perform.                                                         |
|--query                        | Search string for event lookups.                                              |
|-f                             | Full path to log file for ingestion.                                          |

## References

