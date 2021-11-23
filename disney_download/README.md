# DisneyDownload

## Synopsis
Test program to see if it would be possible to get metadata information for Disney+ content or download media.

## Assumptions
* There may or may not be a public API for metadata and media content.
* There is a need to retrieve this information for legal purposes.
* The endpoint is known due to Disney+ being a JavaScript application.

## Requirements
* DisneyDownload will allow users to list metadata information for media.
* DisneyDownload will allow users to download media content.
* DisneyDownload will allow users to scan for additional media content.

## Constraints
* DisneyDownload must make calls to only a development sandbox.
* DisneyDownload must not impact Disney servers.
* DisneyDownload must be lightwight and easily portable.
* DisneyDownload must be compiled.
* DisneyDownload must be compatible with Windows 10+.
* DisneyDownload must be compatible with SUSE 15+.
* DisneyDownload must be compatible with macOS 10.15+.

## Retrospective
* Based on the metadata from the browser, even if there was a public API, access would be cut short due to a script CORS policy.
    * API access is clearly closed to internal access
* Disney+ uses other API's, which doesn't surprise me, but applicationTokens and licenses can clearly be seen in the JavaScript.
* Possible API endpoints:

|Endpoint|Description|
|--|--|
|https://prod-ripcut-delivery.disney-plus.net/v1/|Static content used for images.|
|https://prod-static.disney-plus.net/us-east-1/disneyPlus/app/builds/ce3f05347df82f40e2b31fa593b994e3359143c9d|Deployed build to EC2 instance.|

## Conclusion
Disney+ is safe and secure in terms of data leaks (SECURITY), but users are not aware of how their data might be used (PRIVACY).  Overall, another reason why it's good to be a Disney fan.

## References
* https://flixed.io/disney-api-for-developers/
* https://disneyapi.dev/about

