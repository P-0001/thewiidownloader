package tls

import (
	tlsclient "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
)

func GetClient() tlsclient.HttpClient {
	jar := tlsclient.NewCookieJar()
	options := []tlsclient.HttpClientOption{
		tlsclient.WithTimeoutSeconds(30),
		tlsclient.WithClientProfile(profiles.Chrome_124),
		//	tls_client.WithNotFollowRedirects(),
		tlsclient.WithCookieJar(jar), // create cookieJar instance and pass it as argument
	}

	client, err := tlsclient.NewHttpClient(tlsclient.NewNoopLogger(), options...)

	if err != nil {
		panic(err)
	}

	return client
}
