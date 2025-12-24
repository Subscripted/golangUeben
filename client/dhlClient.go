package client

type DHLClient struct {
	clientId     string
	clientSecret string
}

func construct(clientId string, clientSecret string) *DHLClient {
	return &DHLClient{
		clientId:     "",
		clientSecret: "",
	}
}

func callAPI(Data []string) string {

}
