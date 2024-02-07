# Direct7 Go SDK

Go SDK to seamlessly incorporate communication features into your Go applications via the Direct7 REST API. This SDK empowers you to effortlessly initiate SMS,Whatsapp, Slack, Viber  messages and 2 factor authentication features.

## Documentation
The documentation for the Direct7 REST API can be found here [Direct7 API Reference](https://d7networks.com/docs/).


## Installation

Make sure your project is using Go Modules (it will have a go.mod file in its root if it already is):

```bash
go mod init
```
Then, reference direct7-go-sdk in a Go program with import:

```bash
import (
	"github.com/d7networks/direct7-go-sdk"
)
```

Run any of the normal go commands (build/install/test). The Go toolchain will resolve and fetch the direct7-go module automatically.

Alternatively, you can also explicitly go get the package into a project:

```bash
go get -u github.com/d7networks/direct7-go-sdk
```

## Usage

To get started you need to have an active Direct7 account, If you haven't yet registered, please proceed to [Sign up](https://app.d7networks.com/signup?tag="direct7-go-sdk")

### Authentication

In order to initiate API requests, create a client object using your Direct7 API token. To obtain an API token, kindly visit the following link: https://app.d7networks.com/api-tokens.

## Examples

  - [SMS](#sms)
  - [Verify](#verify)
  - [Whatsapp](#whatsapp)
  - [Number Lookup](#number-lookup)
  - [Viber](#viber)
  - [Slack](#slack)

### SMS


### Send SMS

```go
    package direct7
	apiToken := "Your Api Token"
	client := NewClient(apiToken)
	params := Message{
		Recipients:  []string{"+919999XXXXXX"},
		Content:     "Greetings from D7 API",
		Unicode:     "false",
	}
    response, err := client.sms.SendMessages([]Message{params}, "Sender", "https://the_url_to_receive_delivery_report.com", "2024-02-05T09:48:42+0000")
```

### Send an Unicode SMS

```go
    package direct7
	apiToken := "Your Api Token"
	client := NewClient(apiToken)
	params := Message{
		Recipients:  []string{"+919999XXXXXX"},
		Content:     "مرحبا بالعالم!",
		Unicode:     "true",
	}
    response, err := client.sms.SendMessages([]Message{params}, "Sender", "https://the_url_to_receive_delivery_report.com", "2024-02-05T09:48:42+0000")
```

### Check SMS Request Status

```go
    package direct7
    apiToken := "Your Api Token"
    client := NewClient(apiToken)
    requestID := "001ff613-de30-4f82-81f6-1fe944b8f61b"
    statusResponse, err := client.sms.GetStatus(requestID)
```

### Verify

### Send OTP

```go
package direct7
apiToken := "Your Api Token"
client := NewClient(apiToken)
originator := "SignOTP"
recipient := "+919999XXXXXX"
content := "Greetings from D7 API, your mobile verification code is: {}"
dataCoding := "text"
expiry := 600 
response, err := client.verify.SendOTP(originator, recipient, content, dataCoding, expiry)
```

### Re-Send OTP

```go
package direct7
apiToken := "Your Api Token"
client := NewClient(apiToken)
otpID := "aeffa23f-1204-4e17-bb91-adf6de2cf826"
response, err := client.verify.ResendOTP(otpID)
```

### Verify OTP

```go
package direct7
apiToken := "Your Api Token"
client := NewClient(apiToken)
otpID := "32549451-0eb6-4788-8e91-32b2eb9c4260"
otpCode := "803053"
response, err := client.verify.VerifyOTP(otpID, otpCode)

```

### Check Verify Request Status

```go
    package direct7
    apiToken := "Your Api Token"
    client := NewClient(apiToken)
    requestID := "001ff613-de30-4f82-81f6-1fe944b8f61b"
    statusResponse, err := client.verify.GetStatus(requestID)
```

### Whatsapp

### Send Whatsapp Free-form Message (Contact Details)

```go
package direct7
apiToken := "Your Api Token"
client := NewClient(apiToken)
originator := "+9190XXXXXXXX"
recipient := "+919999XXXXXX"
messageType := "TEXT"
optParams := &OptionalParams{messageText: "HI"}
response, err := client.whatsapp.SendWhatsAppFreeformMessage(originator, recipient, messageType, optParams)

```

### Send Whatsapp Templated Message.

```go
package direct7
apiToken := "Your Api Token"
client := NewClient(apiToken)
originator := "+919061525574"
recipient := "+919999XXXXXX"
templateID := "marketing_media_image"
optParams := &OptionalParams{mediaType: "image", mediaURL: "https://25428574.fs1.hubspotusercontent-eu1.net/hubfs/25428574/D7%20Logo%20rect.webp", bodyParameterValues: map[string]interface{}{
	"0": "Anu",
}}

response, err := client.whatsapp.SendWhatsAppTemplatedMessage(originator, recipient, templateID, optParams)

```

### Check Whatsapp Request Status

```go
package direct7
apiToken := "Your Api Token"
client := NewClient(apiToken)
requestID := "001ff613-de30-4f82-81f6-1fe944b8f61b"
statusResponse, err := client.whatsapp.GetStatus(requestID)
```

### Number Lookup

### Search Phone Number Details

```go
package direct7
apiToken := "Your Api Token"
client := NewClient(apiToken)
recipient := "+919999XXXXXX"
response, err := client.numberLookup.SearchNumberDetails(recipient)
```

### Viber

### Send Viber Message

```go
package direct7
apiToken := "Your Api Token"
client := NewClient(apiToken)
recipients := []string{"+919999XXXXXX"}
content := "Test Viber message"
label := "PROMOTION"
originator := "INFO2WAY"
callBackURL := "https://example.com/callback"

response, err := client.viber.SendViberMessage(recipients, content, label, originator, callBackURL)

```


### Check Viber Request Status

```go
package direct7
apiToken := "Your Api Token"
client := NewClient(apiToken)
requestID := "001ff613-de30-4f82-81f6-1fe944b8f61b"
statusResponse, err := client.viber.GetStatus(requestID)
```

### Slack

### Send Slack Message

```go
package direct7
apiToken := "Your Api Token"
client := NewClient(apiToken)
content := "Test message content"
workSpaceName := "D7-dev"
channelName := "random"
reportURL := "https://example.com/report"

response, err := client.slack.SendSlackMessage(content, workSpaceName, channelName, reportURL)

```


### Check Slack Request Status

```go
    package direct7
    apiToken := "Your Api Token"
    client := NewClient(apiToken)
    requestID := "001ff613-de30-4f82-81f6-1fe944b8f61b"
    statusResponse, err := client.slack.GetStatus(requestID)
```


## FAQ

### How do I get my API token?

You can get your API token from the Direct7 dashboard. If you don't have an account yet, you can create one for free.

### Supported Go versions

The SDK supports go 1.18 and higher.

### Supported APIs

As of now, the SDK supports the following APIs:

| API                    |        Supported?        |
|------------------------|:------------------------:|
| SMS API                |            ✅             |
| Verify API             |            ✅             |
| Whatsapp API           |            ✅             |
| Number Lookup API      |            ✅             |
| Viber API              |            ✅             |
| Slack API              |            ✅             |

### How do I get started?

You can find the platform documentation @ [Direct7 Docs](https://d7networks.com/docs/).

### How do I get help?

If you need help using the SDK, you can create an issue on GitHub or email to support@d7networks.com

## Contributing

We welcome contributions to the Direct7 Go SDK. If you have any ideas for improvements or bug fixes, please feel
free to create an issue on GitHub.
