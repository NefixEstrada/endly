package ses

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/viant/endly"
	"github.com/viant/endly/system/cloud/aws"
)



var clientKey = (*ses.SES)(nil)

func setClient(context *endly.Context, rawRequest map[string]interface{}) error {
	config, err := aws.GetOrCreateAwsConfig(context, rawRequest, clientKey)
	if err != nil || config == nil {
		return err
	}
	session := session.Must(session.NewSession())
	client :=  ses.New(session, config)
	return context.Put(clientKey, client)
}


func getClient(context *endly.Context) (interface{}, error)  {
	client :=  &ses.SES{}
	if !context.GetInto(clientKey, &client) {
		return nil, fmt.Errorf("unable to locate client %T, please add Credentials atribute ", client)
	}
	return client, nil
}
